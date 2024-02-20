package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/wI2L/jsondiff"
)

type ApiResponseDifference struct {
	OriginalApiResponse string `json:"endpoint_1_response"`
	ComparedApiResponse string `json:"endpoint_2_response"`
	Cause               string `json:"difference"`
}

func ResponseCheck(firstApi string, secondApi string) (ApiResponseDifference, error) {
	var response ApiResponseDifference
	jsonForValue := json.RawMessage(firstApi)
	jsonForValue2 := json.RawMessage(secondApi)

	diff, err := jsondiff.CompareJSON(jsonForValue, jsonForValue2)
	if err != nil {
		return response, err
	}

	if diff.String() != "" {
		response.Cause = diff.String()
		response.OriginalApiResponse = firstApi
		response.ComparedApiResponse = secondApi
		return response, nil
	}

	return ApiResponseDifference{}, nil
}

func GetApiResponse(endpoint string, c chan string) {
	urlStruct, err := url.Parse(endpoint)
	if err != nil {
		log.Fatalf("parse error: %s", err.Error())
	}
	request, err := http.NewRequest("GET", urlStruct.String(), nil)
	if err != nil {
		log.Fatalf("issue with request: %s", err.Error())
	}

	client := &http.Client{}

	// Send the request
	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("call error: %s", err.Error())
	}

	defer response.Body.Close()

	// Expecting response with 200 status code
	if response.StatusCode != 200 {
		errMsg := fmt.Sprintf("status code of %s is %d", endpoint, response.StatusCode)
		log.Fatalln(errMsg)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("response issue: %s", err.Error())

	}

	c <- string(body)
}
