package rush

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

type ListenConfig struct {
}

var figletRush = `
 ______     __  __     ______     __  __    
/\  == \   /\ \/\ \   /\  ___\   /\ \_\ \   
\ \  __<   \ \ \_\ \  \ \___  \  \ \  __ \  
 \ \_\ \_\  \ \_____\  \/\_____\  \ \_\ \_\ 		%s
  \/_/ /_/   \/_____/   \/_____/   \/_/\/_/			Listen in port: %v`

func (app *App) Listen(configs ...ListenConfig) {
	port := fmt.Sprintf(":%v", app.Config.Port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error binding to %s: %v", port, err)
		return
	}
	log.Println(fmt.Sprintf(figletRush, version, port))
	if err := http.Serve(listener, app); err == nil {
		log.Fatal(err)
	}
}
