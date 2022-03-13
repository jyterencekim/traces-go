package visits

import (
	"time"

	"github.com/google/uuid"
)

type Place struct {
	Id uuid.NullUUID
}

type Visitor interface {
	Enter(place Place, time time.Time) Presence
}

type Presence struct {
	In Place
	At time.Time
	By Visitor
}

type Move struct {
	From Presence
	To   Presence
}
