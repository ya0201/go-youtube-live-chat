package main

import (
	"fmt"

	"github.com/ya0201/go-youtube-live-chat"
)

func main() {
	client := youtube_live_chat.NewSimpleLiveChatClient()
	client.OnMessage(func(msg *youtube_live_chat.SimpleLiveChatMessage) error {
		fmt.Printf("%s\n\n", msg.Msg)
		return nil
	})
	client.Connect("INPUT THE CHANNEL ID YOU WANT TO VIEW COMMENTS")
}
