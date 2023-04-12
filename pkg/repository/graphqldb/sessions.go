package graphqldb

import (
	"context"
	"errors"
	"github.com/machinebox/graphql"
	"log"
	"strconv"
)

type GraphqlRepository struct {
	graphqlClient *graphql.Client
}

func NewGraphqlRepository(graphqlClient *graphql.Client) *GraphqlRepository {
	return &GraphqlRepository{graphqlClient: graphqlClient}
}

var getTokenRequest = graphql.NewRequest(`
	query GetToken($chat_id: String!) {
  		login(chat_id:$chat_id) {
    		id
    		access_token
    		refresh_token
  		}
	}
`)

var getSessionsRequest = graphql.NewRequest(`
	query GetSessions {
  		loginsList {
			items {
    			chat_id
			}
  		}
	}
`)

var createSessionRequest = graphql.NewRequest(`
	mutation CreateSession($chat_id: String!, $access_token: String!, $refresh_token: String!, $user_id: String!) {
  		loginCreate(data: {chat_id:$chat_id, access_token:$access_token, refresh_token:$refresh_token, user_id:$user_id}) {
    		chat_id
    		access_token
    		refresh_token
			user_id
  		}
	}
`)

var removeSessionRequest = graphql.NewRequest(`
	mutation RemoveSession($chat_id: String!) {
  		loginDelete(filter: {chat_id:$chat_id}) {
    		success
  		}
	}
`)

func (g *GraphqlRepository) GetSessions() ([]int64, error) {
	ctx := context.Background()

	var respData map[string]map[string][]map[string]string
	err := g.graphqlClient.Run(ctx, getSessionsRequest, &respData)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	chatIds := []int64{}

	for _, item := range respData["loginsList"]["items"][0] {
		chatId, _ := strconv.ParseInt(item, 10, 64)
		chatIds = append(chatIds, chatId)
	}

	return chatIds, nil
}

func (g *GraphqlRepository) CreateSession(chatId int64, accessToken string, refreshToken string, userId string) error {
	ctx := context.Background()

	createSessionRequest.Var("chat_id", strconv.FormatInt(chatId, 10))
	createSessionRequest.Var("access_token", accessToken)
	createSessionRequest.Var("refresh_token", refreshToken)
	createSessionRequest.Var("user_id", userId)

	var respData map[string]map[string]string
	err := g.graphqlClient.Run(ctx, createSessionRequest, &respData)
	if err != nil {
		return err
	}

	return nil
}

func (g *GraphqlRepository) RemoveSession(chatId int64) error {
	ctx := context.Background()

	removeSessionRequest.Var("chat_id", strconv.FormatInt(chatId, 10))

	var respData any
	err := g.graphqlClient.Run(ctx, removeSessionRequest, &respData)
	if err != nil {
		return err
	}

	return nil
}

func (g *GraphqlRepository) GetAccessToken(chatId int64) (string, error) {
	ctx := context.Background()

	getTokenRequest.Var("chat_id", strconv.FormatInt(chatId, 10))

	var respData map[string]map[string]string
	err := g.graphqlClient.Run(ctx, getTokenRequest, &respData)
	if err != nil {
		return "", err
	}

	var token string

	for _, item := range respData {
		token = item["access_token"]
	}

	if token == "" {
		return "", errors.New("unauthorized user")
	}

	return token, nil
}
