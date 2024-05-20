package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type AGiXTSDK struct {
	baseURI string
	headers map[string]string
}

func NewAGiXTSDK(baseURI, apiKey string) *AGiXTSDK {
	sdk := &AGiXTSDK{
		baseURI: baseURI,
		headers: make(map[string]string),
	}

	if baseURI == "" {
		sdk.baseURI = "http://localhost:7437"
	}

	if apiKey != "" {
		apiKey = strings.TrimPrefix(apiKey, "Bearer ")
		apiKey = strings.TrimPrefix(apiKey, "bearer ")
		sdk.headers["Authorization"] = apiKey
	}
	sdk.headers["Content-Type"] = "application/json"

	if sdk.baseURI[len(sdk.baseURI)-1] == '/' {
		sdk.baseURI = sdk.baseURI[:len(sdk.baseURI)-1]
	}

	return sdk
}

type AGiXTError struct {
	Message string
}

func (e *AGiXTError) Error() string {
	return e.Message
}

func (sdk *AGiXTSDK) handleError(err error) error {
	fmt.Printf("Error: %v\n", err)
	return &AGiXTError{Message: "Unable to retrieve data."}
}

func (sdk *AGiXTSDK) GetProviders() ([]string, error) {
	resp, err := http.Get(sdk.baseURI + "/api/provider")
	if err != nil {
		return nil, sdk.handleError(err)
	}
	defer resp.Body.Close()

	var data map[string][]string
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, sdk.handleError(err)
	}

	return data["providers"], nil
}

// Implement other methods similarly...

func main() {
	sdk := NewAGiXTSDK("", "")
	providers, err := sdk.GetProviders()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(providers)
}
