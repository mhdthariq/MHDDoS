//go:build linux
// +build linux

package tools

import (
	"os"
	"testing"
)

// TestLinuxProcNetDevExists tests that /proc/net/dev exists on Linux
func TestLinuxProcNetDevExists(t *testing.T) {
	_, err := os.Stat("/proc/net/dev")
	if err != nil {
		t.Skipf("Skipping test: /proc/net/dev not accessible: %v", err)
	}
}

// TestLinuxNetworkStatsFromProcNetDev tests reading from /proc/net/dev
func TestLinuxNetworkStatsFromProcNetDev(t *testing.T) {
	// Check if /proc/net/dev exists
	_, err := os.Stat("/proc/net/dev")
	if err != nil {
		t.Skipf("Skipping test: /proc/net/dev not accessible: %v", err)
	}

	// Reset stats
	currentStats = NetworkStats{}
	
	// Call updateNetworkStats
	updateNetworkStats()

	// On Linux, we should get actual values from /proc/net/dev
	// The values should be greater than 0 if there's been any network activity
	t.Logf("Linux stats - BytesSent: %d, BytesReceived: %d, PacketsSent: %d, PacketsRecv: %d",
		currentStats.BytesSent, currentStats.BytesReceived, currentStats.PacketsSent, currentStats.PacketsRecv)

	// Verify that we're getting reasonable values
	// At minimum, there should be some network activity (even loopback is excluded, 
	// so we should see some real interface data)
	if currentStats.BytesSent < 0 || currentStats.BytesReceived < 0 {
		t.Error("Network stats should not be negative")
	}
}

// TestLinuxProcNetDevParsing tests the parsing logic for /proc/net/dev
func TestLinuxProcNetDevParsing(t *testing.T) {
	// Check if /proc/net/dev exists
	file, err := os.Open("/proc/net/dev")
	if err != nil {
		t.Skipf("Skipping test: /proc/net/dev not accessible: %v", err)
	}
	defer file.Close()

	// Reset and get first reading
	currentStats = NetworkStats{}
	updateNetworkStats()
	firstReading := currentStats

	// Get second reading after a small delay
	updateNetworkStats()
	secondReading := currentStats

	// Verify readings are consistent (monotonically non-decreasing)
	if secondReading.BytesSent < firstReading.BytesSent {
		t.Errorf("BytesSent decreased: %d -> %d", firstReading.BytesSent, secondReading.BytesSent)
	}
	if secondReading.BytesReceived < firstReading.BytesReceived {
		t.Errorf("BytesReceived decreased: %d -> %d", firstReading.BytesReceived, secondReading.BytesReceived)
	}
	if secondReading.PacketsSent < firstReading.PacketsSent {
		t.Errorf("PacketsSent decreased: %d -> %d", firstReading.PacketsSent, secondReading.PacketsSent)
	}
	if secondReading.PacketsRecv < firstReading.PacketsRecv {
		t.Errorf("PacketsRecv decreased: %d -> %d", firstReading.PacketsRecv, secondReading.PacketsRecv)
	}
}

// TestLinuxLoopbackExclusion tests that loopback interface is properly excluded
func TestLinuxLoopbackExclusion(t *testing.T) {
	// This test verifies that the loopback interface (lo) is excluded from stats
	// We can't easily test this without mocking, but we can at least verify
	// the function runs without errors
	currentStats = NetworkStats{}
	updateNetworkStats()
	
	// The function should complete without panic
	t.Log("Successfully updated network stats on Linux")
}
