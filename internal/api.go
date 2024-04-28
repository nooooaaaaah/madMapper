package api

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/grandcat/zeroconf"
	"github.com/nooooaaaaah/madMapper/config"
)

type Device struct {
	IPAddress  string   `json:"ip_address"`
	Hostname   string   `json:"hostname"`
	ServiceURL string   `json:"service_url"`
	DeviceType string   `json:"device_type"`
	TXTRecords []string `json:"txt_records"` // Store raw TXT records for more details
}

func DiscoverMatterDevices(serviceType string, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // Clean up resources after the timeout.

	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		config.LogError("Failed to initialize resolver: ", err)
		return
	}

	entries := make(chan *zeroconf.ServiceEntry)
	go processEntries(entries)

	// Allow dynamic service type browsing
	if err = resolver.Browse(ctx, serviceType, "local.", entries); err != nil {
		config.LogError("Failed to browse for services: ", err)
	}

	<-ctx.Done() // Block until the context expires
}

func processEntries(results <-chan *zeroconf.ServiceEntry) {
	for entry := range results {
		if len(entry.AddrIPv4) == 0 {
			config.LogInfo("No IPv4 address found for entry: ", entry.HostName)
			continue
		}

		device := Device{
			IPAddress:  entry.AddrIPv4[0].String(),
			Hostname:   entry.HostName,
			ServiceURL: extractServiceURL(entry.Text),
			DeviceType: extractDeviceType(entry.Text),
			TXTRecords: entry.Text,
		}

		deviceJSON, err := json.Marshal(device)
		if err != nil {
			config.LogError("Error marshalling device info: ", err)
			continue
		}

		config.LogInfo("Discovered device: ", string(deviceJSON))
	}
}

func extractServiceURL(txtRecords []string) string {
	for _, txt := range txtRecords {
		if strings.HasPrefix(txt, "url=") {
			return strings.SplitN(txt, "=", 2)[1]
		}
	}
	return "unknown"
}

func extractDeviceType(txtRecords []string) string {
	for _, txt := range txtRecords {
		if strings.HasPrefix(txt, "type=") {
			return strings.SplitN(txt, "=", 2)[1]
		}
	}
	return "unknown"
}
