package main

import (
	"os"
	"log"
	"strings"
	utils "../utils"
	db "../db"
	"github.com/akamensky/argparse"
)

func main() {
	log.Println("Running setup ...")

	parser := argparse.NewParser("parser", "")

	filename := parser.String("f", "file", &argparse.Options{Required: true, Help: "Insert the CSV file path"})
	
	err := parser.Parse(os.Args)
	
	if err != nil {
		log.Fatal(parser.Usage(err))
	}

	log.Println("Reading CSV file:", *filename)

	rows, err := utils.ReadCSV(*filename)

	if err != nil { 
		log.Fatal(err)
		return
	}

	log.Println("Creating companies table ...")

	db.CreateCompanyTable()

	for i := 0; i < len(rows); i++ {
		var name = rows[i].Name
		var zip = rows[i].Zip

		if (utils.ValidateName(name) && utils.ValidateZip(zip)) {
			name := strings.ToUpper(name)
			db.InsertCompany(name, zip)
		}
	}

	log.Println("Your database has been populated successfully.")
}