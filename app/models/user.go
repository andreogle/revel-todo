package models

import (
	"fmt"

	"github.com/revel/revel"
)

type User struct {
	Name string
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Name)
}

func (u *User) Validate(v *revel.Validation) {
	v.Required(u.Name).Message("A username is required")
	v.MinSize(u.Name, 3).Message("Username must be at least 3 characters")
	v.MaxSize(u.Name, 16).Message("Username cannot be longer than 16 characters")
}
