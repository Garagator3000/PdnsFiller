package main

import (
	"flag"
	"os"
)

func init() {
	help := flag.Bool("help", false, "usage")
	apiString := flag.String("apistring", ApiString, "pdns server api sring")
	apiKey := flag.String("apikey", ApiKey, "pdns server api secret key")
	domain := flag.String("domain", Domain, "base domain for dns records")
	countOfRecords := flag.Int("count", CountOfRecords, "count of records")
	thirdLvlDomainPrefix := flag.String("3lvlname", ThirdLvlDomainPrefix, "third-level domain")
	maxRecordPerReq := flag.Int("MaxRPR", MaxRecordPerReq, "max records per request")

	flag.Parse()
	func(needHelp bool) {
		if needHelp {
			flag.Usage()
			os.Exit(0)
		}
	}(*help)

	ApiString = *apiString
	ApiKey = *apiKey
	Domain = *domain
	CountOfRecords = *countOfRecords
	ThirdLvlDomainPrefix = *thirdLvlDomainPrefix
	MaxRecordPerReq = *maxRecordPerReq
}
