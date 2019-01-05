package approvaltests_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	// this is a mock testing.T for documentation purposes
	t = &failing{}
)

// failing is a mock struct that is only there for documentation conveniance,
// showing the developer how they would be passing a *testing.T pointer in their
// normal tests.
type failing struct{}

// Fail implements approvaltest.Fail
func (f *failing) Fail() {}

// documentation helper just for the example
func printFileContent(path string) {
	approvedPath := strings.Replace(path, ".received.", ".approved.", 1)
	content, err := ioutil.ReadFile(approvedPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("This produced the file %s\n", path)
	fmt.Printf("It will be compared against the %s file\n", approvedPath)
	fmt.Println("and contains the text:")
	fmt.Println()
	// sad sad hack because go examples trim blank middle lines
	fmt.Println(strings.Replace(string(content), "\n\n", "\n", -1))
}
