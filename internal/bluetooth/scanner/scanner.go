package bt

import (
	"time"
	bt "weightApi/internal/bluetooth"
	"weightApi/internal/models"

	"tinygo.org/x/bluetooth" // The correct package for Bluetooth Low Energy functionality
)

// ScanDevices scans for nearby BLE devices for the given duration (in seconds)
func ScanDevices(durationSec int) ([]models.BLEDevice, error) {
	var devices []models.BLEDevice

	// Initialize the Bluetooth adapter
	adapter, err := bt.InitAdapter() // Assuming InitAdapter is in the bt package
	if err != nil {
		return nil, err
	}

	// Start scanning for devices
	err = adapter.Scan(func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
		// Add found devices to the list
		devices = append(devices, models.BLEDevice{
			Name:    result.LocalName(),
			Address: result.Address.String(),
			RSSI:    result.RSSI,
		})
	})
	if err != nil {
		return nil, err
	}

	// Sleep for duration of scan, then stop
	time.Sleep(time.Duration(durationSec) * time.Second)
	adapter.StopScan()

	// Return the list of devices
	return devices, nil
}
