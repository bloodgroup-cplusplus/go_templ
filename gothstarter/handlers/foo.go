package handlers

import (
	"fmt"
	"net/http"
	"github.com/bloodgroup-cplusplus/go_templ/gothstarter/views/foo"
)

func HandleFoo(w http.ResponseWriter, r *http.Request) error {
	return Render(w,r,foo.Index())
}
