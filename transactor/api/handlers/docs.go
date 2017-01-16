package handlers

import (
	"net/http"
	"io"
)

type DocsHandler struct {}

func NewDocsHandler() *DocsHandler {
	return &DocsHandler{}
}

func (handler *DocsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Here be documentation")
}
