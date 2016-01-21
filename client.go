package logpackerandroid

import "errors"

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
		return c, errors.New("ClusterURL must contain host:port for your LogPacker Cluster")
	}

	// TODO: ping

	return c, nil
}
