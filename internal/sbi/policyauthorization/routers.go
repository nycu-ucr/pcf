/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package policyauthorization

import (
	"github.com/nycu-ucr/gonet/http"
	"strings"

	"github.com/nycu-ucr/gin"

	"github.com/nycu-ucr/pcf/internal/logger"
	"github.com/nycu-ucr/pcf/pkg/factory"
	logger_util "github.com/nycu-ucr/util/logger"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := logger_util.NewGinWithLogrus(logger.GinLog)
	AddService(router)
	return router
}

func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group(factory.PcfPolicyAuthResUriPrefix)

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		case "PATCH":
			group.PATCH(route.Pattern, route.HandlerFunc)
		}
	}
	return group
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"HTTPPostAppSessions",
		strings.ToUpper("Post"),
		"/app-sessions",
		HTTPPostAppSessions,
	},

	{
		"HTTPDeleteEventsSubsc",
		strings.ToUpper("Delete"),
		"/app-sessions/:appSessionId/events-subscription",
		HTTPDeleteEventsSubsc,
	},

	{
		"HTTPUpdateEventsSubsc",
		strings.ToUpper("Put"),
		"/app-sessions/:appSessionId/events-subscription",
		HTTPUpdateEventsSubsc,
	},

	{
		"HTTPDeleteAppSession",
		strings.ToUpper("Post"),
		"/app-sessions/:appSessionId/delete",
		HTTPDeleteAppSession,
	},

	{
		"HTTPGetAppSession",
		strings.ToUpper("Get"),
		"/app-sessions/:appSessionId",
		HTTPGetAppSession,
	},

	{
		"HTTPModAppSession",
		strings.ToUpper("Patch"),
		"/app-sessions/:appSessionId",
		HTTPModAppSession,
	},
}
