package forwarder

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/datawire/dlib/dlog"
)

type udp struct {
	base
}

func newUDP(listen *net.UDPAddr, targetHost string, targetPort uint16) Interceptor {
	return &udp{base: base{
		listenAddr: listen,
		targetHost: targetHost,
		targetPort: targetPort,
	}}
}

func (f *udp) Serve(ctx context.Context) error {
	f.mu.Lock()
	// Set up listener lifetime (same as the overall forwarder lifetime)
	f.ctx, f.lCancel = context.WithCancel(ctx)
	f.ctx = dlog.WithField(f.ctx, "lis", f.listenAddr.String())
	targetHost := f.targetHost
	targetPort := f.targetPort
	ctx = f.ctx
	a := f.listenAddr
	f.mu.Unlock()

	dlog.Infof(ctx, "Forwarding udp from %s", f.listenAddr)
	defer dlog.Infof(ctx, "Done forwarding udp from %s", f.listenAddr)

	lc := net.ListenConfig{}
	agentConn, err := lc.ListenPacket(ctx, a.Network(), a.String())
	if err != nil {
		return err
	}

	targetAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", targetHost, targetPort))
	if err != nil {
		return fmt.Errorf("error on resolve(%s:%d): %w", targetHost, targetPort, err)
	}

	dialer := net.Dialer{}
	tcc, err := dialer.DialContext(ctx, "udp", targetAddr.String())
	if err != nil {
		return fmt.Errorf("error on dial: %w", err)
	}

	targetConn := tcc.(net.PacketConn)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer dlog.Debugf(ctx, "closing target->client")
		ch := make(chan udpReadResult)
		go udpReader(ctx, targetConn, ch)
		for {
			select {
			case <-ctx.Done():
				return
			case rr := <-ch:
				_, err = agentConn.WriteTo(rr.payload, rr.addr)
				if err != nil {
					dlog.Errorf(ctx, "writing to %s: %v", rr.addr, err)
					return
				}
			}
		}
	}()
	go func() {
		defer wg.Done()
		defer dlog.Debugf(ctx, "closing client->target")
		defer targetConn.Close()
		ch := make(chan udpReadResult)
		go udpReader(ctx, agentConn, ch)
		for {
			select {
			case <-ctx.Done():
				return
			case rr := <-ch:
				f.mu.Lock()
				intercept := f.intercept
				f.mu.Unlock()
				if intercept != nil {
					_ = agentConn.SetReadDeadline(time.Time{})
					err = f.interceptConn(ctx, agentConn.(net.Conn), rr.addr, rr.payload, intercept)
				} else {
					_, err := targetConn.WriteTo(rr.payload, rr.addr)
					if err != nil {
						dlog.Errorf(ctx, "writing to %s: %v", rr.addr, err)
						return
					}
				}
			}
		}
	}()
	wg.Wait()
	return nil
}

type udpReadResult struct {
	payload []byte
	addr    net.Addr
}

func udpReader(ctx context.Context, conn net.PacketConn, ch chan<- udpReadResult) {
	isTimeout := func(err error) bool {
		var opErr *net.OpError
		return errors.As(err, &opErr) && opErr.Timeout()
	}

	buf := [0x10000]byte{}
	for ctx.Err() == nil {
		_ = conn.SetReadDeadline(time.Now().Add(time.Second))
		n, addr, err := conn.ReadFrom(buf[:])
		if n > 0 {
			ch <- udpReadResult{buf[:n], addr}
		}
		if err != nil && !isTimeout(err) {
			if !errors.Is(err, net.ErrClosed) {
				dlog.Errorf(ctx, "reading from %s: %v", conn.LocalAddr(), err)
			}
			break
		}
	}
}
