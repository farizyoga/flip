package entity

type Health struct {
	Worker struct {
		Total     int `json:"total"`
		Busy      int `json:"busy"`
		Available int `json:"available"`
	} `json:"worker"`
	Retry struct {
		Inflight int `json:"inflight"`
	} `json:"retry"`
	Usage struct {
		CurrentAllocation uint64 `json:"current_allocation"`
		TotalAllocation   uint64 `json:"total_allocation"`
	} `json:"usage"`
}
