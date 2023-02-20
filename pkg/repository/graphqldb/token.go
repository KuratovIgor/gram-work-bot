package graphqldb

import (
	"context"
	"github.com/machinebox/graphql"
)

type GraphqlRepository struct {
	graphqlClient *graphql.Client
}

func NewGraphqlRepository(graphqlClient *graphql.Client) *GraphqlRepository {
	return &GraphqlRepository{graphqlClient: graphqlClient}
}

var req = graphql.NewRequest(`
	query MyQuery {
  		login(chat_id: "12") {
    		id
    		access_token
    		refresh_token
  		}
	}
`)

func (g *GraphqlRepository) GetAccessToken(chatId int64) (string, error) {
	ctx := context.Background()
	var respData map[string]map[string]string
	err := g.graphqlClient.Run(ctx, req, &respData)
	if err != nil {
		return "nil", err
	}

	var token string

	for _, item := range respData {
		token = item["access_token"]
	}

	return token, nil
}
