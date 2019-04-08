package chat

import (
	"log"
	"testing"
	"github.com/tucnak/telebot"
	"github.com/alxshelepenok/bittrexmarketschecker/store"

)

func init() {
	cs, err := store.New("../../tmp/chat_db_test", false, false)
	if err != nil {
		log.Fatal(err)
	}

	ChatStore = cs
}

func TestChatNew(t *testing.T) {
	log.Print("Testing chat.New()...")
	
	for i := 0; i < 100; i++ {
		chat := telebot.Chat{ID: int64(i+10), FirstName: "John", LastName: "Doe", Username: "john_doe"}
		_, err := New(chat)
		if err != nil {
			log.Print(err)
		}
	}
}

func TestChatFind(t *testing.T) {
	log.Print("Testing chat.Find()...")

	for i := 0; i < 100; i++ {
		_, err := Find(int64(i+10))
		if err != nil {
			log.Print(err)
		}
	}
}

func BenchmarkChatFind(b *testing.B) {
	log.Print("Benchmarking chat.Find()...")

	for i := 0; i < b.N; i++ {
		_, err := Find(int64(i+10))
		if err != nil {
			log.Print(err)
		}
	}
}

func BenchmarkChatNew(b *testing.B) {
	log.Print("Benchmarking chat.New()...")

	for i := 0; i < b.N; i++ {
		chat := telebot.Chat{ID: int64(i+10), FirstName: "John", LastName: "Doe", Username: "john_doe"}
		_, err := New(chat)
		if err != nil {
			log.Print(err)
		}
	}
}