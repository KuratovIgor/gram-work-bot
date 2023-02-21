package graphqldb

import (
	"context"
	"errors"
	"github.com/machinebox/graphql"
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

var createSessionRequest = graphql.NewRequest(`
	mutation CreateSession($chat_id: String!, $access_token: String!, $refresh_token: String!) {
  		loginCreate(data: {chat_id:$chat_id, access_token:$access_token, refresh_token:$refresh_token}) {
    		chat_id
    		access_token
    		refresh_token
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

func (g *GraphqlRepository) CreateSession(chatId int64, accessToken string, refreshToken string) error {
	ctx := context.Background()

	createSessionRequest.Var("chat_id", strconv.FormatInt(chatId, 10))
	createSessionRequest.Var("access_token", accessToken)
	createSessionRequest.Var("refresh_token", refreshToken)

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
