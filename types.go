package snips

type Language string

const (
	LanguageEnglish  Language = "en"
	LanguageGerman   Language = "de"
	LanguageSpanish  Language = "es"
	LanguageFrench   Language = "fr"
	LanguageJapanese Language = "ja"
	LanguageKorean   Language = "ko"
)

type Dataset struct {
	Intents  map[string]Intent `json:"intents"`
	Entities map[string]Entity `json:"entities"`
	Language Language          `json:"language"`
}

type Utterance struct {
	Text     string  `json:"text"`
	Entity   *string `json:"entity,omitempty"`
	SlotName *string `json:"slot_name,omitempty"`
}

type Utterances struct {
	Data []Utterance `json:"data"`
}

type Intent struct {
	Utterances Utterances `json:"utterances"`
}

type Entity struct {
	Data                    []EntityData `json:"data"`
	UseSynonyms             bool         `json:"use_synonyms"`
	AutomaticallyExtensible bool         `json:"use_synonyms"`
}

type EntityData struct {
	Value    string   `json:"value"`
	Synonyms []string `json:"synonyms"`
}
