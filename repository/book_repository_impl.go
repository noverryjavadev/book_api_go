package repository

import (
	"context"
	"database/sql"
	"errors"
	"fresh_api/helper"
	"fresh_api/model"
)

type BookRepositoryImpl struct {
	Db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &BookRepositoryImpl{Db: db}
}

// Delete implements BookRepository
func (b *BookRepositoryImpl) Delete(ctx context.Context, bookId int) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "delete from book where id =$1"
	_, errExec := tx.ExecContext(ctx, SQL, bookId)
	helper.PanicIfError(errExec)
}

func (b *BookRepositoryImpl) FindAll(ctx context.Context) []model.Book {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id,name from book"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var books []model.Book

	for result.Next() {
		book := model.Book{}
		err := result.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)

		books = append(books, book)
	}

	return books
}

func (b *BookRepositoryImpl) FindById(ctx context.Context, bookId int) (model.Book, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id,name from book where id=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, bookId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	book := model.Book{}

	if result.Next() {
		err := result.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)
		return book, nil
	} else {
		return book, errors.New("book id not found")
	}
}

func (b *BookRepositoryImpl) Save(ctx context.Context, book model.Book) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into book(name) values ($1)"
	_, err = tx.ExecContext(ctx, SQL, book.Name)
	helper.PanicIfError(err)
}

func (b *BookRepositoryImpl) Update(ctx context.Context, book model.Book) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update book set name=$1 where id=$2"
	_, err = tx.ExecContext(ctx, SQL, book.Name, book.Id)
	helper.PanicIfError(err)
}
