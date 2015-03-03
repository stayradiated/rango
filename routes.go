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
			"GET", "/dir/{path:.*}", h.ReadDir,
		},
		Route{
			"CreateDir",
			"POST", "/dir/{path:.*}", h.CreateDir,
		},
		Route{
			"UpdateDir",
			"PUT", "/dir/{path:.*}", h.UpdateDir,
		},
		Route{
			"DestroyDir",
			"DELETE", "/dir/{path:.*}", h.DestroyDir,
		},

		// pages
		Route{
			"ReadPage",
			"GET", "/page/{path:.*}", h.ReadPage,
		},
		Route{
			"CreatePage",
			"POST", "/page/{path:.*}", h.CreatePage,
		},
		Route{
			"UpdatePage",
			"PUT", "/page/{path:.*}", h.UpdatePage,
		},
		Route{
			"DestroyPage",
			"DELETE", "/page/{path:.*}", h.DestroyPage,
		},

		// config
		Route{
			"ReadConfig",
			"GET", "/config", h.ReadConfig,
		},
		Route{
			"UpdateConfig",
			"PUT", "/config", h.UpdateConfig,
		},

		// assets
		Route{
			"CreateAsset",
			"POST", "/asset/{path:.*}", h.CreateAsset,
		},

		// misc
		Route{
			"PublishSite",
			"POST", "/site/publish", h.PublishSite,
		},
	}
}
