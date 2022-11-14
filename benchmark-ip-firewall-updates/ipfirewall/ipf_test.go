package ipfirewall

import (
	"net"
	"sync/atomic"
	"testing"
	"time"
)

/****************/
/*     Tests    */
/****************/

func TestNewIPFirewallIsActive(t *testing.T) {
	i := NewIPFirewall()
	if i.IsActive() {
		t.Fatalf("TestNewIPFirewallIsActive: The default state of the firewall must be `%s`. Found `%s`.", disabledStr, i.mode)
	}
}

func TestVersionIncrements(t *testing.T) {
	ip := NewIPFirewall()
	var versionIncrements uint64 = 1000
	// incremement versions
	var i uint64
	for i = 0; i < versionIncrements; i++ {
		ip.IncVersion()
	}
	// test increments
	if ip.versionNumber.Load() != versionIncrements {
		t.Fatalf("TestVersionIncrements: Unexpected number of version increments. Expected `%d`. Found `%d`.", versionIncrements, ip.versionNumber.Load())
	}
}

/****************/
/*  Benchmarks  */
/****************/

// Function Benchmarks

func BenchmarkReadVersion(b *testing.B) {

	ip := NewIPFirewall()
	ip.versionNumber.Store(uint64(time.Now().Unix()))

	for i := 0; i < b.N; i++ {
		ip.ReadVersion()
	}
}

func BenchmarkParallelReadVersion(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		ip := NewIPFirewall()
		ip.versionNumber.Store(uint64(time.Now().Unix()))
		for pb.Next() {
			ip.ReadVersion()
		}
	})
}

func BenchmarkReadEventuallyConsistentVersion(b *testing.B) {

	ip := NewIPFirewall()
	ip.versionNumber.Store(uint64(time.Now().Unix()))

	for i := 0; i < b.N; i++ {
		ip.ReadEventuallyConsistentVersion()
	}
}

func BenchmarkParallelReadEventuallyConsistentVersion(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		ip := NewIPFirewall()
		ip.versionNumber.Store(uint64(time.Now().Unix()))
		for pb.Next() {
			ip.ReadEventuallyConsistentVersion()
		}
	})
}

// Integer Benchmarks
func BenchmarkParallelUint64Struct(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		i := &atomic.Uint64{}
		i.Store(uint64(time.Now().Unix()))
		for pb.Next() {
			i.Load()
		}
	})
}

func BenchmarkParallelUint64Function(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		i := uint64(time.Now().Unix())
		for pb.Next() {
			atomic.LoadUint64(&i)
		}
	})
}

// Pointer Benchmarks
func BenchmarkParallelIPListPointer(b *testing.B) {
	_, ipnet, _ := net.ParseCIDR("2001:db8::/32")
	b.RunParallel(func(pb *testing.PB) {
		ipnetptr := &atomic.Pointer[net.IPNet]{}
		ipnetptr.Store(ipnet)
		for pb.Next() {
			ipnetptr.Load()
		}
	})

}

// Context Benchmarks

// Channel Benchmarks

// RW Mutex Benchmarks

// Mutex Benchmarks
