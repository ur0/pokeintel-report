package db

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/ur0/pokeintel-report/models"
	"time"
)

func StoreEncounters(encounters []models.Encounter) {
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "PokeIntel",
		Precision: "s",
	})

	for _, encounter := range encounters {
		tags := map[string]string{"pokemon_id": fmt.Sprint(encounter.PokemonId)}

		fields := map[string]interface{}{
			"latitude":  encounter.Latitude,
			"longitude": encounter.Longitude,
		}

		encounterTime, _ := time.Parse(time.RFC3339, encounter.EncounterTime)

		pt, _ := client.NewPoint("encounters", tags, fields, encounterTime)
		bp.AddPoint(pt)
	}

	Client.Write(bp)
	log.WithFields(log.Fields{"package": "db", "action": "StoreEncounters", "status": "success", "count": len(encounters)}).Info("Stored encounters")
}
