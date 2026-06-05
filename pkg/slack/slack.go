package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpc"
)

type slackWriter struct {
	client httpc.Service
	url    string
	queue  chan []byte
}

type slackMessage struct {
	Text string `json:"text"`
}

func NewSlackWriter(url string) *slackWriter {
	slackWriter := &slackWriter{
		client: httpc.NewService("slackAsync"),
		url:    url,
		queue:  make(chan []byte, 128),
	}
	go slackWriter.send()
	return slackWriter
}

func (s *slackWriter) Write(p []byte) (n int, err error) {
	if s.url == "" {
		return len(p), nil
	}
	pc := bytes.Clone(p)
	s.queue <- pc
	return len(p), nil
}

func (s *slackWriter) Close() error {
	close(s.queue)
	return nil
}

func (s *slackWriter) send() {
	for msg := range s.queue {
		entry := map[string]any{}
		if err := json.Unmarshal(msg, &entry); err != nil {
			continue
		}
		if entry["level"] == "error" {
			text := fmt.Sprintf("[%v] %v", entry["server"], entry["content"])
			m := slackMessage{
				Text: text,
			}
			ctx := context.Background()
			fmt.Printf("%s, %v", s.url, m)
			_, _ = s.client.Do(ctx, http.MethodPost, s.url, &m)
		}
	}
}

func SendTo(url, message string) {
	if url == "" || message == "" {
		return
	}
	m := slackMessage{
		Text: message,
	}
	_, _ = httpc.Do(context.Background(), http.MethodPost, url, &m)
}
