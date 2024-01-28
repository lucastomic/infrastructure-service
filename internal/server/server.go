package server

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
)

type Server struct {
	listenAddr  string
	service     webprocessor.WebProcessor
	inputParser input.Reader
	logging     logging.Logger
}

func New(
	listenAddr string,
	service webprocessor.WebProcessor,
	inputParser input.Reader,
	logging logging.Logger,
) Server {
	return Server{
		listenAddr,
		service,
		inputParser,
		logging,
	}
}

func (s *Server) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleGenerationReq))
	http.ListenAndServe(s.listenAddr, nil)
}

func makeHTTPHandlerFunc(apiFn types.APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(10000000))
	return func(w http.ResponseWriter, r *http.Request) {
		apiFn(ctx, w, r)
	}
}

func (s Server) handleGenerationReq(
	ctx context.Context,
	writer http.ResponseWriter,
	req *http.Request,
) error {
	pageData, err := s.inputParser.RetrieveInput(*req)
	if err != nil {
		writeJSON(writer, http.StatusBadRequest, map[string]any{"error": err.Error()})
	}
	err = s.service.Process(ctx, pageData)
	if err != nil {
		writeJSON(writer, http.StatusBadRequest, map[string]any{"error": err.Error()})
	}
	writeJSON(writer, 200, map[string]any{"message": "Web generated successfully"})
	return nil
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
