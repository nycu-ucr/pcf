package httpcallback

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
	group := engine.Group(factory.PcfCallbackResUriPrefix)
	// https://localhost:29507/{factory.PcfCallbackResUriPrefix}/route
	for _, route := range routes {
		switch route.Method {
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
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
		"HTTPUdrPolicyDataChangeNotify",
		strings.ToUpper("Post"),
		"/nudr-notify/policy-data/:supi",
		HTTPUdrPolicyDataChangeNotify,
	},

	{
		"HTTPUdrInfluenceDataUpdateNotify",
		strings.ToUpper("Post"),
		"/nudr-notify/influence-data/:supi/:pduSessionId",
		HTTPUdrInfluenceDataUpdateNotify,
	},

	{
		"HTTPAmfStatusChangeNotify",
		strings.ToUpper("Post"),
		"/amfstatus",
		HTTPAmfStatusChangeNotify,
	},
}
