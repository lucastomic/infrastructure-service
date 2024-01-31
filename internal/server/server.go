package server

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/lucastomic/infrastructure-/internal/infrastructure"
	"github.com/lucastomic/infrastructure-/internal/logging"
	"github.com/lucastomic/infrastructure-/internal/types"
	"github.com/lucastomic/infrastructure-/internal/unzipper"
)

type Server struct {
	listenAddr string
	service    infrastructure.InfrastructureService
	logging    logging.Logger
}

func New(
	listenAddr string,
	service infrastructure.InfrastructureService,
	logging logging.Logger,
) Server {
	return Server{
		listenAddr,
		service,
		logging,
	}
}

func (s *Server) Run() {
	http.HandleFunc("/", s.makeHTTPHandlerFunc(s.handleGenerationReq))
	http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) makeHTTPHandlerFunc(apiFn types.APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(10000000))
	return func(w http.ResponseWriter, r *http.Request) {
		err := apiFn(ctx, w, r)
		if err != nil {
			s.logging.Request(ctx, r, http.StatusBadRequest)
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func (s Server) handleGenerationReq(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) error {
	req.ParseMultipartForm(10 << 20) // 10 MB
	name := req.FormValue("name")
	file, header, err := req.FormFile("webFiles")
	if err != nil {
		return err
	}
	defer file.Close()
	err = unzipper.Unzip(file, header.Size, "../../tmp/")
	if err != nil {
		return err
	}
	deployment := s.service.LocalDeploy(types.WebData{
		Name:      name,
		FilesPath: "../../tmp/",
	})
	if err != nil {
		return err
	}
	writeJSON(
		writer,
		http.StatusBadRequest,
		map[string]any{"message": deployment.Addr},
	)
	return nil
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
