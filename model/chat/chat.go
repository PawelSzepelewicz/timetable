package chat

import (
	"fmt"
	"time"
	"github.com/tucnak/telebot"
	"gopkg.in/vmihailenco/msgpack.v2"
	"github.com/alxshelepenok/timetable/store"
)

var ChatStore *store.Store

type Chat struct {
	Id int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Username string
	FirstName string
	LastName string
	Status int64
	Data map[string]interface{}
}

func New(chat telebot.Chat) (*Chat, error) {
	if c, err := Find(chat.ID); err == nil {
		return c, nil
	}

	c := &Chat{
		Id: int64(chat.ID),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username: chat.Username,
		FirstName: chat.FirstName,
		LastName: chat.LastName,
		Status: 0,
		Data: make(map[string]interface{}),
	}

	key := fmt.Sprintf("chat:%d", c.Id)
	bsValue, err := msgpack.Marshal(c)
	if err != nil {
		return c, err
	}

	err = ChatStore.Put([]byte(key), bsValue)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (c *Chat) Subscribe() error {
	c.Status = 1
	key := fmt.Sprintf("chat:%d", c.Id)
	bsValue, err := msgpack.Marshal(c)
	if err != nil {
		return err
	}

	err = ChatStore.Put([]byte(key), bsValue)
	if err != nil {
		return err
	}

	return nil
}

func (c *Chat) Unsubscribe() error {
	c.Status = 0
	key := fmt.Sprintf("chat:%d", c.Id)
	bsValue, err := msgpack.Marshal(c)
	if err != nil {
		return err
	}

	err = ChatStore.Put([]byte(key), bsValue)
	if err != nil {
		return err
	}

	return nil
}

func (c *Chat) SetData(k string, v interface{}) {
	c.Data[k] = v
}

func (c *Chat) GetData(k string) interface{} {
	return c.Data[k]
}

func (c *Chat) HasData(k string) bool {
	if c.Data[k] != nil {
		return true
	}

	return false
}

func Find(id int64) (*Chat, error) {
	var c *Chat
	key := fmt.Sprintf("chat:%d", id)
	bsValue, err := ChatStore.Get([]byte(key))
	if err != nil {
		return c, err
	}

	if err := msgpack.Unmarshal(bsValue, &c); err != nil {
			return c, err
	}

	return c, nil
}

func FindAll() ([]*Chat, error) {
	iterator := ChatStore.Iterator()

	var chats []*Chat
	for iterator.Next() {
		var c *Chat
		if err := msgpack.Unmarshal(iterator.Value(), &c); err != nil {
			return chats, err
		}

		chats = append(chats, c)
	}

	iterator.Release()

	return chats, nil
}