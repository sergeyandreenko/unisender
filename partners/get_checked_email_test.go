package partners_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/sergeyandreenko/unisender/api"
	"github.com/sergeyandreenko/unisender/partners"
	"github.com/sergeyandreenko/unisender/test"
)

func TestGetCheckedEmailRequest_Email(t *testing.T) {
	expectedLogin := test.RandomString(12, 36)

	expectedEmail := test.RandomString(12, 36)
	var givenEmail string

	expectedResult := []partners.GetCheckedEmailResponse{
		{
			Email:     test.RandomString(12, 36),
			Name:      test.RandomString(12, 36),
			IsChecked: true,
		},
	}

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenEmail = req.FormValue("email")

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	_, err := partners.GetCheckedEmail(req, expectedLogin).
		Email(expectedEmail).
		Execute()

	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if expectedEmail != givenEmail {
		t.Fatalf(`Email should be "%s", "%s" given`, expectedEmail, givenEmail)
	}
}

func TestGetCheckedEmailRequest_Execute(t *testing.T) {
	expectedLogin := test.RandomString(12, 36)
	var givenLogin string

	expectedResult := []partners.GetCheckedEmailResponse{
		{
			Email:     test.RandomString(12, 36),
			Name:      test.RandomString(12, 36),
			IsChecked: true,
		},
	}

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		givenLogin = req.FormValue("login")

		result := api.Response{
			Result: expectedResult,
		}

		response, _ := json.Marshal(result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenResult, err := partners.GetCheckedEmail(req, expectedLogin).
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
