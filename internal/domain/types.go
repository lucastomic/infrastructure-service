package domain

import (
	"archive/zip"
	"context"
	"net/http"
)

type LocalDeployment struct {
	Addr string
}

type WebData struct {
	Name      string
	FilesPath *zip.Reader
}
type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error
