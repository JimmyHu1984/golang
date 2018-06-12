package ServerMonitor

import (
	"fmt"
	"encoding/binary"
	"net"
	"bytes"
	"./GameFormatPackage"
	"./MAIN_CMD"
	"./MAIN_CMD/MDM_KN_COMMAND"
	"../ConfigReader"
	"time"
	"sync"
)

var ConnMutex *sync.Mutex
var AlertMessageChannel chan string
const (
	CONNECT_STATUS_OFFLINE = 0
	CONNECT_STATUS_ONLINE = 1
)

type UserPackage struct{
	AccountInfo 	GameFormatPackage.AccountUserInfo
	Conn 			*net.TCPConn
	ConnectStatus	int
	EchoNum			uint16
	SendBuffer		map[uint16] bool //for Echo Test
}

func (user *UserPackage) SetConn(conn *net.TCPConn){
	user.Conn = conn
}

func (user *UserPackage)CloseConnection() {
	if user.ConnectStatus == CONNECT_STATUS_ONLINE {
		ConnMutex.Lock()
		user.Conn.Close()
		user.Conn = nil
		user.ConnectStatus = CONNECT_STATUS_OFFLINE
		ConnMutex.Unlock()
	}
}

func (user *UserPackage)HeartBit(){
	tmpPackage := GameFormatPackage.GetPackageCMD(MAIN_CMD.MDM_KN_COMMAND, MDM_KN_COMMAND.DETECT_SOCKET, 0)
	ackBuffer := bytes.NewBuffer([]byte{})
	binary.Write(ackBuffer, binary.LittleEndian, tmpPackage)
	user.Conn.Write(ackBuffer.Bytes())
}

func (user *UserPackage)OnMessageReceived(conn *net.TCPConn){
	user.HeartBit()

	defer func() {
		if user.ConnectStatus == CONNECT_STATUS_ONLINE {
			user.CloseConnection()
		}
	}()

	for {
		ackBuffer := make([]byte, 1024)
		readLen, err := conn.Read(ackBuffer)
		newBuffer := bytes.NewBuffer(ackBuffer[0:readLen])

		if err != nil {
			break
		}
		mainCMD := binary.LittleEndian.Uint16(ackBuffer[4:6])
		//subCMD := binary.LittleEndian.Uint16(ackBuffer[6:8])

		switch mainCMD {
		case MAIN_CMD.MDM_KN_COMMAND:
			var receivePackage GameFormatPackage.ReceivedHeartBit
			binary.Read(newBuffer, binary.LittleEndian, &receivePackage)
			user.HeartBit()
		case MAIN_CMD.EVENT_TCP_NETWORK_ECHO:
			//Echo Test
			EchoValue := binary.LittleEndian.Uint16(ackBuffer[8:10])
			//	fmt.Println(user.AccountInfo.CustomerName + "(Recv): " + strconv.Itoa(int(EchoValue)))
			_, ok := user.SendBuffer[EchoValue]
			if ok{
				delete(user.SendBuffer, EchoValue)
			}else{
				fmt.Println("Recv Wrong Value: ", EchoValue)
			}
			for key := range user.SendBuffer{
				fmt.Println(user.AccountInfo.CustomerName + " Rest Value: ", key)
			}
			//user.EchoTest()
		default:
			alertMessage := "Unknown Message"
			AlertMessageChannel <- alertMessage
		}
	}
}

func (user *UserPackage)StartMonitor(ip string, port string){
	serverAddress := ip + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddress)
	if err != nil{
		outStr := "Connect error: " + serverAddress
		AlertMessageChannel <- outStr
		return
	}

	var conn = new (net.TCPConn)
	conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		outStr := "Connect "+ tcpAddr.String() + " Fail"
		AlertMessageChannel <- outStr
		return
	}else{
		ConnMutex.Lock()
		user.SetConn(conn)
		ConnMutex.Unlock()
		user.ConnectStatus = CONNECT_STATUS_ONLINE
	}
	time.Sleep(time.Duration(time.Millisecond * 20))
	//user.EchoTest()
	user.OnMessageReceived(conn)
}

func (user *UserPackage)EchoTest(){
	user.SendBuffer = make(map[uint16] bool) //for Echo Test
	user.EchoTest()
}

func (*UserPackage)Start(tmpPara *ConfigReader.ServerMonitorConfiguration, tmpChannel *chan string){
	AlertMessageChannel = *tmpChannel
	for _, monitorServer := range tmpPara.ServerList {
		monitor := UserPackage{}
		go monitor.StartMonitor(monitorServer.IP, monitorServer.Port)
	}
}

