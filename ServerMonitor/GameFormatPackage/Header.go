package GameFormatPackage

import "unsafe"

type Header struct {
	FirstHeader 	uint8
	SecondHeader	uint8
	PacketSize		uint16
	MainCmdID		uint16
	SubCmdID		uint16
}

func GetHeader() Header{
	newHeader := Header{0x05, 0, 0, 0, 0}
	return newHeader
}

func GetPackageCMD(mainCMD uint16, subCMD uint16, packetSize uint16) Header{
	newPackage := GetHeader()
	newPackage.MainCmdID = mainCMD
	newPackage.SubCmdID= subCMD
	newPackage.PacketSize = packetSize + uint16 (unsafe.Sizeof(Header{}))
	return newPackage
}