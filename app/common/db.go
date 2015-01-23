package db

import (
  "fmt"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "os"
)

func connect() (sess mgo.Session) {
  uri := "mongodb://172.17.42.1:27017"
  if uri == "" {
    fmt.Println("no connection string provided")
    os.Exit(1)
  }
 
  sess, err := mgo.Dial(uri)
  if err != nil {
    fmt.Printf("Can't connect to mongo, go error %v\n", err)
    os.Exit(1)
  }
  defer sess.Close()
 
  sess.SetSafe(&mgo.Safe{})
  return
}