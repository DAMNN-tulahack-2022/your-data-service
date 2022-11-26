package repository

import (
	"context"

	"github.com/damnn/tulahack/your-supadmin-service/tools"
)

type articles struct{}

const (
	getArticle = `select * from article`
)

type Article struct {
	Id          uint32      `db:"id"`
	AuthorId    uint32      `db:"author_id"`
	TotalViewed uint32      `db:"total_viewed"`
	Exp         uint32      `db:"exp"`
	Title       string      `db:"title"`
	Description string      `db:"description"`
	SkillsIds   Uint32Array `db:"skills_ids"`
}

func (articles) List(ctx context.Context) ([]*Article, error) {
	result := make([]*Article, 0)
	err := tools.DB.SelectContext(ctx, &result, getArticle)
	return result, err
}
