package datkey

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
)

func sample() {
	// sample smoke test ...

	var (
		updateKeyResponse *http.Response
		samplePayload     UpdateKeyPayload
		config            *Config
		resb              []byte
		err               error
	)
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	config = NewConfig("sample apikey", client)
	samplePayload = UpdateKeyPayload{}
	if updateKeyResponse, err = config.UpdateKey(samplePayload); err != nil {
		log.Printf("error updating key:'%s\n'", err)
		return
	}
	if resb, err = io.ReadAll(updateKeyResponse.Body); err != nil {
		log.Printf("error from DATKEY client:'%s\n'", err)
		return
	}
	fmt.Println("updated key!", string(resb))
}
