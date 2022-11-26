package repository

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"

	"github.com/damnn/tulahack/your-supadmin-service/tools"
)

type matricies struct{}

const (
	getMatricies = `select * from grade_progress`
	getMatrix    = `select * from grade_progress where id = $1`

	addMattrix = `insert into grade_progress (vacancy_id, grades_ids)
						values (:vacancy_id, :grades_ids)
							RETURNING *`

	editMatrix = `update grade_progress set 
					vacancy_id=:vacancy_id
						where id=:id
							returning *`

	removeMatrix = `delete from grade_progress where id = $1`
)

type Uint32Array []uint32

func (ua Uint32Array) Value() (driver.Value, error) {
	return json.Marshal(ua)
}

func (ua *Uint32Array) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &ua)
}

type Matrix struct {
	Id        uint32      `db:"id"`
	VacancyId uint32      `db:"vacancy_id"`
	GradesIds Uint32Array `db:"grades_ids"`
}

func (matricies) List(ctx context.Context) ([]*Matrix, error) {
	result := make([]*Matrix, 0)
	err := tools.DB.SelectContext(ctx, &result, getMatricies)
	return result, err
}

func (matricies) Get(ctx context.Context, id uint32) (*Matrix, error) {
	result := new(Matrix)
	log.Print(id)
	err := tools.DB.GetContext(ctx, result, getMatrix, id)
	return result, err
}

func (matricies) Add(ctx context.Context, v *Matrix) (*Matrix, error) {
	stmt, err := tools.DB.PrepareNamed(addMattrix)
	if err != nil {
		return nil, err
	}

	if err = stmt.GetContext(ctx, v, v); err != nil {
		return nil, err
	}

	return v, err
}

func (matricies) Edit(ctx context.Context, v *Matrix) (*Matrix, error) {
	stmt, err := tools.DB.PrepareNamed(editMatrix)
	if err != nil {
		return nil, err
	}

	if err = stmt.GetContext(ctx, v, v); err != nil {
		return nil, err
	}

	return v, err
}

func (matricies) Remove(ctx context.Context, id uint32) error {
	_, err := tools.DB.ExecContext(ctx, removeMatrix, id)
	return err
}
