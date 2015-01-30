package models

import (
	"github.com/revel/revel"
	"time"
)

type User struct {
	Id 					string 		`bson:"_id"`
	Email 				string 		`bson:"email"`
	FirstName 			string 		`bson:"firstName"`
	LastName 			string 		`bson:"lastName"`
	Role 				Role 		`bson:"role"`
	Password 			string 		`bson:"password"`
	ConfirmPassword 	string 		`bson:"confirmPassword"`
	Joined 				time.Time 	`bson:"joined"`
	Updated 			time.Time 	`bson:"updated"`
}

func (this User) Validate(v *revel.Validation) {
	v.Required(this.Email)
	v.Required(this.FirstName)
	v.Required(this.LastName)
}
