package examples

import (
	"errors"
	"github.com/aamirlatif1/httpclient/gohttp"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCrateRepo(t *testing.T) {
	gohttp.StartMockServer()
	t.Run("timeoutFromGithub", func(t *testing.T) {
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Url:         "https://api.github.com/user/repos",
			Method:      http.MethodPost,
			RequestBody: `{"name":"test-repo","private":true}`,
			Error:       errors.New("timeout from github"),
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		repo, err := CreateRepo(repository)

		assert.Nil(t, repo)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "timeout from github")
	})

	t.Run("noError", func(t *testing.T) {
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Url:         "https://api.github.com/user/repos",
			Method:      http.MethodPost,
			RequestBody: `{"name":"test-repo","private":true}`,
			ResponseStatusCode: http.StatusCreated,
			ResponseBody: `{"id":123,"name":"test-repo"}`,
		})

		repository := Repository{
			Name:    "test-repo",
			Private: true,
		}

		repo, err := CreateRepo(repository)

		assert.Nil(t, err)
		assert.NotNil(t, repo)
		assert.Equal(t, repository.Name, repo.Name)
	})

}
