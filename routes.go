package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	// directories
	Route{
		"ReadDir",
		"GET", "/api/dir/{path:.*}", readDir,
	},
	Route{
		"CreateDir",
		"POST", "/api/dir/{path:.*}", createDir,
	},
	Route{
		"UpdateDir",
		"PUT", "/api/dir/{path:.*}", updateDir,
	},
	Route{
		"DestroyDir",
		"DELETE", "/api/dir/{path:.*}", destroyDir,
	},

	// pages
	Route{
		"ReadPage",
		"GET", "/api/page/{path:.*}", readPage,
	},
	Route{
		"CreatePage",
		"POST", "/api/page/{path:.*}", createPage,
	},
	Route{
		"UpdatePage",
		"PUT", "/api/page/{path:.*}", updatePage,
	},
	Route{
		"DestroyPage",
		"DELETE", "/api/page/{path:.*}", destroyPage,
	},

	// config
	Route{
		"ReadConfig",
		"GET", "/api/config", readConfig,
	},
	Route{
		"UpdateConfig",
		"PUT", "/api/config", updateConfig,
	},

	// misc
	Route{
		"PublishSite",
		"POST", "/site/publish", publishSite,
	},
}
