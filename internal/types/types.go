package types

import "mime/multipart"

type LocalDeployment struct {
	Addr string
}

type WebData struct {
	Name  string
	Files []multipart.File
}
