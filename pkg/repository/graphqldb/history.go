package graphqldb

import (
	"context"
	"github.com/machinebox/graphql"
	"log"
)

var updateResponseStatusRequest = graphql.NewRequest(`
	mutation MyMutation($vacancy_id: String!, $user_id: String!, $status: String!) {
  		response_historyUpdateByFilter(data: {status: {set: $status}}, filter: {vacancy_id: {equals: $vacancy_id}, AND: {user_id: {equals: $user_id}}}) {
    		items {
      			status
    		}
  		}
	}	
`)

var getResponsesRequest = graphql.NewRequest(`
	query GetResponses($user_id: String!) {
  		response_history(user_id:$user_id) {
			ID
    		user_id
			vacancy_name
			employer
			alternate_url
			area
			status
			salary_from
			salary_to
			date
			vacancy_id
  		}
	}
`)

var saveApplyToJobRequest = graphql.NewRequest(`
	mutation SaveApplyToJob($user_id: String!, $vacancy_name: String!, $employer: String!, $alternate_url: String!, $area: String!, $status: String!, $salary_from: Int!, $salary_to: Int!, $date: String!, $vacancy_id: String!) {
  		response_historyCreate(data: {user_id:$user_id, vacancy_name:$vacancy_name, employer:$employer, alternate_url:$alternate_url, area:$area, status:$status, salary_from:$salary_from, salary_to:$salary_to, date:$date, vacancy_id:$vacancy_id}) {
    		user_id
			vacancy_name
			employer
			alternate_url
			area
			status
			salary_from
			salary_to
			date
			vacancy_id
  		}
	}
`)

func (g *GraphqlRepository) SaveApplyToJob(userId string, vacancyId string, vacancyName string, employer string, alternateUrl string, area string, status string, salaryFrom int, salaryTo int, date string) error {
	ctx := context.Background()

	saveApplyToJobRequest.Var("vacancy_id", vacancyId)
	saveApplyToJobRequest.Var("user_id", userId)
	saveApplyToJobRequest.Var("vacancy_name", vacancyName)
	saveApplyToJobRequest.Var("employer", employer)
	saveApplyToJobRequest.Var("alternate_url", alternateUrl)
	saveApplyToJobRequest.Var("area", area)
	saveApplyToJobRequest.Var("status", status)
	saveApplyToJobRequest.Var("salary_from", salaryFrom)
	saveApplyToJobRequest.Var("salary_to", salaryTo)
	saveApplyToJobRequest.Var("date", date)

	var respData any
	err := g.graphqlClient.Run(ctx, saveApplyToJobRequest, &respData)
	if err != nil {
		return err
	}

	return nil
}

func (g *GraphqlRepository) UpdateResponseStatus(vacancyId string, userId string, status string) error {
	ctx := context.Background()

	updateResponseStatusRequest.Var("user_id", userId)
	updateResponseStatusRequest.Var("vacancy_id", vacancyId)
	updateResponseStatusRequest.Var("status", status)

	var respData any
	err := g.graphqlClient.Run(ctx, updateResponseStatusRequest, &respData)
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
