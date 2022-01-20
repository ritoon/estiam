package moke

import "github.com/ritoon/estiam/model"

// go to DB/moke folder
var ListUser map[string]*model.User = map[string]*model.User{
	"abcd": {FirstName: "Bob", LastName: "Pike"},
}
