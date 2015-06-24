package mocks

import (
	"code.google.com/p/go-uuid/uuid"
)

// uuid.UUID is a slice and cannot be used as a map key.
type UID [16]byte

func (u *UID) Interface() uuid.UUID {
	return uuid.UUID(u[:])
}

func (u *UID) IsNIL() bool {
	for i := 0; i < 16; i++ {
		if i != 0 {
			return false
		}
	}
	return true
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

type UIDList []UID
