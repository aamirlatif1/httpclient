package examples

import (
	"errors"
	"github.com/aamir/httpclient/gohttp"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	gohttp.StartMockServer()
	os.Exit(m.Run())
}

func TestGetEndpoints(t *testing.T) {

	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		//Given
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		})

		//When
		endpoints, err := GetEndpoints()

		//Then
		assert.Nil(t, endpoints)
		assert.NotNil(t, err)
		assert.Equal(t, "timeout getting github endpoints", err.Error())
	})
	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		//Given
		//When
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseBody:       `{"current_user_url":123}`,
			ResponseStatusCode: http.StatusOK,
		})

		//Then
		endpoints, err := GetEndpoints()
		assert.Nil(t, endpoints)
		assert.NotNil(t, err)
		assert.Equal(t, "json: cannot unmarshal number into Go struct field Endpoints.current_user_url of type string", err.Error())
	})
	t.Run("TestSuccess", func(t *testing.T) {
		//Given
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseBody:       `{"current_user_url":"https://api.github.com/user"}`,
			ResponseStatusCode: http.StatusOK,
		})

		//When
		endpoints, err := GetEndpoints()

		//Then
		assert.Nil(t, err)
		assert.NotNil(t, endpoints)
		assert.Equal(t, "https://api.github.com/user", endpoints.CurrentUserUrl)
	})

}
