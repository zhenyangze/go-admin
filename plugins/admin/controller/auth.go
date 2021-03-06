package controller

import (
	"github.com/chenhg5/go-admin/modules/auth"
	"fmt"
	"github.com/chenhg5/go-admin/context"
	"bytes"
	"net/http"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/modules/menu"
)

func Auth(ctx *context.Context) {

	password := ctx.Request.FormValue("password")
	username := ctx.Request.FormValue("username")

	if user, ok := auth.Check(password, username); ok {

		auth.SetCookie(ctx, user)

		menu.Unlock()

		ctx.Json(http.StatusOK, map[string]interface{}{
			"code": 200,
			"msg": "登录成功",
			"url": Config.PREFIX + Config.INDEX,
		})
		return
	}
	ctx.Json(http.StatusBadRequest, map[string]interface{}{
		"code": 400,
		"msg": "登录失败",
	})
	return
}

func Logout(ctx *context.Context) {
	auth.DelCookie(ctx)
	ctx.Response.Header.Add("Location", Config.PREFIX + "/login")
	ctx.SetStatusCode(302)
}

func ShowLogin(ctx *context.Context) {

	defer GlobalDeferHandler(ctx)

	tmpl, name := template.GetComp("login").GetTemplate()
	buf := new(bytes.Buffer)
	fmt.Println(tmpl.ExecuteTemplate(buf, name, struct {
		AssertRootUrl string
	}{Config.PREFIX}))
	ctx.WriteString(buf.String())

	ctx.Response.Header.Add("Content-Type", "text/html; charset=utf-8")
}
