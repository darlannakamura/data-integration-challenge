# Data integration challenge

## Setup
Follow the steps bellow to setup the project. The first step is to setup [PostgreSQL](https://www.postgresql.org/). If you already have a local PostgreSQL instance on your computer, skip this step. 

## Starting PostgreSQL service with Docker container
 We are considering that you have Docker installed on your computer. If you don't have it yet, follow these steps to install [Docker](https://docs.docker.com/engine/install/) and [Docker Compose](https://docs.docker.com/compose/install/).

After successfully installed, run the following command:

    docker-compose up -d --build

This command will start two services: [PgAdmin](https://www.pgadmin.org/) and [PostgreSQL](https://www.postgresql.org/). PgAdmin is a powerful web UI for accessing Postgres databases. PostgreSQL is the Database Management System used for this application.

### Running Application Locally
This means that you want to run the application directly from your OS system. To do this, you must have [Golang](https://golang.org/) (version >= 1.15.6) installed on your computer.

#### Setting Database Crendetials
You must define database credentials as environment variables. If you have a Linux-based OS, run the commands:

```bash
export DB_HOST=YOUR_DATABASE_HOST
export DB_PORT=YOUR_DATABASE_PORT
export DB_NAME=YOUR_DATABASE_NAME
export DB_USER=YOUR_DATABASE_USER
export DB_PASS=YOUR_DATABASE_PASSWORD
```

Replace `YOUR_DATABASE_HOST` for the specific host. For example, if you are running postgres locally, it will be `localhost` and the `YOUR_DATABASE_PORT` probably will be the default `5432`. `YOUR_DATABASE_NAME` will be the database name, e.g: `yawoen` and lastly, `YOUR_DATABASE_USER` and `YOUR_DATABASE_PASSWORD` are the database credentials.

#### Populating the database

Now all the settings are done, and we can populate the database. We can compile our setup source-code running:

    go build bin/setup.go

And run with the support file provided:

    ./setup -f test-files/q1_catalog.csv

This program will create a table `companies` in your database, then populating with the valid companies in the CSV file. Feel free to choose another CSV file, with other companies if you want.

#### Running the API
Compile the API source-code and run:

    go build bin/api.go && ./api

Or using `Makefile`:

    make start

**Note**: *you can specify the port that the API will run, by default is 8080, but you can change using the `-p` or `--port` parameter.* E.g, to run at 9000, you can just run: `./api -p 9000`. 

## API Endpoints

### /upload - Upload CSV file

The endpoint `/upload` will receive a CSV file, updating the company website when the company already exists in the database.

You can test this endpoint sending a request via [cURL](https://curl.se/):

    curl http://localhost:8080/upload -F "fileupload=@test-files/q2_clientData.csv" -vvv

### /company - Retrieve Company Data

The endpoint `/company/{name}/{zip_code}` will return the company data in JSON.

For example: send a `GET` to `http://localhost:8081/company/tim/53115` and it will return:

```json
{
    "id": 29,
    "name": "TIM DIEBALL",
    "zip": "53115",
    "website": "http://motorsport-coatings.com"
}
```

## Directory Structure

```bash
data-integration-challenge/
├── bin/ # executable files
│   ├── setup # setup file to populate database
│   ├── api # starts the API server
│
├── pkgs/ # packages
│   ├── db # database connection package
│   ├── files # package for handling files
│   ├── utils # package for validation and logging
│
├── test-files/ # required files to run tests
```

## Testing

Run the tests using `Makefile`:

    make check