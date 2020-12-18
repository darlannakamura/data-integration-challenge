setup:
	go build bin/setup.go
	./setup -f test-files/q1_catalog.csv
start:
	go build bin/api.go
	./api
check:
	go test ./pkgs/utils
	go test ./pkgs/db
