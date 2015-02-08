package controllers

import (
	"code.google.com/p/go.net/websocket"
	"github.com/revel/revel"
	"gogo/app/models"
	"gogo/app/videochat"
)

type VideoChat struct {
	BaseController
}

func (this VideoChat) Index() revel.Result {
	return this.Render()
}

func (this VideoChat) IndexSocket(ws *websocket.Conn) revel.Result {
	subscription := videochat.Subscribe()
	defer subscription.Cancel()

	user := this.RenderArgs["loggedInUser"].(*models.User)

	videochat.Join(user.FirstName + " " + user.LastName)
	defer videochat.Leave(user.FirstName + " " + user.LastName)

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
			videochat.Say(user.FirstName+" "+user.LastName, msg)
		}
	}
	return nil
}
