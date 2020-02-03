package saml2

import (
    "time"
    "github.com/russellhaering/gosaml2/types"
    "encoding/xml"
)

// LogoutRequest is the go struct representation of a logout request
type LogoutRequest struct {
	XMLName             xml.Name             `xml:"urn:oasis:names:tc:SAML:2.0:protocol LogoutRequest"`
	ID                  string               `xml:"ID,attr"`
	Version             string               `xml:"Version,attr"`
	//ProtocolBinding     string          `xml:",attr"`

	IssueInstant        time.Time            `xml:"IssueInstant,attr"`

	Destination         string               `xml:"Destination,attr"`
	Issuer              *types.Issuer              `xml:"Issuer"`

	NameID              *types.NameID         `xml:"NameID"`
	SignatureValidated  bool                 `xml:"-"` // not read, not dumped
}
