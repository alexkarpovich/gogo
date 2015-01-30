package controllers 

import (
	"github.com/revel/revel"
)

func init() {
	revel.InterceptMethod((*BaseController).Startup, revel.BEFORE)
	revel.InterceptMethod((*BaseController).CheckLoggedIn, revel.BEFORE)
	revel.InterceptMethod((*BaseController).CheckAccess, revel.BEFORE)
	revel.InterceptMethod((*BaseController).Shutdown, revel.AFTER)	
}
