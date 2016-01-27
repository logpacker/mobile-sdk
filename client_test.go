package logpackerandroid

import "testing"

func TestNewClient(t *testing.T) {
	_, err := NewClient("")
	if err == nil {
		t.Errorf("Error must be returned and client must be nil")
	}

	clusterURL := "1234567890"
	c, err := NewClient(clusterURL)
	if err == nil {
		t.Errorf("Error must be returned and client must be nil")
	}
	if c.ClusterURL != clusterURL {
		t.Errorf("c.ClusterURL != " + clusterURL)
	}
}
