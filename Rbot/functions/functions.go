package functions

import (
	"fmt"
	"sort"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Atoi(num string) int {
	i, _ := strconv.Atoi(num)
	return i
}

//initializing of keys from a map of students
func FillStudKeys(studKeys []int) []int {

	for k := range Students {
		studKeys = append(studKeys, k)
	}
	//deleting zeroValue
	//studKeys = append(studKeys[:0], studKeys[1:]...)
	return studKeys
}

//sorting slice in ascending order
func StreamlineStudentsMap(studKeys []int) []int {
	sort.Slice(studKeys, func(i, j int) bool {
		return studKeys[i] < studKeys[j]
	})
	return studKeys
}

func ShowStudents(update tgbotapi.Update, bot *tgbotapi.BotAPI, studKeys []int) {
	fmt.Println(studKeys)
	fmt.Println(Students[1])
	fmt.Println(Students[16])

	for i := 0; i < len(studKeys); i++ {
		msgConfig := tgbotapi.NewMessage(
			update.Message.Chat.ID,
			strconv.Itoa(studKeys[i])+
				"    "+Students[studKeys[i]])
		bot.Send(msgConfig)
	}

}
