package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mostlygeek/iot-recorder/eagle"
	"github.com/mostlygeek/iot-recorder/storage"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("Usage: server <address:port> <path/to/database>")
		return
	}

	listenAddr := args[0]
	dbfilename := args[1]

	db, err := storage.New(dbfilename)
	if err != nil {
		fmt.Println("Error opening/creating database: ", err.Error())
		return
	}
	defer db.Close()

	router := gin.New()
	router.Use(gin.Recovery())

	// reduce the number of demand records
	lastDemand := time.Now()

	lastDel := float32(0)
	lastRec := float32(0)

	router.POST("/submit", func(c *gin.Context) {
		if c.Request.Body == nil {
			c.String(http.StatusBadRequest, "No POST body")
			return
		}

		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			msg := fmt.Sprintf("Failed reading POST body: %s", err.Error())
			c.String(http.StatusInternalServerError, msg)

			// it'll be invisible when set to the eagle
			fmt.Println("Error: ", err)
			return
		}

		v := eagle.Message{}
		err = xml.Unmarshal([]byte(body), &v)
		if err != nil {
			msg := fmt.Sprintf("Failed parsing XML: %s", err.Error())
			// this will probably be invisible ... since the eagle drops messages
			c.String(http.StatusInternalServerError, msg)
			fmt.Println("Error: ", err)
			return
		}

		// currently only saves data from two types
		mType := v.Type()
		if mType == eagle.InstantaneousDemand {

			now := time.Now()

			// reduce demand records to one every approx 30 seconds
			if now.Sub(lastDemand) < 30*time.Second {
				c.String(http.StatusTooManyRequests, "--")
				return
			}

			lastDemand = now

			message, _ := v.InstantaneousDemand()
			fmt.Printf("%s Current demand: %d watts\n", message.Timestamp.Format("2006-01-02 15:04:05"), message.Demand)
			db.RecordDemand(message.Timestamp, int32(message.Demand))
		} else if mType == eagle.CurrentSummationDelivered {
			message, _ := v.CurrentSummation()
			del, rec := message.KWh()

			// reduce duplicate records
			if lastDel == del && lastRec == rec {
				c.String(http.StatusTooManyRequests, "--")
				return
			}

			lastDel = del
			lastRec = rec

			fmt.Printf("%s Power used: %.3fKWh, generated: %.3fKWh\n", message.Timestamp.Format("2006-01-02 15:04:05"), del, rec)
			db.RecordSummation(message.Timestamp, del, rec)
		}

		c.String(http.StatusOK, "ok.")
	})

	err = router.Run(listenAddr)
	if err != nil {
		fmt.Println("Error running server: ", err.Error())
	}
}
