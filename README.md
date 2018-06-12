1. import third part package
    install GIT
    command: go get URL (ex: go get github.com/gorilla/mux)
    if it can`t find the package, please check the GoPath and GoRoot
    

2. Config

<?xml version="1.0" encoding="utf-8"?>  
<configuration>
    Which disk do you want to monitor
  <diskPath>C:\</diskPath>
    Warrring disk volume(GB)
  <warringDiskVolume>2</warringDiskVolume>
    Time iterval(min)
  <durationTime>1m</durationTime>
    Robot Token
  <robotToken>416341999:AAFscSbjolhx1Yp6Z_I3VDV1WZHQPVLm37w</robotToken>
    RoomID(Ex: 254706576)
  <chatRoomID></chatRoomID>
    Edit Text
  <warringMessage>This is 172.18.191.51</warringMessage>
    Following are not Support in this version
  <serverList>
    <serverCount>1</serverCount>
    <ip>127.0.0.1</ip>
    <port>5455</port>
  </serverList>
</configuration>
