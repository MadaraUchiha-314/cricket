package cricketers

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Cricketer struct {
	Identifier        string `json:"identifier"`
	Name              string `json:"name"`
	UniqueName        string `json:"unique_name"`
	KeyBCCI           string `json:"key_bcci"`
	KeyBigBash        string `json:"key_bigbash"`
	KeyCricbuzz       string `json:"key_cricbuzz"`
	KeyCricHeroes     string `json:"key_cricheroes"`
	KeyCricHQ         string `json:"key_crichq"`
	KeyCricingif      string `json:"key_cricingif"`
	KeyCricketArchive string `json:"key_cricketarchive"`
	KeyCricketWorld   string `json:"key_cricketworld"`
	KeyCricinfo       string `json:"key_cricinfo"`
	KeyNVPlay         string `json:"key_nvplay"`
	KeyOpta           string `json:"key_opta"`
	KeyPulse          string `json:"key_pulse"`
}

var cricketers []Cricketer = nil

func GetCricketersInfoFromFile() {
	/*
	 * This function will read from a local data file present in "./data/people/people.csv".
	 * The file is a CSV file
	 *
	 * The fields on each row (after the header) include, but are not limited to, the following:
	 *
	 * identifier
	 *   The Cricsheet identifier for the person, as used in people.csv
	 * name
	 *   The name used for the person in Cricsheet data. This can (and will) be used for multiple people of the same name
	 * unique_name
	 *   The unique name used for the person in Cricsheet data. Guaranteed not to be used for another person
	 * key_bcci
	 *   The person's identifier on BCCI
	 * key_bigbash
	 *   The person's identifier on Big Bash
	 * key_cricbuzz
	 *   The person's identifier on Cricbuzz
	 * key_cricheroes
	 *   The person's identifier on CricHeroes
	 * key_crichq
	 *   The person's identifier on CricHQ
	 * key_cricingif
	 *   The person's identifier on Cricingif
	 * key_cricketarchive
	 *   The person's identifier on CricketArchive
	 * key_cricketworld
	 *   The person's identifier on Cricketworld
	 * key_cricinfo
	 *   The person's identifier on ESPNcricinfo
	 * key_nvplay
	 *   The person's identifier on NV Play
	 * key_opta
	 *   The person's identifier on Opta
	 * key_pulse
	 *   The person's identifier on Pulse
	 */
	file, err := os.Open("../data-store/people/people.csv")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("failed to read CSV file: %s", err)
	}

	for _, record := range records[1:] { // Skip header row
		cricketer := Cricketer{
			Identifier:        record[0],
			Name:              record[1],
			UniqueName:        record[2],
			KeyBCCI:           record[3],
			KeyBigBash:        record[4],
			KeyCricbuzz:       record[5],
			KeyCricHeroes:     record[6],
			KeyCricHQ:         record[7],
			KeyCricingif:      record[8],
			KeyCricketArchive: record[9],
			KeyCricketWorld:   record[10],
			KeyCricinfo:       record[11],
			KeyNVPlay:         record[12],
			KeyOpta:           record[13],
			KeyPulse:          record[14],
		}
		cricketers = append(cricketers, cricketer)
	}
}

func GetCricketers(w http.ResponseWriter, r *http.Request) {
	if cricketers == nil {
		GetCricketersInfoFromFile()
	}
	// Convert the array of structs to a json response string
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(cricketers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
