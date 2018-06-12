package ConfigReader

import (
	"io/ioutil"
	"fmt"
	"github.com/tidwall/gjson"
)

type ServerMonitorConfiguration struct{
	DiskPath 		string
	WarringVolume	int
	DurationTime 	string
	RobotToken		string
	ChatRoomID		string
	WarringMessage	string
	ServerList	 	[]ServerListStructure
}
type ServerListStructure struct{
	IP 			string
	Port		string
}

func ReadConfigJSONFile() (ServerMonitorConfiguration, bool){
	var contentJSON ServerMonitorConfiguration
	configData, err := ioutil.ReadFile("./JSON/config.json")
	if err!=nil {
		fmt.Println(err)
		return contentJSON, false
	}

	jsonString:=string(configData)

	contentJSON.DiskPath = gjson.Get(jsonString, "DiskPath").String()
	contentJSON.WarringVolume = int (gjson.Get(jsonString, "WarringVolume").Int())
	contentJSON.DurationTime = gjson.Get(jsonString, "DurationTime").String()
	contentJSON.RobotToken = gjson.Get(jsonString, "RobotToken").String()
	contentJSON.ChatRoomID = gjson.Get(jsonString, "ChatRoomID").String()
	contentJSON.WarringMessage = gjson.Get(jsonString, "WarringMessage").String()

	serverArray := gjson.Get(jsonString,"ServerList").Array()
	for _, tmpServerInfo := range serverArray{
		var serverElement ServerListStructure
		serverElement.IP = gjson.Get(tmpServerInfo.String(), "IP").String()
		serverElement.Port = gjson.Get(tmpServerInfo.String(), "Port").String()
		contentJSON.ServerList = append(contentJSON.ServerList, serverElement)
	}

	return contentJSON, true
}