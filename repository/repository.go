package repository

import (
	"gorm.io/gorm"

	"github.com/lunmy/go-api-core/service/query"
)

// GenericRepository interface for common database operations
type GenericRepository interface {
	Create(value interface{}) error
	FindOne(dest interface{}, args ...interface{}) error
	FindAll(dest interface{}, queryParams *query.Params) (int64, error)
	Update(value interface{}, updates interface{}) error
	Delete(value interface{}, args ...interface{}) error
}

// genericRepository implements GenericRepository
type genericRepository struct {
	db *gorm.DB
}

// NewGenericRepository creates a new GenericRepository
func NewGenericRepository(db *gorm.DB) GenericRepository {
	return &genericRepository{db: db}
}

// Create a new record
func (r *genericRepository) Create(value interface{}) error {
	return r.db.Create(value).Error
}

// FindOne finds a single record by query
func (r *genericRepository) FindOne(dest interface{}, args ...interface{}) error {
	return r.db.First(dest, args...).Error
}

// FindAll finds all records with pagination, sorting, and filtering
func (r *genericRepository) FindAll(dest interface{}, queryParams *query.Params) (int64, error) {
	var total int64

	// Apply filters
	db := r.db.Model(dest)
	if queryParams.Filters != nil {
		for _, filter := range queryParams.Filters {
			db = db.Where(filter.Field+" = ?", filter.Value)
		}
	}

	// Count total items
	db.Count(&total)

	// Apply pagination
	offset := (queryParams.Pagination.Page - 1) * queryParams.Pagination.ItemsPerPage
	db = db.Offset(offset).Limit(queryParams.Pagination.ItemsPerPage)

	// Apply sorting
	if queryParams.Sort != nil {
		for _, sort := range queryParams.Sort {
			db = db.Order(sort.Field + " " + sort.Direction)
		}
	}

	return total, db.Find(dest).Error
}

// Update updates a record
func (r *genericRepository) Update(value interface{}, updates interface{}) error {
	return r.db.Model(value).Updates(updates).Error
}

// Delete deletes a record by query
func (r *genericRepository) Delete(value interface{}, args ...interface{}) error {
	return r.db.Delete(value, args...).Error
}