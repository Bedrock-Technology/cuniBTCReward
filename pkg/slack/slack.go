package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpc"
	"go.opentelemetry.io/otel/trace"
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
			if traceId, ok := entry["trace"].(string); ok {
				if spanId, ok := entry["span"].(string); ok {
					traceID, _ := trace.TraceIDFromHex(traceId)
					spanID, _ := trace.SpanIDFromHex(spanId)
					spanContext := trace.NewSpanContext(trace.SpanContextConfig{
						TraceID:    traceID,
						SpanID:     spanID,
						TraceFlags: trace.FlagsSampled,
					})
					ctx = trace.ContextWithSpanContext(ctx, spanContext)
				}
			}
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
