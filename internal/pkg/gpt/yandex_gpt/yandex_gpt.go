package yandex

import (
	"booky-back/internal/config"
	"booky-back/internal/pkg/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type YandexGPT struct {
	config *config.GptConfig
}

func NewYandexGPT(config *config.GptConfig) *YandexGPT {
	return &YandexGPT{config: config}
}

type Message struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

type Alternative struct {
	Message Message `json:"message"`
	Status  string  `json:"status"`
}

type Usage struct {
	InputTextTokens  string `json:"inputTextTokens"`
	CompletionTokens string `json:"completionTokens"`
	TotalTokens      string `json:"totalTokens"`
}

type Result struct {
	Alternatives []Alternative `json:"alternatives"`
	Usage        Usage         `json:"usage"`
	ModelVersion string        `json:"modelVersion"`
}

type Response struct {
	Result Result `json:"result"`
}

func (g *YandexGPT) GetImprovedNote(note string) (string, error) {
	prompt := map[string]interface{}{
		"modelUri": g.config.YandexGPT.ModelUri,
		"completionOptions": map[string]interface{}{
			"stream":      false,
			"temperature": 0.6,
			"maxTokens":   2000,
		},
		"messages": []map[string]string{
			{
				"role": "system",
				"text": g.config.NoteImprovementPrompt,
			},
			{
				"role": "user",
				"text": note,
			},
		},
	}

	jsonData, err := json.Marshal(prompt)
	if err != nil {
		logger.Error("Error marshaling prompt:", err)
		return "", err
	}

	req, err := http.NewRequest("POST", g.config.YandexGPT.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Error("Error creating request:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Api-Key %s", g.config.YandexGPT.ApiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error sending request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body:", err)
		return "", err
	}

	logger.Info("Response body:", string(body))

	// Unmarshal the response body into the Response struct
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		logger.Error("Error unmarshaling response:", err)
		return "", err
	}

	// Check if there are any alternatives and return the first message text
	if len(response.Result.Alternatives) > 0 {
		textMessage := response.Result.Alternatives[0].Message.Text
		return textMessage, nil
	}

	logger.Error("No alternatives found in the response")
	return "", nil
}
