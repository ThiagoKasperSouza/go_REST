package main

import (
	"fmt"
	middlewares "go_Rest/src/middlewares"
	routes "go_Rest/src/routes"
	"net/http"

	secure "github.com/srikrsna/security-headers"
)

/*
Copyright 2024 Thiago Kasper de Souza

This file is part of go_rest.

go_rest is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

go_rest is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with go_rest.  If not, see <https://www.gnu.org/licenses/>
*/

func main() {

	router := http.NewServeMux()

	s := &secure.Secure{
		ContentTypeNoSniff: true,
	}
	csp := &secure.CSP{
		Value:    `object-src 'none'; script-src {{nonce}} 'strict-dynamic'; base-uri 'self';style-src 'self' https: 'unsafe-inline';upgrade-insecure-requests`,
		ByteSize: 8,
	}
	routes.RegisterRoutes("/donations", router)
	routes.RegisterRoutes("/isps_info", router)
	routes.RegisterRoutes("/water_level", router)
	routes.RegisterRoutes("/shelters", router)
	routes.RegisterRoutes("/blocked_roads", router)
	routes.RegisterRoutes("/news", router)

	stack := middlewares.CreateStack(
		middlewares.Headers,
		middlewares.Logging,
		csp.Middleware(),
		s.Middleware(),
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	fmt.Println("Server listening on http://localhost:8080")
	server.ListenAndServe()

}
