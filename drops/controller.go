package drops

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Viva-con-Agua/echo-pool/auth"
	"github.com/Viva-con-Agua/echo-pool/resp"
	"github.com/labstack/echo"
)




/** 
	* Redirect to drops to get the oauth_code 
	*/
func DropsLogin(c echo.Context) (err error) {
	url, err := RedirectCodeUrl()
	return c.Redirect(http.StatusMovedPermanently, url)
}


/**
	* Redirect_Uri
	*/
func DropsCode(c echo.Context) (err error) {
	
	// get code from url params
	code := c.QueryParam("code")

	// create url for request token
	url, err := RequestTokenUrl(code)

	// request for get token
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}
	if resp.StatusCode == 200 {
		//initial struct for json body
		var token auth.AccessToken
		// read body
		err = json.NewDecoder(resp.Body).Decode(&token)
		if err != nil {
			log.Print(err)
		}
		url, err = RequestUser(token.AccessToken)
		resp, err = http.Get(url)
		if err != nil {
			log.Print(err)
		}
		if resp.StatusCode == 200 {
			body := new(DropsUser)
			//var body interface{}
			err = json.NewDecoder(resp.Body).Decode(&body)
			//body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Print(err)
			}
			response := body.PoolUser()

			LogoutList.Delete(response.Uuid)
			auth.SetSession(c, response, &token)
			return c.JSON(http.StatusOK, response)
		}
	}
	resp.Body.Close()
	return c.JSON(http.StatusUnauthorized, "bla")
}

func TestDrops(c echo.Context) (err error) {
	user, contains := auth.GetUser(c)
	if contains != true {
		return c.JSON(http.StatusInternalServerError, resp.InternelServerError)
	}

	return c.JSON(http.StatusOK, user)
}
