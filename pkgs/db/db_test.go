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

func TestUpdateCompany(t *testing.T) {
	var name = "random new company"
	var zip = "12345"
	var website = "https://random.company.com"

	err := InsertCompany(name, zip)

	if err != nil {
		t.Errorf("%s('%s','%s') FAILED, got an unexpected error: %s", "InsertCompany", name, zip, err)
	}

	id, err := GetCompanyIdByNameAndZip(name, zip)

	if err != nil {
		t.Errorf("%s('%s','%s') FAILED, got an unexpected error: %s", "GetCompanyIdByNameAndZip", name, zip, err)
	}

	if id == 0 {
		t.Errorf("%s('%s','%s') FAILED, got 0 in company id", "GetCompanyIdByNameAndZip", name, zip)
	}

	updateError := UpdateCompanyWebsite(id, website)

	if updateError != nil {
		t.Errorf("%s('%s','%s') FAILED, got an unexpected error: %s", "UpdateCompanyWebsite", name, zip, updateError)
	}

	delError := DeleteCompany(name, zip)

	if delError != nil {
		t.Errorf("%s('%s','%s') FAILED, got an unexpected error: %s", "DeleteCompany", name, zip, delError)
	}
}

func TestGetCompanyByNameAndZip(t *testing.T) {
	var name = "super random company"
	var zip = "00001"

	insertErr := InsertCompany(name, zip)

	if insertErr != nil {
		t.Errorf("%s('%s','%s') FAILED, got an unexpected error: %s", "InsertCompany", name, zip, insertErr)
	}

	company, err := GetCompanyByNameAndZip(name, zip)

	if err != nil {
		t.Errorf("%s('%s','%s') FAILED, got an unexpected error: %s", "GetCompanyByNameAndZip", name, zip, err)
	}

	if company == nil {
		t.Errorf("%s('%s','%s') FAILED, got nil when getting existent company", "GetCompanyByNameAndZip", name, zip)
	}

	delErr := DeleteCompany(name, zip)

	if delErr != nil {
		t.Errorf("%s('%s','%s') FAILED, got an unexpected error: %s", "DeleteCompany", name, zip, delErr)
	}
}

func TestGetInexistentCompanyByNameAndZip(t *testing.T) {
	var name = "inexistentcompany"
	var zip = "00009"

	company, err := GetCompanyByNameAndZip(name, zip)

	if err != nil {
		t.Errorf("%s('%s','%s') FAILED, got an unexpected error: %s", "GetCompanyByNameAndZip", name, zip, err)
	}

	if company != nil {
		t.Errorf("%s('%s','%s') FAILED, got a company when getting an inexistent company", "GetCompanyByNameAndZip", name, zip)
	}
}