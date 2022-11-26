package repository

import (
	"context"

	"github.com/damnn/tulahack/your-supadmin-service/tools"
)

type projects struct{}

const (
	getProjects = `select * from project`
)

type Project struct {
	Id          uint32      `db:"id"`
	Title       string      `db:"title"`
	Description string      `db:"description"`
	SkillsIds   Uint32Array `db:"skills_ids"`
	UserIds     Uint32Array `db:"user_ids"`
	TeamLeadId  uint32      `db:"team_lead_id"`
	Exp         uint32      `db:"exp"`
}

func (projects) List(ctx context.Context) ([]*Project, error) {
	result := make([]*Project, 0)
	err := tools.DB.SelectContext(ctx, &result, getProjects)
	return result, err
}
