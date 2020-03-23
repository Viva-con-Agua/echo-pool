/** Implements the drops oauth request for user backend authentication
	*	
	*
	*/



package pool

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"

)


type (
	City struct {
		Country string `json:"country"`
		Name string `json:"name"` 
	}
	Crew struct {
		Id string `json:"id"`
		Name string `json:"name"`
		Cities []City `json:"cities"`
	}
	Pillar struct {
		Pillar string `json:"pillar"`
	}
	Role struct {
		Crew Crew `json:"crew"`
		Name string `json:"name"`
		Pillar Pillar `json:"pillar"`
	}
	Role2 struct {
		Role string `json:"role"`
	}
	Supporter struct {
		FullName string `json:"fullName"`
		Roles []Role `json:"roles"`
	}
	Profile struct {
		Email string `json:"email"`
		Supporter Supporter `json:"supporter"`
	}
	Address struct {
		Additional string `json:"additional"`
		City string `json:"city"`
		Country string `json:"country"`
		PublicId string `json:"publicId"`
		Street string `json:"street"`
		Zip string `json:"zip"`
	}
	DropsUser struct {
		Id	string `json:"id"`
		Profiles []Profile `json:"profiles"`
		Roles	[]Role2 `json:"roles"`
	}
	PoolRole struct {
		Name string `json:"name"`
		CrewId string `json:"crewId"`
		CrewName string `json:"crewName"`
		Pillar string `json:"pillar"`
	}

	SessionUser struct {
		PoolUser PoolUser
		AccessToken string
	}
	Permission struct {
		Role string
		Pillar string
	}
	AccessToken struct {
		AccessToken string `json:"access_token"`
	}
	StringList []string
)

var LogoutList StringList

func (l StringList) Add(uuid string) StringList {
	return append(l, uuid)
}

func (l StringList) Delete(uuid string) StringList {
	for i, v := range l {
		if v == uuid {
			l[i] = l[len(l)-1]
			l[len(l)-1] = ""
			return l[:len(l)-1]
		}
	}
	return l
}

func (l StringList) Contains(uuid string) bool {
	for _, v := range l {
		if v == uuid {
			return true
		}
	}
	return false
}


func NatsLogout() {
	Nats.Subscribe("LOGOUT", func(s string){
		LogoutList = LogoutList.Add(s)
	})
}


func (u* DropsUser) PoolUser() *PoolUser {
	user := new(PoolUser)
	user.Uuid = u.Id
	for _, p := range u.Profiles {
		user.Email = p.Email
		user.Name = p.Supporter.FullName
		var roles []PoolRole
		for _, r := range u.Roles {
			role := new(PoolRole)
			role.Name = r.Role
			roles = append(roles, *role)
		}
		var ddRoles = p.Supporter.Roles
		for _, r := range ddRoles {
			role := new(PoolRole)
			role.Name = r.Name
			role.Pillar = r.Pillar.Pillar
			role.CrewId = r.Crew.Id
			role.CrewName = r.Name
			roles = append(roles, *role)
		}
		user.Roles = roles
	}
	return user
} 

/**
	* Generate the url for get authorization_code. 
	*/
func RedirectCodeUrl () (uri string, err error) {
	// url for get oauth code from config.yml
	u, err := url.Parse(Config.Drops.Url.Code)
	// seperate the query
	q, err := url.ParseQuery(u.RawQuery)
	
	// add params to query
	q.Add("client_id", Config.Drops.Oauth.ClientId)
	q.Add("ajax", "true")
	q.Add("response_type", "code")
	q.Add("state", "")
	q.Add("redirect_uri", Config.Drops.Url.Redirect)
	
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
	u, err := url.Parse(Config.Drops.Url.Token)
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
	q.Add("client_id", Config.Drops.Oauth.ClientId)
	q.Add("ajax", "true")
	q.Add("code", code)
	q.Add("state", "")
	q.Add("redirect_uri", Config.Drops.Url.Redirect)
	q.Add("grant_type", "authorization_code")

	u.RawQuery = q.Encode()
	return u.String() , err
}

func RequestUser(token string) (uri string, err error) {
	u, err := url.Parse(Config.Drops.Url.User)
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
		var token AccessToken
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
			SetSession(c, response, &token)
			return c.JSON(http.StatusOK, response)
		}
	}
	resp.Body.Close()
	return c.JSON(http.StatusUnauthorized, "bla")
}

func TestDrops(c echo.Context) (err error) {
	sess, _ := session.Get("session", c)
	sess.Save(c.Request(), c.Response())
	val := sess.Values["user"]
	var user = &PoolUser{}
	user, _ = val.(*PoolUser)
	return c.JSON(http.StatusOK, user)
}
