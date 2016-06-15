package util

import (
	"encoding/json"
	"io"
	"log"
)

//LogErrOrResp logs if error or print if resp
func LogErrOrResp(resp interface{}, err error) {
	if err != nil {
		log.Print(err)
	} else {
		log.Print(resp)
	}
}

//LogErr logs if error exists
func LogErr(err error) {
	if err != nil {
		log.Print(err)
	}
}

//UnmarshalResponse unmarshals web response
func UnmarshalResponse(body io.Reader, input interface{}) {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&input)
	LogErr(err)
}
