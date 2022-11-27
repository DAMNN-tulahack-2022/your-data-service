package repository

import (
	"context"

	"github.com/damnn/tulahack/your-supadmin-service/tools"
)

type users struct{}

const (
	getUsers          = `select * from your_user`
	getUser           = `select * from your_user where id = $1`
	updateUserExp     = `update your_user set total_exp=total_exp + $1 where id = $2 returning *`
	updateUserRole    = `update your_user set role=$1 where id = $2 returning *`
	updateUserVacancy = `update your_user set vacancy_id=$1
	 						where id = $2
						 		returning *`

	updateUserViewedMetrics = `update your_user
									set total_exp=:total_exp,
									viewed_ids=:viewed_ids
					 					where id=:id`
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

func (users) Get(ctx context.Context, id uint32) (*User, error) {
	result := new(User)
	err := tools.DB.GetContext(ctx, result, getUser, id)
	return result, err
}

func (users) UpExp(ctx context.Context, id, exp uint32) (*User, error) {
	result := new(User)
	err := tools.DB.GetContext(ctx, result, updateUserExp, exp, id)
	return result, err
}

func (users) ChangeRole(ctx context.Context, id uint32, role string) (*User, error) {
	result := new(User)
	err := tools.DB.GetContext(ctx, result, updateUserRole, role, id)
	return result, err
}

func (users) AssignVacancy(ctx context.Context, id, vacancyId uint32) (*User, error) {
	result := new(User)
	err := tools.DB.GetContext(ctx, result, updateUserVacancy, vacancyId, id)
	return result, err
}

func (users) ReadPost(ctx context.Context, tx *Transaction, user *User) (*User, error) {
	stmt, err := tx.Tx.PrepareNamed(updateUserViewedMetrics)
	if err != nil {
		return nil, err
	}

	if _, err = stmt.ExecContext(ctx, user); err != nil {
		return nil, err
	}

	return user, err
}
