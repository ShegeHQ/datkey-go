package datkey

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var baseUrl = "https://api.datkey.dev"

type Config struct {
	apiKey string
}

type GenerateKeyPayload struct {
	ApiId             string                  `json:"api_id"`
	Name              string                  `json:"name"`
	Prefix            *string                 `json:"prefix,omitempty"`
	Length            int64                   `json:"length"`
	Meta              *map[string]interface{} `json:"meta"`
	ExpiresAt         *int64                  `json:"expires_at,omitempty"`
	VerificationLimit *int64                  `json:"verification_limit,omitempty"`
}

type UpdateKeyPayload struct {
	ExpiresAt         *int64 `json:"expires_at,omitempty"`
	VerificationLimit *int64 `json:"verification_limit,omitempty"`
}

type VerifyKeyPayload struct {
	ApiId string `json:"api_id"`
	Key   string `json:"key"`
}

func Init(apiKey string) Config {
	return Config{apiKey: apiKey}
}

func (c Config) GenerateKey(payload GenerateKeyPayload) (*http.Response, error) {
	body, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/keys", baseUrl), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)

	client := http.Client{}
	generateKey, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return generateKey, nil
}

func (c Config) GetKey(keyId string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/keys/%s", baseUrl, keyId), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", c.apiKey)

	client := http.Client{}
	getKey, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return getKey, nil
}

func (c Config) RevokeKey(keyId string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/keys/%s", baseUrl, keyId), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", c.apiKey)

	client := http.Client{}
	deleteKey, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return deleteKey, nil
}

func (c Config) VerifyKey(payload VerifyKeyPayload) (*http.Response, error) {
	body, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/keys/verify", baseUrl), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)

	client := http.Client{}
	verifyKey, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return verifyKey, nil
}

func (c Config) UpdateKey(payload UpdateKeyPayload) (*http.Response, error) {
	body, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/keys", baseUrl), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)

	client := http.Client{}
	updateKey, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return updateKey, nil
}
