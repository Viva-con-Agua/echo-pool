package drops

import (
	"log"
	"net/url"
	"github.com/Viva-con-Agua/echo-pool/config"
)

/**
	* Generate the url for get authorization_code. 
	*/
func RedirectCodeUrl () (uri string, err error) {
	// url for get oauth code from config.yml
	u, err := url.Parse(config.Config.Drops.Url.Code)
	// seperate the query
	q, err := url.ParseQuery(u.RawQuery)
	
	// add params to query
	q.Add("client_id", config.Config.Drops.Oauth.ClientId)
	q.Add("ajax", "true")
	q.Add("response_type", "code")
	q.Add("state", "")
	q.Add("redirect_uri", config.Config.Drops.Url.Redirect)
	
	// handle errors
	if err != nil {
		log.Print(err)
		return u.String(), err
	}
	// add query to url
	u.RawQuery = q.Encode()
	// return url
	return u.String() , err
}

func RequestTokenUrl (code string) (uri string, err error) {
	// url for get oauth code
	u, err := url.Parse(config.Config.Drops.Url.Token)
		if err != nil {
		log.Print(err)
		return u.String(), err
	}

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		log.Print(err)
		return u.String(), err
	}

	//add client_id
	q.Add("client_id", config.Config.Drops.Oauth.ClientId)
	q.Add("ajax", "true")
	q.Add("code", code)
	q.Add("state", "")
	q.Add("redirect_uri", config.Config.Drops.Url.Redirect)
	q.Add("grant_type", "authorization_code")

	u.RawQuery = q.Encode()
	return u.String() , err
}

func RequestUser(token string) (uri string, err error) {
	u, err := url.Parse(config.Config.Drops.Url.User)
	if err != nil {
		log.Print(err)
		return u.String(), err
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		log.Print(err)
		return u.String(), err
	}
	q.Add("access_token", token)
	u.RawQuery = q.Encode()
	return u.String(), err
}

