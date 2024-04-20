package main

import (
	_ "embed"
	"flag"
	"github.com/oschwald/geoip2-golang"
	"log"
	"log/syslog"
	"net"
	"os"
)

//go:embed src/GeoLite2-Country.mmdb
var mmdb []byte

func main() {
	logger, err := syslog.New(syslog.LOG_AUTHPRIV|syslog.LOG_NOTICE, "ssh-geo-ip-blocker")
	if err != nil {
		os.Exit(1)
	}
	log.SetOutput(logger)

	flag.Parse()
	argsIp := flag.Arg(0)
	if argsIp == "" {
		log.Printf("Error:Arg\n")
		os.Exit(0)
	}

	db, err := geoip2.FromBytes(mmdb)
	if err != nil {
		log.Printf("Error:geoip2.FromBytes\n")
		os.Exit(0)
	}

	ip := net.ParseIP(argsIp)
	if ip == nil {
		log.Printf("Error:ParseIP\n")
		os.Exit(0)
	}

	record, err := db.Country(ip)
	if err != nil {
		os.Exit(0)
	}

	haystack := []string{"JP", ""}
	response := inArray(haystack, record.Country.IsoCode)

	if response {
		//許可
		log.Printf("%s sshd connection from %s (Country:%s)\n", access(response), argsIp, record.Country.IsoCode)
		os.Exit(0)
	}

	//不許可
	log.Printf("%s sshd connection from %s (Country:%s)\n", access(response), argsIp, record.Country.IsoCode)
	os.Exit(1)
}

func access(b bool) string {
	if b {
		return "Allow"
	}
	return "Deny"
}

func inArray(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
