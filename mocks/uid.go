package mocks

import (
	"github.com/pborman/uuid"
)

// UID is defined since uuid.UUID is a slice and cannot be used as a map key.
type UID [16]byte

// Interface converts the fixed length array representation to the uuid package
// type.
func (u *UID) Interface() uuid.UUID {
	return uuid.UUID(u[:])
}

// IsNIL returns true if the UID value is 0.
func (u *UID) IsNIL() bool {
	return *u == UID{}
}

func parseUID(idStr string) UID {
	var uid UID
	copy(uid[:], uuid.Parse(idStr))
	return uid
}

func makeUID(id uuid.UUID) UID {
	var uid UID
	copy(uid[:], id)
	return uid
}

// Compare UID values
func Compare(lhs, rhs UID) int {
	for i := 0; i < 16; i++ {
		if lhs[i] < rhs[i] {
			return -1
		} else if lhs[i] > rhs[i] {
			return 1
		}
	}
	return 0
}

// UIDList is a slice of UID values
type UIDList []UID
