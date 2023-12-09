package datkey

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://api.datkey.dev"

type ConfigInterface interface {
	GenerateKey(payload GenerateKeyPayload) (*http.Response, error)
	GetKey(KeyId string) (*http.Response, error)
	RevokeKey(KeyId string) (*http.Response, error)
	VerifyKey(payload VerifyKeyPayload) (*http.Response, error)
	UpdateKey(payload UpdateKeyPayload) (*http.Response, error)
}

var _ ConfigInterface = &Config{}

type Config struct {
	apiKey     string
	httpClient http.Client
	headers    map[string]string
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

func NewConfig(apikey string, httpclient http.Client) *Config {
	return &Config{
		apiKey:     apikey,
		httpClient: httpclient,
		headers: map[string]string{
			"Content-Type": "application/json",
			"x-api-key":    apikey,
		},
	}
}

func (c *Config) makeRequest(path string, method string, requestBody []byte, headers map[string]string) (*http.Request, error) {
	var (
		url     = fmt.Sprintf("%s/%s", baseUrl, path)
		body    io.Reader
		request *http.Request
		err     error
	)
	if requestBody != nil {
		body = bytes.NewBuffer(requestBody)
	}

	for k, v := range headers {
		request.Header.Set(k, v)
	}

	request, err = http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func (c *Config) doRequest(request *http.Request) (*http.Response, error) {
	var (
		response *http.Response
		err      error
	)
	response, err = c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Config) GenerateKey(payload GenerateKeyPayload) (*http.Response, error) {
	var (
		req         *http.Request
		generateKey *http.Response
		body        []byte
		err         error
	)

	body, err = json.Marshal(&payload)
	if err != nil {
		return nil, err
	}
	req, err = c.makeRequest("/keys", http.MethodPost, body, c.headers)
	if err != nil {
		return nil, err
	}
	generateKey, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer generateKey.Body.Close()
	return generateKey, nil
}

func (c *Config) GetKey(KeyId string) (*http.Response, error) {
	var (
		req    *http.Request
		getKey *http.Response
		body   []byte
		err    error
	)

	req, err = c.makeRequest("/keys/"+KeyId, http.MethodGet, body, c.headers)
	if err != nil {
		return nil, err
	}
	getKey, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer getKey.Body.Close()
	return getKey, nil
}

func (c *Config) RevokeKey(KeyId string) (*http.Response, error) {
	var (
		req       *http.Request
		deleteKey *http.Response
		body      []byte
		err       error
	)

	req, err = c.makeRequest("/keys/"+KeyId, http.MethodDelete, body, c.headers)
	if err != nil {
		return nil, err
	}
	deleteKey, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer deleteKey.Body.Close()
	return deleteKey, nil
}

func (c *Config) VerifyKey(payload VerifyKeyPayload) (*http.Response, error) {
	var (
		req       *http.Request
		verifyKey *http.Response
		body      []byte
		err       error
	)

	req, err = c.makeRequest("/keys/verify", http.MethodPost, body, c.headers)
	if err != nil {
		return nil, err
	}
	verifyKey, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer verifyKey.Body.Close()
	return verifyKey, nil
}

func (c *Config) UpdateKey(payload UpdateKeyPayload) (*http.Response, error) {
	var (
		req         *http.Request
		generateKey *http.Response
		body        []byte
		err         error
	)

	body, err = json.Marshal(&payload)
	if err != nil {
		return nil, err
	}
	req, err = c.makeRequest("/keys", http.MethodPut, body, c.headers)
	if err != nil {
		return nil, err
	}
	generateKey, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer generateKey.Body.Close()
	return generateKey, nil
}
