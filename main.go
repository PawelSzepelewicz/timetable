package main

import (
	"os"
	"log"
	"fmt"
	"time"
	"strings"
	"net/http"
	"github.com/alxshelepenok/timetable/bot"
	"github.com/alxshelepenok/timetable/store"
	"github.com/alxshelepenok/timetable/config"
	"github.com/alxshelepenok/timetable/timetable"
	"github.com/alxshelepenok/timetable/students"
	"github.com/alxshelepenok/timetable/model/chat"
)

var (
	currentTimetable *timetable.Timetable
	currentStudents *students.Students
	telegramBot *bot.Bot
)

func createBot(c *config.Config) {
	if len(c.TelegramToken) == 0 {
		log.Fatal("#error Token not specified!")
	}

	telegramBot = bot.New(c.TelegramToken)
}

func createStore(c *config.Config) {
	chatStore, err := store.New(c.DatabasePath, false, false)
	if err != nil {
		log.Fatal(err)
	}

	chat.ChatStore = chatStore
}

func startMessageListen(c *config.Config) {
	for message := range telegramBot.Messages {
		messageText := strings.TrimSpace(message.Text)
		messageChat := message.Chat

		switch messageText {
			case "/today@AA111Bot", "/today":
				subjects, count := currentTimetable.Today()
				week := currentTimetable.TodayWeek()

				if (count > 0) {
					replacer := strings.NewReplacer("%_WEEK_%", fmt.Sprintf("%d", week), "%_SUBJECTS_%", subjects)
					message := replacer.Replace(c.TodayTemplate)
					telegramBot.SendMessage(messageChat, message)
				} else {
					replacer := strings.NewReplacer("%_WEEK_%", fmt.Sprintf("%d", week), "%_SUBJECTS_%", subjects)
					message := replacer.Replace(c.TodayWeekdayTemplate)
					telegramBot.SendMessage(messageChat, message)
				}

			case "/nextday@AA111Bot", "/nextday":
				subjects, count := currentTimetable.Nextday()
				week := currentTimetable.NextdayWeek()

				if (count > 0) {
					replacer := strings.NewReplacer("%_WEEK_%", fmt.Sprintf("%d", week), "%_SUBJECTS_%", subjects)
					message := replacer.Replace(c.NextdayTemplate)
					telegramBot.SendMessage(messageChat, message)
				} else {
					replacer := strings.NewReplacer("%_WEEK_%", fmt.Sprintf("%d", week))
					message := replacer.Replace(c.NextdayWeekdayTemplate)
					telegramBot.SendMessage(messageChat, message)
				}

			case "/oracle@AA111Bot", "/oracle":
					list := currentStudents.ShuffleString()
					replacer := strings.NewReplacer("%_LIST_%", fmt.Sprintf("%s", list))
					message := replacer.Replace(c.OracleTemplate)
					telegramBot.SendMessage(messageChat, message)

			case "/start@AA111Bot", "/start":
				chatEntity, err := chat.New(messageChat)
				if err != nil {
					log.Print(err)
					telegramBot.SendMessage(messageChat, "К сожалению, что-то пошло не так.")

					continue
				}

				err = chatEntity.Subscribe()
				if err != nil {
					log.Print(err)
					telegramBot.SendMessage(messageChat, "К сожалению, что-то пошло не так.")

					continue
				}

				telegramBot.SendMessage(messageChat, "Привет, я бот группы АА-111, я буду уведомлять вас о расписании.")
			case "/stop@AA111Bot", "/stop":
				chatEntity, err := chat.Find(messageChat.ID)
				if err != nil {
					log.Print(err)
					telegramBot.SendMessage(messageChat, "Что бы подписаться на рассылку расписания отправьте команду /start@AA111Bot.")

					continue
				}

				err = chatEntity.Unsubscribe()
				if err != nil {
					log.Print(err)
					telegramBot.SendMessage(messageChat, "К сожалению, что-то пошло не так.")

					continue
				}

				telegramBot.SendMessage(messageChat, "Окей! Уведомления о расписании больше не будут приходить.")
			case "/help@AA111Bot", "/help":
				telegramBot.SendMessage(messageChat, "Всё очень просто, есть несколько команд: \n/start - активировать бот.\n/stop - остановить бот.\n/help - список основных команд.\n/today - расписание на сегодня.\n/nextday - расписание на завтра. \n\nЕсли чего-то не хватает, всегда можно предложить идею @alxshelepenok в Telegram.")
		}
	}
}

func MainHandler(resp http.ResponseWriter, _ *http.Request) {
    resp.Write([]byte("Hi there! I'm AA111Bot!"))
}


func main() {
	fmt.Println("#info Load configuration...")
	c, err := config.New("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("#info Load timetable...")
	t, err := timetable.New("./timetable.json")
	if err != nil {
		log.Fatal(err)
	}

	currentTimetable = t

	fmt.Println("#info Load students...")
	s, err := students.New("./students.json")
	if err != nil {
		log.Fatal(err)
	}

	currentStudents = s

	fmt.Println("#info Starting bot...")
	createBot(c)
	createStore(c)

	go startMessageListen(c)

	http.HandleFunc("/", MainHandler)
    go http.ListenAndServe(":" + os.Getenv("PORT"), nil)

	telegramBot.Start(3 * time.Millisecond)
}
