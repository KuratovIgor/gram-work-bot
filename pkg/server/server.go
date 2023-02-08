package server

import (
	"bytes"
	"github.com/KuratovIgor/gram-work-bot/pkg/config"
	"github.com/KuratovIgor/gram-work-bot/pkg/repository"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strconv"
)

type AuthorizationServer struct {
	server          *http.Server
	tokenRepository repository.TokenRepository
	redirectURL     string
	config          *config.Config
}

func NewAuthorizationServer(tokenRepository repository.TokenRepository, redirectURL string, config *config.Config) *AuthorizationServer {
	return &AuthorizationServer{tokenRepository: tokenRepository, redirectURL: redirectURL, config: config}
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

	s.Authorize(s.config, chatID, authorizationCode)

	w.Header().Add("Location", s.redirectURL)
	w.WriteHeader(http.StatusMovedPermanently)
}

func (s *AuthorizationServer) Authorize(config *config.Config, chatID string, authCode string) {
	const authURL = "https://hh.ru/oauth/token"

	urlParameters := "grant_type=authorization_code&client_id=" + config.ClientID + "&client_secret=" + config.ClientSecret + "&redirect_uri=" + config.RedirectURI + "/?chat_id=" + chatID + "&code=" + authCode
	postData := []byte(urlParameters)

	r, _ := http.Post(authURL, "application/x-www-form-urlencoded", bytes.NewBuffer(postData))

	defer r.Body.Close()
	a, _ := io.ReadAll(r.Body)

	accessToken := gjson.Get(string(a), "access_token").String()
	refreshToken := gjson.Get(string(a), "refreshToken").String()

	chatIdInt, _ := strconv.ParseInt(chatID, 10, 64)

	s.tokenRepository.Save(chatIdInt, accessToken, repository.AccessTokens)
	s.tokenRepository.Save(chatIdInt, refreshToken, repository.RefreshToken)
}
