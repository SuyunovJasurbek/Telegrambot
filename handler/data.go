package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"telegram_bot/models"
	"time"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) Pong(c *gin.Context) {
	c.JSON(http.StatusOK, "Assalomu Alaykum ")
}

func (h *Handler) TelegramHandler() {

	var numericKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🗳 Ovoz berish"),
			tgbotapi.NewKeyboardButton("💰 Balans"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🧾Botdan foydalanish"),
			tgbotapi.NewKeyboardButton("📌 To'lovlar"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("📝 Fikr va takliflar"),
		),
	)

	var backKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("⬅️ Ortga"),
		),
	)
	var amountKeyboard = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("⬅️ Ortga"),
			tgbotapi.NewKeyboardButton("💸 To'lash"),
		),
	)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 40

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			// Bu yerda Telegram kanalga malumotlarni yuborib turishadi, 
			msgChannel := tgbotapi.NewMessage(-1001890330009, update.Message.Text)
			sucChannel := tgbotapi.NewMessage(-1001859400914, update.Message.Text)
			status := h.s.GetStatus(int(update.Message.Chat.ID))
			if status == 0 {

				switch update.Message.Text {
				case "/start":
					{
						var data = models.DataCreate{
							ChatID:   int(update.Message.Chat.ID),
							Phone:    0,
							FullName: update.Message.Chat.FirstName,
							LastName: update.Message.Chat.LastName,
							Status:   0,
							Type:     0,
							Time:     0,
						}
						h.s.CreateData(data)
						msg.Text = "Assalomu Alaykum <b> " + update.Message.Chat.FirstName + ". </b>Bo'limlardan birini tanlang"
						msg.ParseMode = "HTML"
						msg.ReplyMarkup = numericKeyboard
						if _, err := bot.Send(msg); err != nil {
							fmt.Println(err)
						}
					}
				case "🗳 Ovoz berish":
					{
						msg.Text = "<i>Telefon raqamni quyidagi formatda kiriting: 941234567 </i>"
						msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
						msg.ParseMode = "HTML"
						msg.ReplyMarkup = backKeyboard
						h.s.UpdateStatus(1, int(update.Message.Chat.ID))
						if _, err := bot.Send(msg); err != nil {
							fmt.Println(err)
						}
					}
				case "💰 Balans":
					{
						amount, _ := h.s.GetAmount(int(update.Message.Chat.ID))

						if amount == 0 {
							msg.Text = "<strong>Sizning balansingiz " + strconv.Itoa(amount) + " so'm</strong>"
							msg.ParseMode = "HTML"
							if _, err := bot.Send(msg); err != nil {
								fmt.Println(err)
							}
							msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
							msg.ReplyMarkup = backKeyboard
							continue
						}
						msg.Text = "<strong>Sizning balansingiz " + strconv.Itoa(amount) + " so'm</strong>"
						msg.ParseMode = "HTML"
						h.s.UpdateStatus(4, int(update.Message.Chat.ID))
						msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
						msg.ReplyMarkup = amountKeyboard
						if _, err := bot.Send(msg); err != nil {
							fmt.Println(err)
						}
					}
				case "🧾Botdan foydalanish":
					{
						msg.Text = `
						1.🗳 Ovoz berish  

						2.📞Ovoz bermoqchi bulgan telefon nomer kiritish 
					   
						3.✔️ Matematik misol qiymatini topish 
					   
						4.📤 Telefon raqamga kelgan sms ni kiritish 
					   
						5.💰 Balans  buyrug'ini bosib xisobingizga yig'ilgan summani kurish. 
					   
					    6.📌 Paynet qilish uchun telefon raqami kiritish`
						msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
						msg.ReplyMarkup = numericKeyboard
						if _, err := bot.Send(msg); err != nil {
							fmt.Println(err)
						}
					}
				case "📌 To'lovlar":
					{
						msg.Text = `https://t.me/openbudjet_tulov_users`
						if _, err := bot.Send(msg); err != nil {
							fmt.Println(err)
						}
						msg.Text = string("Bu kanalda xar bir foydalanuvchi ovoz berganligi uchun tulab berilgan mablag'lar tashlanib boriladi")
						if _, err := bot.Send(msg); err != nil {
							fmt.Println(err)
						}
					}
				case "📝 Fikr va takliflar":
					{
						msg.Text = "Fikr va takliflar uchun: @openbudjet_ad"
						if _, err := bot.Send(msg); err != nil {
							fmt.Println(err)
						}
					}
				default:
					{
						msg.Text = "<i>Bo'limlardan birini tanlang</i>"
						msg.ParseMode = "HTML"
						msg.ReplyMarkup = numericKeyboard
						if _, err := bot.Send(msg); err != nil {
							fmt.Println(err)
						}
					}
				}

			} else if status == 1 {
				if update.Message.Text == "/start" {
					var data = models.DataCreate{
						ChatID:   int(update.Message.Chat.ID),
						Phone:    0,
						FullName: update.Message.Chat.FirstName,
						LastName: update.Message.Chat.LastName,
						Status:   0,
						Type:     0,
						Time:     0,
					}
					h.s.CreateData(data)
					msg.Text = "Assalomu Alaykum <b> " + update.Message.Chat.FirstName + ". </b>Bo'limlardan birini tanlang"
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = numericKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
				if update.Message.Text == "⬅️ Ortga" {
					var data = models.DataCreate{
						ChatID:   int(update.Message.Chat.ID),
						Phone:    0,
						FullName: update.Message.Chat.FirstName,
						LastName: update.Message.Chat.LastName,
						Status:   0,
						Type:     0,
						Time:     0,
					}
					h.s.CreateData(data)
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = numericKeyboard
					msg.Text = "<b>Bo'limlardan birini tanlang</b>"
					msg.ParseMode = "HTML"
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
				// phone check
				nom, err := strconv.Atoi(update.Message.Text)
				fmt.Println(err)
				if err != nil {
					msg.Text = "<i>Telefon raqamni quyidagi formatda kiriting: 941234567 </i>"
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = backKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
				if len(update.Message.Text) != 9 {
					msg.Text = "<i>Telefon raqamni quyidagi formatda kiriting: <b>941234567</b>. Siz ortiqcha raqam kiritdingiz.</i>"
					msg.ParseMode = "HTML"
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
				ph := update.Message.Text
				if ph[:2] == "33" {
					msg.Text = "<b>Humans nomerdan ovoz berishni iloji yuq.</b>"
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = numericKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
				if ph[:2] == "88" || ph[:2] == "90" || ph[:2] == "91" || ph[:2] == "93" || ph[:2] == "94" || ph[:2] == "95" || ph[:2] == "97" || ph[:2] == "98" || ph[:2] == "99" {
					_, err := h.s.OpenbudgetCheck(nom, int(update.Message.Chat.ID))
					h.s.UpdatePhone(nom, int(update.Message.Chat.ID))

					if err != nil {
						msg.Text = "Openbudgetda sayti javob bermayapti"
						if _, err := bot.Send(msg); err != nil {
							fmt.Println(err)
						}
						continue
					}
					// file path qaytadu
					//pathdagi string image ga utkaziladi
					chat_id_string := strconv.Itoa(int(update.Message.Chat.ID))
					file := tgbotapi.FilePath("./uploads/" + chat_id_string + ".png")
					photo := tgbotapi.NewPhoto(update.Message.Chat.ID, file)
					photo.Caption = "<i>Suratdagi misolni yechib, javobni yuboring</i>"
					photo.ParseMode = "HTML"
					if _, err = bot.Send(photo); err != nil {
						log.Fatalln(err)
					}
					os.Remove("./uploads/" + chat_id_string + ".png")
					h.s.UpdateStatus(2, int(update.Message.Chat.ID))
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = backKeyboard
					continue
				} else {
					msg.Text = "Telefon raqamni quyidagi formatda kiriting: <b>941234567</b>. Siz raqamni noto'g'ri kiritdingiz."
					msg.ParseMode = "HTML"
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}

			} else if status == 2 {
				if update.Message.Text == "/start" {
					var data = models.DataCreate{
						ChatID:   int(update.Message.Chat.ID),
						Phone:    0,
						FullName: update.Message.Chat.FirstName,
						LastName: update.Message.Chat.LastName,
						Status:   0,
						Type:     0,
						Time:     0,
					}
					h.s.CreateData(data)
					h.s.DeleteChatId(int(update.Message.Chat.ID))
					msg.ReplyMarkup = numericKeyboard
					msg.Text = "Assalomu Alaykum <b> " + update.Message.Chat.FirstName + ". </b>Bo'limlardan birini tanlang"
					msg.ParseMode = "HTML"
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
				if update.Message.Text == "⬅️ Ortga" {
					var data = models.DataCreate{
						ChatID:   int(update.Message.Chat.ID),
						Phone:    0,
						FullName: update.Message.Chat.FirstName,
						LastName: update.Message.Chat.LastName,
						Status:   0,
						Type:     0,
						Time:     0,
					}
					h.s.CreateData(data)
					h.s.DeleteChatId(int(update.Message.Chat.ID))
					msg.ReplyMarkup = numericKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
				answer, err := strconv.Atoi(update.Message.Text)
				if err != nil {
					msg.Text = "<i>Iltimos! <b>raqam</b> kiriting<i>"
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = backKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}

					continue
				}
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				msg.ReplyMarkup = backKeyboard
				res, _ := h.s.OpenbudgetPhoneCheck(int(update.Message.Chat.ID), answer)
				if res == 0 {
					msg.Text = " <i>Kiritgan raqamingizga sms yuborildi.1 daqiqa ichida sms kodni kiritng</i>"
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = backKeyboard
					h.s.UpdateStatus(3, int(update.Message.Chat.ID))
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				} else if res == -1 {
					msg.Text = "<b>Bu nomerdan avval foydalanilgan. Boshqa raqam kiritib ko'ring</b>"
					msg.ParseMode = "HTML"
					h.s.UpdateStatus(0, int(update.Message.Chat.ID))
					h.s.UpdatePhone(0, int(update.Message.Chat.ID))
					h.s.DeleteChatId(int(update.Message.Chat.ID))
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = numericKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				} else if res == 1 {
					msg.Text = "<i>Siz misolni notug'ri yechdingiz.😁</i>"
					msg.ParseMode = "HTML"
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					// file path qaytadi
					_, err := h.s.CheckResult(int(update.Message.Chat.ID))
					if err != nil {
						fmt.Println(err)
					}
					//pathdagi string image ga utkaziladi
					chat_id_string := strconv.Itoa(int(update.Message.Chat.ID))
					file := tgbotapi.FilePath("./uploads/" + chat_id_string + ".png")
					photo := tgbotapi.NewPhoto(update.Message.Chat.ID, file)
					photo.Caption = "Iltimos, misolni yechib javobni yuboring"
					if _, err = bot.Send(photo); err != nil {
						log.Fatalln(err)
					}
					os.Remove("./uploads/" + chat_id_string + ".png")
					continue
				}
			} else if status == 3 {
				if update.Message.Text == "/start" {
					var data = models.DataCreate{
						ChatID:   int(update.Message.Chat.ID),
						Phone:    0,
						FullName: update.Message.Chat.FirstName,
						LastName: update.Message.Chat.LastName,
						Status:   0,
						Type:     0,
						Time:     0,
					}
					h.s.CreateData(data)
					h.s.DeleteChatId(int(update.Message.Chat.ID))
					msg.ReplyMarkup = numericKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
				if update.Message.Text == "⬅️ Ortga" {
					var data = models.DataCreate{
						ChatID:   int(update.Message.Chat.ID),
						Phone:    0,
						FullName: update.Message.Chat.FirstName,
						LastName: update.Message.Chat.LastName,
						Status:   0,
						Type:     0,
						Time:     0,
					}
					h.s.CreateData(data)
					h.s.DeleteChatId(int(update.Message.Chat.ID))
					msg.ReplyMarkup = numericKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
				kod, err := strconv.Atoi(update.Message.Text)
				fmt.Println(err)
				if err != nil {
					msg.Text = "</i>Kod <b>6</> ta sondan iborat bo'lishi kerak. Qaytadan kiriting</i>"
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = backKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
				if len(update.Message.Text) != 6 {
					msg.Text = "<i>Kod 6 ta sondan iborat bo'lishi kerak. Qaytadan kiriting</i>"
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = backKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
				fullname := update.Message.Chat.FirstName + " " + update.Message.Chat.LastName
				res, err, pho := h.s.FinalSMSCode(int(update.Message.Chat.ID), kod, fullname)
				if err != nil {
					fmt.Println(err)
				}
				if res == 1 {
					h.s.UpdateStatus(3, int(update.Message.Chat.ID))
					msg.Text = "<b>Kod noto'g'ri kiritildi. Iltimos qayta kiritng </b>"
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = backKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
				} else if res == 0 {
					msg.Text = "<b>Tabriklaymiz  🎉🎉🎉 siz kiritgan kod to'g'ri. 💰 Balans 5000 so'mga oshdi. </b>"
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = backKeyboard
					sucChannel.Text = "500 sum ketdi.... Saytga ovoz bergan nomer :" + strconv.Itoa(pho) + update.Message.Chat.FirstName + " " + update.Message.Chat.LastName + " " + update.Message.Chat.UserName + " " + strconv.Itoa(int(update.Message.Chat.ID))
					if _, err := bot.Send(sucChannel); err != nil {
						fmt.Println(err)
					}
					var data = models.DataCreate{
						ChatID:   int(update.Message.Chat.ID),
						Phone:    0,
						FullName: update.Message.Chat.FirstName,
						LastName: update.Message.Chat.LastName,
						Status:   0,
						Type:     0,
						Time:     0,
					}
					h.s.CreateData(data)
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
				} else if res == -1 {
					msg.Text = "<b>Kod noto'g'ri kiritildi. Urinishlar soni tugadi. iltimos boshidan boshlang.</b>"
					msg.ParseMode = "HTML"
					var data = models.DataCreate{
						ChatID:   int(update.Message.Chat.ID),
						Phone:    0,
						FullName: update.Message.Chat.FirstName,
						LastName: update.Message.Chat.LastName,
						Status:   0,
						Type:     0,
						Time:     0,
					}
					h.s.CreateData(data)
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = numericKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
				} else if res == -2 {
					msg.Text = "<b>Kodni kiritish muddati tugadi. iltimos boshidan boshlang. </b>"
					msg.ParseMode = "HTML"
					var data = models.DataCreate{
						ChatID:   int(update.Message.Chat.ID),
						Phone:    0,
						FullName: update.Message.Chat.FirstName,
						LastName: update.Message.Chat.LastName,
						Status:   0,
						Type:     0,
						Time:     0,
					}
					h.s.CreateData(data)
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = numericKeyboard
					h.s.DeleteChatId(int(update.Message.Chat.ID))
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
				} else {
					msg.Text = "<b>Iltimos boshidan boshlang.</b>"
					msg.ParseMode = "HTML"
					var data = models.DataCreate{
						ChatID:   int(update.Message.Chat.ID),
						Phone:    0,
						FullName: update.Message.Chat.FirstName,
						LastName: update.Message.Chat.LastName,
						Status:   0,
						Type:     0,
						Time:     0,
					}
					h.s.CreateData(data)
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = numericKeyboard
					h.s.DeleteChatId(int(update.Message.Chat.ID))
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
				}

			} else if status == 4 {
				if update.Message.Text == "/start" {
					var data = models.DataCreate{
						ChatID:   int(update.Message.Chat.ID),
						Phone:    0,
						FullName: update.Message.Chat.FirstName,
						LastName: update.Message.Chat.LastName,
						Status:   0,
						Type:     0,
						Time:     0,
					}
					h.s.CreateData(data)
					h.s.DeleteChatId(int(update.Message.Chat.ID))
					msg.Text = "Assalomu Alaykum <b> " + update.Message.Chat.FirstName + ". </b>Bo'limlardan birini tanlang"
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = numericKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				} else if update.Message.Text == "⬅️ Ortga" {
					var data = models.DataCreate{
						ChatID:   int(update.Message.Chat.ID),
						Phone:    0,
						FullName: update.Message.Chat.FirstName,
						LastName: update.Message.Chat.LastName,
						Status:   0,
						Type:     0,
						Time:     0,
					}
					h.s.CreateData(data)
					msg.Text = "Assalomu Alaykum <b> " + update.Message.Chat.FirstName + ". </b>Bo'limlardan birini tanlang"
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = numericKeyboard

					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				} else if update.Message.Text == "💸 To'lash" {
					msg.Text = "<b> To'lov qilmoqchi bulgan raqamni <i>941234567</i> formatda kiritng </b>"
					msg.ParseMode = "HTML"
					h.s.UpdateStatus(5, int(update.Message.Chat.ID))
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = backKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				} else {
					msg.Text = "<b>Iltimos birini tanlang </b>"
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = amountKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
			} else if status == 5 {
				if update.Message.Text == "/start" {
					var data = models.DataCreate{
						ChatID:   int(update.Message.Chat.ID),
						Phone:    0,
						FullName: update.Message.Chat.FirstName,
						LastName: update.Message.Chat.LastName,
						Status:   0,
						Type:     0,
						Time:     0,
					}
					h.s.CreateData(data)
					h.s.DeleteChatId(int(update.Message.Chat.ID))
					msg.Text = "Assalomu Alaykum <b> " + update.Message.Chat.FirstName + ". </b>Bo'limlardan birini tanlang"
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = numericKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
				if update.Message.Text == "⬅️ Ortga" {
					var data = models.DataCreate{
						ChatID:   int(update.Message.Chat.ID),
						Phone:    0,
						FullName: update.Message.Chat.FirstName,
						LastName: update.Message.Chat.LastName,
						Status:   0,
						Type:     0,
						Time:     0,
					}
					h.s.CreateData(data)
					msg.Text = "Assalomu Alaykum <b> " + update.Message.Chat.FirstName + ". </b>Bo'limlardan birini tanlang"
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = numericKeyboard

					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
				nom, err := strconv.Atoi(update.Message.Text)
				if err != nil {
					msg.Text = "Telefon raqamni quyidagi formatda kiriting: <b>941234567</b>  Siz raqamni noto'g'ri kiritdingiz."
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = backKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
				if len(update.Message.Text) != 9 {
					msg.Text = "Telefon raqamni quyidagi formatda kiriting: <b>941234567</b>  Siz raqamni noto'g'ri kiritdingiz."
					msg.ParseMode = "HTML"
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
				ph := update.Message.Text
				if ph[:2] == "33" {
					msg.Text = "<b>Humans</b> <i>nomerga xozircha tulov qilishni iloji yuq. Iltimos boshqa nomer kiriting</i>"
					msg.ParseMode = "HTML"
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = backKeyboard
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
				if ph[:2] == "88" || ph[:2] == "90" || ph[:2] == "91" || ph[:2] == "93" || ph[:2] == "94" || ph[:2] == "95" || ph[:2] == "97" || ph[:2] == "98" || ph[:2] == "99" {
					phone_comp, _ := strconv.Atoi(ph[:2])
					fullname := update.Message.Chat.FirstName + " " + update.Message.Chat.LastName
					res, price := h.s.PriceFunding(phone_comp, nom, int(update.Message.Chat.ID), fullname)
					if res[0:25] == "https://payme.uz/checkout" {
						msg.Text = "<b>" + update.Message.Chat.FirstName + "</b>" + " Ovoz berganingiz uchun rahmat. To'lov cheki : 👇👇"
						msg.ParseMode = "HTML"
						if _, err := bot.Send(msg); err != nil {
							fmt.Println(err)
						}
						t, _ := time.LoadLocation("Asia/Tashkent")
						time.Local = t
						msgChannel.Text = "<b>♻️ Paynet orqali o'tkazma</b>\n<b>📞</b> <b>Telefon raqami:</b> <i>" + "+998" + ph[:2] + "****" + ph[6:9] + "\n" + "</i> <b>💴Summasi:</b> <i>" + strconv.Itoa(price) + "so'm" + "</i>" + "\n" + " <b>👤</b><b>Foydalanuvchi:</b> <i>" + update.Message.Chat.FirstName + " " + update.Message.Chat.LastName + "</i>\n📤 Xolati : <b>✅✅✅</b>\n⏰<b> " + time.Now().Format("2006-01-02 15:04:05") + "</b>"
						msgChannel.ParseMode = "HTML"
						if _, err := bot.Send(msgChannel); err != nil {
							fmt.Println(err)
						}
						msg.Text = res
						if _, err := bot.Send(msg); err != nil {
							fmt.Println(err)
						}
						msg.Text = "<b>Yaqinlaringizni ovoz berishga taklif qilishni unutmang. Yoki ularni nomeridan o'zingiz ovoz bering 🔖. Va sizga to'lov qilishda davom etamiz 🤝</b>"
						msg.ParseMode = "HTML"
						if _, err := bot.Send(msg); err != nil {
							fmt.Println(err)
						}
						msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
						msg.ReplyMarkup = numericKeyboard
						var data = models.DataCreate{
							ChatID:   int(update.Message.Chat.ID),
							Phone:    0,
							FullName: update.Message.Chat.FirstName,
							LastName: update.Message.Chat.LastName,
							Status:   0,
							Type:     0,
							Time:     0,
						}
						h.s.CreateData(data)
						continue
					}
					msg.Text = res
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					msg.ReplyMarkup = numericKeyboard
					var data = models.DataCreate{
						ChatID:   int(update.Message.Chat.ID),
						Phone:    0,
						FullName: update.Message.Chat.FirstName,
						LastName: update.Message.Chat.LastName,
						Status:   0,
						Type:     0,
						Time:     0,
					}
					h.s.CreateData(data)
					continue

				} else {
					msg.Text = "<b>Telefon raqamni formati noto'gri. Iltimos e'tiborliroq bulib, quyidagi formatda kiriting: <i>941234567</i></b>"
					msg.ParseMode = "HTML"
					if _, err := bot.Send(msg); err != nil {
						fmt.Println(err)
					}
					continue
				}
			}
		}
	}
}