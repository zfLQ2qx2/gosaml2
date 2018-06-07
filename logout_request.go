package saml2

import (
    "time"
    "github.com/russellhaering/gosaml2/types"
)

// LogoutRequest is the go struct representation of a logout request
type LogoutRequest struct {
	ID                  string          `xml:",attr"`
	Version             string          `xml:",attr"`
	ProtocolBinding     string          `xml:",attr"`

	IssueInstant        time.Time       `xml:",attr"`

	Destination         string          `xml:",attr"`
	Issuer              string

	NameID              *types.NameID         `xml:"NameID"`
}
