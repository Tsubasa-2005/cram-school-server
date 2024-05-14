package handler

import (
	"cram-school-reserve-server/front/pkg"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	item := struct {
		Title   string
		Message string
		Account string
		ID      string
	}{
		Title:   " ",
		Message: "下のLoginからログインしてください。アカウントがない方は、Signupから登録してください。",
		Account: " ",
		ID:      "login",
	}

	er := pkg.Page("home").Execute(c.Writer, item)
	if er != nil {
		return
	}
}
