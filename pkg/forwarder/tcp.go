package forwarder

import (
	"context"
	"fmt"
	"io"
	"net"

	"github.com/datawire/dlib/dlog"
	"github.com/telepresenceio/telepresence/rpc/v2/manager"
)

type tcp struct {
	base
	tCtx    context.Context
	tCancel context.CancelFunc
}

func newTCP(listen net.Addr, targetHost string, targetPort uint16) Interceptor {
	return &tcp{
		base: base{
			listenAddr: listen,
			targetHost: targetHost,
			targetPort: targetPort,
		},
	}
}

func (f *tcp) Serve(ctx context.Context) error {
	listener, err := f.Listen(ctx)
	if err != nil {
		return err
	}
	defer listener.Close()

	dlog.Debugf(ctx, "Forwarding from %s", f.listenAddr.String())
	defer dlog.Debugf(ctx, "Done forwarding from %s", f.listenAddr.String())

	go func() {
		<-ctx.Done()
		listener.Close()
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		conn, err := listener.AcceptTCP()
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			dlog.Infof(ctx, "Error on accept: %+v", err)
			continue
		}
		go func() {
			if err := f.forwardConn(conn); err != nil {
				dlog.Error(ctx, err)
			}
		}()
	}
}

func (f *tcp) Listen(ctx context.Context) (*net.TCPListener, error) {
	f.mu.Lock()

	// Set up listener lifetime (same as the overall forwarder lifetime)
	f.ctx, f.lCancel = context.WithCancel(ctx)
	f.ctx = dlog.WithField(f.ctx, "lis", f.listenAddr.String())

	// Set up target lifetime
	f.tCtx, f.tCancel = context.WithCancel(f.ctx)
	listenAddr := f.listenAddr

	f.mu.Unlock()
	return net.ListenTCP("tcp", listenAddr.(*net.TCPAddr))
}

func (f *tcp) SetIntercepting(intercept *manager.InterceptInfo) {
	f.base.SetIntercepting(intercept)
	f.mu.Lock()

	// Drop existing connections
	f.tCancel()

	// Set up new target and lifetime
	f.tCtx, f.tCancel = context.WithCancel(f.ctx)
	f.intercept = intercept
	f.mu.Unlock()
}

func (f *tcp) forwardConn(clientConn *net.TCPConn) error {
	f.mu.Lock()
	ctx := f.tCtx
	targetHost := f.targetHost
	targetPort := f.targetPort
	intercept := f.intercept
	f.mu.Unlock()
	if intercept != nil {
		return f.interceptConn(ctx, clientConn, clientConn.RemoteAddr(), nil, intercept)
	}

	targetAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", targetHost, targetPort))
	if err != nil {
		return fmt.Errorf("error on resolve(%s:%d): %w", targetHost, targetPort, err)
	}

	ctx = dlog.WithField(ctx, "client", clientConn.RemoteAddr().String())
	ctx = dlog.WithField(ctx, "target", targetAddr.String())

	dlog.Debug(ctx, "Forwarding...")
	defer dlog.Debug(ctx, "Done forwarding")

	defer clientConn.Close()

	targetConn, err := net.DialTCP("tcp", nil, targetAddr)
	if err != nil {
		return fmt.Errorf("error on dial: %w", err)
	}
	defer targetConn.Close()

	done := make(chan struct{})

	go func() {
		if _, err := io.Copy(targetConn, clientConn); err != nil {
			dlog.Debugf(ctx, "Error clientConn->targetConn: %+v", err)
		}
		_ = targetConn.CloseWrite()
		done <- struct{}{}
	}()
	go func() {
		if _, err := io.Copy(clientConn, targetConn); err != nil {
			dlog.Debugf(ctx, "Error targetConn->clientConn: %+v", err)
		}
		_ = clientConn.CloseWrite()
		done <- struct{}{}
	}()

	// Wait for both sides to close the connection
	for numClosed := 0; numClosed < 2; {
		select {
		case <-ctx.Done():
			return nil
		case <-done:
			numClosed++
		}
	}
	return nil
}
