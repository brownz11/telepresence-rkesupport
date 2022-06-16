package integration_test

import (
	"fmt"
	"net"
	"time"

	"github.com/datawire/dlib/dlog"
	"github.com/telepresenceio/telepresence/v2/integration_test/itest"
)

// Test_MultipleUnnamedServicePorts tests the use-case where multiple services using
// unnamed ports targets different container ports in the same workload and the same
// container.
func (s *connectedSuite) Test_MultipleUnnamedServicePorts() {
	ctx := s.Context()
	dep := "echo-double-one-unnamed"
	s.ApplyTestApp(ctx, dep, "deploy/"+dep)
	defer s.Kubectl(ctx, "delete", "deploy", dep)

	require := s.Require()
	require.NoError(s.Kubectl(ctx, "expose", "deploy", dep, "--port", "80", "--target-port", "8080", "--name", dep+"-"+"80"))
	defer s.Kubectl(ctx, "delete", "svc", dep+"-"+"80")
	require.NoError(s.Kubectl(ctx, "expose", "deploy", dep, "--port", "81", "--target-port", "8081", "--name", dep+"-"+"81"))
	defer s.Kubectl(ctx, "delete", "svc", dep+"-"+"81")
	require.NoError(s.Kubectl(ctx, "expose", "deploy", dep, "--port", "82", "--target-port", "8082", "--name", dep+"-"+"82"))
	defer s.Kubectl(ctx, "delete", "svc", dep+"-"+"82")

	portTest := func(svcPort, targetPort string) {
		ctx := s.Context()
		svc := dep + "-" + svcPort
		localPort, cancel := itest.StartLocalHttpEchoServer(ctx, svc)
		defer cancel()
		itest.TelepresenceOk(ctx, "intercept", "-n", s.AppNamespace(), "-p", fmt.Sprintf("%d:%s", localPort, svcPort), dep)
		defer itest.TelepresenceOk(ctx, "leave", dep+"-"+s.AppNamespace())
		itest.PingInterceptedEchoServer(ctx, svc, svcPort)
	}
	s.Run("port 80", func() {
		portTest("80", "8080")
	})
	s.Run("port 81", func() {
		portTest("81", "8081")
	})
	s.Run("port 82", func() {
		// Must fail. The container doesn't expose port 8082
		ctx := s.Context()
		require := s.Require()
		svcPort := "82"
		svc := dep + "-" + svcPort

		localPort, cancel := itest.StartLocalHttpEchoServer(ctx, svc)
		defer cancel()
		_, _, err := itest.Telepresence(ctx, "intercept", "-n", s.AppNamespace(), "-p", fmt.Sprintf("%d:%s", localPort, svcPort), dep)
		require.Error(err)
	})
}

// Test_MultipleUnnamedServicePorts tests the use-case where multiple services using
// unnamed ports targets different container ports in the same workload and the same
// container.
func (s *connectedSuite) Test_NoContainerPort() {
	ctx := s.Context()
	dep := "echo-no-containerport"
	s.ApplyTestApp(ctx, dep, "deploy/"+dep)
	defer s.Kubectl(ctx, "delete", "deploy", dep)
	require := s.Require()
	require.NoError(s.Kubectl(ctx, "expose", "deploy", dep, "--port", "80", "--target-port", "8080", "--name", dep+"-"+"80"))
	defer s.Kubectl(ctx, "delete", "svc", dep+"-"+"80")
	require.NoError(s.Kubectl(ctx, "expose", "deploy", dep, "--port", "81", "--target-port", "8081", "--name", dep+"-"+"81"))
	defer s.Kubectl(ctx, "delete", "svc", dep+"-"+"81")

	portTest := func(svcPort, targetPort string) {
		ctx := s.Context()
		svc := dep + "-" + svcPort

		localPort, cancel := itest.StartLocalHttpEchoServer(ctx, svc)
		defer cancel()
		itest.TelepresenceOk(ctx, "intercept", "-n", s.AppNamespace(), "-p", fmt.Sprintf("%d:%s", localPort, svcPort), dep)
		defer itest.TelepresenceOk(ctx, "leave", dep+"-"+s.AppNamespace())
		itest.PingInterceptedEchoServer(ctx, svc, svcPort)
	}
	s.Run("port 80", func() {
		portTest("80", "8080")
	})
	s.Run("port 81", func() {
		portTest("81", "8081")
	})
}

