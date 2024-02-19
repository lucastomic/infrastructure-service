package controller

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/lucastomic/infrastructure-/internal/controller/apitypes"
	"github.com/lucastomic/infrastructure-/internal/domain"
	"github.com/lucastomic/infrastructure-/internal/errs"
	"github.com/lucastomic/infrastructure-/internal/infrastructure"
	"github.com/lucastomic/infrastructure-/internal/logging"
)

// InfraController represents a controller that handles infrastructure operations.
// It embeds a CommonController for common functionalities across controllers
// and uses an InfrastructureService for performing operations on infrastructure.
type InfraController struct {
	logger           logging.Logger                       // logger is used for logging messages in the context of infrastructure operations.
	srv              infrastructure.InfrastructureService // srv is the service responsible for performing infrastructure operations.
	CommonController                                      // CommonController provides common functionalities for all controllers.
}

// Router configures and returns the routes associated with the InfraController.
// It defines endpoints for handling specific infrastructure operations.
func (c InfraController) Router() apitypes.Router {
	return []apitypes.Route{
		{
			"/",
			"POST",
			c.handleDeploy,
		},
	}
}

// handleDeploy handles the deployment request for infrastructure.
// It validates and parses the incoming request, performs the deployment
// using the infrastructure service, and returns a response indicating the outcome.
func (c InfraController) handleDeploy(w http.ResponseWriter, r *http.Request) apitypes.Response {
	webData, err := c.validateAndParseRequest(r)
	if err != nil {
		return c.ParseError(r.Context(), r, w, err)
	}
	localDeployment, err := c.srv.LocalDeploy(webData)
	if err != nil {
		return c.ParseError(r.Context(), r, w, err)
	}
	return apitypes.Response{
		Status:  200,
		Content: fmt.Sprintf("Successfully deployed on %s port", localDeployment.Addr),
	}
}

// validateAndParseRequest validates the incoming HTTP request for deployment
// and parses it to extract relevant data needed for the deployment operation.
// It returns an error if validation fails or parsing is unsuccessful.
func (c InfraController) validateAndParseRequest(
	r *http.Request,
) (domain.WebData, error) {
	name := r.FormValue("name")
	if name == "" {
		return domain.WebData{}, fmt.Errorf("%w:%s", errs.ErrInvalidInput, "No name provided")
	}
	zipReader, err := c.getFileFromReq(r)
	if err != nil {
		return domain.WebData{}, fmt.Errorf("%w:%s", errs.ErrInvalidInput, "Error retrieving file")
	}

	return domain.WebData{
		Name:      name,
		FilesPath: zipReader,
	}, nil
}

// getFileFromReq extracts the zip file from the incoming HTTP request
// intended for deployment. It validates the file's presence and attempts
// to read and parse the zip file, returning a zip.Reader for further processing.
// An error is returned if the file is missing, unreadable, or parsing fails.
func (c InfraController) getFileFromReq(r *http.Request) (*zip.Reader, error) {
	zipFile, _, err := r.FormFile("web")
	if err != nil {
		return &zip.Reader{}, errs.ErrinternalError
	}
	if zipFile == nil {
		return &zip.Reader{}, fmt.Errorf("%w:%s", errs.ErrInvalidInput, "No file provided")
	}
	zipContent, err := io.ReadAll(zipFile)
	if err != nil {
		c.logger.Error(r.Context(), "Error trying to read zip file: %s", err.Error())
		return &zip.Reader{}, errs.ErrinternalError
	}

	zipReader, err := zip.NewReader(bytes.NewReader(zipContent), int64(len(zipContent)))
	if err != nil {
		c.logger.Error(r.Context(), "Error trying to create zip reader: %s", err.Error())
		return &zip.Reader{}, errs.ErrinternalError
	}
	return zipReader, nil
}
