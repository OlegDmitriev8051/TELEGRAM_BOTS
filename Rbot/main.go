package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	f "github.com/TELEGRAM/TELEGRAM_BOTS/Rbot/functions"
)

var mainMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Начать консультацию"),
	),
)
var deleteMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Да", "Да"),
		tgbotapi.NewInlineKeyboardButtonData("Нет", "Нет"),
	),
)

func main() {

	var (
		bot        *tgbotapi.BotAPI
		err        error
		updChannel tgbotapi.UpdatesChannel
		update     tgbotapi.Update
		updConfig  tgbotapi.UpdateConfig
		//slice of keys from a map of students
		studKeys = make([]int, 0, len(f.Students))
	)

	bot, err = tgbotapi.NewBotAPI(tgbotapiKey)
	if err != nil {
		panic("bot init error: %s\n" + err.Error())

	}

	updConfig.Timeout = 60
	updConfig.Limit = 1
	updConfig.Offset = 0

	updChannel, err = bot.GetUpdatesChan(updConfig)
	if err != nil {
		panic("update channel error: %s\n" + err.Error())
	}
	// data, err := ioutill.Readfile("Karpach3.opus")
	// type File interface {
	// 	"github.com/TELEGRAM/TELEGRAM_BOTS/Rbot/Karpach3.opus",
	// 	io.Readfile

	//}

	for {

		update = <-updChannel
		// msgConfig := tgbotapi.NewMessage(
		// 	update.Message.Chat.ID,
		// 	f.Greeting+
		// 		update.Message.From.FirstName+"\n"+f.Img)
		// bot.Send(msgConfig)

		if update.Message != nil {
			if update.Message.IsCommand() {
				cmdText := update.Message.Command()
				if cmdText == "help" {
					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID,
						f.Help)
					bot.Send(msgConfig)
				} else if cmdText == "menu" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вот что я могу:")
					msg.ReplyMarkup = mainMenu
					bot.Send(msg)

				}
			} else {
				// if user pressed "Начать консультацию"
				if update.Message.Text == mainMenu.Keyboard[0][0].Text {
					fmt.Printf("Mesage: %s\n", update.Message.Text)

					studKeys = f.FillStudKeys(studKeys)
					studKeys = f.StreamlineStudentsMap(studKeys)

					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID,
						"Я взял на карандаш следующих студентов:")

					//Hide keyboard
					msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					bot.Send(msgConfig)

					f.ShowStudents(update, bot, studKeys)
					//Show keyboard
					msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)

					for i := 1; i <= len(f.Students); i++ {
						//New keyboard
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Кого-нибудь пощадим?")
						msg.ReplyMarkup = deleteMenu
						bot.Send(msg)
						if update.Message.Text == deleteMenu.InlineKeyboard[0][1].Text {
							msgConfig := tgbotapi.NewMessage(
								update.Message.Chat.ID,
								"Хороший выбор,"+update.Message.From.FirstName+
									"Скольких надо отчислить?")
							bot.Send(msgConfig)
							break
							// if user pressed "Да"
						} else if update.Message.Text == deleteMenu.InlineKeyboard[0][0].Text {
							msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
							msgConfig := tgbotapi.NewMessage(
								update.Message.Chat.ID,
								"Введи порядковый номер студента")
							bot.Send(msgConfig)
							fmt.Printf("message: %s\n", update.Message.Text)
						}

					}
				}

				// Who is it? Text of message
				fmt.Printf("from: %s;chatID :%v; message: %s\n",
					update.Message.From.FirstName,
					update.Message.Chat.ID,
					update.Message.Text)
			}
		}
	}

	bot.StopReceivingUpdates()
}
