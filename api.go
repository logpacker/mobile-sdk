package logpackerandroid

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
	ID       string // Auto-generated unique hash per event
	Message  string // String for 1 log message
	Source   string // Filename or Module name
	Time     string // Unix Timestamp
	TagName  string // Can be set manually
	LogLevel int    // NoticeLogLevel|DebugLogLevel|InfoLogLevel|WarnLogLevel|ErrorLogLevel|FatalLogLevel
}

// Result will be returned from Cluster (in JSON)
type Result struct {
	Code  int
	Error string
}

// Send sends error to the LogPacker Cluster
func (c *Client) Send(message *Message) (*Result, error) {
	return nil, nil
}
