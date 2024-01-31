package types

import (
	"context"
	"net/http"
)

type LocalDeployment struct {
	Addr string
}

type WebData struct {
	Name      string
	FilesPath string
}
type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error
