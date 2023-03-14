package argparser

import (
	"flag"
	"fmt"
	"jwt-server/logger"
	"os"
)

type Args struct {
	Port string
}

var parsedArgs *Args

func init() {
	parsedArgs = parse()
}
func GetArgs() *Args {
	return parsedArgs
}

func parse() *Args {
	portEnv := os.Getenv("LISTEN_PORT")
	portFlag := flag.Int("port", 8080, "The port to listen on")

	// Parse command-line flags
	flag.Parse()

	var port string
	if portEnv != "" {
		port = portEnv
	} else {
		port = fmt.Sprintf("%d", *portFlag)
	}
	logger.Logger.Println("Setup variables")
	logger.Logger.Println("Port: ", port)

	return &Args{
		Port: port,
	}
}
