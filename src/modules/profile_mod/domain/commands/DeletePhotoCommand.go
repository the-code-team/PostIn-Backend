package commands

type DeletePhotoCommand struct {
	Email   string
	PhotoId int16
}

func (q *DeletePhotoCommand) Type() string {
	return "DeletePhotoCommand"
}
