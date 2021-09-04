package main

import (
	"gopkg.in/macaron.v1"

	"log"
	"os"
	"strconv"

	config "github.com/jeffprestes/test-bbolt/conf"
	conf "github.com/jeffprestes/test-bbolt/conf/app"
)

// application entrypoint
func main() {
	app := macaron.New()
	conf.SetupMiddlewares(app)
	conf.SetupRoutes(app)
	/*
		Generated using http://www.kammerl.de/ascii/AsciiSignature.php - (Font: 'starwars')
		All signatures are made with FIGlet (c) 1991, 1993, 1994 Glenn Chappell and Ian Chai
		All fonts are taken from figlet.org and jave.de.
		Please check for Font Credits the figlet font database!
		Figlet Frontend - Written by Julius Kammerl - 2005
	*/
	log.Println("______  ________________________________  _____________________  _________		")
	log.Println("___   |/  /__  ____/__  __ \\_  ____/_  / / /__  __ \\___  _/_  / / /_  ___/	")
	log.Println("__  /|_/ /__  __/  __  /_/ /  /    _  / / /__  /_/ /__  / _  / / /_____ \\ 	")
	log.Println("_  /  / / _  /___  _  _, _// /___  / /_/ / _  _, _/__/ /  / /_/ / ____/ / 		")
	log.Println("/_/  /_/  /_____/  /_/ |_| \\____/  \\____/  /_/ |_| /___/  \\____/  /____/  ")

	app.Run(port())
}

// configure http port
func port() int {
	port, err := config.Cfg.Section("").Key("http_port").Int()
	if err != nil {
		log.Fatal(err)
	}

	if forceLocal, _ := config.Cfg.Section("").Key("force_local_http_port").Bool(); forceLocal == false {
		if i, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
			port = i
		}
	}

	return port
}
