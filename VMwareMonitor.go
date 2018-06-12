package main

import (
	"./ConfigReader"
	"./TelegramRobot"
	"./DiskMonitor"
	"./ServerMonitor"
	"time"
	"sync"
)

var ConnMutex *sync.Mutex
var AlertMessageChannel chan string

const (
  DISK_MONITOR = "diskMonitor"
  SERVER_MONITOR = "serverMonitor"
  TG_ROBOT  = "tgRobot"
)
type ModuleInterface interface {
    Start(*ConfigReader.ServerMonitorConfiguration, *chan string)
}

func ModuleSelector(module string) ModuleInterface {
  switch module {
  case DISK_MONITOR:
  	return new(DiskMonitor.DiskMonitor)
  case SERVER_MONITOR:
  	return new(ServerMonitor.UserPackage)
  case TG_ROBOT:
  	return new(TelegramRobot.Robot)
  default:
    return nil
  }
}

func initGlobalParameter(){
	ConnMutex = new (sync.Mutex)
	AlertMessageChannel = make(chan string, 1)
}

func main(){
	configStructure, result := ConfigReader.ReadConfigJSONFile()
	if !result {
		panic("Parameter Error")
		return
	}
	initGlobalParameter()
	ModuleSelector(DISK_MONITOR).Start(&configStructure, &AlertMessageChannel)
	ModuleSelector(SERVER_MONITOR).Start(&configStructure, &AlertMessageChannel)
	ModuleSelector(TG_ROBOT).Start(&configStructure, &AlertMessageChannel)

	for{
		time.Sleep(1 * time.Second)
	}
}
