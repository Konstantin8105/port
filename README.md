# pscan

[![Coverage Status](https://coveralls.io/repos/github/Konstantin8105/pscan/badge.svg?branch=master)](https://coveralls.io/github/Konstantin8105/pscan?branch=master)
[![Build Status](https://travis-ci.org/Konstantin8105/pscan.svg?branch=master)](https://travis-ci.org/Konstantin8105/pscan)
[![Go Report Card](https://goreportcard.com/badge/github.com/Konstantin8105/pscan)](https://goreportcard.com/report/github.com/Konstantin8105/pscan)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/Konstantin8105/pscan/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/Konstantin8105/pscan?status.svg)](https://godoc.org/github.com/Konstantin8105/pscan)

simple port scanner

Example of using:
```golang
func ExampleScanAddress() {
	address := "localhost"
	ps, err := pscan.ScanAddress(address)
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
