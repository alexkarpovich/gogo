package app

import (
	"github.com/revel/revel"
	"gogo/common/db"
	"gogo/app/models"
	"gopkg.in/mgo.v2/bson"
	"os"
)

func checkLoggedIn(this *revel.Controller) revel.Result {
  if id,ok := this.Session["user"]; ok {
    connection,err := db.Connect()
    if err != nil {
      os.Exit(1)
    }

    defer connection.Close()

    var loggedInUser *models.User

    err = connection.DB("blog").C("users").Find(bson.M{"_id":id}).One(&loggedInUser)
    if err != nil {
      os.Exit(1)
    }

    this.RenderArgs["loggedInUser"] = loggedInUser

    return nil
  }

  delete(this.RenderArgs, "loggedInUser")

  return nil
}

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}
	revel.InterceptFunc(checkLoggedIn, revel.BEFORE, &revel.Controller{})

	revel.TemplateFuncs["equal"] = func(a, b interface{}) bool { return a == b }
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
