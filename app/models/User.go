package models

import (
	"github.com/revel/revel"
)

type User struct {
	Id string `bson:"_id"`
	Email string `bson:"email"`
	FirstName string `bson:"firstName"`
	LastName string `bson:"lastName"`
	Password string `bson:"password"`
	ConfirmPassword string `bson:"confirmPassword"`
}

func (this User) Validate(v *revel.Validation) {
	v.Required(this.Email)
	v.Required(this.FirstName)
	v.Required(this.LastName)
	v.Required(this.Password)
	v.Required(this.ConfirmPassword)
	v.Required(this.Password==this.ConfirmPassword).Message("Not matching passwords")
}
