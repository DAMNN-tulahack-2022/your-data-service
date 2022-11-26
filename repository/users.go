package repository

import (
	"context"

	"github.com/damnn/tulahack/your-supadmin-service/tools"
)

type users struct{}

const (
	getUsers = `select * from your_user`
)

type User struct {
	Id                   uint32      `db:"id"`
	TotalExp             uint32      `db:"total_exp"`
	VacancyId            uint32      `db:"vacancy_id"`
	Login                string      `db:"login"`
	FirstName            string      `db:"first_name"`
	LastName             string      `db:"last_name"`
	MiddleName           string      `db:"middle_name"`
	AvaterURL            string      `db:"avatar_uri"`
	CurrentProjectId     uint32      `db:"current_project_id"`
	ViewedPostIds        Uint32Array `db:"viewed_ids"`
	CompletedProjectsIds Uint32Array `db:"completed_project_ids"`
	SkillsIds            Uint32Array `db:"skills_ids"`
	Role                 string      `db:"role"`
}

func (users) List(ctx context.Context) ([]*User, error) {
	result := make([]*User, 0)
	err := tools.DB.SelectContext(ctx, &result, getUsers)
	return result, err
}
