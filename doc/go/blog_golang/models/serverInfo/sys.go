package serverInfo

type Sys struct {
	ComputerName string `json:"computerName"`
	ComputerIp   string `json:"computerIp"`
	UserDir      string `json:"userDir"`
	OsName       string `json:"osName"`
	OsArch       string `json:"osArch"`
}
