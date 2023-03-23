package partners_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"testing"

	"github.com/sergeyandreenko/unisender/api"
	"github.com/sergeyandreenko/unisender/partners"
	"github.com/sergeyandreenko/unisender/test"
)

func TestGetSenderDomainListRequest_Domain(t *testing.T) {
	expectedLogin := test.RandomString(12, 36)

	expectedDomain := test.RandomString(12, 36)
	var givenDomain string

	expectedResult := randomGetSenderDomainListResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenDomain = req.FormValue("domain")

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := partners.GetSenderDomainList(req, expectedLogin).
		Domain(expectedDomain).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedDomain != givenDomain {
		t.Fatalf(`Domain should be "%s", "%s" given`, expectedDomain, givenDomain)
	}
}

func TestGetSenderDomainListRequest_Limit(t *testing.T) {
	expectedLogin := test.RandomString(12, 36)

	expectedLimit := test.RandomInt(1, 99)
	var givenLimit int

	expectedResult := randomGetSenderDomainListResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLimit, _ = strconv.Atoi(req.FormValue("limit"))

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := partners.GetSenderDomainList(req, expectedLogin).
		Limit(expectedLimit).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLimit != givenLimit {
		t.Fatalf("Limit should be %d, %d given", expectedLimit, givenLimit)
	}
}

func TestGetSenderDomainListRequest_Offset(t *testing.T) {
	expectedLogin := test.RandomString(12, 36)

	expectedOffset := test.RandomInt(1, 99)
	var givenOffset int

	expectedResult := randomGetSenderDomainListResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenOffset, _ = strconv.Atoi(req.FormValue("offset"))

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := partners.GetSenderDomainList(req, expectedLogin).
		Offset(expectedOffset).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedOffset != givenOffset {
		t.Fatalf("Offset should be %d, %d given", expectedOffset, givenOffset)
	}
}

func TestGetSenderDomainListRequest_Execute(t *testing.T) {
	expectedLogin := test.RandomString(12, 36)
	var givenLogin string

	expectedResult := randomGetSenderDomainListResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLogin = req.FormValue("username")

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := partners.GetSenderDomainList(req, expectedLogin).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedLogin != givenLogin {
		t.Fatalf(`Login should be "%s", "%s" given`, expectedLogin, givenLogin)
	}

	if !reflect.DeepEqual(expectedResult, givenResult) {
		t.Fatalf("Results should be equal")
	}
}

func randomGetSenderDomainListResult() *partners.GetSenderDomainListResult {
	return &partners.GetSenderDomainListResult{
		Status: test.RandomString(12, 36),
		Domains: []partners.GetSenderDomainListResultDomain{
			{
				Domain: test.RandomString(12, 36),
				Status: test.RandomString(12, 36),
				Key:    test.RandomString(12, 36),
			},
		},
	}
}
