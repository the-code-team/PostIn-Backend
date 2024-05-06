package commands

import (
	"bytes"
)

type UpdatePhotosCommand struct {
	Email  string
	Photos []bytes.Buffer
}

func (q *UpdatePhotosCommand) Type() string {
	return "UpdatePhotosCommand"
}
