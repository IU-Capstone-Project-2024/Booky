package yandex

import (
	"booky-back/internal/logger"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type YandexGPT struct {
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
		"modelUri": "gpt://b1gf6f4qqesojkcpfk9a/yandexgpt-lite",
		"completionOptions": map[string]interface{}{
			"stream":      false,
			"temperature": 0.6,
			"maxTokens":   2000,
		},
		"messages": []map[string]string{
			{
				"role": "system",
				"text": "You are a highly professional text editor specialized in academic content. Your task is to receive a text note related to a university course, analyze its content for grammar, punctuation, and formatting errors, and then improve the text. The improved text should be presented in Markdown format with proper headers, paragraph divisions, and bullet points, to enhance readability and comprehension. Ensure that the content maintains its original intent while enhancing its clarity and structure. Whenever you receive a note for editing, immediately return the improved text formatted in Markdown without any additional dialogue.",
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

	url := "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Error("Error creating request:", err)
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Api-Key AQVN1YKRO5y39LFOhy2_izkfD0k2AbKRPTx1l0_c")

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
