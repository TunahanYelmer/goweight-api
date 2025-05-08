package models

import (
	"time"
)

type Client struct {
	ID         string          `json:"id"` // Unique identifier for the client
	MACAddress string          `json:"mac_address"`
	CreatedAt  time.Time       `json:"created_at"`
	Readings   []WeightReading `json:"readings"` // History of weight readings
}

func NewClient(id string, address string, readings []WeightReading) *Client {

	return &Client{
		ID:         id,
		MACAddress: address,
		CreatedAt:  time.Now(),
		Readings:   readings, // History of weight readings
	}

}
func (c *Client) UpdateClient(id string, address string, readings []WeightReading) {
	c.ID = id
	c.MACAddress = address
	c.Readings = readings

}
