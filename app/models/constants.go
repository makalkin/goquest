package models

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"github.com/revel/revel"
	"github.com/maxwellhealth/bongo"
	"fmt"
	"log"
)

var (
	redirectUrl string = "https://goquest.herokuapp.com/App/Auth"
	FACEBOOK *oauth2.Config
)

func InitConstants() {
	FACEBOOK = &oauth2.Config{
		ClientID:	"1179783535435590",
		ClientSecret:	"767a8dac05cfa56cf4043be2b075644c",
		Scopes: 	[]string{},
		Endpoint:	facebook.Endpoint,
		RedirectURL:	revel.Config.StringDefault("facebook.RedirectUrl", redirectUrl),
	}
}

var (
	DB *bongo.Connection
)

func InitDB()  {
	config := &bongo.Config{
		ConnectionString: revel.Config.StringDefault("db.spec",""),
		Database:         revel.Config.StringDefault("", "goquest_db"),
	}
	var err error
	fmt.Print(revel.Config.StringDefault("db",""), revel.Config.StringDefault("", "goquest_db"))
	DB, err = bongo.Connect(config)

	if err != nil {
		log.Fatal(err)
	}
}