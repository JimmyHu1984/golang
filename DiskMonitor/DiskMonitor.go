package DiskMonitor

import (
	"time"
	"log"
	"github.com/shirou/gopsutil/disk"
	"strconv"
	"../ConfigReader"
)
var AlertMessageChannel chan string

const (
	Byte  = 1
	KB = 1024 * Byte
	MB = 1024 * KB
	GB = 1024 * MB
)

type DiskMonitor struct{
}

func (monitor *DiskMonitor)getDiskUsage(path string) float64{
	usage ,err := disk.Usage(path)
	if err != nil{
		log.Printf("Get Disk Info Fail")
	}
	return float64(usage.Free)/(GB)
}

//func (monitor *DiskMonitor)Start(path string, durationTime string, warringVolume int, warringMessage string){
func (monitor *DiskMonitor)Start(tmpPara *ConfigReader.ServerMonitorConfiguration, tmpChannel *chan string){
	AlertMessageChannel = *tmpChannel
	tmpTimeDuration, err := time.ParseDuration(tmpPara.DurationTime)
	if err != nil{
		log.Fatal("Parse DurationTime Error")
		return
	}

	go func(){
		for{
			capacity := monitor.getDiskUsage(tmpPara.DiskPath)
			strCap := strconv.FormatFloat(capacity, 'f', 3, 64)+" GB"
			if capacity < float64(tmpPara.WarringVolume) && capacity > 1.0{
				strCap =  tmpPara.WarringMessage + " Warring Disk Rest: " + strCap
				AlertMessageChannel <- strCap
				log.Printf(strCap)
			}else if capacity < 1.0{
				strCap = tmpPara.WarringMessage + " Error Disk Rest: " + strCap
				AlertMessageChannel <- strCap
				log.Printf(strCap)
			}else{
				strCap = tmpPara.WarringMessage + " Monitor Disk Rest: "+ strCap
				log.Println(strCap)
			}
			time.Sleep(time.Nanosecond * tmpTimeDuration)
		}
	}()
}