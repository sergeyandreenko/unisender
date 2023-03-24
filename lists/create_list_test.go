package lists_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/sergeyandreenko/unisender/api"
	"github.com/sergeyandreenko/unisender/lists"
	"github.com/sergeyandreenko/unisender/test"
)

func TestCreateListRequest_BeforeSubscribeUrl(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)
	expectedTitle := test.RandomString(12, 36)

	expectedBeforeSubscribeUrl := test.RandomString(12, 36)
	var givenBeforeSubscribeUrl string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenBeforeSubscribeUrl = req.FormValue("before_subscribe_url")

		result := api.Response{
			Result: &lists.CreateListResult{
				ID: expectedListID,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := lists.CreateList(req, expectedTitle).
		BeforeSubscribeUrl(expectedBeforeSubscribeUrl).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if givenBeforeSubscribeUrl != expectedBeforeSubscribeUrl {
		t.Errorf(`Param "before_subscribe_url" should be "%s", "%s" given`, expectedBeforeSubscribeUrl, givenBeforeSubscribeUrl)
	}
}

func TestCreateListRequest_AfterSubscribeUrl(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)
	expectedTitle := test.RandomString(12, 36)

	expectedAfterSubscribeUrl := test.RandomString(12, 36)
	var givenAfterSubscribeUrl string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenAfterSubscribeUrl = req.FormValue("after_subscribe_url")

		result := api.Response{
			Result: &lists.CreateListResult{
				ID: expectedListID,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := lists.CreateList(req, expectedTitle).
		AfterSubscribeUrl(expectedAfterSubscribeUrl).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if givenAfterSubscribeUrl != expectedAfterSubscribeUrl {
		t.Errorf(`Param "after_subscribe_url" should be "%s", "%s" given`, expectedAfterSubscribeUrl, givenAfterSubscribeUrl)
	}
}

func TestCreateListRequest_Execute(t *testing.T) {
	expectedListID := test.RandomInt64(9999, 999999)

	expectedTitle := test.RandomString(12, 36)
	var givenTitle string

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenTitle = req.FormValue("title")

		result := api.Response{
			Result: &lists.CreateListResult{
				ID: expectedListID,
			},
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenListID, err := lists.CreateList(req, expectedTitle).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedListID != givenListID {
		t.Fatalf("List ID should be %d, %d given", expectedListID, givenListID)
	}

	if expectedTitle != givenTitle {
		t.Fatalf(`Title should be "%s", "%s" given`, expectedTitle, givenTitle)
	}
}
