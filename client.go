package youtube_live_chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

type SimpleLiveChatClient interface {
	Connect(channelId string) error
	OnMessage(onMessage func(msg *SimpleLiveChatMessage) error)
}

// simpleLiveChatClientはSimpleLiveChatClient interfaceを実装している
var _ SimpleLiveChatClient = (*simpleLiveChatClient)(nil)

type simpleLiveChatClient struct {
	onMessage func(msg *SimpleLiveChatMessage) error
}

func NewSimpleLiveChatClient() *simpleLiveChatClient {
	return &simpleLiveChatClient{}
}

func (this *simpleLiveChatClient) OnMessage(onMessage func(msg *SimpleLiveChatMessage) error) {
	this.onMessage = onMessage
}

func parseInnertubeApiKey(livePageSource []byte) string {
	pattern := regexp.MustCompile(`"innertubeApiKey":"(.*?)"`)
	res := pattern.FindSubmatch(livePageSource)

	if len(res) == 2 {
		return string(res[1])
	} else {
		return ""
	}
}

func parseContinuation(livePageSource []byte) string {
	pattern := regexp.MustCompile(`"continuation":"(.*?)"`)
	res := pattern.FindSubmatch(livePageSource)

	if len(res) == 2 {
		return string(res[1])
	} else {
		return ""
	}
}

func fetchLivePageSource(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

// https://qiita.com/pasta04/items/ec7ed6ded14d5af1663e
func (this *simpleLiveChatClient) Connect(channelId string) error {
	if this.onMessage == nil {
		return fmt.Errorf("Fatal error occured when invoking simpleLiveChatClient.Connect(): callback function OnMessage is nil.")
	}

	liveUrl := "https://www.youtube.com/channel/" + channelId + "/live"

	source, err := fetchLivePageSource(liveUrl)
	if err != nil {
		return err
	}

	innertubeApiKey := parseInnertubeApiKey(source)
	continuation := parseContinuation(source)

	for {
		time.Sleep(1 * time.Second)
		liveChat, _ := fetchLiveChat(innertubeApiKey, continuation)
		for _, msg := range liveChat.ContinuationContents.LiveChatContinuation.Actions {
			msgruns := msg.AddChatItemAction.Item.LiveChatTextMessageRenderer.Message.Runs
			if len(msgruns) > 0 {
				msgtxt := msgruns[0].Text
				simpleLiveChatMessage := &SimpleLiveChatMessage{Msg: msgtxt}
				this.onMessage(simpleLiveChatMessage)
			}
		}
		continuation = parseNextContinuation(liveChat)
	}

	return nil
}

func parseNextContinuation(liveChat *LiveChat) string {
	res := liveChat.ContinuationContents.LiveChatContinuation.Continuations[0].TimedContinuationData.Continuation
	if res == "" {
		res = liveChat.ContinuationContents.LiveChatContinuation.Continuations[0].InvalidationContinuationData.Continuation
	}

	return res
}

func fetchLiveChat(innertubeApiKey, continuation string) (*LiveChat, error) {
	tmpl := `
{
  "context": {
    "client": {
      "clientName": "WEB",
      "clientVersion": "2.20210126.08.02",
      "timeZone": "Asia/Tokyo",
      "utcOffsetMinutes": 540,
      "mainAppWebInfo": {
        "graftUrl": "https://www.youtube.com/live_chat?continuation="
      }
    },
    "request": {
      "useSsl": true
    }
  },
  "continuation": "%s"
}
`

	url := "https://www.youtube.com/youtubei/v1/live_chat/get_live_chat?key=" + innertubeApiKey
	reqbody := fmt.Sprintf(tmpl, continuation)

	resp, err := http.Post(url, "application/json", bytes.NewBufferString(reqbody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)

	var liveChat LiveChat
	err = json.Unmarshal(byteArray, &liveChat)
	if err != nil {
		return nil, err
	}

	return &liveChat, nil
}
