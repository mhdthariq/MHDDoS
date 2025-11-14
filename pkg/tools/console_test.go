package tools

import (
	"testing"
)

// TestNetworkStats tests the NetworkStats struct
func TestNetworkStats(t *testing.T) {
	stats := NetworkStats{
		BytesSent:     1024,
		BytesReceived: 2048,
		PacketsSent:   10,
		PacketsRecv:   20,
	}

	if stats.BytesSent != 1024 {
		t.Errorf("NetworkStats.BytesSent = %d; want 1024", stats.BytesSent)
	}

	if stats.BytesReceived != 2048 {
		t.Errorf("NetworkStats.BytesReceived = %d; want 2048", stats.BytesReceived)
	}

	if stats.PacketsSent != 10 {
		t.Errorf("NetworkStats.PacketsSent = %d; want 10", stats.PacketsSent)
	}

	if stats.PacketsRecv != 20 {
		t.Errorf("NetworkStats.PacketsRecv = %d; want 20", stats.PacketsRecv)
	}
}

// TestIPInfo tests the IPInfo struct
func TestIPInfo(t *testing.T) {
	info := IPInfo{
		Success: true,
		IP:      "192.168.1.1",
		Country: "United States",
		City:    "New York",
		Region:  "NY",
		ISP:     "Example ISP",
		Org:     "Example Org",
	}

	if !info.Success {
		t.Error("IPInfo.Success should be true")
	}

	if info.IP != "192.168.1.1" {
		t.Errorf("IPInfo.IP = %s; want 192.168.1.1", info.IP)
	}

	if info.Country != "United States" {
		t.Errorf("IPInfo.Country = %s; want United States", info.Country)
	}

	if info.City != "New York" {
		t.Errorf("IPInfo.City = %s; want New York", info.City)
	}

	if info.Region != "NY" {
		t.Errorf("IPInfo.Region = %s; want NY", info.Region)
	}

	if info.ISP != "Example ISP" {
		t.Errorf("IPInfo.ISP = %s; want Example ISP", info.ISP)
	}

	if info.Org != "Example Org" {
		t.Errorf("IPInfo.Org = %s; want Example Org", info.Org)
	}
}

// TestUpdateNetworkStats tests the updateNetworkStats function
func TestUpdateNetworkStats(t *testing.T) {
	// Reset stats
	currentStats = NetworkStats{}

	// Call updateNetworkStats
	updateNetworkStats()

	// Check that stats were updated
	if currentStats.BytesSent == 0 {
		t.Error("updateNetworkStats() should update BytesSent")
	}

	if currentStats.BytesReceived == 0 {
		t.Error("updateNetworkStats() should update BytesReceived")
	}

	if currentStats.PacketsSent == 0 {
		t.Error("updateNetworkStats() should update PacketsSent")
	}

	if currentStats.PacketsRecv == 0 {
		t.Error("updateNetworkStats() should update PacketsRecv")
	}

	// Save the current stats
	oldStats := currentStats

	// Call again and verify stats increased
	updateNetworkStats()

	if currentStats.BytesSent <= oldStats.BytesSent {
		t.Error("updateNetworkStats() should increase BytesSent on subsequent calls")
	}

	if currentStats.BytesReceived <= oldStats.BytesReceived {
		t.Error("updateNetworkStats() should increase BytesReceived on subsequent calls")
	}

	if currentStats.PacketsSent <= oldStats.PacketsSent {
		t.Error("updateNetworkStats() should increase PacketsSent on subsequent calls")
	}

	if currentStats.PacketsRecv <= oldStats.PacketsRecv {
		t.Error("updateNetworkStats() should increase PacketsRecv on subsequent calls")
	}
}

// TestGlobalStats tests the global stats variables
func TestGlobalStats(t *testing.T) {
	// Reset global variables
	lastNetStats = NetworkStats{}
	currentStats = NetworkStats{}
	statsRecorded = false

	// Verify initial state
	if statsRecorded {
		t.Error("statsRecorded should be false initially")
	}

	if lastNetStats.BytesSent != 0 || lastNetStats.BytesReceived != 0 {
		t.Error("lastNetStats should be zero initially")
	}

	if currentStats.BytesSent != 0 || currentStats.BytesReceived != 0 {
		t.Error("currentStats should be zero initially")
	}

	// Set statsRecorded
	statsRecorded = true

	if !statsRecorded {
		t.Error("statsRecorded should be true after setting")
	}
}

// BenchmarkUpdateNetworkStats benchmarks the updateNetworkStats function
func BenchmarkUpdateNetworkStats(b *testing.B) {
	for i := 0; i < b.N; i++ {
		updateNetworkStats()
	}
}
