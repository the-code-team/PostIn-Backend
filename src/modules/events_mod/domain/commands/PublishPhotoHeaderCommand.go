package commands

import (
	"bytes"
	"github.com/google/uuid"
)

type PublishPhotoHeaderCommand struct {
	EventId uuid.UUID
	Photos  []bytes.Buffer
}

func (q *PublishPhotoHeaderCommand) Type() string {
	return "PublishPhotoHeaderCommand"
}
