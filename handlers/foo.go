package handlers

import (
	"my-blog-goth/views/hello"
	"net/http"
)

func HandleFoo(w http.ResponseWriter, r *http.Request) {
	component := hello.Hello("Hiang")
	component.Render(r.Context(), w)
}
