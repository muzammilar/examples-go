/*
 * The package implements a basic examle of an http server
 */

package httpserver

/*
 * HTTP Server Package
 */
import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"

	"github.com/muzammilar/examples-go/geometry-grpc/protos/shape"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

/*
 * Constants
 */

// Create an HTTP Server and register all the required endpoints
func StartServer(wg *sync.WaitGroup, addr string, ctx context.Context, logger *logrus.Logger) {

	// if there's a wait group implemented, then notify about the thread finishing
	if wg != nil {
		defer wg.Done()
	}

	// create a new HTTP Mux handler
	mux := http.NewServeMux()

	// a basic helloworld handler
	mux.HandleFunc("/hello", hellogrpc)

	// a basic hellojson handler
	mux.HandleFunc("/json", hellojson)

	// prometheus handler
	mux.Handle("/metrics", promhttp.Handler())

	// create an http server with mux handler
	var server *http.Server = &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// setup http shutdown handler
	go func() {
		// Wait for context to be done before shutting down
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			logger.Warn("HTTP Server failed to shutdown: %#v", err)
		}
	}()

	// start the http server and ignore 'server closed' errors
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Warn("HTTP Server failed to listen and serve: %#v", err)
	}
	// server shutdown is complete
}

/*
 * Private Functions
 */

// An example of handling a basic hello world response
func hellogrpc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello gRPC!\n")
}

// An example of handling a basic hello world response with json (and protobufs)
func hellojson(w http.ResponseWriter, req *http.Request) {
	// create a random cuboid structure
	cuboid := &shape.Cuboid{
		Id: &shape.Identifier{
			Id: int64(rand.Uint32()),
		},
		Length: int64(10 + rand.Uint32()%25),
		Width:  int64(1 + rand.Uint32()%10),
		Height: int64(1 + rand.Uint32()%25),
	}
	// The performance of structs to json is generally slow since the json package is slow (and reflection is often involved)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cuboid)
}
