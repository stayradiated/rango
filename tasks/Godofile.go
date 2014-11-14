package main

import (
	. "gopkg.in/godo.v1"
)

func Tasks(p *Project) {
	p.Task("default", D{"server"})

	p.Task("server", W{"**/*.go"}, Debounce(3000), func() {
		// Start recompiles and restarts on changes when watching
		Start("server.go")
	})
}

func main() {
	Godo(Tasks)
}
