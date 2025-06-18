package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/elghazx/perpustakaan/domain"
)

type bookStock struct {
	db *goqu.Database
}

func NewBookStock(con *sql.DB) domain.BookStockRepository {
	return &bookStock{
		db: goqu.New("default", con),
	}
}

func (b bookStock) FindByBookId(ctx context.Context, id string) (bookStocks []domain.BookStock, err error) {
	dataset := b.db.From("book_stocks").Where(goqu.C("book_id").Eq(id))
	err = dataset.ScanStructsContext(ctx, &bookStocks)
	return
}

func (b bookStock) FindByBookAndCode(ctx context.Context, id string, code string) (book domain.BookStock, err error) {
	dataset := b.db.From("book_stocks").Where(
		goqu.C("book_id").Eq(id),
		goqu.C("code").Eq(code),
	)
	_, err = dataset.ScanStructContext(ctx, &book)
	return
}

func (b bookStock) Save(ctx context.Context, data []domain.BookStock) error {
	executor := b.db.Insert("book_stocks").
		Rows(data).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (b bookStock) Update(ctx context.Context, stock *domain.BookStock) error {
	executor := b.db.Update("book_stocks").
		Where(goqu.C("code").Eq(stock.Code)).
		Set(stock).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (b bookStock) DeletedByBookId(ctx context.Context, id string) error {
	executor := b.db.Delete("book_stocks").
		Where(goqu.C("book_id").Eq(id)).
		Executor()

	_, err := executor.ExecContext(ctx)
	return err
}

func (b bookStock) DeletedByCodes(ctx context.Context, codes []string) error {
	executor := b.db.Delete("book_stocks").
		Where(goqu.C("code").In(codes)).
		Executor()

	_, err := executor.ExecContext(ctx)
	return err
}
