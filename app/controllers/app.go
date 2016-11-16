package controllers

import (
	"github.com/revel/revel"
	"golang.org/x/oauth2"
	"strconv"
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

	user := c.connected()
	user.AccessToken = tok.AccessToken
	return c.Redirect(App.Index)
}

func (c App) Logout() revel.Result {
	u := c.connected()
	me := map[string]interface{}{}
	if u != nil{
		u.AccessToken = ""
	}
	return c.Redirect(App.Index, me)
}

func setuser(c *revel.Controller) revel.Result {
	var user *models.User
	if _, ok := c.Session["uid"]; ok {
		uid, _ := strconv.ParseInt(c.Session["uid"], 10, 0)
		user = models.GetUser(int(uid))
	}
	if user == nil {
		user = models.NewUser()
		c.Session["uid"] = fmt.Sprintf("%d", user.Uid)
	}
	c.RenderArgs["user"] = user
	return nil
}

func getuser(c *revel.Controller) revel.Result  {
	u := connected(c)
	me := map[string]interface{}{}
	if u != nil && u.AccessToken != "" {
		resp, _ := http.Get("https://graph.facebook.com/me?access_token=" +
			url.QueryEscape(u.AccessToken))
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&me); err != nil {
			revel.ERROR.Println(err)
		}
		revel.INFO.Println(me)
	}
	authUrl := models.FACEBOOK.AuthCodeURL("state", oauth2.AccessTypeOffline)
	c.RenderArgs["me"] = me
	c.RenderArgs["authUrl"] = authUrl
	return nil
}

func init()  {
	revel.InterceptFunc(setuser, revel.BEFORE, &App{})
	revel.InterceptFunc(getuser, revel.AFTER, &App{})
}

func connected(c *revel.Controller) *models.User {
	return c.RenderArgs["user"].(*models.User)
}

func(c App) connected() *models.User {
	return c.RenderArgs["user"].(*models.User)
}