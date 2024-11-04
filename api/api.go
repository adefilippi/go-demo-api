package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"net/http"
)

type ApiError struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func HandleError(err error) (int, ApiError) {
	var response ApiError
	response.Message = err.Error()
	var code int

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		code = http.StatusNotFound
	case errors.Is(err, gorm.ErrInvalidTransaction):
		code = http.StatusInternalServerError
	case errors.Is(err, gorm.ErrNotImplemented):
		code = http.StatusNotImplemented
	case errors.Is(err, gorm.ErrMissingWhereClause):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrUnsupportedRelation):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrPrimaryKeyRequired):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrModelValueRequired):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrModelAccessibleFieldsRequired):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrSubQueryRequired):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrInvalidData):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrUnsupportedDriver):
		code = http.StatusInternalServerError
	case errors.Is(err, gorm.ErrRegistered):
		code = http.StatusConflict
	case errors.Is(err, gorm.ErrInvalidField):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrEmptySlice):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrDryRunModeUnsupported):
		code = http.StatusNotImplemented
	case errors.Is(err, gorm.ErrInvalidDB):
		code = http.StatusInternalServerError
	case errors.Is(err, gorm.ErrInvalidValue):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrInvalidValueOfLength):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrPreloadNotAllowed):
		code = http.StatusBadRequest
	case errors.Is(err, gorm.ErrDuplicatedKey):
		code = http.StatusConflict
	case errors.Is(err, gorm.ErrForeignKeyViolated):
		code = http.StatusConflict
	case errors.Is(err, gorm.ErrCheckConstraintViolated):
		code = http.StatusConflict

	default:
		code = http.StatusInternalServerError
	}

	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch pgErr.Code {
		case "23505": // unique_violation
			code = http.StatusConflict
		case "23503": // foreign_key_violation
			code = http.StatusConflict
		case "23502": // not_null_violation
			code = http.StatusBadRequest
		case "23514": // check_violation
			code = http.StatusConflict
		case "42703": // undefined_column
			code = http.StatusBadRequest
		case "42883": // undefined_function
			code = http.StatusBadRequest
		case "42601": // syntax_error
			code = http.StatusBadRequest
		case "23508": // exclusion_violation
			code = http.StatusConflict
		case "22P02": // invalid_text_representation (e.g. invalid input syntax for type)
			code = http.StatusBadRequest
		case "22007": // invalid_datetime_format
			code = http.StatusBadRequest
		case "42P01": // undefined_table
			code = http.StatusNotFound
		case "42P07": // duplicate_table
			code = http.StatusConflict
		case "40P01": // deadlock_detected
			code = http.StatusInternalServerError
		default:
			code = http.StatusInternalServerError
		}

		response.Message = pgErr.Message
		response.Detail = pgErr.Detail

	} else {
		// Generic error handling for other error types
		code = http.StatusInternalServerError
		response.Message = "An unexpected error occurred"
		response.Detail = err.Error()
	}

	return code, response
}

func Home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Ok")
}
