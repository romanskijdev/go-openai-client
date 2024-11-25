package openaiclient

import "encoding/json"

// Структура для JSON-схемы
type HoroscopeSchema struct {
	Type                 string         `json:"type"`
	Properties           map[string]any `json:"properties"`
	Required             []string       `json:"required"`
	AdditionalProperties bool           `json:"additionalProperties"` // Поле для дополнительных свойств
}

// Реализуем метод MarshalJSON для HoroscopeSchema
func (hs HoroscopeSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type                 string         `json:"type"`
		Properties           map[string]any `json:"properties"`
		Required             []string       `json:"required"`
		AdditionalProperties bool           `json:"additionalProperties"`
	}{
		Type: "object",
		Properties: map[string]any{
			"horoscopes": map[string]any{
				"type": "array",
				"items": map[string]any{
					"type": "object",
					"properties": map[string]any{
						"text": map[string]any{
							"type":        "string",
							"description": "Horoscope text",
						},
						"sign": map[string]any{
							"type":        "string",
							"description": "Zodiac sign",
						},
					},
					"required":             []string{"text", "sign"},
					"additionalProperties": false, // Устанавливаем для элементов массива
				},
			},
		},
		Required:             []string{"horoscopes"},
		AdditionalProperties: false, // Добавляем на верхнем уровне
	})
}

// Структура для JSON-схемы
type CompatibilitySchema struct {
	Type                 string         `json:"type"`
	Properties           map[string]any `json:"properties"`
	Required             []string       `json:"required"`
	AdditionalProperties bool           `json:"additionalProperties"` // Поле для дополнительных свойств
}

// Реализуем метод MarshalJSON для CompatibilitySchema
func (cs CompatibilitySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type                 string         `json:"type"`
		Properties           map[string]any `json:"properties"`
		Required             []string       `json:"required"`
		AdditionalProperties bool           `json:"additionalProperties"`
	}{
		Type: "object",
		Properties: map[string]any{
			"compatibility_result": map[string]any{
				"type": "object",
				"properties": map[string]any{
					"compatibility_text": map[string]any{
						"type":        "string",
						"description": "Текст основной совместимости",
					},
					"first_sphere_compatibility": map[string]any{
						"type":        "string",
						"description": "Совместимость по первой сфере",
					},
					"second_sphere_compatibility": map[string]any{
						"type":        "string",
						"description": "Совместимость по второй сфере",
					},
					"first_additional": map[string]any{
						"type":        "string",
						"description": "Первая рекомендация",
					},
					"second_additional": map[string]any{
						"type":        "string",
						"description": "Вторая рекомендация",
					},
					"first_sphere_percent": map[string]any{
						"type":        "number",
						"description": "Совместимость % первой сферы",
					},
					"second_sphere_percent": map[string]any{
						"type":        "number",
						"description": "Совместимость % второй сферы",
					},
					"first_sphere_name": map[string]any{
						"type":        "string",
						"description": "Название первой сферы",
					},
					"second_sphere_name": map[string]any{
						"type":        "string",
						"description": "Название второй сферы",
					},
				},
				"required":             []string{"compatibility_text", "first_sphere_compatibility", "second_sphere_compatibility", "first_additional", "second_additional", "first_sphere_percent", "second_sphere_percent", "first_sphere_name", "second_sphere_name"},
				"additionalProperties": false, // Запрещаем добавление лишних свойств
			},
		},
		Required:             []string{"compatibility_result"},
		AdditionalProperties: false, // Запрещаем добавление лишних свойств на верхнем уровне
	})
}

// Структура для JSON-схемы
type PersonalHoroscopeSchema struct {
	Type                 string         `json:"type"`
	Properties           map[string]any `json:"properties"`
	Required             []string       `json:"required"`
	AdditionalProperties bool           `json:"additionalProperties"` // Поле для дополнительных свойств
}

// Реализуем метод MarshalJSON для PersonalHoroscopeSchema
func (phs PersonalHoroscopeSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type                 string         `json:"type"`
		Properties           map[string]any `json:"properties"`
		Required             []string       `json:"required"`
		AdditionalProperties bool           `json:"additionalProperties"`
	}{
		Type: "object",
		Properties: map[string]any{
			"horoscope_text": map[string]any{
				"type":        "string",
				"description": "Текст гороскопа",
			},
			"first_sphere_horoscope": map[string]any{
				"type":        "string",
				"description": "Гороскоп по первой сфере",
			},
			"second_sphere_horoscope": map[string]any{
				"type":        "string",
				"description": "Гороскоп по второй сфере",
			},
			"first_additional": map[string]any{
				"type":        "string",
				"description": "Первый совет",
			},
			"second_additional": map[string]any{
				"type":        "string",
				"description": "Второй совет",
			},
			"first_additional_name": map[string]any{
				"type":        "string",
				"description": "Название первой сферы одним словом",
			},
			"second_additional_name": map[string]any{
				"type":        "string",
				"description": "Название второй сферы одним словом",
			},
			"first_sphere_percent": map[string]any{
				"type":        "number",
				"description": "Процент успеха в первой сфере",
			},
			"second_sphere_percent": map[string]any{
				"type":        "number",
				"description": "Процент успеха во второй сфере",
			},
		},
		Required:             []string{"horoscope_text", "first_sphere_horoscope", "second_sphere_horoscope", "first_additional", "second_additional", "first_additional_name", "second_additional_name", "first_sphere_percent", "second_sphere_percent"},
		AdditionalProperties: false,
	})
}

// Структура для JSON-схемы
type ChattingSchema struct {
	Type                 string         `json:"type"`
	Properties           map[string]any `json:"properties"`
	Required             []string       `json:"required"`
	AdditionalProperties bool           `json:"additionalProperties"` // Поле для дополнительных свойств
}

// Реализуем метод MarshalJSON для HoroscopeSchema
func (hs ChattingSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type                 string         `json:"type"`
		Properties           map[string]any `json:"properties"`
		Required             []string       `json:"required"`
		AdditionalProperties bool           `json:"additionalProperties"`
	}{
		Type: "object",
		Properties: map[string]any{
			"response_text": map[string]any{
				"type":        "string",
				"description": "Ответ на сообщение",
			},
		},
		Required:             []string{"response_text"},
		AdditionalProperties: false, // Добавляем на верхнем уровне
	})
}
