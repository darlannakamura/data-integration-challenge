package db

import "testing"

func TestCreateDBConnection(t *testing.T) {
	_, err := CreateDBConnection(DBCredentials{
		Host: "localhost",
		Port: 5432,
		Name: "yawoen",
		Username: "admin",
		Password: "admin",
	})

	if err != nil {
		t.Errorf("CreateDBConnection() FAILED, got an unexpected error: %s", err)
	}
}

func TestCreateCompanyTable(t *testing.T) {
	err := CreateCompanyTable()

	if err != nil {
		t.Errorf("CreateCompanyTable() FAILED, got an unexpected error: %s", err)
	}
}

func TestInsertCompany(t *testing.T) {
	var name = "RANDOM NEW COMPANY"
	var zip = "12345"
	var quantityBeforeInsert, getErr = GetQuantityOfCompanies()

	if getErr != nil {
		t.Errorf("GetQuantityOfCompanies() FAILED, got an unexpected error: %s", getErr)
	}

	err := InsertCompany(name, zip)

	var quantityAfterInsert, getAfterErr = GetQuantityOfCompanies()

	if getAfterErr != nil {
		t.Errorf("GetQuantityOfCompanies() FAILED, got an unexpected error: %s", err)
	}

	if quantityBeforeInsert + 1 != quantityAfterInsert {
		t.Errorf("InsertCompany('%s', '%s') FAILED, assertion failed when quantity before insert +1 was different from after insert.", name, zip)
	}

	if err != nil {
		t.Errorf("InsertCompany('%s', '%s') FAILED, got an unexpected error: %s", name, zip, err)
	} else{
		err := DeleteCompany(name, zip)

		if err != nil {
			t.Errorf("DeleteCompany('%s', '%s') FAILED, got an unexpected error: %s", name, zip, err)
		}
	}
}
