package main

import (
	"strconv"

	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	f "github.com/TELEGRAM/TELEGRAM_BOTS/Rbot/functions"
)

var mainMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Начать жеребьевку на отчисление"),
	),
)
var deleteMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Да"),
		tgbotapi.NewKeyboardButton("Нет"),
	),
)
var nextButton = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Далее"),
	),
)

func main() {

	var (
		bot        *tgbotapi.BotAPI
		err        error
		updChannel tgbotapi.UpdatesChannel
		update     tgbotapi.Update
		updConfig  tgbotapi.UpdateConfig
		studSlice  []string //временное хранилище студентов
		stateKey   int      //1 - Начать жеребьевку на отчисление
		//2 - Исключаем из жеребьевки.Читаем из чата имена студентов
		//3 - Определились с участниками жеребьевки
		//4 - Читаем из чата число студентов, которых надо отчислить
	)

	//создание бота
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

	//принимаем сообщения
	for {
		update = <-updChannel

		if update.Message != nil {

			//информация об отправителе и сообщение
			fmt.Printf("from: %s;chatID :%v; message: %s\n",
				update.Message.From.FirstName,
				update.Message.Chat.ID,
				update.Message.Text)

			//является ли сообщение командой?
			if update.Message.IsCommand() {
				cmdText := update.Message.Command()
				switch {
				// если ввели команду /help
				case cmdText == "help":
					msgConfig := tgbotapi.NewMessage(update.Message.Chat.ID, f.Help)
					bot.Send(msgConfig)

				// если ввели команду /menu
				case cmdText == "menu":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вот что я могу:")
					msg.ReplyMarkup = mainMenu
					bot.Send(msg)
				}
			} else {
				switch {
				//Начать жеребьевку на отчисление
				case update.Message.Text == mainMenu.Keyboard[0][0].Text:
					stateKey = 1
					fmt.Println("stateKey = ", stateKey)

					//добавляем список студентов в временное хранилище
					studSlice = append(studSlice, f.StudentsRK9...)
					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID,
						"Я взял на карандаш следующих студентов:")

					// спрятать клавиатуру
					msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					bot.Send(msgConfig)

					// Карпачев показывает список студентов
					f.ShowSlice(update, bot, studSlice, len(studSlice))

					// Карпачев предлагает кого-нибудь пощадить
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Кого-нибудь исключем из жеребьевки?")
					msg.ReplyMarkup = deleteMenu
					bot.Send(msg)
				// Никого не исключаем из жеребьевки
				case update.Message.Text == deleteMenu.Keyboard[0][1].Text && stateKey == 1:
					stateKey = 3
					fmt.Println("stateKey = ", stateKey)

					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID,
						"Хороший выбор, "+update.Message.From.FirstName+"!\n"+
							"Скольких студентов надо отчислить? Введи число от 1 до 16")
					msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					bot.Send(msgConfig)

				// Исключаем кого-то из жеребьевки: устанавливаем состояние "2 - Исключаем из жеребьевки"
				case update.Message.Text == deleteMenu.Keyboard[0][0].Text && stateKey == 1:
					stateKey = 2
					fmt.Println("stateKey = ", stateKey)

					msgConfig := tgbotapi.NewMessage(
						update.Message.Chat.ID,
						"Перешли мне студентов, которые не будут учавствовать в жеребьевке")
					msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					bot.Send(msgConfig)

				// Исключаем кого-то из жеребьевки: отлавливаем сообщения с именами студентов
				case stateKey == 2 && f.IsContains(update.Message.Text, studSlice):
					fmt.Println(update.Message.Text + " was deleted")

					// Удаление студентов из массива
					studSlice = f.RemoveStudents(update.Message.Text, studSlice)

					// Кнопка nextButton меняет состояние программы на 3

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text+", вычеркул")
					msg.ReplyMarkup = nextButton
					bot.Send(msg)
				// Переключение на состояние 3
				case update.Message.Text == nextButton.Keyboard[0][0].Text && stateKey == 2:
					stateKey = 3
					fmt.Println("stateKey = ", stateKey)

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "кхе..кхее...\n"+
						"Скольких студентов надо отчислить? Введи число от 1 до 16")
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					bot.Send(msg)

					//Скольких надо отчислить.
				case stateKey == 3:

					//Число студентов, которых надо будет удалить
					numDS, err := strconv.Atoi(update.Message.Text)
					if err != nil {
						msgConfig := tgbotapi.NewMessage(
							update.Message.Chat.ID,
							"Глупец, что ты ввел?\nДаю тебе последний шанс")
						bot.Send(msgConfig)
						break
					}

					if numDS > 16 || numDS < 1 {
						msgConfig := tgbotapi.NewMessage(
							update.Message.Chat.ID,
							"Идиот, ты не попал в диапазон\nУ тебя последняя попытка")
						bot.Send(msgConfig)
						break
					}
					fmt.Println(numDS)
					// в этом case меняется состояние в конце программы
					// чтобы у пользователя была возможность исправить неверно введенное число
					stateKey = 4
					fmt.Println("stateKey = ", stateKey)

				}

				// ! Очистить слайс!
			}

			//когда пришло не текстовое сообщение
		} else {
			fmt.Printf("not message: %+v\n", update)
		}
	}

	bot.StopReceivingUpdates()
}
