package models

import (
	"fmt"

	"github.com/revel/revel"
)

type Todo struct {
	Name      string
	Completed bool
	Number    int
	Creator   string
}

func (t Todo) String() string {
	return fmt.Sprintf("Todo %s is done: %t", t.Name, t.Completed)
}

func (t *Todo) Validate(v *revel.Validation) {
	v.Required(t.Name).Message("A todo must have a name!")
}
