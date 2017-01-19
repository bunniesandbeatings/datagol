package api

import (
	"github.com/tedsuo/rata"
)

const (
	Docs   = "Docs"
	Assert = "Assert"
	Accumulate = "Accumulate"
)

var Routes = rata.Routes{
	{Path: "/entities", Method: rata.POST, Name: Assert},
	{Path: "/", Method: rata.GET, Name: Docs},
	{Path: "/docs", Method: rata.GET, Name: Docs},
	{Path: "/entities/:entity_id", Method: rata.PUT, Name: Accumulate},
}
