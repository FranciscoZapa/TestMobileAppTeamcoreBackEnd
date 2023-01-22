package configs

import (
	"github.com/FranciscoZapa/TestMobileAppTeamcoreBackEnd/src/handlers"
)

var BaseUrl = handlers.GetEnvVariable("BASEURL")
var SecurityToken = handlers.GetEnvVariable("SECURITY_TOKEN")
var GetDataUrl = "/api/questions"
