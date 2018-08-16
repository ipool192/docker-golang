package main

import (
	"fmt"
	"ipool192/docker-golang/infrastructures"
	"ipool192/docker-golang/routes"
	"net/http"

	config "github.com/spf13/viper"
	"github.com/urfave/negroni"
	"github.com/rs/cors"
)

var text = `
 _          _ _                 _        
| |__   ___| | | ___        ___| |__   ___  _ __  
| '_ \ / _ \ | |/ _ \ _____/ __| '_ \ / _ \| '_ \ 
| | | |  __/ | | (_) |_____\__ \ | | | (_) | |_) |
|_| |_|\___|_|_|\___/      |___/_| |_|\___/| .__/ 
                                           |_|

HTTP Server running on :
- name    : %s
- port    : %s`

func main() {

	// initialize configuration
	appConfig := new(infrastructures.Configuration)
	appConfig.InitConfig()
	// spalsh text
	fmt.Println(fmt.Sprintf(text, config.GetString("app.name"), config.GetString("app.port")))
	// initialize route
	routes := new(routes.Route)
	router := routes.Init()
	// cors
	c := cors.Default().Handler(router)
	// set negroni
	n := negroni.Classic()
	n.UseHandler(c)
	// initialize serve
	server := &http.Server{
		Addr:    ":" + config.GetString("app.port"),
		Handler: n,
	}
	// running the server
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		fmt.Println("Listen: ", err.Error())
	}
}
