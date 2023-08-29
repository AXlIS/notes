package speller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type SpellClient struct {
}

func NewSpellClient() *SpellClient {
	return &SpellClient{}
}

const serviceURL = "https://speller.yandex.net/services/spellservice.json/checkText"

func (s *SpellClient) CheckText(text string) (SpellResponse, error) {
	var response SpellResponse

	params := make(url.Values)
	params.Add("text", text)
	resp, err := http.PostForm(serviceURL, params)
	if err != nil {
		return response, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	if err = json.Unmarshal(body, &response); err != nil {
		return response, err
	}

	if resp.StatusCode != http.StatusOK {
		return response, errors.New(fmt.Sprint("Response status: ", resp.Status))
	}

	return response, nil
}
