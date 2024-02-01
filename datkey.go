package datkey

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

const requestTimeout = 60 * time.Second

var baseUrl = "https://api.datkey.dev"

type Config struct {
	apiKey     string
	httpClient *http.Client
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
	return Config{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

func (c Config) GenerateKey(payload GenerateKeyPayload) (*http.Response, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&payload); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/keys", baseUrl), &body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)

	generateKey, err := c.httpClient.Do(req)
	if err != nil {
		if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
			log.Printf("Request to %s timed out", req.URL)
		} else {
			log.Printf("Request to %s failed: %v", req.URL, err)
		}
		return nil, err
	}

	return generateKey, nil
}

func (c Config) GetKey(keyId string) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/keys/%s", baseUrl, keyId), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", c.apiKey)

	getKey, err := c.httpClient.Do(req)
	if err != nil {
		if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
			log.Printf("Request to %s timed out", req.URL)
		} else {
			log.Printf("Request to %s failed: %v", req.URL, err)
		}
		return nil, err
	}

	return getKey, nil
}

func (c Config) RevokeKey(keyId string) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "DELETE", fmt.Sprintf("%s/keys/%s", baseUrl, keyId), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", c.apiKey)

	deleteKey, err := c.httpClient.Do(req)
	if err != nil {
		if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
			log.Printf("Request to %s timed out", req.URL)
		} else {
			log.Printf("Request to %s failed: %v", req.URL, err)
		}
		return nil, err
	}

	return deleteKey, nil
}

func (c Config) VerifyKey(payload VerifyKeyPayload) (*http.Response, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&payload); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/keys/verify", baseUrl), &body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)

	verifyKey, err := c.httpClient.Do(req)
	if err != nil {
		if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
			log.Printf("Request to %s timed out", req.URL)
		} else {
			log.Printf("Request to %s failed: %v", req.URL, err)
		}
		return nil, err
	}

	return verifyKey, nil
}

func (c Config) UpdateKey(payload UpdateKeyPayload) (*http.Response, error) {
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&payload); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "PUT", fmt.Sprintf("%s/keys", baseUrl), &body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)

	updateKey, err := c.httpClient.Do(req)
	if err != nil {
		if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
			log.Printf("Request to %s timed out", req.URL)
		} else {
			log.Printf("Request to %s failed: %v", req.URL, err)
		}
		return nil, err
	}

	return updateKey, nil
}
