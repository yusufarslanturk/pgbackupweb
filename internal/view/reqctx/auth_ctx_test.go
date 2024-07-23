package reqctx

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAuthCtxFuncs(t *testing.T) {
	testUser := dbgen.User{
		ID:    uuid.New(),
		Email: "user@example.com",
		Name:  "John",
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	t.Run("Create authentication values in context", func(t *testing.T) {
		authData := AuthCtx{
			IsAuthed: true,
			User:     testUser,
		}

		SetAuthCtx(c, authData)
		auth := GetAuthCtx(c)

		assert.True(t, auth.IsAuthed)
		assert.Equal(t, testUser, auth.User)
		assert.Equal(t, testUser.Email, auth.User.Email)
	})

	t.Run("Create authentication values in context with only IsAuthed", func(t *testing.T) {
		authData := AuthCtx{
			IsAuthed: true,
		}

		SetAuthCtx(c, authData)
		auth := GetAuthCtx(c)

		assert.True(t, auth.IsAuthed)
		assert.Empty(t, auth.User)
	})
}
