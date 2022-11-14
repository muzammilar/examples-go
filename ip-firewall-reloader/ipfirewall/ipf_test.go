package ipfirewall

import (
	"testing"
	"time"
)

/****************/
/*     Tests    */
/****************/

func TestNewIPFirewallIsActive(t *testing.T) {
	i := new(IPFirewall)
	if i.IsActive() {
		t.Fatalf("TestNewIPFirewallIsActive: The default state of the firewall must be `%s`. Found `%s`.", disabledStr, i.mode)
	}
}

func TestVersionIncrements(t *testing.T) {
	ip := new(IPFirewall)
	var versionIncrements uint64 = 1000
	// incremement versions
	var i uint64
	for i = 0; i < versionIncrements; i++ {
		ip.IncVersion()
	}
	// test increments
	if ip.versionNumber != versionIncrements {
		t.Fatalf("TestVersionIncrements: Unexpected number of version increments. Expected `%d`. Found `%d`.", versionIncrements, ip.versionNumber)
	}
}

/****************/
/*  Benchmarks  */
/****************/
func BenchmarkReadVersion(b *testing.B) {

	ip := new(IPFirewall)
	ip.versionNumber = uint64(time.Now().Unix())

	for i := 0; i < b.N; i++ {
		ip.ReadVersion()
	}
}

func BenchmarkParallelReadVersion(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		ip := new(IPFirewall)
		ip.versionNumber = uint64(time.Now().Unix())
		for pb.Next() {
			ip.ReadVersion()
		}
	})
}

func BenchmarkReadEventuallyConsistentVersion(b *testing.B) {

	ip := new(IPFirewall)
	ip.versionNumber = uint64(time.Now().Unix())

	for i := 0; i < b.N; i++ {
		ip.ReadEventuallyConsistentVersion()
	}
}

func BenchmarkParallelReadEventuallyConsistentVersion(b *testing.B) {

	b.RunParallel(func(pb *testing.PB) {
		ip := new(IPFirewall)
		ip.versionNumber = uint64(time.Now().Unix())
		for pb.Next() {
			ip.ReadEventuallyConsistentVersion()
		}
	})
}
