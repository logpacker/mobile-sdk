package logpackermobilesdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var (
	Server     *httptest.Server
	VersionURL string
	SaveURL    string
)

// Route - route struct
type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Response struct
type Response struct {
	Code  int
	Error string
	Data  interface{}
}

// Routes - slice of routes
type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/version",
		restVersion,
	},
	Route{
		"POST",
		"/save",
		restSave,
	},
	Route{
		"OPTIONS",
		"/save",
		restSave,
	},
}

func init() {
	router := mux.NewRouter()
	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
	}

	Server = httptest.NewServer(router)
	VersionURL = fmt.Sprintf("%s/version", Server.URL)
	SaveURL = fmt.Sprintf("%s/save", Server.URL)
}

func sendResponse(w http.ResponseWriter, r *http.Request, response Response) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func restVersion(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Code:  200,
		Error: "",
		Data:  "Test Server Version N",
	}

	sendResponse(w, r, response)
}

func restSave(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Code:  200,
		Error: "",
		Data:  []string{"ID-1001"},
	}
	sendResponse(w, r, response)
}

func TestNewClient(t *testing.T) {
	_, err := NewClient("", "", "")
	if err == nil {
		t.Errorf("Error must be returned and client must be nil")
	}

	c, err := NewClient("1234567890", "", "")
	if err == nil {
		t.Errorf("Error must be returned and client must be nil")
	}

	c, err = NewClient(Server.URL, "production", "Nexus")
	if err != nil {
		t.Errorf("This URL is correct")
	}
	if c.ClusterURL != Server.URL {
		t.Errorf("ClusterURL is not set")
	}
	if c.Environment != "production" {
		t.Errorf("Environment is not set")
	}
	if c.Agent != "Nexus" {
		t.Errorf("Agent is not set")
	}
}

func TestSend(t *testing.T) {
	c := Client{
		ClusterURL: Server.URL,
	}

	result, err := c.Send(&Message{
		Message: "",
	})
	if err == nil || result != nil {
		t.Errorf("Error must be returned and result must be nil")
	}

	result, err = c.Send(&Message{
		Message:  "Crash message!",
		LogLevel: 100,
	})
	if err == nil || result != nil {
		t.Errorf("Error must be returned and result must be nil")
	}

	result, err = c.Send(&Message{
		Message: "Crash message!",
	})
	if err != nil || result == nil {
		t.Errorf("Result must be not nil")
	}
	if result.ID != "ID-1001" {
		t.Errorf("Invalid ID returned")
	}
}

func TestNewMessage(t *testing.T) {
	c := Client{
		ClusterURL: Server.URL,
	}

	m := c.NewMessage()
	if m == nil || m.LogLevel != InfoLogLevel {
		t.Errorf("Message is not initialized correctly")
	}
}
