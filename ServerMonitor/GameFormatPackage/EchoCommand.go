package GameFormatPackage

import "../MAIN_CMD"

type EchoCommand struct {
	Head 		Header
	EchoNum		uint16
}

func (echo *EchoCommand) SetEchoCommand(tmpNum uint16){
	echo.Head = GetPackageCMD(MAIN_CMD.EVENT_TCP_NETWORK_ECHO, 0, 2)
	echo.EchoNum = tmpNum
}

