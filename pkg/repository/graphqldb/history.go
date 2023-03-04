package graphqldb

import (
	"context"
	"github.com/machinebox/graphql"
)

var saveApplyToJobRequest = graphql.NewRequest(`
	mutation SaveApplyToJob($user_id: String!, $vacancy_name: String!, $employer: String!, $alternate_url: String!, $area: String!, $status: String!, $salary_from: Int!, $salary_to: Int!, $date: String!) {
  		response_historyCreate(data: {user_id:$user_id, vacancy_name:$vacancy_name, employer:$employer, alternate_url:$alternate_url, area:$area, status:$status, salary_from:$salary_from, salary_to:$salary_to, date:$date}) {
    		user_id
			vacancy_name
			employer
			alternate_url
			area
			status
			salary_from
			salary_to
			date
  		}
	}
`)

func (g *GraphqlRepository) SaveApplyToJob(userId string, vacancyName string, employer string, alternateUrl string, area string, status string, salaryFrom int, salaryTo int, date string) error {
	ctx := context.Background()

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
