package logpackerandroid

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
	_, err := NewClient("")
	if err == nil {
		t.Errorf("Error must be returned and client must be nil")
	}

	c, err := NewClient("1234567890")
	if err == nil {
		t.Errorf("Error must be returned and client must be nil")
	}

	c, err = NewClient(Server.URL)
	if err != nil {
		t.Errorf("This URL is correct")
	}
	if c.ClusterURL != Server.URL {
		t.Errorf("ClusterURL is not set")
	}
}

func TestSend(t *testing.T) {
	c := Client{
		ClusterURL: Server.URL,
	}

	result, err := c.Send(&Message{
		UserID:      "android-1",
		Agent:       "android",
		TagName:     "android",
		UserName:    "John",
		Message:     "",
		Source:      "paymentmodule",
		LogLevel:    FatalLogLevel,
		Environment: "production",
	})
	if err == nil || result != nil {
		t.Errorf("Error must be returned and result must be nil")
	}

	result, err = c.Send(&Message{
		UserID:      "android-1",
		UserName:    "John",
		Message:     "Crash message!",
		Source:      "paymentmodule",
		LogLevel:    100,
		Environment: "production",
	})
	if err == nil || result != nil {
		t.Errorf("Error must be returned and result must be nil")
	}

	result, err = c.Send(&Message{
		UserID:      "android-1",
		Agent:       "Android 4.4",
		TagName:     "MyApp",
		UserName:    "John",
		Message:     "Crash message!",
		Source:      "paymentmodule",
		LogLevel:    ErrorLogLevel,
		Environment: "production",
	})
	if err != nil || result == nil {
		t.Errorf("Result must be not nil")
	}
	if result.ID != "ID-1001" {
		t.Errorf("Invalid ID returned")
	}
}
