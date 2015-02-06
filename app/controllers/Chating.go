package controllers

import (
	"code.google.com/p/go.net/websocket"
	"github.com/revel/revel"
	"gogo/app/chat"
	"gogo/app/models"
)

type WebSocket struct {
	BaseController
}

func (this WebSocket) Room() revel.Result {
	return this.Render()
}

func (this WebSocket) RoomSocket(ws *websocket.Conn) revel.Result {
	subscription := chat.Subscribe()
	defer subscription.Cancel()

	user := this.RenderArgs["loggedInUser"].(*models.User)

	chat.Join(user.FirstName + " " + user.LastName)
	defer chat.Leave(user.FirstName + " " + user.LastName)

	for _, event := range subscription.Archive {
		if websocket.JSON.Send(ws, &event) != nil {
			return nil
		}
	}

	newMessages := make(chan string)

	go func() {
		var msg string
		for {
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				close(newMessages)
				return
			}
			newMessages <- msg
		}
	}()

	for {
		select {
		case event := <-subscription.New:
			if websocket.JSON.Send(ws, &event) != nil {
				return nil
			}
		case msg, ok := <-newMessages:
			if !ok {
				return nil
			}
			chat.Say(user.FirstName+" "+user.LastName, msg)
		}
	}
	return nil
}
