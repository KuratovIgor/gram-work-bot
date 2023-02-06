package server

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/repository"
	"net/http"
)

type AuthorizationServer struct {
	server          *http.Server
	tokenRepository repository.TokenRepository
	redirectURL     string
}

func NewAuthorizationServer(tokenRepository repository.TokenRepository, redirectURL string) *AuthorizationServer {
	return &AuthorizationServer{tokenRepository: tokenRepository, redirectURL: redirectURL}
}

func (s *AuthorizationServer) Start() error {
	s.server = &http.Server{
		Addr:    ":80",
		Handler: s,
	}

	return s.server.ListenAndServe()
}

func (s *AuthorizationServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	authorizationCode := r.URL.Query().Get("code")
	if authorizationCode == "" {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
}
