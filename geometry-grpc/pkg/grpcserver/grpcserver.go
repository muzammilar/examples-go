/*
	The package grpcserver provides useful set of function to create grpc servers
*/
package grpcserver

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/muzammilar/examples-go/geometry-grpc/pkg/serverstats"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// CreateServerWithStatsAndTLS creates a simple grpc server with TLS and stats collection enabled.
func CreateServerWithStatsAndTLS(certFile string, keyFile string, l *logrus.Logger) *grpc.Server {

	var opts []grpc.ServerOption

	// TLS
	// Generally panic is not a good way to handle errors. Allow it cos PoC
	creds, err := credentials.NewClientTLSFromFile(certFile, keyFile)
	if err != nil {
		panic(err)
	}
	opts = append(opts, grpc.Creds(creds))

	// Stats (by default no stats handler is configured)
	// Note: Stats handlers can be very expensive and slow down grpc streams (especially for small messages)
	statsHandler := serverstats.NewGRPCStats(l)
	opts = append(opts, grpc.StatsHandler(statsHandler))

	return grpc.NewServer(opts...)

}

// ShutDownServerWithTimeout provides a user with a way to shutdown the server safely (if possible).
// It first tries graceful shutdown, if that fails, a shutdown is forced after timeout
func ShutDownServerWithTimeout(s *grpc.Server, t time.Duration) {

	// Have channel to track graceful shutdown
	gracefulClose := make(chan struct{})

	// Start the graceful shutdown in a subroutine
	go func() {
		s.GracefulStop()
		close(gracefulClose)
	}()

	timer := time.NewTimer(t)
	running := true
	// Either wait for timer to trigger or graceful shutdown to complete. Otherwise wait
	for running {
		select {
		// Force stop after timeout
		case <-timer.C:
			s.Stop()
			running = false
		// If the gracefulClose channel is closed
		case <-gracefulClose:
			running = false
		// Frequently re-check
		default:
			time.Sleep(t / 50)
		}
	}
}

// TCPListener creates a TCP listener and returns it
func TCPListener(host string, port int) *net.TCPListener {
	// create a tcp list
	tcpAddr := net.JoinHostPort(host, strconv.Itoa(port))
	// Generally panic is not a good way to handle errors. Allow it cos PoC
	lis, err := net.Listen("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	// panic if it's not a TCP listener
	/*
		switch nLis := lis.(type) {
		case *net.TCPListener:
			return nLis
		default:
			panic(fmt.Errorf("Listener %+v can not be converted to `*net.Listener`", lis))
		}
	*/
	if tLis, ok := lis.(*net.TCPListener); ok {
		return tLis
	}
	panic(fmt.Errorf("Listener %+v can not be converted to `*net.Listener`", lis))
}
