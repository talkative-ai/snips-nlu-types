package snips

type Language string

const (
	LanguageEnglish Language = "en"
)

type Dataset struct {
	Intents map[string]Intent
}

type Utterance struct {
	Text     string
	Entity   *string
	SlotName *string
}

type Utterances struct {
	Data []Utterance
}

type Intent struct {
	Utterances
	Entities map[string]Entity
	Language
}

type Entity struct {
	Data                    []EntityData
	UseSynonyms             bool
	AutomaticallyExtensible bool
}

type EntityData struct {
	Value    string
	Synonyms []string
}
