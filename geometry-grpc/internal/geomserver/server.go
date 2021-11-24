/*
 * The package implements a basic examle of a geometry grpc server
 */

package geomserver

/*
 * HTTP Server Package
 */
import (
	"context"
	"sync"
	"time"

	"github.com/muzammilar/examples-go/geometry-grpc/pkg/grpcserver"
	"github.com/muzammilar/examples-go/geometry-grpc/protos/shapecalc"
	"github.com/sirupsen/logrus"
)

/*
 * Public Functions
 */

// Create an HTTP Server and register all the required endpoints
func Serve(wg *sync.WaitGroup, ctx context.Context, port int, certFile string, keyFile string, logger *logrus.Logger) {

	// if there's a wait group implemented, then notify about the thread finishing
	if wg != nil {
		defer wg.Done()
	}

	// create a grpc server
	serverRegistrar := grpcserver.CreateServerWithStatsAndTLS(certFile, keyFile, logger)

	// create a handler for geometry server
	geometryHandler := &GeometryServer{
		logger: logger,
	}

	// register the geometry handler with the server
	shapecalc.RegisterGeometryServer(serverRegistrar, geometryHandler)

	// create a tcp listener
	listener := grpcserver.TCPListener("", port)

	// register the shutdown handler
	go func() {
		<-ctx.Done() // blocking wait
		grpcserver.ShutDownServerWithTimeout(serverRegistrar, 20*time.Second)
	}()

	// start the registrar server/registrar
	if err := serverRegistrar.Serve(listener); err != nil {
		// TODO: log error
	}
}

/*
 * Private Functions
 */
