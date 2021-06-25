package main

import (
	"go_course_thinknetika/14_search_engine/pkg/webapp"
)

func main() {
	webapp.PerformScan([]string{"https://golang.org", "https://go.dev"}, 2)
	webapp.ListenAndServe()
}
