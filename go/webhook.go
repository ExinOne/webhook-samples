package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "os"
    "log"
)

// Configuration
type Configuration struct {
	LogFile			string
    WebhookUrl    	string
    AccessToken   	string
}

// Request post body
type PostBody struct {
	Category	string `json:"category"`
	Data		string `json:"data"`
}

// Resonse value
type Response struct {
    Success bool `json:"success"`
}

// Send message to Mixin Messenger based on Webhook.
//
// Parameters: content string, message content.
//
// Return: bool, status of sending message.
//
// Examples:
//
//  sendMixin("Hello World")
//
func sendMixin(content string) bool {

	var success = false

	// Load configuration files
	configFile, _ := os.Open("conf.json")
	defer configFile.Close()
	decoder := json.NewDecoder(configFile)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
	  fmt.Println("error:", err)
	}

	webhookUrl := configuration.WebhookUrl
    accessToken := configuration.AccessToken

    // Create logger
	fileName := configuration.LogFile
	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
	   	log.Fatalf("Failed opening file: %s", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "[Info] ", log.LstdFlags)

	// Initialize post body
	postBody := PostBody {
		Category:	"PLAIN_TEXT",
		Data:		content,
	}

    jsonStr, err := json.Marshal(postBody)

    if err != nil {
		logger.SetPrefix("[Error] ")
    	logger.Println(err)
    }

    // Initialize request
    url := fmt.Sprintf("%s%s", webhookUrl, accessToken)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
		logger.SetPrefix("[Error] ")
    	logger.Println(err)
    }
    defer resp.Body.Close()

    header := resp.Header
    body, _ := ioutil.ReadAll(resp.Body)

    response := Response{}
    err = json.Unmarshal(body, &response)

    if err != nil {
    	logger.SetPrefix("[Error] ")
    	logger.Println(err)
	}

    success = response.Success

    logger.Println(string(body))
    logger.Println(header)

    return success
}

func main() {
	var success = false
	success = sendMixin("Hello World")

	if success {
		log.Printf("Send mixin successfully.")
	} else {
		log.Fatalf("Send mixin failed.")
	}
}