func (s *connectedSuite) Test_UnnamedUdpAndTcpPort() {
	ctx := s.Context()
	dep := "echo-udp-tcp-unnamed"
	s.ApplyTestApp(ctx, dep, "deploy/"+dep)
	defer s.Kubectl(ctx, "delete", "deploy", dep)

	require := s.Require()
	require.NoError(s.Kubectl(ctx, "expose", "deploy", dep, "--port", "80", "--protocol", "UDP", "--target-port", "8080", "--name", "echo-udp"))
	defer s.Kubectl(ctx, "delete", "svc", "echo-udp")
	require.NoError(s.Kubectl(ctx, "expose", "deploy", dep, "--port", "80", "--protocol", "TCP", "--target-port", "8080", "--name", "echo-tcp"))
	defer s.Kubectl(ctx, "delete", "svc", "echo-tcp")

	svcPort := "80"
	s.Run("TCP port 80", func() {
		ctx := s.Context()
		localPort, cancel := itest.StartLocalHttpEchoServer(ctx, "echo-tcp")
		defer cancel()
		itest.TelepresenceOk(ctx, "intercept", "-n", s.AppNamespace(), "--service", "echo-tcp", "-p", fmt.Sprintf("%d:%s", localPort, svcPort), dep)
		defer itest.TelepresenceOk(ctx, "leave", dep+"-"+s.AppNamespace())
		itest.PingInterceptedEchoServer(ctx, "echo-tcp", svcPort)
	})

	s.Run("UDP port 80", func() {
		ctx := s.Context()
		require := s.Require()
		localPortCh := make(chan int, 1)

		// Start a local interceptor service that prints some relevant messages
		go func() {
			lc := net.ListenConfig{}
			pc, err := lc.ListenPacket(ctx, "udp", ":0")
			addr := pc.LocalAddr().(*net.UDPAddr)
			localPortCh <- addr.Port
			close(localPortCh)
			require.NoError(err)
			buf := [0x100]byte{}
			for err == nil {
				var rr net.Addr
				var n int
				n, rr, err = pc.ReadFrom(buf[:])
				if n > 0 {
					dlog.Infof(ctx, "Local UDP server received %q", string(buf[0:n]))
					_, werr := pc.WriteTo([]byte(fmt.Sprintf("received %d bytes", n)), rr)
					require.NoError(werr)
				}
			}
		}()
		var localPort int
		select {
		case <-ctx.Done():
			return
		case localPort = <-localPortCh:
		}

		itest.TelepresenceOk(ctx, "loglevel", "trace")
		itest.TelepresenceOk(ctx, "intercept", "-n", s.AppNamespace(), "--service", "echo-udp", "-p", fmt.Sprintf("%d:%s", localPort, svcPort), dep)
		s.CapturePodLogs(ctx, dep, "traffic-agent", s.AppNamespace())
		defer itest.TelepresenceOk(ctx, "leave", dep+"-"+s.AppNamespace())

		var conn net.Conn
		require.Eventually(
			func() bool {
				var err error
				conn, err = net.Dial("udp", fmt.Sprintf("%s.%s:80", "echo-udp", s.AppNamespace()))
				return err == nil
			},
			12*time.Second, // waitFor
			3*time.Second,  // polling interval
			`dial never succeeds`)
		defer conn.Close()

		pingPong := func(msg string) {
			_ = conn.SetDeadline(time.Now().Add(10 * time.Second))
			bm := []byte(msg)
			n, err := conn.Write(bm)
			require.NoError(err)
			require.Equal(len(bm), n)
			buf := [0x100]byte{}
			n, err = conn.Read(buf[:])
			require.NoError(err)
			require.Equal(fmt.Sprintf("received %d bytes", len(bm)), string(buf[0:n]))
		}

		pingPong("12345678")
		pingPong("a slightly longer message")
	})
}
