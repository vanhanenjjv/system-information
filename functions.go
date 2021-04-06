package sysinfo

import (
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

func findSystemLoad(s string) *float64 {
	e := regexp.MustCompile(`System load:\ +[0-9]+.[0-9]+`)
	sm := e.FindStringSubmatch(s)

	if sm == nil {
		return nil
	}

	l, err := strconv.ParseFloat(sm[1], 64)

	if err != nil {
		log.Fatal(err)
	}

	return &l
}

func findMemoryUsage(s string) *int64 {
	e := regexp.MustCompile(`Memory usage:\ +([0-9]+)%`)
	sm := e.FindStringSubmatch(s)

	if sm == nil {
		return nil
	}

	u, err := strconv.ParseInt(sm[1], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	return &u
}

func findSwapUsage(s string) *int64 {
	e := regexp.MustCompile(`Swap usage:\ +([0-9]+)%`)
	sm := e.FindStringSubmatch(s)

	if sm == nil {
		return nil
	}

	u, err := strconv.ParseInt(sm[1], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	return &u
}

func findProcessCount(s string) *int64 {
	e := regexp.MustCompile(`Processes:\ +([0-9]+)`)
	sm := e.FindStringSubmatch(s)

	if sm == nil {
		return nil
	}

	c, err := strconv.ParseInt(sm[1], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	return &c
}

func findUsersLoggedInCount(s string) *int64 {
	e := regexp.MustCompile(`Users logged in:\ +([0-9]+)`)
	sm := e.FindStringSubmatch(s)

	if sm == nil {
		return nil
	}

	c, err := strconv.ParseInt(sm[1], 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	return &c
}

func findAllNetworkDevice(s string) []NetworkDevice {
	e := regexp.MustCompile(`IPv[46] address for (.+):\s+(.+)`)
	sms := e.FindAllStringSubmatch(s, -1)

	var ds []NetworkDevice

	for _, sm := range sms {
		ds = append(ds, NetworkDevice{
			Interface: sm[1],
			Address:   sm[2],
		})
	}

	return ds
}

func Read() (*SystemInformation, error) {
	b, err := exec.Command("landscape-sysinfo").Output()

	if err != nil {
		return nil, err
	}

	s := string(b)

	i := SystemInformation{
		SystemLoad:         findSystemLoad(s),
		MemoryUsage:        findMemoryUsage(s),
		SwapUsage:          findSwapUsage(s),
		ProcessCount:       findProcessCount(s),
		UsersLoggedInCount: findUsersLoggedInCount(s),
		NetworkDevices:     findAllNetworkDevice(s),
	}

	return &i, nil
}
