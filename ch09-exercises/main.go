/*
Look at the code in the sample_code/exercise directory in the Chapter 9 repository.
You are going to modify this code in each of these exercises.
It works correctly, but improvements should be made to its error handling.
https://github.com/learning-go-book-2e/ch09

1.
Create a sentinel error to represent an invalid ID.
In main, use errors.Is to check for the sentinel error, and print a message when it is found.

2.
Define a custom error type to represent an empty field error.
This error should include the name of the empty Employee field.
In main, use errors.As to check for this error.
Print out a message that includes the field name.

3.
Rather than returning the first error found, return back a single error that
contains all errors discovered during validation.
Update the code in main to properly report multiple errors.
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	d := json.NewDecoder(strings.NewReader(data))
	count := 0
	for d.More() {
		count++
		var emp Employee
		err := d.Decode(&emp)
		if err != nil {
			fmt.Printf("record %d: %v\n", count, err)
			continue
		}
		// tehtävä 3 ratkaisu löytyy tuolta kirjan reposta. En itse olisi ikinä saanut väsättyä tuota...
		err = ValidateEmployee(emp)
		if err != nil {
			if errors.Is(err, ErrInvalidID){
				fmt.Printf("invalid id: record %d: %+v error: %v\n", count, emp, err)
				continue
			}
			var emptyErr EmptyFieldErr
			if errors.As(err, &emptyErr){
				fmt.Printf("record %d: %+v empty field error: %v\n", count, emp, err)
				continue
			}
			fmt.Printf("record %d: %+v error: %v\n", count, emp, err)
			continue
		}
		fmt.Printf("record %d: %+v\n", count, emp)
	}
}

const data = `
{
	"id": "ABCD-123",
	"first_name": "Bob",
	"last_name": "Bobson",
	"title": "Senior Manager"
}
{
	"id": "XYZ-123",
	"first_name": "Mary",
	"last_name": "Maryson",
	"title": "Vice President"
}
{
	"id": "BOTX-263",
	"first_name": "",
	"last_name": "Garciason",
	"title": "Manager"
}
{
	"id": "HLXO-829",
	"first_name": "Pierre",
	"last_name": "",
	"title": "Intern"
}
{
	"id": "MOXW-821",
	"first_name": "Franklin",
	"last_name": "Watanabe",
	"title": ""
}
{
	"id": "",
	"first_name": "Shelly",
	"last_name": "Shellson",
	"title": "CEO"
}
{
	"id": "YDOD-324",
	"first_name": "",
	"last_name": "",
	"title": ""
}
`

type Employee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Title     string `json:"title"`
}

var (
	validID = regexp.MustCompile(`\w{4}-\d{3}`)
)

func ValidateEmployee(e Employee) error {
	var allErrors []error

	if len(e.ID) == 0 {
		allErrors = append(allErrors, EmptyFieldErr{FieldName: "ID"})
	}
	if !validID.MatchString(e.ID) {
		//return errors.New("invalid ID")
		allErrors = append(allErrors, ErrInvalidID)
	}
	if len(e.FirstName) == 0 {
		allErrors = append(allErrors, EmptyFieldErr{FieldName: "FirstName"})
	}
	if len(e.LastName) == 0 {
		allErrors = append(allErrors, EmptyFieldErr{FieldName: "LastName"})
	}
	if len(e.Title) == 0 {
		allErrors = append(allErrors, EmptyFieldErr{FieldName: "Title"})
	}

	switch len(allErrors){
	case 0:
		return nil
	case 1:
		return allErrors[0]
	default:
		return errors.Join(allErrors...)
	}
}

var ErrInvalidID = errors.New("invalid ID")

type EmptyFieldErr struct {
	FieldName string
}

func (e EmptyFieldErr) Error() string{
	return e.FieldName
}
