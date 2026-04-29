package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type slackWriter struct {
	client *http.Client
	url    string
	queue  chan []byte
}

type slackMessage struct {
	Text string `json:"text"`
}

func NewSlackWriter(url string) *slackWriter {
	slackWriter := &slackWriter{
		client: &http.Client{},
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
			dataJson, _ := json.Marshal(&m)
			req, _ := http.NewRequest("POST", s.url, bytes.NewReader(dataJson))
			_, _ = s.client.Do(req)
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
	dataJson, _ := json.Marshal(&m)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(dataJson))
	_, _ = http.DefaultClient.Do(req)
}
