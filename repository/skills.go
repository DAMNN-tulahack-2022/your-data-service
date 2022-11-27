package repository

import (
	"context"
	"log"

	"github.com/damnn/tulahack/your-supadmin-service/tools"
)

type skills struct{}

const (
	getSkills = `select * from skills`
	getSkill  = `select * from skills where id = $1`
	addSkill  = `insert into skills (label) values ($1) RETURNING *`

	editSkill = `update skills set 
					label=:label
						where id=:id
							returning *`

	removeSkill = `delete from skills where id = $1`
)

type Skill struct {
	Id    uint32 `db:"id"`
	Label string `db:"label"`
}

func (skills) List(ctx context.Context) ([]*Skill, error) {
	result := make([]*Skill, 0)
	log.Printf("db is nil - [%v]", tools.DB == nil)
	err := tools.DB.SelectContext(ctx, &result, getSkills)
	return result, err
}

func (skills) Add(ctx context.Context, label string) (*Skill, error) {
	result := new(Skill)
	err := tools.DB.GetContext(ctx, result, addSkill, label)
	return result, err
}

func (skills) Edit(ctx context.Context, skill *Skill) (*Skill, error) {
	err := tools.DB.GetContext(ctx, skill, editSkill, skill)
	return skill, err
}

func (skills) Remove(ctx context.Context, id uint32) error {
	_, err := tools.DB.ExecContext(ctx, removeSkill, id)
	return err
}
