package domain

type ChatPrompt struct {
	Message string
}

type AIProviderPort interface {
	Generate(string) string
}

type AicServicePort interface {
	Generate(string) string
}
