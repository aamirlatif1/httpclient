package examples

type Endpoints struct {
	CurrentUserUrl                  string `json:"current_user_url"`
	CurrentUserAuthorizationHtmlUrl string `json:"current_user_authorizations_html_url"`
	AuthorizationsUrl               string `json:"authorizations_url"`
}

func GetEndpoints() (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com")
	if err != nil {
		return nil, err
	}

	var endpoints Endpoints

	if err := response.UnmarshalJson(&endpoints); err != nil {
		return nil, err
	}
	return &endpoints, nil

}
