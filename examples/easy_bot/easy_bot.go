package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/zemlyak-l/vkcringe/bot"
	"github.com/zemlyak-l/vkcringe/object"
)

const (
	token = "Your token"
)

var (
	fruits = []string{
		"apple", "banana", "orange",
		"grape", "lemon", "melon",
	}
)

func main() {
	rand.Seed(time.Now().Unix())

	bot, err := bot.NewBot(token)
	if err != nil {
		panic(err)
	}
	bot.OnPrivateMessage("/eat", func(message object.NewMessage) {
		randIndex := rand.Intn(len(fruits))
		randValue := fruits[randIndex]
		answer := fmt.Sprintf("You ate the %s!", randValue)
		bot.Api.MessagesSend(&object.Message{
			PeerID: message.PeerID,
			Text:   answer,
		})
	})
	bot.Routes.MessageTypingState = func(event object.MessageTypingStateObject) {
		users := &object.UsersGetResponse{}
		strUserID := strconv.Itoa(event.FromID)
		bot.Api.UsersGet(
			object.UsersGetObject{UserIDs: strUserID},
			users,
		)
		user := users.UsersList[0]
		fmt.Printf("%s %s печатает...\n", user.FirstName, user.LastName)
	}
	bot.RunSync()
}
