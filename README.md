# Datkey API Client for Go

This Go package provides a client for interacting with the Datkey API. It allows you to perform operations such as
generating keys, retrieving keys, revoking keys, verifying keys, and updating keys.

## Installation

```bash
go get -u github.com/ShegeHQ/datkey
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/ShegeHQ/datkey"
)

func main() {
	// Initialize the Datkey API client with your workspace API key
	config := datkey.Init("your-api-key")

	// Example: Generate a key
	generateKeyPayload := datkey.GenerateKeyPayload{
		ApiId:  "your-api-id",
		Name:   "MyKey",
		Length: 16,
		// Other payload fields as needed
	}

	generateKeyResponse, err := config.GenerateKey(generateKeyPayload)

	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	body, err := ioutil.ReadAll(generateKeyResponse.Body)

	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	fmt.Println("Generated Key:", string(body))

	// Similar usage for other operations
}
```

## API Operations

* **Generate Key:** Generate a new key.
* **Get Key:** Retrieve information about a specific key.
* **Revoke Key:** Revoke (delete) a key.
* **Verify Key:** Verify the validity of a key.
* **Update Key:** Update the properties of a key.

## Configuration

**ApiKey:** Your Datkey workspace API key.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or
submit a pull request.
