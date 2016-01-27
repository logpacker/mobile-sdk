package logpackerandroid

import "testing"

func TestSend(t *testing.T) {
	c := Client{
		ClusterURL: "localhost:9995",
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
	if err == nil || result != nil {
		t.Errorf("Error must be returned and result must be nil")
	}
}
