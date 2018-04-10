package pscan

import (
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"
)

func Scan(address string, port int) (used bool, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Error scan %v:%v. %v",
				address, port, err)
		}
	}()
	if port < 1 {
		err = fmt.Errorf("port `%v` cannot less 1", port)
		return
	}
	if port > 65535 {
		err = fmt.Errorf("port `%v` cannot be more 65535", port)
		return
	}

	u, err := url.Parse(address)
	if err != nil || len(u.Path) == 0 {
		err = fmt.Errorf("Undefined host `%v` or error = %v",
			u.Path, err)
		return
	}
	if strings.ContainsAny(u.Path, ": ") {
		err = fmt.Errorf("Not acceptable char in address path")
		return
	}

	conn, err := net.DialTimeout(
		"tcp", fmt.Sprintf("%v:%v", address, port), 10*time.Millisecond)
	if err == nil {
		// error ignored, because the error is
		// not identify the problem
		_ = conn.Close()
		used = true
	}
	err = nil
	return
}

func ScanAddress(address string) (usedPort []int, err error) {
	for port := 1; port <= 65535; port++ {
		var used bool
		used, err = Scan(address, port)
		if err != nil {
			return
		}
		if used {
			usedPort = append(usedPort, port)
		}
	}
	return
}
