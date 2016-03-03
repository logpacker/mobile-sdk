package logpackermobilesdk

import (
	"errors"
	"net/http"
)

// Client will be initialized 1 time
// ClusterURL is a host:port to the LogPacker cluster
type Client struct {
	ClusterURL  string
	Environment string
	Agent       string
}

// NewClient returns Client object and error
func NewClient(clusterURL string, environment string, agent string) (*Client, error) {
	if agent == "" {
		agent = "mobile"
	}
	if environment == "" {
		environment = "development"
	}

	c := &Client{
		ClusterURL:  clusterURL,
		Environment: environment,
		Agent:       agent,
	}

	if clusterURL == "" {
		return c, errors.New("ClusterURL must contain host:port for your LogPacker Cluster")
	}

	// Ping cluster Public API
	_, err := http.Get(clusterURL + "/version")
	if err != nil {
		return c, errors.New("ClusterURL " + clusterURL + " isn't reachable")
	}

	return c, nil
}
