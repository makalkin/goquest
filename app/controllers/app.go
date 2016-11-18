package controllers

import (
	"github.com/revel/revel"
	"golang.org/x/oauth2"
	"fmt"
	"github.com/makalkin/goquest/app/models"
	"net/http"
	"net/url"
	"encoding/json"
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

	// Not we have to retrieve user
	userData := map[string]interface{}{}
	resp, err := http.Get("https://graph.facebook.com/me?fields=id,name&access_token=" +
		url.QueryEscape(tok.AccessToken))
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		revel.ERROR.Println(err)
	}

	service := models.UserService{}
	user := &models.User{}

	if err := service.GetUser(map[string]interface{}{"fid": userData["id"]}, user); err == nil {
		user.AccessToken = tok.AccessToken
	} else {
		user = &models.User{
			AccessToken: tok.AccessToken,
			Name: userData["name"].(string),
			Fid: userData["id"].(string),
		}
	}

	if err := service.AddUser(user); err != nil {
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
	userService := models.UserService{}

	if _, ok := c.Session["uid"]; ok {
		uid, _ := c.Session["uid"]
		user := &models.User{}
		if err := userService.GetUser(map[string]interface{}{"fid": uid}, user); err == nil {
			c.RenderArgs["user"] = user
		}
	}
	c.RenderArgs["authUrl"] = models.FACEBOOK.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return nil
}

func init()  {
	revel.InterceptFunc(setuser, revel.BEFORE, &App{})
}