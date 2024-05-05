package models

import "database/sql/driver"

type Tag string

const (
	Art         Tag = "Art"
	Books       Tag = "Books"
	Cooking     Tag = "Cooking"
	Design      Tag = "Design"
	Development Tag = "Development"
	Exercise    Tag = "Exercise"
	Games       Tag = "Games"
	Movies      Tag = "Movies"
	Music       Tag = "Music"
	Photography Tag = "Photography"
	Travel      Tag = "Travel"
)

func (t *Tag) Scan(value interface{}) error {
	*t = Tag(value.([]byte))
	return nil
}

func (t Tag) Value() (driver.Value, error) {
	return string(t), nil
}
