package logger

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// discordHook implements zapcore.WriteSyncer interface
type discordHook struct {
	webhookURL string
}

func NewDiscordHook(webhookURL string) *discordHook {
	return &discordHook{webhookURL: webhookURL}
}

func (hook *discordHook) Write(p []byte) (n int, err error) {
	webhookMessage := map[string]interface{}{
		"content": string(p),
		// You can customize this payload as needed
	}

	payloadBytes, err := json.Marshal(webhookMessage)
	if err != nil {
		return 0, err
	}

	resp, err := http.Post(hook.webhookURL, "application/json", bytes.NewReader(payloadBytes))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return len(p), nil
}
