package client

import (
	"context"
	"fmt"
	"net/http"
)

// Translation represents the result for the translate call.
type Translation struct {
	Score       float64 `json:"score"`
	Translation string  `json:"translation"`
	Model       string  `json:"model"`
	Status      string  `json:"status"`
}

// Translate represents the result for the translate call.
type Translate struct {
	ID                   string        `json:"id"`
	Object               string        `json:"object"`
	Created              Time          `json:"created"`
	BestTranslation      string        `json:"best_translation"`
	BestTranslationModel string        `json:"best_translation_model"`
	Score                float64       `json:"best_score"`
	Translations         []Translation `json:"translations"`
}

// Translate converts text from one language to another.
func (cln *Client) Translate(ctx context.Context, text string, source Language, target Language, useThirdPartyEngine bool) (Translate, error) {
	url := fmt.Sprintf("%s/translate", cln.host)

	body := struct {
		Text                string   `json:"text"`
		Source              Language `json:"source_lang"`
		Target              Language `json:"target_lang"`
		UseThirdPartyEngine bool     `json:"use_third_party_engine"`
	}{
		Text:                text,
		Source:              source,
		Target:              target,
		UseThirdPartyEngine: useThirdPartyEngine,
	}

	var resp Translate
	if err := cln.do(ctx, http.MethodPost, url, body, &resp); err != nil {
		return Translate{}, err
	}

	return resp, nil
}
