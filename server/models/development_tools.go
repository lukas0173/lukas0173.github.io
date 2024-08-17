package models

type Style struct {
	Span      int    `json:"span"`
	TextColor string `json:"textColor"`
}

type Icons struct {
	Path string   `json:"path"`
	Name []string `json:"name"`
}

// DevelopmentTools the development tools data model
type DevelopmentTools struct {
	Id          int    `json:"id"`
	Field       string `json:"field"`
	Description string `json:"description"`
	Style       Style  `json:"style"`
	Icons       Icons  `json:"icons"`
	Order       int    `json:"order"`
}
