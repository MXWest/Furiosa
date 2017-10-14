package app

type HealthStatusMsg struct {
	Status  string `json:"state"`
	Message string `json:"message"`
}
