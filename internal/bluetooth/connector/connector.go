package bt

import (
	"errors"
	"log"
	bt "weightApi/internal/bluetooth"
	"weightApi/internal/models"

	"tinygo.org/x/bluetooth"
)

// Connector struct to manage a Bluetooth connection to a BLE device.
type Connector struct {
	Device   models.BLEDevice    // The device we are connected to
	Conn     *bluetooth.Device   // The actual connection to the BLE device
	Services []bluetooth.Service // List of services the device provides (optional)
}

// Connect establishes a connection to a specific BLE device by address.
func (c *Connector) Connect(device models.BLEDevice) error {
	adapter, err := bt.InitAdapter()
	if err != nil {
		return err
	}

	deviceAddress, err := bluetooth.ParseAddress(device.Address)
	if err != nil {
		return err
	}

	c.Conn, err = adapter.Connect(deviceAddress)
	if err != nil {
		return err
	}

	c.Device = device
	log.Printf("Connected to device: %s [%s]", device.Name, device.Address)
	return nil
}

// ConnectByAddress connects to a device by a known address string.
func (c *Connector) ConnectByAddress(address string) (models.BLEDevice, error) {
	adapter, err := bt.InitAdapter()
	if err != nil {
		return models.BLEDevice{}, err
	}

	deviceAddress, err := bluetooth.ParseAddress(address)
	if err != nil {
		return models.BLEDevice{}, err
	}

	conn, err := adapter.Connect(deviceAddress)
	if err != nil {
		return models.BLEDevice{}, err
	}

	c.Conn = conn
	c.Device = models.BLEDevice{
		Name:    "", // Optionally update if you know the name
		Address: address,
		RSSI:    0,
	}

	log.Printf("Connected to device: [%s]", address)
	return c.Device, nil
}

// Disconnect disconnects from the currently connected BLE device.
func (c *Connector) Disconnect() error {
	if c.Conn != nil {
		err := c.Conn.Disconnect()
		if err != nil {
			return err
		}
		log.Printf("Disconnected from device: %s", c.Device.Name)
	}
	return nil
}

// GetServices retrieves the services provided by the connected BLE device.
func (c *Connector) GetServices() ([]bluetooth.Service, error) {
	if c.Conn == nil {
		return nil, errors.New("no device connected")
	}

	services, err := c.Conn.DiscoverServices()
	if err != nil {
		return nil, err
	}

	c.Services = services
	return services, nil
}
