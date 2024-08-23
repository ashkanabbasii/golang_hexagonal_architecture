package serr_test

import (
	"b2bapi/internal/serr"
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDBError(t *testing.T) {
	// DB errors' status must be 500 to avoid error leaking unless it's a not found error
	t.Run("internal error", func(t *testing.T) {
		t.Parallel()
		err := errors.New("random err")
		dbErr := serr.NewDBError("Test", "Test", err)
		assert.Equal(t, dbErr.(*serr.ServiceError).Code, 500)
	})
	t.Run("sql not found db error", func(t *testing.T) {
		t.Parallel()
		dbErr := serr.NewDBError("Test", "Test", sql.ErrNoRows)
		assert.Equal(t, dbErr.(*serr.ServiceError).Code, 404)
	})
}

func TestIsDBNoRows(t *testing.T) {
	t.Run("sql no rows", func(t *testing.T) {
		t.Parallel()
		assert.True(t, serr.IsDBNoRows(sql.ErrNoRows))
	})
}
