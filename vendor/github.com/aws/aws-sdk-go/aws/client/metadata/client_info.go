// Copyright IBM Corp. 2015, 2026

package metadata

// ClientInfo wraps immutable data from the client.Client structure.
type ClientInfo struct {
	ServiceName   string
	APIVersion    string
	Endpoint      string
	SigningName   string
	SigningRegion string
	JSONVersion   string
	TargetPrefix  string
}
