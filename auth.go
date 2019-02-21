// NMHS package's main function that start the application.
// The root of the application, the application starts here.
package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"nhms.com/auth/authweb"
)

const (
	// PublicDir is public directory for statics files.
	PublicDir = "/public"
	// StaticPrefix is the prefix of files to remove
	StaticPrefix = "/statics"
	// DefaultAddress is the server's default address and port.
	DefaultAddress = ":8080"
)

func main() {

	router := httprouter.New()

	//files := http.FileServer(http.Dir(PublicDir))
	//router.Handle(StaticPrefix, http.StripPrefix(StaticPrefix, files))
	router.POST("/auth/public/v1/authenticate", authenticate)
	router.POST("/auth/public/v1/sessions", getSession)
	router.GET("/auth/secure/v1/users/:id", authweb.GetUser)
	router.POST("/auth/secure/v1/users", authweb.CreateUser)
	router.PUT("/auth/secure/v1/users", authweb.UpdateUser)

	log.Fatal(http.ListenAndServe(DefaultAddress, router))
}
