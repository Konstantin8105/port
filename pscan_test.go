package pscan_test

import (
	"testing"

	"github.com/Konstantin8105/pscan"
)

func TestScanAddressLocalhost(t *testing.T) {
	address := "localhost"
	ps, err := pscan.ScanAddress(address)
	if err != nil {
		t.Errorf("Error in scanning address `%v`. %v", address, err)
	}
	for _, p := range ps {
		t.Logf("Port %3d is open", p)
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
