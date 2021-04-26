package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	location := flag.String("xml", "", "Location of JUnit.xml file to parse")
	duration := flag.Int("d", 500, "Threshold duration for tests in milliseconds")
	flag.Parse()

	xmlFile, err := os.Open(*location)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer xmlFile.Close()

	bytes, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	testFile, err := Parse(bytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fmt.Sprintf("Total test duration = %v", testFile.Duration))

	tests := DurationsGreaterThan(testFile, *duration)

	fmt.Println(fmt.Sprintf("%v tests found with duration greater than %v", len(tests), *duration))
}
