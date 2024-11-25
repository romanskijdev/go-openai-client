package openaiclient

type OpenAIConfig struct {
	APIKey      string
	AssistantID string
}

// Структура для гороскопов
type Horoscope struct {
	Text string `json:"text"`
	Sign string `json:"sign"`
}

// Структура ответа
type HoroscopeResponse struct {
	Horoscopes []Horoscope `json:"horoscopes"`
}
