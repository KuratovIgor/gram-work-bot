package graphqldb

import (
	"context"
	"github.com/machinebox/graphql"
	"strconv"
)

var createUserRequest = graphql.NewRequest(`
	mutation CreateUser($chat_id: String!, $name: String!, $lastname: String!, $middlename: String!, $email: String!, $phone: String!, $user_id: String!) {
  		auth_userCreate(data: {chat_id:$chat_id, name:$name, lastname:$lastname, middlename:$middlename, email:$email, phone:$phone, user_id:$user_id}) {
    		chat_id
    		name
			lastname
			middlename
			email
			phone
			user_id
  		}
	}
`)

func (g *GraphqlRepository) CreateUser(chatId int64, name string, lastname string, middlename string, email string, phone string, userId string) error {
	ctx := context.Background()

	createUserRequest.Var("chat_id", strconv.FormatInt(chatId, 10))
	createUserRequest.Var("name", name)
	createUserRequest.Var("lastname", lastname)
	createUserRequest.Var("middlename", middlename)
	createUserRequest.Var("email", email)
	createUserRequest.Var("phone", phone)
	createUserRequest.Var("user_id", userId)

	var respData map[string]map[string]string
	err := g.graphqlClient.Run(ctx, createUserRequest, &respData)
	if err != nil {
		return err
	}

	return nil
}
