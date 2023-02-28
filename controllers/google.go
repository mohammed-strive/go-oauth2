package controllers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mohammed-strive/go-oauth2/config"
	"google.golang.org/appengine/log"
)

const GOOGLE_USERINFO = "https://www.gooleapis.com/oauth2/v2/userinfo"

func GoogleLogin(ctx *fiber.Ctx) error {
	url := config.AppConfig.GoogleLoginConfig.AuthCodeURL("randomstate")
	ctx.Status(fiber.StatusSeeOther)
	ctx.Redirect(url)
	return ctx.JSON(url)
}
func GoogleCallback(ctx *fiber.Ctx) error {
	state := ctx.Query("state")
	if state != "randomstate" {
		return ctx.SendString("states do not match!!")
	}

	code := ctx.Query("code")
	googleConfig := config.GoogleConfig()

	token, err := googleConfig.Exchange(context.Background(), code)
	if err != nil {
		return ctx.SendString("auth code exchange failed")
	}

	userinfoRequest, err := http.NewRequest("GET", GOOGLE_USERINFO, nil)
	if err != nil {
		return ctx.SendString("unable to formulate userinfo request")
	}

	userinfoRequest.Header = http.Header{
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}

	res, err := http.DefaultClient.Do(userinfoRequest)
	if err != nil {
		log.Errorf(ctx.Context(), "unable to fetch userinfo: %v", err)
		ctx.SendString("unable to access userinfo")
	}

	defer res.Body.Close()

	userData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf(ctx.Context(), "unable to read user data: %v", err)
		ctx.SendString("json parsing failed")
	}

	return ctx.SendString(string(userData))
}
