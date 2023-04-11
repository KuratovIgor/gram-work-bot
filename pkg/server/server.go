package server

import (
	"github.com/KuratovIgor/gram-work-bot/pkg/repository"
	headhunter "github.com/KuratovIgor/head_hunter_sdk"
	"net/http"
	"strconv"
)

type AuthorizationServer struct {
	server            *http.Server
	graphqlRepository repository.GraphqlRepository
	redirectURL       string
	client            *headhunter.Client
}

func NewAuthorizationServer(graphqlRepository repository.GraphqlRepository, redirectURL string, client *headhunter.Client) *AuthorizationServer {
	return &AuthorizationServer{graphqlRepository: graphqlRepository, redirectURL: redirectURL, client: client}
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

	userId, creationErr := s.createUser(chatID, response.AccessToken)
	if creationErr != nil {
		return creationErr
	}

	sessionErr := s.graphqlRepository.CreateSession(chatID, response.AccessToken, response.RefreshToken, userId)
	if sessionErr != nil {
		return sessionErr
	}

	filterError := s.graphqlRepository.CreateDefaultFilters(chatID)
	if filterError != nil {
		return filterError
	}

	return nil
}

func (s *AuthorizationServer) createUser(chatID int64, token string) (string, error) {
	response, err := s.client.GetInfoAboutMe(token)
	if err != nil {
		return "", err
	}

	err = s.graphqlRepository.CreateUser(chatID, response.Name, response.LastName, response.MiddleName, response.Email, response.Phone, response.UserID)
	if err != nil {
		return "", err
	}

	return response.UserID, nil
}
