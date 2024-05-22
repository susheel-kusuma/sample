package main

import (
	"encoding/json"
	"io"

	"github.com/beego/beego/v2/server/web"
)

type User struct {
	web.Controller
}

func (u *User) GetUser() {
	u.Data["json"] = map[string]string{
		"Name":   "Ganesh",
		"UserID": "JNFJNS7DF",
	}
	u.ServeJSON()
}

func (u *User) PostUser() {
	// Read the request body
	body, _ := io.ReadAll(u.Ctx.Request.Body)
	defer u.Ctx.Request.Body.Close()
	data := make(map[string]interface{})
	json.Unmarshal(body, &data)
	data["UserID"] = "DSH7DOF6OSDGF"

	// return resp
	u.Data["json"] = data
	u.ServeJSON()
}

func main() {
	web.Router("/user/post", &User{}, "post:PostUser")
	web.Router("/user/get", &User{}, "get:GetUser")
	web.Run()
}
