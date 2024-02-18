package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Url     string
	Payload string
	Headers map[string]string
}

func (r *Request) MakeRequest(method string) (bodyString string) {
	client := &http.Client{}
	req, err := http.NewRequest(method, r.Url, bytes.NewBuffer([]byte(r.Payload)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	bodyString = string(body)
	return
}

func (r *Request) CallAPI(method string) string {
	config, error := LoadConfig()
	if error != nil {
		fmt.Println("Error loading config:", error)
		return ""
	}
	r.Headers = map[string]string{"Authorization": "Bearer " + config.TOKEN}
	r.Url = fmt.Sprintf("%s/%s", config.API_URL, r.Url)
	return r.MakeRequest(method)
}

func ParseJson(jsonString string, target interface{}) {
	err := json.Unmarshal([]byte(jsonString), target)
	if err != nil {
		fmt.Println("Error parsing json:", err)
	}
}

func ParsePayload(payload map[string]interface{}) string {
	payloadString, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error parsing payload:", err)
	}
	return string(payloadString)
}
