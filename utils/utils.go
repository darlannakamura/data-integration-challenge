package utils

import (
	"encoding/csv"
	"io"
	"os"
	"regexp"
	"math/rand"
	"time"
)

type CSVRow struct {
	Name string
	Zip string
}

type ReadCSVError struct{
	message string
}

func (e *ReadCSVError) Error() string {  
    return e.message
}

func NewReadCSVError(text string) error {
	return &ReadCSVError{text}
}

func ReadCSV(filename string) ([]CSVRow, error) {
	rows := []CSVRow{}

	csvfile, err := os.Open(filename)
	
	if err != nil {
		return rows, NewReadCSVError("Couldn't open the csv file")
	}

	reader := csv.NewReader(csvfile)
	reader.Comma = ';'

	var rowIndex int = 0

	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return rows, NewReadCSVError("Couldn't read csv file")
		}

		if len(row) != 2 {
			return rows, NewReadCSVError("CSV file with wrong comma separator. Please, check if is ; and try again.")
		}

		if rowIndex != 0 {
			rows = append(rows, CSVRow{Name: row[0], Zip: row[1]})
		}

		rowIndex++
	}
	
	return rows, nil
}

func ValidateName(name string) bool {
	if len(name) <= 255 {
		return true
	}

	return false
}

func ValidateZip(zip string) bool {
	if len(zip) == 5 {
		re := regexp.MustCompile("[0-9]+")
		arr := re.FindAllString(zip, -1)

		if (len(arr) == 1 && len(arr[0]) == 5) {
			return true
		}
	}

	return false
}

func StringWithCharset(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
	  b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	return StringWithCharset(length, charset)
}