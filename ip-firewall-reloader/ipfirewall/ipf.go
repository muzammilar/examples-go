package ipfirewall

import (
	"net"
	"sync/atomic"
)

/****************/
/*     Mode     */
/****************/
type FWMode int

const (
	ModeDisabled FWMode = iota
	ModeAllow           // for Allowlists
	ModeBlock           // or ModeDeny for Blocklists
)

const (
	disabledStr = "disabled"
	allowStr    = "allow"
	blockStr    = "block/deny"
)

func (f FWMode) String() string {
	return [...]string{disabledStr, allowStr, blockStr}[f]
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

// NewIPFirewall creates a new IP Firewall with a given mode
func NewIPFirewall(m FWMode) *IPFirewall {
	return &IPFirewall{
		mode: m,
	}
}

// IsActive checks if the firewall is either in allow mode or deny mode
func (i *IPFirewall) IsActive() bool {
	return i.mode != ModeDisabled
}

// IncVersion is a thread-safe way to increment the version number of the Firewall mode.
// Note that the version must be incremented after mode updates or updates to the ipList (but not before the updates)
func (i *IPFirewall) IncVersion() {
	atomic.AddUint64(&i.versionNumber, 1)
}

// ReadVersion reads the version number using atomic instructions.
func (i *IPFirewall) ReadVersion() uint64 {
	return atomic.LoadUint64(&i.versionNumber)
}

// ReadEventuallyConsistentVersion reads the version number without using atomic instructions.
// This is not recommended especially for multi-cpu architecture, however, it could be an allowable solution for some applications.
func (i *IPFirewall) ReadEventuallyConsistentVersion() uint64 {
	return i.versionNumber
}
