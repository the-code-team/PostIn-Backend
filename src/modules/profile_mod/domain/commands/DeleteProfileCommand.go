package commands

type DeleteProfileCommand struct {
	Email string
}

func (q *DeleteProfileCommand) Type() string {
	return "DeleteProfileCommand"
}
