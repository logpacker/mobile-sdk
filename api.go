package logpackerandroid

import (
	"bytes"
	"encoding/json"
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
	UserID   string // Username or ID
	Message  string // String for 1 log message
	Source   string // Filename or Module name
	LogLevel int    // NoticeLogLevel|DebugLogLevel|InfoLogLevel|WarnLogLevel|ErrorLogLevel|FatalLogLevel
}

// Result will be returned from Cluster (in JSON)
type Result struct {
	Code  int
	Error string
}

// Send sends error to the LogPacker Cluster
func (c *Client) Send(msg *Message) (*Result, error) {
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

	result := &Result{
		Code: resp.StatusCode,
	}

	resultForErrorMessage := &Result{}
	err = json.NewDecoder(resp.Body).Decode(resultForErrorMessage)
	if err != nil {
		return nil, err
	}
	result.Error = resultForErrorMessage.Error

	return result, nil
}

func (c *Client) getRequest(payload []byte) (*http.Request, error) {
	buf := bytes.NewBuffer(payload)
	return http.NewRequest("POST", c.ClusterURL, buf)
}

func (c *Client) generatePayload(msg *Message) ([]byte, error) {
	type client struct {
		UserID      string `json:"user_id"`
		PageURL     string `json:"page_url"`
		ErrorLogURL string `json:"error_log_url"`
	}

	type message struct {
		Text      string `json:"message_text"`
		JSFileURL string `json:"js_file_url"`
		Line      int    `json:"line"`
		Column    int    `json:"column"`
		Error     string `json:"error"`
	}

	type payload struct {
		Client   client    `json:"client"`
		Messages []message `json:"messages"`
	}

	payloadData := payload{
		Client: client{
			UserID:      msg.UserID,
			ErrorLogURL: msg.Source,
			PageURL:     msg.Source,
		},
		Messages: []message{
			message{
				Text:      msg.Message,
				JSFileURL: msg.Source,
			},
		},
	}

	json, err := json.Marshal(payloadData)
	if err != nil {
		return []byte(""), err
	}

	return json, nil
}
