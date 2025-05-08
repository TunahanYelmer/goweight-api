package models

type Scale struct {
	ID           string   `json:"id"`   // Unique identifier for the scale
	Name         string   `json:"name"` // Name of the scale (e.g., "Smart Scale")
	MACAddress   string   `json:"mac_address"`
	Capabilities []string `json:"capabilities"` // e.g., ["weight", "body_fat", "bmi"]
	Status       bool     `json:"status"`       // Scale connection status (e.g., "connected", "disconnected")

}

func NewScale(id string, name string, address string, status bool) *Scale {

	return &Scale{
		ID:           id,
		Name:         name,
		MACAddress:   address,
		Capabilities: []string{"weight"},
		Status:       status,
	}

}
func (s *Scale) ConnectScale() bool {
	//TODO:Implement connection logic
	return false

}
func (s *Scale) UpdateScale(id string, name string, address string, status bool, capabilities []string) {
	s.ID = id
	s.Name = name
	s.MACAddress = address
	s.Status = status
	s.Capabilities = capabilities
}
func (s *Scale) CheckCapabilites(status bool) []string {
	// TODO: Implement capability checking based on scale status
	return []string{}
}

func (s *Scale) UpdateCapabilities(capabilities []string) {
	// TODO: Implement capability update logic
}
