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

type UtteranceChunk struct {
	Text     string  `json:"text"`
	Entity   *string `json:"entity,omitempty"`
	SlotName *string `json:"slot_name,omitempty"`
}

type Utterance struct {
	Data []UtteranceChunk `json:"data"`
}

type Intent struct {
	Utterances []Utterance `json:"utterances"`
}

type Entity struct {
	Data                    []EntityData `json:"data"`
	UseSynonyms             bool         `json:"use_synonyms"`
	AutomaticallyExtensible bool         `json:"automatically_extensible"`
}

type EntityData struct {
	Value    string   `json:"value"`
	Synonyms []string `json:"synonyms"`
}

func (i *Intent) AddUtterance(val Utterance) {
	i.Utterances = append(i.Utterances, val)
}

func (u *Utterance) AddChunk(val UtteranceChunk) {
	u.Data = append(u.Data, val)
}

type Result struct {
	Input  string       `json:"input"`
	Intent ResultIntent `json:"intent"`
	Slots  []Slot       `json:"slots"`
}

type ResultIntent struct {
	Name        string  `json:"intentName"`
	Probability float32 `json:"probability"`
}

type Slot struct {
	Range struct {
		Start int `json:"start"`
		End   int `json:"end"`
	} `json:"range"`
	RawValue string                 `json:"rawValue"`
	Value    map[string]interface{} `json:"value"`
	Entity   string                 `json:"entity"`
	Name     string                 `json:"slotName"`
}

func (r *Result) SlotsMappedByName() map[string]*Slot {
	result := map[string]*Slot{}
	for _, slot := range r.Slots {
		result[slot.Name] = &slot
	}
	return result
}
