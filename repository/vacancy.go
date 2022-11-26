package repository

import (
	"context"

	"github.com/damnn/tulahack/your-supadmin-service/tools"
)

type vacancies struct{}

const (
	getVacancies = `select * from vacancy`
	addVacancy   = `insert into vacancy (label) values ($1) RETURNING *`

	editVacancy = `update vacancy set 
					label=:label
						where id=:id
							returning *`

	removeVacancy = `delete from vacancy where id = $1`
)

type Vacancy struct {
	Id    uint32 `db:"id"`
	Label string `db:"label"`
}

func (vacancies) List(ctx context.Context) ([]*Vacancy, error) {
	result := make([]*Vacancy, 0)
	err := tools.DB.SelectContext(ctx, &result, getVacancies)
	return result, err
}

func (vacancies) Add(ctx context.Context, label string) (*Vacancy, error) {
	result := new(Vacancy)
	err := tools.DB.GetContext(ctx, result, addVacancy, label)
	return result, err
}

func (vacancies) Edit(ctx context.Context, vacancy *Vacancy) (*Vacancy, error) {
	err := tools.DB.GetContext(ctx, vacancy, editSkill, editVacancy)
	return vacancy, err
}

func (vacancies) Remove(ctx context.Context, id uint32) error {
	_, err := tools.DB.ExecContext(ctx, removeVacancy, id)
	return err
}
