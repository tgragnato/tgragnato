---
title: Elasticsearch Log Analysis
---

```go
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/oschwald/geoip2-golang"
)

type logline struct {
	IP              net.IP    `json:"ip"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	Country         string    `json:"country"`
	ASN             int       `json:"asn"`
	ASNOrganization string    `json:"asn_organization"`
	Timestamp       time.Time `json:"timestamp"`
	Frontend        string    `json:"frontend"`
	Backend         string    `json:"backend"`
	Server          string    `json:"server"`
	ResponseCode    int       `json:"response_code"`
	Host            string    `json:"host"`
	URL             string    `json:"url"`
	UserAgent       string    `json:"user_agent"`
	Method          string    `json:"method"`
	Protocol        string    `json:"protocol"`
}

func main() {
	file, err := os.Open("haproxy.log")
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	dbasn, err := geoip2.Open("GeoLite2-ASN.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer dbasn.Close()

	dbcity, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer dbcity.Close()

	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			"https://<host_1>",
			"https://<host_2>",
            "...",
			"https://<host_n>",
		},
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	_, err = es.Ping()
	if err != nil {
		log.Fatalf("Error pinging the elasticsearch servers: %s", err)
	}

	es.Indices.Create("incident")

	for scanner.Scan() {
		log := scanner.Text()
		temp := strings.SplitN(log, ":", 4)
		if len(temp) < 4 {
			continue
		}
		log = strings.TrimLeft(temp[3], " ")

		// IP
		temp = strings.SplitN(log, ":", 2)
		if len(temp) < 2 {
			continue
		}
		ip := net.ParseIP(temp[0])
		temp2 := strings.SplitN(temp[1], "[", 2)
		if len(temp2) < 2 {
			continue
		}
		log = temp2[1]

		// Country, Latitude, Longitude
		record, err := dbcity.City(ip)
		if err != nil {
			fmt.Println(err.Error())
		}
		country := record.Country.IsoCode
		latitude := record.Location.Latitude
		longitude := record.Location.Longitude

		// Autonomous System Number, Autonomous System Organization
		recordASN, err := dbasn.ASN(ip)
		if err != nil {
			fmt.Println(err.Error())
		}
		asn := recordASN.AutonomousSystemNumber
		asnOrg := recordASN.AutonomousSystemOrganization

		// Timestamp
		temp = strings.SplitN(log, "]", 2)
		layout := "02/Jan/2006:15:04:05.000"
		timestamp, err := time.Parse(layout, temp[0])
		if err != nil {
			fmt.Println(err.Error())
		}
		log = strings.TrimLeft(temp[1], " ")

		// Frontend
		temp = strings.SplitN(log, " ", 2)
		frontend := strings.TrimRight(temp[0], "~")
		log = temp[1]

		// Backend/Server
		temp = strings.SplitN(log, " ", 3)
		temp2 = strings.SplitN(temp[0], "/", 2)
		if len(temp2) < 2 {
			continue
		}
		backend := temp2[0]
		server := temp2[1]
		log = temp[2]

		// Response Code
		temp = strings.SplitN(log, " ", 7)
		responseCode, err := strconv.Atoi(temp[0])
		if err != nil {
			continue
		}
		log = temp[6]

		// {Host|URL|UserAgent}
		temp = strings.SplitN(log, "{", 2)
		if len(temp) < 2 {
			continue
		}
		log = temp[1]
		temp = strings.SplitN(log, "}", 2)
		temp2 = strings.SplitN(temp[0], "|", 3)
		host := temp2[0]
		userAgent := temp2[2]
		log = strings.TrimLeft(temp[1], " ")

		// Method Url Protocol
		temp = strings.SplitN(log, " ", 3)
		if len(temp) < 3 {
			continue
		}
		method := strings.TrimLeft(temp[0], "\"")
		url := strings.TrimLeft(temp[1], "\"")
		protocol := strings.TrimRight(temp[2], "\"")

		logline := logline{
			IP:              ip,
			Latitude:        latitude,
			Longitude:       longitude,
			Country:         country,
			ASN:             int(asn),
			ASNOrganization: asnOrg,
			Timestamp:       timestamp,
			Frontend:        frontend,
			Backend:         backend,
			Server:          server,
			ResponseCode:    responseCode,
			Host:            host,
			URL:             url,
			UserAgent:       userAgent,
			Method:          method,
			Protocol:        protocol,
		}

		data, _ := json.Marshal(logline)
		es.Index("incident", strings.NewReader(string(data)))
	}

	if err := scanner.Err(); err != nil {
		log.Println(err.Error())
	}
}
```
