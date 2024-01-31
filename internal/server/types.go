package server

import (
	"archive/zip"
)

type LocalDeployRequest struct {
	WebFiles zip.File
	Name     string
}
