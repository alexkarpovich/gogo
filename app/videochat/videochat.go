package videochat

import (
	"container/list"
	"time"
)

type Event struct {
	Type      string
	User      string
	Timestamp int
	Text      string
}

type Subscription struct {
	Archive []Event
	New     <-chan Event
}

func (this Subscription) Cancel() {
	unsubscribe <- this.New
	drain(this.New)
}

func newEvent(typ, user, msg string) Event {
	return Event{typ, user, int(time.Now().Unix()), msg}
}

func Subscribe() Subscription {
	resp := make(chan Subscription)
	subscribe <- resp
	return <-resp
}

func Join(user string) {
	publish <- newEvent("join", user, "")
}

func Say(user, message string) {
	publish <- newEvent("message", user, message)
}

func Leave(user string) {
	publish <- newEvent("leave", user, "")
}

const archiveSize = 10

var (
	subscribe   = make(chan (chan<- Subscription), 10)
	unsubscribe = make(chan (<-chan Event), 10)
	publish     = make(chan Event, 10)
)

func chat() {
	archive := list.New()
	subscribers := list.New()

	for {
		select {
		case ch := <-subscribe:
			var events []Event
			for e := archive.Front(); e != nil; e = e.Next() {
				events = append(events, e.Value.(Event))
			}
			subscriber := make(chan Event, 10)
			subscribers.PushBack(subscriber)
			ch <- Subscription{events, subscriber}

		case event := <-publish:
			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				ch.Value.(chan Event) <- event
			}
			if archive.Len() >= archiveSize {
				archive.Remove(archive.Front())
			}
			archive.PushBack(event)

		case unsub := <-unsubscribe:
			for ch := subscribers.Front(); ch != nil; ch = ch.Next() {
				if ch.Value.(chan Event) == unsub {
					subscribers.Remove(ch)
					break
				}
			}
		}
	}
}

func init() {
	go chat()
}

func drain(ch <-chan Event) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
		default:
			return
		}
	}
}
