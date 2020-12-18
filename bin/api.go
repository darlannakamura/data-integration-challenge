package main

import (
	"os"
	"log"
	"errors"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/akamensky/argparse"
	utils "../utils"
	db "../db"
	files "../files"
)

const STATUS_HTTP_UNPROCESSABLE_ENTITY = 422

type Response struct {
	Message string `json:"message"`
}

func Greetings(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(Response{Message: "API is up"})
	w.Write(res)

	utils.LoggingRequest(r, http.StatusOK)
}

func handleUnprocessableEntity(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(STATUS_HTTP_UNPROCESSABLE_ENTITY)
	res, _ := json.Marshal(Response{Message: err.Error()})
	w.Write(res)

	utils.LoggingRequest(r, STATUS_HTTP_UNPROCESSABLE_ENTITY)
}

func UploadCsv(w http.ResponseWriter, r *http.Request) {
	filename, err := files.SaveUploadedCsv(r)
	
	if err != nil {
		handleUnprocessableEntity(w, r, err)
		return
	}

	var filepath = "upload-files/"+filename

	rows, readError := utils.ReadCSV(filepath)

	if readError != nil {
		handleUnprocessableEntity(w, r, readError)
		return
	}
	
	for i := 0; i < len(rows); i++ {
		var name = rows[i].Name
		var zip = rows[i].Zip
		var website = rows[i].Website

		if utils.ValidateFields(name, zip, website) {
			id, err := db.GetCompanyIdByNameAndZip(name, zip)

			if err != nil {
				log.Fatal(err)
			}

			if id != 0 {
				db.UpdateCompanyWebsite(id, website)
			}
		}
	}

	files.DeleteFile(filepath)

	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(Response{Message: filename})
	w.Write(res)

	utils.LoggingRequest(r, http.StatusOK)
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var name = vars["name"]
	var zip = vars["zip"]

	if (utils.ValidateName(name) && utils.ValidateZip(zip)) {
		company, err := db.GetCompanyByNameAndZip(name, zip)

		if err != nil {
			handleUnprocessableEntity(w, r, err)
		}

		if company == nil {
			w.WriteHeader(http.StatusNotFound)
			res, _ := json.Marshal(Response{Message: "company not found"})
			w.Write(res)

			utils.LoggingRequest(r, http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
			res, _ := json.Marshal(company)
			w.Write(res)
	
			utils.LoggingRequest(r, http.StatusOK)			
		}

	} else {
		handleUnprocessableEntity(w, r, errors.New("Wrong parameters. Name must be lower than 256 characters and zip code must have 5 digits"))
	}

}
 
func main() {
	parser := argparse.NewParser("parser", "")

	var port *int = parser.Int("p", "port", &argparse.Options{Required: false, Help: "Port to run API", Default: 8080})
	
	err := parser.Parse(os.Args)
	
	if err != nil {
		log.Print(parser.Usage(err))
	}

	log.Println("API running at http://localhost:"+strconv.Itoa(*port))

	router := mux.NewRouter()

	router.Path("/").HandlerFunc(Greetings)
	router.Path("/upload").Methods("POST").HandlerFunc(UploadCsv)
	router.Path("/company/{name}/{zip:[0-9]+}").HandlerFunc(GetCompany)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), router))
}