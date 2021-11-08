package book

import (
	"context"
	"strconv"

	"github.com/pkg/errors"
	"github.com/soncaodb/model"
	"gorm.io/gorm"
)

type pgRepository struct {
	getClient func(ctx context.Context) *gorm.DB
}

func NewPGRepository(getClient func(ctx context.Context) *gorm.DB) Repository {
	return &pgRepository{getClient}
}

func (r *pgRepository) GetById(ctx context.Context, ID int64) (*model.Book, error) {
	db := r.getClient(ctx)
	book := &model.Book{}

	err := db.Where("id = ?", ID).
		First(book).Error

	if err != nil {
		return nil, errors.Wrap(err, "get book by id")
	}

	return book, nil
}

func (r *pgRepository) GetAll(ctx context.Context) ([]model.Book, error) {
	db := r.getClient(ctx)
	listBook := []model.Book{}

	db.Find(&listBook)

	return listBook, nil
}

func (r *pgRepository) Insert(ctx context.Context, book *model.Book) (*model.Book, error) {
	db := r.getClient(ctx)
	err := db.Create(book).Error

	return book, errors.Wrap(err, "create book")
}

func (r *pgRepository) Delete(ctx context.Context, id int64) error {
	db := r.getClient(ctx)
	err := db.Where("id = ?", id).Delete(&model.Book{}).Error

	return errors.Wrap(err, "delete book fail")
}

func (r *pgRepository) Update(ctx context.Context, book *model.Book) (*model.Book, error) {
	db := r.getClient(ctx)
	err := db.Save(book).Error

	return book, errors.Wrap(err, "update book fail")
}

func (r *pgRepository) GetTitle(ctx context.Context, tile string) (*model.Book, error) {
	db := r.getClient(ctx)
	book := &model.Book{}

	err := db.Where("title = ?", tile).
		First(book).Error

	if err != nil {
		return nil, errors.Wrap(err, "get book by title")
	}

	return book, nil
}

func (r *pgRepository) SearchBook(ctx context.Context,
	paginator *model.Paginator,
	searchText string,
	filter *model.BookFilter,
	orders []string) (*model.BookResult, error) {
	db := r.getClient(ctx)
	query := db.Model(&model.Book{})

	//Order
	for _, order := range orders {
		query.Order(order)
	}

	filterAuthor := ""
	filterTitle := ""
	filterCate := ""
	filterRate := ""

	// if filter.AuthorID != 0 {
	// 	filterAuthor = "JOIN book_authors ON book_authors.book_id = books.id AND book_authors.author_id = " + strconv.FormatInt(filter.AuthorID, 10)
	// }

	if filter.CateID != 0 {
		filterCate = "JOIN book_categories ON book_categories.book_id = books.id AND book_categories.category_id = " + strconv.FormatInt(filter.CateID, 10)
	}

	if filter.MinRating != -1 {
		filterRate = "AND rating_average > " + strconv.Itoa(filter.MinRating)
	}

	if searchText != "" {
		filterTitle = "AND title LIKE " + "'%" + searchText + "%'"
	}

	queryStr := "SELECT * FROM books" +
		filterAuthor + " " +
		filterCate + " WHERE isnull(books.deleted_at) " +
		filterRate + " " +
		filterTitle

	//Paging
	var res model.BookResult

	if paginator.Limit >= 0 {
		if paginator.Page <= 0 {
			paginator.Page = 1
		}

		if paginator.Limit == 0 {
			paginator.Limit = model.PageSize
		}
		res.Page = paginator.Page
		res.Limit = paginator.Limit
		query.Count(&res.Total).Scopes(paginator.Paginate())
	}

	lim := strconv.Itoa(res.Limit)
	limit := " LIMIT " + lim
	offsetQuery := strconv.Itoa((res.Page - 1) * res.Limit)
	offset := " OFFSET " + offsetQuery
	queryStr += limit + offset

	err := query.Raw(queryStr).Find(&res.Data).Error

	return &res, err
}
