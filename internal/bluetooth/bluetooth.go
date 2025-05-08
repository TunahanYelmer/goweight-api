package bt

import (
	"tinygo.org/x/bluetooth"
)

// Declare the global Adapter variable as a pointer to bluetooth.Adapter
var Adapter *bluetooth.Adapter = &bluetooth.Adapter{}

// InitAdapter initializes the Bluetooth adapter and returns the adapter and any error encountered.
func InitAdapter() (*bluetooth.Adapter, error) {
	// Enable the Bluetooth adapter
	err := Adapter.Enable()
	if err != nil {
		return nil, err
	}
	return Adapter, nil
}
