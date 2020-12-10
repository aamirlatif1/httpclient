package gohttp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	//Given
	client := httpClient{
		builder: &clientBuilder{},
	}
	ch := make(http.Header)
	ch.Set("Content-Type", "application/json")
	ch.Set("User-Agent", "cool-http-client")
	client.builder.headers = ch

	//When
	rh := make(http.Header)
	rh.Set("X-Request-Id", "ABC-123")
	finalHeaders := client.getRequestHeaders(rh)

	//Then
	assert.Equal(t, 3, len(finalHeaders))
	assert.Contains(t, finalHeaders["Content-Type"], "application/json")
	assert.Contains(t, finalHeaders["User-Agent"], "cool-http-client")
	assert.Contains(t, finalHeaders["X-Request-Id"], "ABC-123")
}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{
		builder: &clientBuilder{},
	}
	requestBody := []string{"one", "two"}

	t.Run("noBodyNilResponse", func(t *testing.T) {
		//When
		body, err := client.getRequestBody("", nil)

		//Then
		assert.Nil(t, err)
		assert.Nil(t, body)
	})

	t.Run("BodyWithJson", func(t *testing.T) {
		//When
		body, err := client.getRequestBody("application/json", requestBody)
		fmt.Println(string(body))
		//Then
		assert.Nil(t, err)
		assert.Equal(t, `["one","two"]`, string(body))
	})

	t.Run("BodyWithXML", func(t *testing.T) {
		//When
		body, err := client.getRequestBody("application/xml", requestBody)
		fmt.Println(string(body))
		//Then
		assert.Nil(t, err)
		assert.Equal(t, `<string>one</string><string>two</string>`, string(body))
	})

	t.Run("BodyWithDefaultJson", func(t *testing.T) {
		//When
		body, err := client.getRequestBody("", requestBody)

		//Then
		assert.Nil(t, err)
		assert.Equal(t, `["one","two"]`, string(body))
	})


}