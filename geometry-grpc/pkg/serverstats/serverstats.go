/*
	The serverstats package implements the grpc/stats.Handler interface and the related read-only functions
*/
package serverstats

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/stats"
)

// NOTE: Since the stats handler GRPCStats is a read only struct, pointer recievers for functions of the interface are NOT allowed.
// GRPCStats implements the grpc/stats.Handler interface: https://pkg.go.dev/google.golang.org/grpc@v1.42.0/stats#Handler
type GRPCStats struct {
	logger *logrus.Logger
}

// Create a new GRPCStats struct. Note that these fields can not be updated/modified in the stats.Handler interface functions
func NewGRPCStats(l *logrus.Logger) GRPCStats {
	return GRPCStats{
		logger: l,
	}
}

// TagRPC can attach some information to the given context.
// The context used for the rest lifetime of the RPC will be derived from
// the returned context.
func (g GRPCStats) TagRPC(ctx context.Context, rtag *stats.RPCTagInfo) context.Context {
	return ctx
}

// HandleRPC processes the RPC stats.
func (g GRPCStats) HandleRPC(context.Context, stats.RPCStats) {}

// TagConn can attach some information to the given context.
// The returned context will be used for stats handling.
// For conn stats handling, the context used in HandleConn for this
// connection will be derived from the context returned.
// For RPC stats handling,
//  - On server side, the context used in HandleRPC for all RPCs on this
// connection will be derived from the context returned.
//  - On client side, the context is not derived from the context returned.
func (g GRPCStats) TagConn(ctx context.Context, ctag *stats.ConnTagInfo) context.Context {
	return ctx
}

// HandleConn processes the Conn stats.
func (g GRPCStats) HandleConn(context.Context, stats.ConnStats) {}
