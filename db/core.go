package db

import (
	log "github.com/Sirupsen/logrus"
	"github.com/influxdata/influxdb/client/v2"
)

var Client client.Client // Allows the client to be accessed outside the package. Clients are threadsafe so this is actually useful

func Connect() (c client.Client) {
	c, err := client.NewHTTPClient(client.HTTPConfig{Addr: "http://localhost:8086"})

	if err != nil {
		log.WithFields(log.Fields{"package": "db"}).Fatal(err)
	}

	log.WithFields(log.Fields{"package": "db"}).Info("Created DB client")

	Client = c

	return
}
