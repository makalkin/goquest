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

func (c *App) Auth(code string) revel.Result {
	tok, err := models.FACEBOOK.Exchange(oauth2.NoContext, code)

	if err != nil {
		revel.ERROR.Println(err)
		return c.Redirect(App.Index)
	}

	// Not we have to retrieve user
	userData := map[string]interface{}{}
	resp, err := http.Get("https://graph.facebook.com/me?fields=id,name&access_token=" +
		url.QueryEscape(tok.AccessToken))
	if err == nil {
		defer resp.Body.Close()
	}

	if err := json.NewDecoder(resp.Body).Decode(&userData); err != nil {
		revel.ERROR.Println(err)
	}

	service := UserService{}
	user := c.connected()


	if !user.IsNew() {
		if err := service.GetUser(bson.M{"fid": userData["id"]}, user); err == nil {
			revel.INFO.Println("FOUND USER", user)
		}
	} else {
		user.Name = userData["name"].(string)
		user.Fid = userData["id"].(string)
	}
	user.AccessToken = tok.AccessToken

	if err := service.AddUser(user); err != nil {
		revel.ERROR.Println(err)
		return c.Redirect(App.Index)
	} else {
		c.RenderArgs["user"] = user
	}

	c.Session["uid"] = fmt.Sprintf("%v", user.Fid)

	return c.Redirect(c.Request.Referer())
}

func (c *App) Logout() revel.Result {
	delete(c.Session, "uid")
	delete(c.RenderArgs, "user")
	return c.Redirect(c.Request.Referer())
}

func setuser(c *revel.Controller) revel.Result {
	userService := UserService{}
	var user *models.User

	if _, ok := c.Session["uid"]; ok {
		uid, _ := c.Session["uid"]
		user = &models.User{}
		if err := userService.GetUser(bson.M{"fid": uid}, user); err == nil {
			c.RenderArgs["user"] = user
		}
	}

	if user == nil {
		user = &models.User{}
		user.SetIsNew(true)
	}
	c.RenderArgs["user"] = user
	c.RenderArgs["authUrl"] = models.FACEBOOK.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return nil
}

func init() {
	revel.InterceptFunc(setuser, revel.BEFORE, &App{})
}

func (c *App) connected() *models.User {
	return c.RenderArgs["user"].(*models.User)
}
