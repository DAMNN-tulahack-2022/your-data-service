package repository

import (
	"context"

	"github.com/damnn/tulahack/your-supadmin-service/tools"
)

type projects struct{}

const (
	getProjects = `select * from project`
	getProject  = `select * from project where id = $1`

	addProject = `insert into project
	 					(id, title, description, skills_ids, user_ids, team_lead_id, exp)
					values (:id, :title, :description, :skills_ids, :user_ids, :team_lead_id, :exp)
						RETURNING *`

	assignProjectUser = `update project
							set user_ids=:user_ids
								where id=:id
									RETURNING *`
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

func (projects) Get(ctx context.Context, projectId uint32) (*Project, error) {
	result := new(Project)
	err := tools.DB.GetContext(ctx, result, getProject, projectId)
	return result, err
}

func (projects) Add(ctx context.Context, project *Project) (*Project, error) {
	stmt, err := tools.DB.PrepareNamed(addProject)
	if err != nil {
		return nil, err
	}

	if err = stmt.GetContext(ctx, project, project); err != nil {
		return nil, err
	}

	return project, err
}

func (projects) AssignUser(ctx context.Context, tx *Transaction, project *Project) (*Project, error) {
	stmt, err := tx.Tx.PrepareNamed(assignProjectUser)
	if err != nil {
		return nil, err
	}

	if _, err = stmt.ExecContext(ctx, project); err != nil {
		return nil, err
	}

	return project, err
}
