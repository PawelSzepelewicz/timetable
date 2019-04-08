package bot

import (
	"log"
	"sync"
	"time"
	"github.com/tucnak/telebot"
)

type Bot struct {
	tb *telebot.Bot
	options *telebot.SendOptions
	pool sync.Pool
	Messages chan telebot.Message
}

func New(token string) *Bot {
	tb, err := telebot.NewBot(token)
	if err != nil {
		log.Fatal(err)
	}

	tb.Messages = make(chan telebot.Message, 100)

	bot := &Bot{
		tb: tb,
		options: &telebot.SendOptions{ParseMode: "HTML"},
		Messages: tb.Messages,
	}

	bot.pool.New = func() interface{} {
		options := &telebot.SendOptions{
			ParseMode: "HTML",
			ReplyMarkup: telebot.ReplyMarkup{
				CustomKeyboard: [][]string{},
				HideCustomKeyboard: true,
			},
			DisableWebPagePreview: false,
		}

		return options
	}

	return bot
}

func (b *Bot) Listen(subscription chan telebot.Message, timeout time.Duration) {
	b.tb.Listen(subscription, timeout)
}

func (b *Bot) Start(timeout time.Duration) {
	b.tb.Start(timeout)
}

func (b *Bot) AnswerInlineQuery(query *telebot.Query, response *telebot.QueryResponse) error {
	return b.tb.AnswerInlineQuery(query, response)
}


func (b *Bot) SendMessage(chat telebot.Chat, text string) {
	defer b.reset()
	
	err := b.tb.SendMessage(chat, text, b.options)
	if err != nil {
		log.Print(err)
	}
}

func (b *Bot) reset() {
	options := b.pool.Get().(*telebot.SendOptions)
	defer b.pool.Put(options)
	b.options.ParseMode = options.ParseMode
	b.options.ReplyMarkup = options.ReplyMarkup
}
