package boats

import (
	storage "boats/pkg/storage/mocks"
	"context"
	"github.com/kilchik/logo/pkg/logo"
	"github.com/pkg/errors"
	assert2 "github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBoatsServer_HandleSuggest(t *testing.T) {
	assert := assert2.New(t)
	logo.Init(false)

	t.Run("happy_path", func(t *testing.T) {
		ctx := context.Background()
		store := storage.NewStorageMock(t)
		srv := NewBoatsServer(store, nil)
		recorder := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/v1/suggest?param=builders&prefix=foo", nil)

		store.FindBuildersByPrefixMock.Expect(ctx, "foo", suggestListLimit).Return([]string{"fooa", "foob"}, nil)
		srv.HandleSuggest(ctx, recorder, req)
		assert.Equal(http.StatusOK, recorder.Code)
		assert.Equal(`["fooa","foob"]`+"\n", recorder.Body.String())
	})

	t.Run("unexpected_query", func(t *testing.T) {
		ctx := context.Background()
		store := storage.NewStorageMock(t)
		srv := NewBoatsServer(store, nil)
		recorder := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/v1/suggest?param=charters&prefix=foo", nil)

		srv.HandleSuggest(ctx, recorder, req)
		assert.Equal(http.StatusBadRequest, recorder.Code)
	})

	t.Run("db_fail", func(t *testing.T) {
		ctx := context.Background()
		store := storage.NewStorageMock(t)
		srv := NewBoatsServer(store, nil)
		recorder := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/v1/suggest?param=builders&prefix=foo", nil)

		store.FindBuildersByPrefixMock.Return(nil, errors.New("db failed"))
		srv.HandleSuggest(ctx, recorder, req)
		assert.Equal(http.StatusInternalServerError, recorder.Code)
	})
}
