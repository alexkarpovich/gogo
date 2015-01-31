package models

import (
	"github.com/revel/revel"
)

type Resource struct {
	Controller string `bson:"controller"`
	Action     string `bson:"action"`
}

func (this Resource) Validate(v *revel.Validation) {
	v.Required(this.Controller)
	v.Required(this.Action)
}
