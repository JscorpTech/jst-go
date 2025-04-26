package providers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/JscorpTech/jst-go/go-generate/internal/domain"
)

type pizzaGptProvider struct {
	BaseUrl string
}
type Prompt struct {
	Question      string `json:"question"`
	Model         string `json:"model"`
	SearchEnabled bool   `json:"searchEnabled"`
}

type Result struct {
	Citations []string `json:"citations"`
	Content   string   `json:"content"`
}

func NewPizzaGptProvider() domain.AIProviderPort {
	return &pizzaGptProvider{
		BaseUrl: "https://www.pizzagpt.it",
	}
}

func (p *pizzaGptProvider) Generate(prompt string) string {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(Prompt{
		Question:      prompt,
		Model:         "gpt-4o-mini",
		SearchEnabled: false,
	})
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest("POST", p.BaseUrl+"/api/chatx-completion", &buf)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-language", "en-US,en;q=0.9,uz;q=0.8,ru;q=0.7")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("origin", "https://www.pizzagpt.it")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("referer", "https://www.pizzagpt.it/en")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"135\", \"Not-A.Brand\";v=\"8\", \"Chromium\";v=\"135\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36")
	req.Header.Add("x-secret", "Marinara")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}
	var data Result
	json.NewDecoder(res.Body).Decode(&data)
	return data.Content
}
