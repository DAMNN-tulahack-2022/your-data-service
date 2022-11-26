package repository

import (
	"context"

	"github.com/damnn/tulahack/your-supadmin-service/tools"
)

type articles struct{}

const (
	getArticles = `select * from article`
	getArticle  = `select * from article where id = $1`

	viewArtice = `update article set
	 					total_viewed=total_viewed+1
							where id = $1
				 				returning *`

	addArticle = `insert into article (author_id, exp, title, description, skills_ids)
			 			values (:author_id, 100, :title, :description, :skills_ids)
							RETURNING *`
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
	err := tools.DB.SelectContext(ctx, &result, getArticles)
	return result, err
}

func (articles) Get(ctx context.Context, id uint32) (*Article, error) {
	result := new(Article)
	err := tools.DB.GetContext(ctx, result, getArticle, id)
	return result, err
}

func (articles) Add(ctx context.Context, article *Article) (*Article, error) {
	stmt, err := tools.DB.PrepareNamed(addArticle)
	if err != nil {
		return nil, err
	}

	if err = stmt.GetContext(ctx, article, article); err != nil {
		return nil, err
	}

	return article, err
}

func (articles) View(ctx context.Context, tx *Transaction, id uint32) (*Article, error) {
	result := new(Article)
	err := tx.Tx.GetContext(ctx, result, viewArtice, id)
	return result, err
}
