package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/elghazx/perpustakaan/domain"
)

type book struct {
	db *goqu.Database
}

func NewBook(con *sql.DB) domain.BookRepository {
	return &book{
		db: goqu.New("default", con),
	}
}

func (br book) FindAll(ctx context.Context) (books []domain.Book, err error) {
	dataset := br.db.From("books").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &books)

	return
}

func (br book) FindById(ctx context.Context, id string) (book domain.Book, err error) {
	dataset := br.db.From("books").
		Where(
			goqu.C("id").Eq(id),
			goqu.C("deleted_at").IsNull(),
		)

	_, err = dataset.ScanStructContext(ctx, &book)
	return
}

func (br book) Save(ctx context.Context, b *domain.Book) error {
	executor := br.db.Insert("books").Rows(b).Executor()
	_, err := executor.ExecContext(ctx)

	return err
}

func (br book) Update(ctx context.Context, b *domain.Book) error {
	executor := br.db.Update("books").
		Where(goqu.C("id").Eq(b.Id)).
		Set(b).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (br book) Delete(ctx context.Context, id string) error {
	executor := br.db.Update("books").
		Where(goqu.C("id").Eq(id)).
		Set(goqu.Record{
			"deleted_at": sql.NullTime{
				Valid: true,
				Time:  time.Now(),
			},
		}).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
