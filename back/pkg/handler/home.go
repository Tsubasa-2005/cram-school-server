package handler

import (
	"cram-school-reserve-server/front/pkg"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter) {
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

	er := pkg.Page("home").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
	return
}
