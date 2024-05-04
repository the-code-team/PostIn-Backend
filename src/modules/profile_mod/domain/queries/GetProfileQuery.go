type GetUserQuery struct {
	userId string
}

func (self *GetUserQuery) execute() (User, error) {
	return GetUser(self.userId)
}