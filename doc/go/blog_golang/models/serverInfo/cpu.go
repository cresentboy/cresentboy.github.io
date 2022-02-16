package serverInfo

type Cpu struct {
	CpuNum int     `json:"cpuNum"`
	Total  float64 `json:"total"`
	Sys    float64 `json:"sys"`
	Used   float64 `json:"used"`
	Wait   float64 `json:"wait"`
	Free   float64 `json:"free"`
}
