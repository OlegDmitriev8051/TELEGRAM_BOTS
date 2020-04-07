package functions

import (
	// "fmt"
	// "sort"
	// "strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//ShowSlice отправляет сообщения с именами студентов,
//n - число студентов, которых надо напечатать
func ShowSlice(update tgbotapi.Update, bot *tgbotapi.BotAPI, studSlice []string, n int) {
	for i := 0; i < n; i++ {
		msgConfig := tgbotapi.NewMessage(
			update.Message.Chat.ID,
			studSlice[i])
		bot.Send(msgConfig)
	}
}

//IsContains проверяет содержится ли студент в списке
func IsContains(message string, studSlice []string) bool {
	for _, student := range studSlice {
		if student == message {
			return true
		}
	}
	return false
}

func RemoveStudents(message string, studSlice []string) []string {

	var rSlice []string
	for i, student := range studSlice {
		if student == message {
			rSlice = append(studSlice[:i], studSlice[i+1:]...)
		}
	}

	return rSlice
}
