package auth

import (
	"errors"

	"golang.org/x/oauth2"
)

type GoogleUser struct {
	Sub   string `json:"sub"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (p *SecurityPrincipal) updateFromGoogle() error {

	// retrieve token
	token, err := p.GetToken()
	if err != nil {
		return err
	}

	if token == nil {
		return errors.New("Google OAuth token is not set")
	}

	// Google client
	oauthConfig := GetOauthConfig(p.AuthProvider, p.Groups != nil)
	client := oauthConfig.Client(oauth2.NoContext, token)

	// get user details
	googleUser := &GoogleUser{}
	err = GetObject(client, "https://www.googleapis.com/oauth2/v3/userinfo", googleUser)
	if err != nil {
		return err
	}

	p.ID = googleUser.Sub
	p.Login = googleUser.Email
	p.Name = googleUser.Name
	p.Email = googleUser.Email

	// groups are not supported for Google yet
	p.Groups = nil

	return nil
}
