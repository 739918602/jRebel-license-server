package handler

import (
	"encoding/json"
	"fmt"
	"jRebel-license-server/util"
	"log"
	"net/http"
	"strconv"
	"time"
)

const leasesStr = `{
    "serverVersion": "3.2.4",
    "serverProtocolVersion": "1.1",
    "serverGuid": "a1b4aea8-b031-4302-b602-670a990272cb",
    "groupType": "managed",
    "id": 1,
    "licenseType": 1,
    "evaluationLicense": false,
    "signature": "OJE9wGg2xncSb+VgnYT+9HGCFaLOk28tneMFhCbpVMKoC/Iq4LuaDKPirBjG4o394/UjCDGgTBpIrzcXNPdVxVr8PnQzpy7ZSToGO8wv/KIWZT9/ba7bDbA8/RZ4B37YkCeXhjaixpmoyz/CIZMnei4q7oWR7DYUOlOcEWDQhiY=",
    "serverRandomness": "H2ulzLlh7E0=",
    "seatPoolType": "standalone",
    "statusCode": "SUCCESS",
    "offline": %t,
    "validFrom": %s,
    "validUntil": %s,
    "company": "Administrator",
    "orderId": "",
    "zeroIds": [
        
    ],
    "licenseValidFrom": %d,
    "licenseValidUntil": %d
	}`
const validateConnStr = `{
    "serverVersion": "3.2.4",
    "serverProtocolVersion": "1.1",
    "serverGuid": "a1b4aea8-b031-4302-b602-670a990272cb",
    "groupType": "managed",
    "statusCode": "SUCCESS",
    "company": "Administrator",
    "canGetLease": true,
    "licenseType": 1,
    "evaluationLicense": false,
    "seatPoolType": "standalone"
}`
const leases1Str = `{
    "serverVersion": "3.2.4",
    "serverProtocolVersion": "1.1",
    "serverGuid": "a1b4aea8-b031-4302-b602-670a990272cb",
    "groupType": "managed",
    "statusCode": "SUCCESS",
    "msg": null,
    "statusMessage": null,
    "company": %s,
}`

func Leases(w http.ResponseWriter, request *http.Request) {
	values := util.GetUrlParams(request)
	randomness := values.Get("randomness")
	username := values.Get("username")
	clientTime := values.Get("clientTime")
	guid := values.Get("guid")
	offline, _ := strconv.ParseBool(values.Get("offline"))
	validFrom := "null"
	validUntil := "null"
	//在线校验
	licenseValidFrom := time.Now().Unix() * 1000
	licenseValidUntil := time.Now().AddDate(3, 0, 0).Unix() * 1000
	if offline {
		//离线校验
		ct, _ := strconv.ParseUint(clientTime, 10, 64)
		clientTimeUntil := ct + uint64((180 * 24 * time.Hour).Milliseconds())
		validFrom = clientTime
		validUntil = strconv.FormatUint(clientTimeUntil, 10)
	}
	data := fmt.Sprintf(leasesStr, offline, validFrom, validUntil, licenseValidFrom, licenseValidUntil)
	var jsonObject map[string]interface{}
	err := json.Unmarshal([]byte(data), &jsonObject)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		w.WriteHeader(500)
		util.WriteJson(w, map[string]string{"error": "internal server error"})
		return
	}
	if randomness == "" || username == "" || guid == "" {
		w.WriteHeader(403)
	} else {
		signature := util.RSA{}.Sign(randomness, guid, offline, validFrom, validUntil)
		jsonObject["signature"] = signature
		jsonObject["company"] = username
		util.WriteJson(w, jsonObject)
	}
}
func ValidateConnection(w http.ResponseWriter, request *http.Request) {
	util.WriteJson(w, validateConnStr)
}
func Leases1(w http.ResponseWriter, request *http.Request) {
	values := util.GetUrlParams(request)

	company := values.Get("username")

	jsonStr := fmt.Sprintf(leases1Str, company)
	util.WriteText(w, jsonStr)
}
