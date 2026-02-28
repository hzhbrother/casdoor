package handler

import (
	"net/http"
	"github.com/casdoor/casdoor/routers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	routers.BeegoApp.ServeHTTP(w, r)
}
