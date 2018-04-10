package pscan_test

import (
	"fmt"
	"testing"

	"github.com/Konstantin8105/pscan"
)

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

func TestScanAddressLocalhost(t *testing.T) {
	address := "localhost"
	ps, err := pscan.ScanAddress(address)
	if err != nil {
		t.Errorf("Error in scanning address `%v`. %v", address, err)
	}
	for _, p := range ps {
		t.Logf("Port %3d is open", p)
	}
	address = "127.0.0.1"
	if len(ps) > 0 {
		used, err := pscan.Scan(address, ps[0])
		if !used {
			t.Errorf("Addresses `localhost` and `127.0.0.1` is different")
		}
		if err != nil {
			t.Errorf("Error in 127.0.0.1. %v", err)
		}
	}
}

func TestWrongAddress(t *testing.T) {
	used, err := pscan.Scan("wrong address", 229)
	if used {
		t.Errorf("Cannot found open port for wrong address.")
	}
	if err == nil {
		t.Errorf("Error cannot be nil for wrong address.")
	}
	used, err = pscan.Scan("", 229)
	if used {
		t.Errorf("Cannot found open port for wrong address.")
	}
	if err == nil {
		t.Errorf("Error cannot be nil for wrong address.")
	}
	_, err = pscan.ScanAddress("")
	if err == nil {
		t.Errorf("Error cannot be nil for wrong address.")
	}
}

func TestWrongIp(t *testing.T) {
	address := "localhost"
	ip := -23
	used, err := pscan.Scan(address, ip)
	if used || err == nil {
		t.Errorf("Not correct for wrong IP `%v`. %v", ip, err)
	}
	ip = 230000
	used, err = pscan.Scan(address, ip)
	if used || err == nil {
		t.Errorf("Not correct for wrong IP `%v`. %v", ip, err)
	}
}
