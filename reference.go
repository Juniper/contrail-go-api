//
// Copyright (c) 2014 Juniper Networks, Inc. All rights reserved.
//

package contrail

type LinkAttribute interface {
}

// A Reference represents a link (and optional associated metadata) between
// two objects.
type Reference struct {
	To []string `json:"to,omitempty"`
	Uuid string `json:"uuid,omitempty"`
	Href string `json:"href,omitempty"`
	Attr LinkAttribute `json:"attr,omitempty"`
}

type ReferenceList []Reference

type ReferencePair struct {
	Object IObject
	Attribute LinkAttribute
}

// POST to http://server:port/ref-update
type ReferenceUpdateMsg struct {
	// object typename
	Type string `json:"type"`
	// object uuid
	Uuid string `json:"uuid"`
	// object field (without the trailing _refs and tr/_/-/)
	RefType string `json:"ref-type"`
	// reference uuid
	RefUuid string `json:"ref-uuid"`
	// reference fqn
	RefFQName []string `json:"ref-fq-name"`
	// ADD, DELETE
	Operation string `json:"operation"`
	// Attribute
	Attr LinkAttribute `json:"attr"`
}
