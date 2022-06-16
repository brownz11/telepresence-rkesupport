package forwarder

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/blang/semver"

	"github.com/datawire/dlib/dlog"
	"github.com/telepresenceio/telepresence/rpc/v2/manager"
	"github.com/telepresenceio/telepresence/v2/pkg/iputil"
	"github.com/telepresenceio/telepresence/v2/pkg/restapi"
	"github.com/telepresenceio/telepresence/v2/pkg/tunnel"
)

type Interceptor interface {
	io.Closer
	SetManager(sessionInfo *manager.SessionInfo, manager manager.ManagerClient, version semver.Version)
	InterceptId() (id string)
	InterceptInfo() *restapi.InterceptInfo
	Serve(ctx context.Context) error
	SetIntercepting(intercept *manager.InterceptInfo)
	Target() (string, uint16)
}

type base struct {
	mu sync.Mutex

	ctx        context.Context
	lCancel    context.CancelFunc
	listenAddr net.Addr

	targetHost string
	targetPort uint16

	manager     manager.ManagerClient
	sessionInfo *manager.SessionInfo

	intercept  *manager.InterceptInfo
	mgrVersion semver.Version
}

func NewInterceptor(addr net.Addr, targetHost string, targetPort uint16) Interceptor {
	switch addr := addr.(type) {
	case *net.TCPAddr:
		return newTCP(addr, targetHost, targetPort)
	case *net.UDPAddr:
		return newUDP(addr, targetHost, targetPort)
	default:
		panic(fmt.Errorf("unsupported net.Addr type %T", addr))
	}
}

func (f *base) SetManager(sessionInfo *manager.SessionInfo, manager manager.ManagerClient, version semver.Version) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.sessionInfo = sessionInfo
	f.manager = manager
	f.mgrVersion = version
}

func (f *base) Close() error {
	f.lCancel()
	return nil
}

func (f *base) Target() (string, uint16) {
	f.mu.Lock()
	defer f.mu.Unlock()

	return f.targetHost, f.targetPort
}

func (f *base) InterceptInfo() *restapi.InterceptInfo {
	ii := &restapi.InterceptInfo{}
	f.mu.Lock()
	if f.intercept != nil {
		ii.Intercepted = true
		ii.Metadata = f.intercept.Metadata
	}
	f.mu.Unlock()
	return ii
}

func (f *base) InterceptId() (id string) {
	f.mu.Lock()
	if f.intercept != nil {
		id = f.intercept.Id
	}
	f.mu.Unlock()
	return id
}

func (f *base) SetIntercepting(intercept *manager.InterceptInfo) {
	f.mu.Lock()
	defer f.mu.Unlock()

	iceptInfo := func(ii *manager.InterceptInfo) string {
		is := ii.Spec
		return fmt.Sprintf("'%s' (%s:%d)", is.Name, is.Client, is.TargetPort)
	}
	if intercept == nil {
		if f.intercept == nil {
			return
		}
		dlog.Debugf(f.ctx, "Forward target changed from intercept %s to %s:%d", iceptInfo(f.intercept), f.targetHost, f.targetPort)
	} else {
		if f.intercept == nil {
			dlog.Debugf(f.ctx, "Forward target changed from %s:%d to intercept %s", f.targetHost, f.targetPort, iceptInfo(intercept))
		} else {
			if f.intercept.Id == intercept.Id {
				return
			}
			dlog.Debugf(f.ctx, "Forward target changed from intercept %s to intercept %q", iceptInfo(f.intercept), iceptInfo(intercept))
		}
	}
	f.intercept = intercept
}

func (f *base) interceptConn(ctx context.Context, conn net.Conn, addr net.Addr, initial []byte, iCept *manager.InterceptInfo) error {
	dlog.Infof(ctx, "Accept got connection from %s", addr)

	srcIp, srcPort, err := iputil.SplitToIPPort(addr)
	if err != nil {
		return fmt.Errorf("failed to parse intercept source address %s", addr)
	}

	spec := iCept.Spec
	destIp := iputil.Parse(spec.TargetHost)
	id := tunnel.NewConnID(tunnel.IPProto(addr.Network()), srcIp, destIp, srcPort, uint16(spec.TargetPort))

	ms, err := f.manager.Tunnel(ctx)
	if err != nil {
		return fmt.Errorf("call to manager.Tunnel() failed. Id %s: %v", id, err)
	}

	s, err := tunnel.NewClientStream(ctx, ms, id, f.sessionInfo.SessionId, time.Duration(spec.RoundtripLatency), time.Duration(spec.DialTimeout))
	if err != nil {
		return err
	}
	if err = s.Send(ctx, tunnel.SessionMessage(iCept.ClientSession.SessionId)); err != nil {
		return fmt.Errorf("unable to send client session id. Id %s: %v", id, err)
	}
	d := tunnel.NewConnEndpoint(s, conn)
	d.Start(ctx)
	if len(initial) > 0 {
		s.Send(ctx, tunnel.NewMessage(tunnel.Normal, initial))
	}
	<-d.Done()
	return nil
}
