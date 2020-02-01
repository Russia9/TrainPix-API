package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"trainpix-api/api"
)

var log = logrus.New()

func main() {
	log.Out = os.Stdout

	log.Info("TrainPix API By Russia9")
	log.Info("More info: https://github.com/Russia9/TrainPix-API")

	switch os.Getenv("API_LEVEL") {
	case "TRACE":
		log.SetLevel(logrus.TraceLevel)
		break
	case "DEBUG":
		log.SetLevel(logrus.DebugLevel)
		break
	case "INFO":
		log.SetLevel(logrus.InfoLevel)
		break
	case "WARN":
		log.SetLevel(logrus.WarnLevel)
		break
	case "ERROR":
		log.SetLevel(logrus.ErrorLevel)
		break
	case "FATAL":
		log.SetLevel(logrus.FatalLevel)
		break
	case "PANIC":
		log.SetLevel(logrus.PanicLevel)
		break
	default:
		log.SetLevel(logrus.InfoLevel)
	}

	port := 8080
	if os.Getenv("API_PORT") != "" {
		port, _ = strconv.Atoi(os.Getenv("API_PORT"))
		log.Trace("Detected PORT: ", port)
	} else {
		log.Trace("Port not specified, using: ", port)
	}

	log.Debug("Starting server at port: ", port)

	api.Api(port, log)
}
