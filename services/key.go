package services

import (
	"strconv"

	"fmt"

	"github.com/dinever/golf"
)

func MainHandler(ctx *golf.Context) {
	name, err := ctx.Session.Get("name")
	ctx.SetHeader("Content-Type", "text/html;charset=UTF-8")
	if err != nil {
		ctx.Send("Hello World! Please <a href=\"/login\">log in</a>. Current sessions: " + strconv.Itoa(ctx.App.SessionManager.Count()))
	} else {
		ctx.Send("Hello " + name.(string) + ". Current sessions: " + strconv.Itoa(ctx.App.SessionManager.Count()))
	}
}

func LoginHandler(ctx *golf.Context) {
	ctx.Loader("default").Render("login.html", make(map[string]interface{}))
}

func LoginHandlerPost(ctx *golf.Context) {
	ctx.Session.Set("name", ctx.Request.FormValue("name"))
	ctx.Send("Hi, " + ctx.Request.FormValue("name"))
}

func GoMockHandler(ctx *golf.Context) {
	ctx.Loader("default").Render("go-mock.html", make(map[string]interface{}))
}

func CreateMockResponse(ctx *golf.Context) {
	ctx.Session.Set("method", ctx.Request.FormValue("method"))
	ctx.Session.Set("key", ctx.Request.FormValue("key"))
	ctx.Session.Set("value", ctx.Request.FormValue("value"))

	err := SetResponse(ctx.Request.FormValue("method"), ctx.Request.FormValue("key"), ctx.Request.FormValue("value"))

	if err != nil {
		ctx.Send(fmt.Sprintln("Error Setting key"))
	} else {
		ctx.Send(fmt.Sprintln("Key: %s, created.", ctx.Request.FormValue("key")))
	}
}

type MockResponse struct {
	Method   string `json:"method"`
	Request  string `json:"request"`
	Response string `json:"response"`
}
