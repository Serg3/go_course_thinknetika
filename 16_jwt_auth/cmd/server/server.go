package main

import (
	"go_course_thinknetika/16_jwt_auth/pkg/api"
	"net/http"
)

func main() {
	a := api.New()
	http.ListenAndServe(":8000", a.Router())
}
