package repository

import (
	"context"
	"log"

	"github.com/damnn/tulahack/your-supadmin-service/tools"
)

type grades struct{}

const (
	getGrades = `select * from grade`
	getGrade  = `select * from grade where id = $1`
	addGrade  = `insert into grade (label, description, exp, needs_approval)
					values (:label, :description, :exp, :needs_approval)
						RETURNING *`

	editGrade = `update grade set 
					label=:label, description=:description,
					exp=:exp, needs_approval=:needs_approval
						where id=:id
							returning *`

	removeGrade = `delete from grade where id = $1`
)

type Grade struct {
	Id            uint32  `db:"id"`
	Label         string `db:"label"`
	Description   string `db:"description"`
	Exp           uint32 `db:"exp"`
	NeedsApproval bool   `db:"needs_approval"`
}

func (grades) List(ctx context.Context) ([]*Grade, error) {
	result := make([]*Grade, 0)
	err := tools.DB.SelectContext(ctx, &result, getGrades)
	return result, err
}

func (grades) Get(ctx context.Context, id uint32) (*Grade, error) {
	result := new(Grade)
	log.Print(id)
	err := tools.DB.GetContext(ctx, result, getGrade, id)
	return result, err
}

func (grades) Add(ctx context.Context, grade *Grade) (*Grade, error) {
	stmt, err := tools.DB.PrepareNamed(addGrade)
	if err != nil {
		return nil, err
	}

	if err = stmt.GetContext(ctx, grade, grade); err != nil {
		return nil, err
	}

	return grade, err
}

func (grades) Edit(ctx context.Context, grade *Grade) (*Grade, error) {
	stmt, err := tools.DB.PrepareNamed(editGrade)
	if err != nil {
		return nil, err
	}

	if err = stmt.GetContext(ctx, grade, grade); err != nil {
		return nil, err
	}

	return grade, err
}

func (grades) Remove(ctx context.Context, id uint32) error {
	_, err := tools.DB.ExecContext(ctx, removeSkill, id)
	return err
}
