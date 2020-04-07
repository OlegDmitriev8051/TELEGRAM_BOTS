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
	//msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)

	//New keyboard
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Кого-нибудь пощадим?")
	msg.ReplyMarkup = deleteMenu
	bot.Send(msg)
}
//	if user pressed "No"
if update.Message.Text == deleteMenu.Keyboard[0][1].Text {
	msgConfig := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"Хороший выбор, "+update.Message.From.FirstName+"!"+
			"ForwardMe")
	msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msgConfig)
	fmt.Printf("message: %s\n", update.Message.Text)
}
// if user pressed "Да"
if update.Message.Text == deleteMenu.Keyboard[0][0].Text {
	msgConfig := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"Введи порядковый номер студента")
	msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msgConfig)
	fmt.Printf("message: %s\n", update.Message.Text)
}

// data, err := ioutill.Readfile("Karpach3.opus")
	// type File interface {
	// 	"github.com/TELEGRAM/TELEGRAM_BOTS/Rbot/Karpach3.opus",
	// 	io.Readfile

	//}

	// msgConfig := tgbotapi.NewMessage(
		// 	update.Message.Chat.ID,
		// 	f.Greeting+
		// 		update.Message.From.FirstName+"\n"+f.Img)
		// bot.Send(msgConfig)