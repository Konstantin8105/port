# Port scanner

[![Coverage Status](https://coveralls.io/repos/github/Konstantin8105/port/badge.svg?branch=master)](https://coveralls.io/github/Konstantin8105/port?branch=master)
[![Build Status](https://travis-ci.org/Konstantin8105/port.svg?branch=master)](https://travis-ci.org/Konstantin8105/port)
[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin8105/port)](https://goreportcard.com/report/github.com/Konstantin8105/port)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Konstantin8105/port/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/Konstantin8105/port?status.svg)](https://godoc.org/github.com/Konstantin8105/port)

simple port scanner

Example of using:
```golang
func ExampleScanAddress() {
	address := "localhost"
	ps, err := port.ScanAddress(address)
	if err != nil {
		return
	}
	if len(ps) > 0 {
		fmt.Println("Found opened ports")
	}

	// Output:
	// Found opened ports
}
```
