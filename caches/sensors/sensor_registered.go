package sensors

type RegisteredSensor struct {
	Uuid     string `json:"uuid"`
	IsActive bool   `json:"active"`
	Type     string `json:"type"`
}
