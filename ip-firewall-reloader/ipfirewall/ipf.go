package ipfirewall

import "net"

/****************/
/*     Mode     */
/****************/
type FWMode int

const (
	Unknown FWMode = iota
	AllowList
	BlockList
)

const (
	unknownStr = "unknown"
	allowStr   = "allow"
	blockStr   = "block"
)

func (f FWMode) String() string {
	return [...]string{unknownStr, allowStr, blockStr}[f]
}

/****************/
/*  IPFirewall  */
/****************/

// IPFirewall is a dummy data structure containing some allow lists and some block lists
// Since this is a POC we use a list of CIDR IP addresses. It is not necessary for IP ranges to fall into a single subnet.
// A better structure for this approach is probably a radix tree/ip tree
type IPFirewall struct {
	ipList        []*net.IPNet
	mode          FWMode
	versionNumber uint64
}
