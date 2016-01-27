package logpackerandroid

import "testing"

func TestSend(t *testing.T) {
	c := Client{
		ClusterURL: "localhost:9995",
	}

	result, err := c.Send(&Message{
		UserID:   "android-1",
		Message:  "",
		Source:   "paymentmodule",
		LogLevel: FatalLogLevel,
	})
	if err == nil || result != nil {
		t.Errorf("Error must be returned and result must be nil")
	}

	result, err = c.Send(&Message{
		UserID:   "android-1",
		Message:  "Crash is here!",
		Source:   "paymentmodule",
		LogLevel: FatalLogLevel,
	})
	if err == nil || result != nil {
		t.Errorf("Error must be returned and result must be nil")
	}
}
