package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func GetRoutes(h *Handlers) Routes {
	return Routes{
		// directories
		Route{
			"ReadDir",
			"GET", "/api/dir/{path:.*}", h.ReadDir,
		},
		Route{
			"CreateDir",
			"POST", "/api/dir/{path:.*}", h.CreateDir,
		},
		Route{
			"UpdateDir",
			"PUT", "/api/dir/{path:.*}", h.UpdateDir,
		},
		Route{
			"DestroyDir",
			"DELETE", "/api/dir/{path:.*}", h.DestroyDir,
		},

		// pages
		Route{
			"ReadPage",
			"GET", "/api/page/{path:.*}", h.ReadPage,
		},
		Route{
			"CreatePage",
			"POST", "/api/page/{path:.*}", h.CreatePage,
		},
		Route{
			"UpdatePage",
			"PUT", "/api/page/{path:.*}", h.UpdatePage,
		},
		Route{
			"DestroyPage",
			"DELETE", "/api/page/{path:.*}", h.DestroyPage,
		},

		// config
		Route{
			"ReadConfig",
			"GET", "/api/config", h.ReadConfig,
		},
		Route{
			"UpdateConfig",
			"PUT", "/api/config", h.UpdateConfig,
		},

		// misc
		Route{
			"PublishSite",
			"POST", "/site/publish", h.PublishSite,
		},
	}
}
