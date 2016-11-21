package controllers

import (
	"encoding/json"
	"fmt"
	. "github.com/makalkin/goquest/app/api/v1/app/services"
	"github.com/makalkin/goquest/app/models"
	"github.com/revel/revel"
	"golang.org/x/oauth2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/url"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Quests() revel.Result {
	return c.Render()
}

func (c App) Auth(code string) revel.Result {
	tok, err := models.FACEBOOK.Exchange(oauth2.NoContext, code)
	if err != nil {
		revel.ERROR.Println(err)
		return c.Redirect(App.Index)
	}
	println(tok.AccessToken)
	println(tok.Expiry.String())
	println(tok.RefreshToken)
	println(tok.TokenType)
	// Not we have to retrieve user
	userData := map[string]interface{}{}
	resp, err := http.Get("https://graph.facebook.com/me?fields=id,name&access_token=" +
		url.QueryEscape(tok.AccessToken))
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		revel.ERROR.Println(err)
	}

	service := UserService{}
	user := &models.User{}

	if err := service.GetUser(bson.M{"fid": userData["id"]}, user); err == nil {
		revel.INFO.Println("FOUND USER", user)
		user.AccessToken = tok.AccessToken
	} else {
		user = &models.User{
			AccessToken: tok.AccessToken,
			Name:        userData["name"].(string),
			Fid:         userData["id"].(string),
		}
	}
	revel.INFO.Println("ERRR", err)
	if err := service.AddUser(user); err != nil {
		revel.ERROR.Println(err)
		return c.Redirect(App.Index)
	} else {
		c.RenderArgs["user"] = user
	}

	c.Session["uid"] = fmt.Sprintf("%v", user.Fid)

	return c.Redirect(App.Index)
}

func (c App) Logout() revel.Result {
	delete(c.Session, "uid")
	return c.Redirect(App.Index)
}

func setuser(c *revel.Controller) revel.Result {
	userService := UserService{}

	if _, ok := c.Session["uid"]; ok {
		uid, _ := c.Session["uid"]
		user := &models.User{}
		if err := userService.GetUser(bson.M{"fid": uid}, user); err == nil {
			c.RenderArgs["user"] = user
		}
	}
	c.RenderArgs["authUrl"] = models.FACEBOOK.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return nil
}

func init() {
	revel.InterceptFunc(setuser, revel.BEFORE, &App{})
}
