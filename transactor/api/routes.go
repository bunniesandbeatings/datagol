package api

import (
	"github.com/tedsuo/rata"
)

const (
	Assert = "Assert"
)

var Routes = rata.Routes{
	{Path: "/entities", Method: rata.POST, Name: Assert},
}
