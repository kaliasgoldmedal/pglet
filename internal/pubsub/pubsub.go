package pubsub

import (
	"github.com/pglet/pglet/internal/cache"
)

func Subscribe(channel string) chan []byte {
	return cache.Subscribe(channel)
}

func Unsubscribe(ch chan []byte) {
	cache.Unsubscribe(ch)
}

func Send(channel string, message []byte) {
	cache.Send(channel, message)
}
