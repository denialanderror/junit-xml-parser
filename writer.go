package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

// WriteToFile writes the test data to a CSV in the specified directory
func WriteToFile(data []Test, dir string) {
	f, err := os.Create(dir + "/junit.csv")
	if err != nil {
		log.Fatal(err)
	}

	w := csv.NewWriter(f)

	for _, t := range data {
		err := w.Write([]string{t.Name, strconv.Itoa(t.Duration), t.Breadcrumbs})
		if err != nil {
			log.Fatal(err)
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
