# junit-xml-parser

Written to parse a large XML file in order to find slow-running tests, and to practice Go.

## Usage

Build with `go build` and then run `junit-parser` with the following arguments:

`-xml=<path/to/xmlfile>`

`-d=<duration in ms, defaulted to 500 if not passed>`

`-out=<output directory for csv file (directory must exist), defaults to current directory>`
