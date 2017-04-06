package loadcsv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	// setup reader
	csvIn, err := os.Open("./path/to/datafile.csv")
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(csvIn)

	// setup writerß
	csvOut, err := os.Create("./where/to/write/resultsfile.csv")
	if err != nil {
		log.Fatal("Unable to open output")
	}
	w := csv.NewWriter(csvOut)
	defer csvOut.Close()

	// handle header
	rec, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
	rec = append(rec, "score")
	if err = w.Write(rec); err != nil {
		log.Fatal(err)
	}

	for {
		rec, err = r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		// get float value
		value := rec[1]
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal("Record, error: %v, %v", value, err)
		}

		// calculate scores; THIS EXTERNAL METHOD CANNOT BE CHANGED
		score := calculateStuff(floatValue)

		scoreString := strconv.FormatFloat(score, 'f', 8, 64)
		rec = append(rec, scoreString)

		if err = w.Write(rec); err != nil {
			log.Fatal(err)
		}
		w.Flush()
	}
}
