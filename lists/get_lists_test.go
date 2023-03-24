package lists_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/sergeyandreenko/unisender/api"
	"github.com/sergeyandreenko/unisender/lists"
	"github.com/sergeyandreenko/unisender/test"
)

func TestGetListsRequest_Execute(t *testing.T) {
	expectedLists := randomGetListsResult()

	req := test.NewRequest(func(req *http.Request) (res *http.Response, err error) {
		result := api.Response{
			Result: expectedLists,
		}

		response, _ := json.Marshal(&result)

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewBuffer(response)),
		}, nil
	})

	givenLists, err := lists.GetLists(req).Execute()
	if err != nil {
		t.Fatalf(`Error should be nil, "%s" given`, err.Error())
	}

	if len(givenLists) != len(expectedLists) {
		t.Fatalf(`Lists slice should have length %d, %d given`, len(expectedLists), len(givenLists))
	}

	if !reflect.DeepEqual(expectedLists, givenLists) {
		t.Fatal("Expected and given lists should be equal")
	}
}

func randomGetListsResult() (slice []lists.GetListsResult) {
	l := test.RandomInt(12, 36)
	for i := 0; i < l; i++ {
		slice = append(slice, lists.GetListsResult{
			ID:    test.RandomInt64(9999, 999999),
			Title: test.RandomString(12, 36),
		})
	}

	return
}
