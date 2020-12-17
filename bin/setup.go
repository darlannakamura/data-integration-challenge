package main

import (
	"os"
	"fmt"
	"strings"
	utils "../utils"
	db "../db"
	"github.com/akamensky/argparse"
)

func main() {
	fmt.Println("Running setup ...")

	parser := argparse.NewParser("parser", "")

	filename := parser.String("f", "file", &argparse.Options{Required: true, Help: "Insert the CSV file path"})
	
	err := parser.Parse(os.Args)
	
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	fmt.Println("Reading CSV file:", *filename)

	rows, err := utils.ReadCSV(*filename)

	if err != nil { 
		fmt.Println(err)
		return
	}

	fmt.Println("Recreating companies table ...")

	db.DropCompanyTable()
	db.CreateCompanyTable()

	for i := 0; i < len(rows); i++ {
		var name = rows[i].Name
		var zip = rows[i].Zip

		if (utils.ValidateName(name) && utils.ValidateZip(zip)) {
			name := strings.ToUpper(name)
			db.InsertCompany(name, zip)
		}
	}

	fmt.Println("Your database has been populated successfully.")
}