package controllers

import (
	"code.google.com/p/go.net/websocket"
	"github.com/revel/revel"
	"gogo/app/chat"
	"gogo/app/models"
	"io"
	"os"
)

type Chat struct {
	BaseController
}

func (this Chat) Room() revel.Result {
	return this.Render()
}

func (this Chat) RoomSocket(ws *websocket.Conn) revel.Result {
	if this.Request.Method == "POST" {
		result := make(map[string][]string)
		fileList := make([]string, 0)
		m := this.Request.MultipartForm
		for fname, _ := range m.File {
			fheaders := m.File[fname]
			for i, _ := range fheaders {
				//for each fileheader, get a handle to the actual file
				file, err := fheaders[i].Open()
				defer file.Close() //close the source file handle on function return
				if err != nil {
					panic(err)
				}
				//create destination file making sure the path is writeable.
				dst_path := "public/img/uploaded/" + fheaders[i].Filename
				dst, err := os.Create(dst_path)
				defer dst.Close()                             //close the destination file handle on function return
				defer os.Chmod(dst_path, (os.FileMode)(0644)) //limit access restrictions
				if err != nil {
					panic(err)
				}
				//copy the uploaded file to the destination file
				if _, err := io.Copy(dst, file); err != nil {
					panic(err)
				}
				fname := ResizeImage(dst_path)
				fileList = append(fileList, fname)
			}
			revel.TRACE.Printf("Upload successful..")
		}
		result["images"] = fileList
		return this.RenderJson(result)
	} else {

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
	}
	return nil
}
