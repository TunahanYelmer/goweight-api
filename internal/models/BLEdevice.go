package models

// BLEDevice represents the information about a Bluetooth Low Energy device.
type BLEDevice struct {
	Name    string `json:"name"`    // The name of the BLE device (e.g., "Smart Scale")
	Address string `json:"address"` // The MAC address of the BLE device
	RSSI    int16  `json:"rssi"`    // Received Signal Strength Indicator (RSSI) in dBm
}
