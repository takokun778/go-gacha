package csv

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
)

const (
	column0 = "name"
	column1 = "rate"
)

func Import(name string) ([][]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	reader := csv.NewReader(f)

	header, err := reader.Read()
	if err != nil {
		return nil, err
	}

	if header[0] != column0 {
		return nil, errors.New("column 0 is not a " + column0)
	}

	if header[1] != column1 {
		return nil, errors.New("column 1 is not a " + column1)
	}

	result := [][]string{}

	for {
		rec, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, rec)
	}

	return result, nil
}
