package GameFormatPackage


type AccountUserInfo struct {
	ConfigCode       string `json:"configCode"`
	Flag             int    `json:"flag"`
	CustomerID       int 	`json:"customerId"`
	CustomerName     string `json:"customerName"`
	Token            string `json:"token"`
	MachineID        string `json:"machineId"`
	GameServerIP     string `json:"gameServerIp"`
	GameServerPort   string `json:"gameServerPort"`
	GameID           string `json:"gameId"`
	RoomType         string `json:"roomType"`
	RoomName         string `json:"roomName"`
	CustomerNickName string `json:"customerNickName"`
	ServerID         string `json:"serverId"`
	AcctType         string `json:"acctType"`
}