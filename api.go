package logpackermobilesdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

// FatalLogLevel var
var FatalLogLevel = 0

// ErrorLogLevel var
var ErrorLogLevel = 1

// WarnLogLevel var
var WarnLogLevel = 2

// InfoLogLevel var
var InfoLogLevel = 3

// DebugLogLevel var
var DebugLogLevel = 4

// NoticeLogLevel var
var NoticeLogLevel = 5

// Message format to be sent to LogPacker
type Message struct {
	Message  string // required, String for 1 log message, can be multiline
	Source   string // optional, Filename or Module name
	LogLevel int    // optional, NoticeLogLevel|DebugLogLevel|InfoLogLevel|WarnLogLevel|ErrorLogLevel|FatalLogLevel
	UserID   string // optional, User ID
	UserName string // optional, Username
}

// Result will be returned from Cluster (in JSON)
type Result struct {
	Code  int
	Error string
	ID    string // Unique message ID after save
}

// NewMessage initializes new message object
func (c *Client) NewMessage() *Message {
	return &Message{
		LogLevel: InfoLogLevel,
	}
}

// Send sends error to the LogPacker Cluster
func (c *Client) Send(msg *Message) (*Result, error) {
	err := c.validate(msg)
	if err != nil {
		return nil, err
	}

	payload, err := c.generatePayload(msg)
	if err != nil {
		return nil, err
	}

	req, err := c.getRequest(payload)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	type apiResultType struct {
		Code  int
		Error string
		Data  []string
	}
	apiResult := &apiResultType{}
	err = json.NewDecoder(resp.Body).Decode(apiResult)
	if err != nil {
		return nil, err
	}

	messageID := ""
	if len(apiResult.Data) > 0 {
		messageID = apiResult.Data[0]
	}
	result := &Result{
		Code:  apiResult.Code,
		Error: apiResult.Error,
		ID:    messageID,
	}

	return result, nil
}

func (c *Client) validate(msg *Message) error {
	if msg.Message == "" {
		return errors.New("Message cannot be empty")
	}
	if msg.LogLevel < FatalLogLevel || msg.LogLevel > NoticeLogLevel {
		return errors.New("LogLevel is invalid. Valid are: 0 - 5")
	}

	return nil
}

func (c *Client) getRequest(payload []byte) (*http.Request, error) {
	buf := bytes.NewBuffer(payload)
	return http.NewRequest("POST", c.ClusterURL+"/save", buf)
}

func (c *Client) generatePayload(msg *Message) ([]byte, error) {
	type client struct {
		UserID   string `json:"user_id"`
		UserName string `json:"user_name"`
		URL      string `json:"url"`
		Env      string `json:"environment"`
		Agent    string `json:"agent"`
	}

	type message struct {
		Message  string `json:"message"`
		Source   string `json:"source"`
		Line     int    `json:"line"`
		Column   int    `json:"column"`
		LogLevel int    `json:"log_level"`
		TagName  string `json:"tag_name"`
	}

	type payload struct {
		Client   client    `json:"client"`
		Messages []message `json:"messages"`
	}

	payloadData := payload{
		Client: client{
			UserID:   msg.UserID,
			UserName: msg.UserName,
			Env:      c.Environment,
			Agent:    c.Agent,
			URL:      "",
		},
		Messages: []message{
			message{
				Message:  msg.Message,
				Source:   msg.Source,
				Line:     0,
				Column:   0,
				LogLevel: msg.LogLevel,
				TagName:  "mobile",
			},
		},
	}

	json, err := json.Marshal(payloadData)
	if err != nil {
		return []byte(""), err
	}

	return json, nil
}
