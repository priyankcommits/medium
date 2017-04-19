package main

import (
	"medium/routers"

	"github.com/urfave/negroni"
)

func main() {
	router := routers.GetRouter()
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3001")
}
