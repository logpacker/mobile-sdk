package logpackerandroid

import (
	"errors"
	"net/http"
)

// Client will be initialized 1 time
// ClusterURL is a host:port to the LogPacker cluster
type Client struct {
	ClusterURL string
}

// NewClient returns Client object and error
func NewClient(clusterURL string) (*Client, error) {
	c := &Client{
		ClusterURL: clusterURL,
	}

	if clusterURL == "" {
		return c, errors.New("ClusterURL must contain host:port for your LogPacker Cluster. Given " + clusterURL)
	}

	// Ping cluster Public API
	_, err := http.Get(clusterURL + "/version")
	if err != nil {
		return c, errors.New("ClusterURL " + clusterURL + " isn't reachable")
	}

	return c, nil
}
