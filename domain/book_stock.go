package domain

import "context"

type BookStock struct {
	Code       string `db:"code"`
	BookId     string `db:"book_id"`
	Status     string `db:"status"`
	BorrowerId string `db:"borrower_id"`
	BorrowedAt string `db:"borrowed_at"`
}

type BookStockRepository interface {
	FindByBookId(ctx context.Context, id string) ([]BookStock, error)
	FindByBookAndCode(ctx context.Context, id string, code string) (BookStock, error)
	Save(ctx context.Context, data []BookStock) error
	Update(ctx context.Context, stock *BookStock) error
	DeletedByBookId(ctx context.Context, id string) error
	DeletedByCodes(ctx context.Context, codes []string) error
}
