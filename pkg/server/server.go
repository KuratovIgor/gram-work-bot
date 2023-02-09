package server

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/repository"
	headhunter "github.com/KuratovIgor/head_hunter_sdk"
	"net/http"
	"strconv"
)

type AuthorizationServer struct {
	server          *http.Server
	tokenRepository repository.TokenRepository
	redirectURL     string
	client          *headhunter.Client
}

func NewAuthorizationServer(tokenRepository repository.TokenRepository, redirectURL string, client *headhunter.Client) *AuthorizationServer {
	return &AuthorizationServer{tokenRepository: tokenRepository, redirectURL: redirectURL, client: client}
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
	chatID := r.URL.Query().Get("chat_id")
	if authorizationCode == "" || chatID == "" {
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	chatIdInt, _ := strconv.ParseInt(chatID, 10, 64)
	err := s.Authorize(chatIdInt, authorizationCode)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		w.Header().Add("Location", s.redirectURL)
		w.WriteHeader(http.StatusMovedPermanently)
	}
}

func (s *AuthorizationServer) Authorize(chatID int64, authCode string) error {
	response, err := s.client.Authorize(chatID, authCode)
	if err != nil {
		return err
	}

	s.tokenRepository.Save(chatID, response.AccessToken, repository.AccessTokens)
	s.tokenRepository.Save(chatID, response.RefreshToken, repository.RefreshToken)

	return nil
}
