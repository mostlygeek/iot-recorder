package main

import (
	"encoding/xml"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/mostlygeek/iot-recorder/eagle"
	"github.com/mostlygeek/iot-recorder/storage"
)

func main() {

	var (
		listen, path string
	)

	flag.StringVar(&listen, "listen", "127.0.0.1:8080", "hostname and port to listen on")
	flag.StringVar(&path, "path", ":memory:", "path to sqlite3 database")
	flag.Parse()

	log.Printf("DB Path: %s\nListen: %s\n\n", path, listen)

	db, err := storage.New(path)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		os.Exit(1)
	} else {
		defer db.Close()
	}

	// accept the uploaded xml fragments from the Eagle
	// ref: https://rainforestautomation.com/wp-content/uploads/2014/07/EAGLE-Uploader-API_06.pdf
	http.ListenAndServe(listen, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Body == nil {
			w.WriteHeader(http.StatusOK)
			return
		}

		body, _ := ioutil.ReadAll(req.Body)

		v := eagle.Message{}
		err := xml.Unmarshal([]byte(body), &v)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		mType := v.Type()
		if mType == eagle.InstantaneousDemand {
			message, _ := v.InstantaneousDemand()
			err := db.RecordDemand(message.Timestamp, message.Demand)
			if err != nil {
				log.Printf("Error, could not record demand: %s\n", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			} else {
				log.Printf("Demand: %d watts\n", message.Demand)
			}
		} else if mType == eagle.CurrentSummationDelivered {
			message, _ := v.CurrentSummation()
			del, rec := message.KWh()
			// turn into watthour
			dWh := int(del * 1000)
			rWh := int(rec * 1000)
			err := db.RecordSummation(message.Timestamp, dWh, rWh)
			if err != nil {
				log.Printf("Error, could not record summation: %s\n", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			} else {
				log.Printf("Sum, delivered: %.3fWh, generated: %.3fKWh\n", del, rec)
			}
		}

		w.WriteHeader(http.StatusOK)
	}))
}
