package logpackerandroid

import "time"

// LogLevel structure
type LogLevel struct {
	Mode int    // From 0 to 5
	Name string // String representation
}

// FatalLogLevel var
var FatalLogLevel = LogLevel{
	Mode: 0,
	Name: "Fatal",
}

// ErrorLogLevel var
var ErrorLogLevel = LogLevel{
	Mode: 1,
	Name: "Error",
}

// WarnLogLevel var
var WarnLogLevel = LogLevel{
	Mode: 2,
	Name: "Warning",
}

// InfoLogLevel var
var InfoLogLevel = LogLevel{
	Mode: 3,
	Name: "Info",
}

// DebugLogLevel var
var DebugLogLevel = LogLevel{
	Mode: 4,
	Name: "Debug",
}

// NoticeLogLevel var
var NoticeLogLevel = LogLevel{
	Mode: 5,
	Name: "Notice",
}

// Message format to be sent to LogPacker
type Message struct {
	ID       string    // Auto-generated unique hash per event
	Message  string    // String for 1 log message
	Source   string    // Filename or Module name
	Time     time.Time // Time when it's occured
	TagName  string    // Can be set manually
	LogLevel LogLevel  // NoticeLogLevel|DebugLogLevel|InfoLogLevel|WarnLogLevel|ErrorLogLevel|FatalLogLevel
}

// Result will be returned from Cluster (in JSON)
type Result struct {
	Messages []string // IDs of saved messages
}

// Send sends error to the LogPacker Cluster
func (c *Client) Send(messages []*Message) (*Result, error) {
	return nil, nil
}
