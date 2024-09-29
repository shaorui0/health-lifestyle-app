package ai

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

type OpenAIRequest struct {
    Model       string  `json:"model"`
    Prompt      string  `json:"prompt"`
    MaxTokens   int     `json:"max_tokens"`
    Temperature float64 `json:"temperature"`
}

type OpenAIResponse struct {
    ID      string `json:"id"`
    Object  string `json:"object"`
    Created int    `json:"created"`
    Choices []struct {
        Text         string      `json:"text"`
        Index        int         `json:"index"`
        Logprobs     interface{} `json:"logprobs"`
        FinishReason string      `json:"finish_reason"`
    } `json:"choices"`
}

func GetAIResponse(prompt string) (string, error) {
    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        return "", fmt.Errorf("OpenAI API key not set in environment variables")
    }

    requestBody := OpenAIRequest{
        Model:       "text-davinci-003",
        Prompt:      prompt,
        MaxTokens:   150,
        Temperature: 0.7,
    }

    jsonData, err := json.Marshal(requestBody)
    if err != nil {
        return "", err
    }

    client := &http.Client{Timeout: time.Second * 30}
    req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(jsonData))
    if err != nil {
        return "", err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+apiKey)

    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var aiResp OpenAIResponse
    err = json.Unmarshal(body, &aiResp)
    if err != nil {
        return "", err
    }

    if len(aiResp.Choices) > 0 {
        return aiResp.Choices[0].Text, nil
    }

    return "", nil
}
