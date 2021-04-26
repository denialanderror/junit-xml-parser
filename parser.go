package main

import (
	"encoding/xml"
	"regexp"
	"strconv"
	"strings"
)

var re *regexp.Regexp

// JUnitFile is a representation of the top level XML
type JUnitFile struct {
	Duration string  `xml:"duration"`
	Suites   []Suite `xml:"suite"`
}

// Suite is a representation of a single test file
type Suite struct {
	Cases       []Case   `xml:"case"`
	Breadcrumbs []string `xml:"enclosingBlockName"`
	Duration    string   `xml:"duration"`
}

// Case is a representation of a single test
type Case struct {
	Name     string `xml:"className"`
	Duration string `xml:"duration"`
}

// Test is a single test with a duration and parent name in breadcrumbs
type Test struct {
	Name        string
	Duration    int
	Breadcrumbs string
}

// Parse unmarshals a byte array into a JUnitFile, or throws error if encountered first
func Parse(bytes []byte) (JUnitFile, error) {
	tests := JUnitFile{}
	err := xml.Unmarshal(bytes, &tests)

	return tests, err
}

// DurationsGreaterThan filters the cases in a JUnitFile that have durations greater than ms
func DurationsGreaterThan(tests JUnitFile, ms int) []Test {
	matching := []Test{}

	for _, s := range tests.Suites {
		tests := []Test{}

		for _, c := range s.Cases {

			d := durationInMs(c.Duration)
			if d > ms {
				tests = append(tests, Test{c.Name, d, strings.Join(s.Breadcrumbs, ">")})
			}
		}
		if len(tests) > 0 {
			matching = append(matching, tests...)
		}
	}

	return matching
}

func durationInMs(d string) int {
	if re == nil {
		re = regexp.MustCompile(`^\d+\.?\d*`)
	}
	d = re.FindString(d)
	f, _ := strconv.ParseFloat(d, 32)
	f = f * 100
	return int(f)
}
