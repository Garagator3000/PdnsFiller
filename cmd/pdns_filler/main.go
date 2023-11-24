package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"strconv"
)

var (
	ApiString            = "http://localhost:8081/api/v1/"
	ApiKey               = "apipass"
	Domain               = "domain.ru"
	CountOfRecords       = 1
	ThirdLvlDomainPrefix = "test"
	MaxRecordPerReq      = 1
)

const (
	CreateRecord = "servers/localhost/zones/"

	HeaderToken = "X-API-Key"
)

func main() {
	Fill_Dns(Prepare_DTO())
}

func Fill_Dns(dataset *Filler_DTO) {
	// for i := range dto.Rrsets {
	// 	for k := range dto.Rrsets[i].Records {
	// 		fmt.Println(dto.Rrsets[i].Records[k].Name + "  --  " + dto.Rrsets[i].Records[k].Content)
	// 	}
	// }
	client := http.Client{}
	for i := range dataset.Rrsets {
		var dto Filler_DTO
		dto.Rrsets = append(dto.Rrsets, dataset.Rrsets[i])
		payload, err := json.Marshal(dto)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(payload))
		req, err := http.NewRequest(http.MethodPatch, ApiString+CreateRecord+Domain+".", bytes.NewReader(payload))
		if err != nil {
			log.Println(err)
		}
		req.Header.Add(HeaderToken, ApiKey)
		req.Header.Add("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(req.URL.String() + " -- " + resp.Status + " -- " + string(body))
		resp.Body.Close()
	}
}

func Prepare_DTO() *Filler_DTO {
	var dto Filler_DTO
	var rrsets []Rrset

	for i := 0; CountOfRecords != 0; i++ {
		var records []Record
		for k := 0; k < MaxRecordPerReq && k < CountOfRecords; k++ {
			record := Record{}
			if MaxRecordPerReq == 1 {
				record.Content = "192.168." + strconv.Itoa(i%255) + "." + strconv.Itoa((rand.Int())%255)
			} else {
				record.Content = "192.168." + strconv.Itoa(i%255) + "." + strconv.Itoa(k%255)
			}
			record.Disabled = false
			record.Name = ThirdLvlDomainPrefix + strconv.Itoa(i) + strconv.Itoa(k) + "." + Domain + "."
			record.Type = "A"
			record.Priority = 0
			records = append(records, record)
			CountOfRecords--
		}
		rrset := Rrset{}
		rrset.Name = ThirdLvlDomainPrefix + strconv.Itoa(i) + "." + Domain + "."
		rrset.Type = "A"
		rrset.Changetype = "REPLACE"
		rrset.TTL = 86400
		rrset.Records = records

		rrsets = append(rrsets, rrset)
	}
	dto.Rrsets = rrsets

	return &dto
}
