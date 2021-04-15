package model

type Device struct {
	// Sensor map[string]string `json:"sensor"`
	// ID     string `json:"id"`
	// IP     string `json:"ip"`
	// Sensor string `json:"sensor"`
	Ip       string `json:"ip"`
	Location string `json:"location"`
	Server   string `json:"server"`
	Type     string `json:"type"`
}

type Sensor struct {
	// ID    string `json:"id"` //device
	// Type  string `json:"type"`
	// Value string `json:"value"`
	Id    string  `json:"id"`
	Ip    string  `json:"ip"`
	Type  string  `json:"type"`
	Value float64 `json:"value"`
}
