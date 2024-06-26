package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Salikhov079/library-management/genprotos"
	"github.com/google/uuid"
)

type BorrowerStorage struct {
	db *sql.DB
}

func NewBorrowerStorage(db *sql.DB) *BorrowerStorage {
	return &BorrowerStorage{db: db}
}

func (p *BorrowerStorage) Create(borrower *pb.BorrowerReq) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO borrowers 
			(id, user_id, book_id, borrow_date, return_date)
		VALUES 
			($1, $2, $3, $4, $5)
	`
	_, err := p.db.Exec(query, id, borrower.UserId, borrower.BookId, borrower.BorrowDate, borrower.ReturnDate)
	return nil, err
}

func (p *BorrowerStorage) GetById(id *pb.ById) (*pb.BorrowerRes, error) {
	query := `
		SELECT
			b.title, a.name, a.biography, g.name, b.summary, br.borrow_date, br.return_date, br.user_id
		FROM 
			borrowers br JOIN books b
		ON 
			br.book_id = b.id
		JOIN 
			authors a
		ON 
			b.author_id = a.id
		JOIN 
			genres g
		ON 
			b.genre_id = g.id
		WHERE 
			br.id = $1 and br.deleted_at = 0
	`

	row := p.db.QueryRow(query, id.Id)
	var genre pb.Genre
	var book pb.BookRes
	var author pb.AuthorRes
	var borrower pb.BorrowerRes
	var user pb.UserRes

	err := row.Scan(
		&book.Title,
		&author.Name,
		&author.Biography,
		&genre.Name,
		&book.Summary,
		&borrower.BorrowDate,
		&borrower.ReturnDate,
		&user.Id)
	if err != nil {
		return nil, err
	}
	book.Author = &author
	book.Genre = &genre
	borrower.Book = &book
	borrower.User = &user

	return &borrower, nil
}

func (p *BorrowerStorage) GetAll(filter *pb.FilterBorrower) (*pb.AllBorrowers, error) {
	borrowers := &pb.AllBorrowers{}
	var arr []interface{}
	count := 1

	query := `
		SELECT
			b.title, a.name, a.biography, g.name, b.summary, br.borrow_date, br.return_date, br.user_id
		FROM 
			borrowers br JOIN books b
		ON 
			br.book_id = b.id
		JOIN 
			authors a
		ON 
			b.author_id = a.id
		JOIN 
			genres g
		ON 
			b.genre_id = g.id
		WHERE 
	`
	if len(filter.BorrowDate) > 0 {
		if count == 1 {
			query += fmt.Sprintf(" br.borrow_date=$%d ", count)
			count++
		} else {
			query += fmt.Sprintf(" and br.borrow_date=$%d ", count)
			count++
		}
		arr = append(arr, filter.BorrowDate)
	}

	if len(filter.ReturnDate) > 0 {
		if count == 1 {
			query += fmt.Sprintf(" br.return_date=$%d ", count)
			count++
		} else {
			query += fmt.Sprintf(" and br.return_date=$%d ", count)
			count++
		}
		arr = append(arr, filter.ReturnDate)
	}

	if len(filter.UserId) > 0 {
		if count == 1 {
			query += fmt.Sprintf(" br.user_id=$%d ", count)
			count++
		} else {
			query += fmt.Sprintf(" and br.user_id=$%d ", count)
			count++
		}
		arr = append(arr, filter.UserId)
	}

	if filter.DeletedAt != "true" {
		if count == 1 {
			query += " br.deleted_at = 0 "
		} else {
			query += " and br.deleted_at = 0 "
		}
	}

	row, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var genre pb.Genre
		var book pb.BookRes
		var author pb.AuthorRes
		var borrower pb.BorrowerRes
		var user pb.UserRes

		err := row.Scan(
			&book.Title,
			&author.Name,
			&author.Biography,
			&genre.Name,
			&book.Summary,
			&borrower.BorrowDate,
			&borrower.ReturnDate,
			&user.Id)
		if err != nil {
			return nil, err
		}
		book.Author = &author
		book.Genre = &genre
		borrower.Book = &book
		borrower.User = &user

		borrowers.Borrowers = append(borrowers.Borrowers, &borrower)
	}

	return borrowers, nil
}

func (p *BorrowerStorage) Update(borrower *pb.BorrowerRes) (*pb.Void, error) {
	query := `
		UPDATE borrowers
		SET user_id = $2, book_id = $3, borrow_date = $4, return_date = $5, updated_at = now()
		WHERE id = $1 and deleted_at = 0 
	`

	_, err := p.db.Exec(query, borrower.Id, borrower.User.Id, borrower.Book.Id, borrower.BorrowDate, borrower.ReturnDate)
	return nil, err
}

func (p *BorrowerStorage) Delete(id *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE borrowers 
		SET deleted_at = $2
		WHERE id = $1 and deleted_at = 0
	`
	_, err := p.db.Exec(query, id.Id, time.Now().Unix())
	return nil, err
}

func (p *BorrowerStorage) GetAllId(*pb.Void) (*pb.AllBorrowers, error) {
	borrowers := &pb.AllBorrowers{}

	query := `
		SELECT
			id, borrow_date, return_date
		FROM 
			borrowers
		WHERE 
			deleted_at = 0
	`

	row, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var borrower pb.BorrowerRes

		err := row.Scan(&borrower.Id, &borrower.BorrowDate, &borrower.ReturnDate)
		if err != nil {
			return nil, err
		}

		borrowers.Borrowers = append(borrowers.Borrowers, &borrower)
	}

	return borrowers, nil
}
