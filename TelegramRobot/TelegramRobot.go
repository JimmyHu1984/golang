package TelegramRobot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"../ConfigReader"
	"fmt"
)

type Robot struct {
}

func (*Robot) Start(tmpPara *ConfigReader.ServerMonitorConfiguration, AlertMessageChannel *chan string){
	bot , err := tgbotapi.NewBotAPI(tmpPara.RobotToken)
	if err != nil{
		log.Panic(err)
	}

	roomID, _ := strconv.ParseInt(tmpPara.ChatRoomID, 10, 64)
	msg := tgbotapi.NewMessage(roomID, tmpPara.WarringMessage + " Start Monitor")
	msg.ParseMode = "markdown"
	bot.Send(msg)

	go func(){
		for{
			select{
			case tmpStr := <- *AlertMessageChannel:
				fmt.Println(tmpStr)
				msg := tgbotapi.NewMessage(roomID, tmpStr)
				msg.ParseMode = "markdown"
				bot.Send(msg)
			}
		}
	}()
}
