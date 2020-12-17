package utils

import "testing"

func SuccessTestReadCSV(t *testing.T) {
	var filepath = "../files/q1_catalog.csv"
	rows, err := ReadCSV(filepath)
	
	if err != nil {
		t.Errorf("ReadCSV('%s') FAILED, expected nil but got error %s", filepath, err)
	}

	if len(rows) != 44 {
		t.Errorf("ReadCSV('%s') FAILED, expected %d but got value %d", filepath, 44, len(rows))
	} else {
		t.Logf("ReadCSV('%s') PASSED, expected %d and got value %d", filepath, 44, len(rows))
	}
}

func TestReadCSVWhenFileIsUnknown(t *testing.T) {
	var filepath = "../files/unknown_file.csv"
	_, err := ReadCSV(filepath)

	if err != nil && err.Error() == "Couldn't open the csv file" {
		t.Logf("ReadCSV('%s') PASSED, got the expected error: '%s'", filepath, err)
	} else {
		t.Errorf("ReadCSV('%s') FAILED, got none or an unexpected error: '%s'", filepath, err)
	}
}

func TestReadCSVWithWrongSeparator(t *testing.T) {
	var filepath = "../test_files/utils/comma-separator.csv"
	_, err := ReadCSV(filepath)

	if err != nil && err.Error() == "CSV file with wrong comma separator. Please, check if is ; and try again." {
		t.Logf("ReadCSV('%s') PASSED, got the expected error: '%s'", filepath, err)
	} else {
		t.Errorf("ReadCSV('%s') FAILED, got none or an unexpected error: '%s'", filepath, err)
	}
}

func TestValidateNameWhenNameIsShort(t *testing.T) {
	var name = "short name"
	res := ValidateName(name)

	if res == true {
		t.Logf("ValidateName('%s') PASSED", name)
	} else {
		t.Errorf("ValidateName('%s') FAILED, got an unexpected result", name)
	}
}

func TestValidateNameWhenLengthIsLowerThan255(t *testing.T) {
	var name = GenerateRandomString(254)
	res := ValidateName(name)

	if res == true {
		t.Logf("ValidateName('%s') PASSED", name)
	} else {
		t.Errorf("ValidateName('%s') FAILED, got an unexpected result", name)
	}
}

func TestValidateNameWhenLengthIsEqualTo255(t *testing.T) {
	var name = GenerateRandomString(255)
	res := ValidateName(name)

	if res == true {
		t.Logf("ValidateName('%s') PASSED", name)
	} else {
		t.Errorf("ValidateName('%s') FAILED, got an unexpected result", name)
	}
}

func TestValidateNameWhenLengthIsGreaterThan255(t *testing.T) {
	var name = GenerateRandomString(256)
	res := ValidateName(name)

	if res == true {
		t.Errorf("ValidateName('%s') FAILED, got an unexpected result", name)
		
	} else {
		t.Logf("ValidateName('%s') PASSED", name)
	}
}


func TestValidateZipWhenValid(t *testing.T) {
	var input = "12345"
	res := ValidateZip(input)

	if res == true {
		t.Logf("ValidateZip('%s') PASSED", input)
	} else {
		t.Errorf("ValidateZip('%s') FAILED, got an unexpected result", input)
	}
}

func TestValidateZipWhenContainsALetter(t *testing.T) {
	var input = "1234f"

	res := ValidateZip(input)

	if res == true {
		t.Errorf("ValidateZip('%s') FAILED, got an unexpected result", input)
	} else {
		t.Logf("ValidateZip('%s') PASSED", input)
	}
}

func TestValidateZipWithSixDigits(t *testing.T) {
	var input = "123456"
	res := ValidateZip(input)

	if res == true {
		t.Errorf("ValidateZip('%s') FAILED, got an unexpected result", input)
	} else {
		t.Logf("ValidateZip('%s') PASSED", input)
	}
}

func TestValidateZipWithJustLetters(t *testing.T) {
	var input = "abcdef"
	res := ValidateZip(input)

	if res == true {
		t.Errorf("ValidateZip('%s') FAILED, got an unexpected result", input)
	} else {
		t.Logf("ValidateZip('%s') PASSED", input)
	}
}

func TestValidateZipWithLettersAndDigits(t *testing.T) {
	var input = "1a2b3c4d5f"
	res := ValidateZip(input)

	if res == true {
		t.Errorf("ValidateZip('%s') FAILED, got an unexpected result", input)
	} else {
		t.Logf("ValidateZip('%s') PASSED", input)
	}
}
