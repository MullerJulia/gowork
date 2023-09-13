package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mikenai/gowork/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestUsers_Create(t *testing.T) {
	type fields struct {
		user UsersService
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
		wantBody []byte
	}{
		{
			name: "success",
			fields: fields{
				user: &UsersServiceMock{
					CreateFunc: func(ctx context.Context, name string) (models.User, error) {
						assert.Equal(t, "mike", name)
						return models.User{Name: name, ID: "1"}, nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name": "mike"}`)),
			},
			wantCode: http.StatusOK,
			wantBody: []byte(`{"id":"1","name":"mike"}` + "\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUsers(tt.fields.user)
			u.Routes().ServeHTTP(tt.args.w, tt.args.r)

			assert.Equal(t, tt.wantCode, tt.args.w.Code)
			assert.Equal(t, tt.args.w.Body.Bytes(), tt.wantBody, "unxpected body")
		})
	}
}

func TestUsers_UpdateUserName(t *testing.T) {
	type fields struct {
		user UsersService
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
		wantBody []byte
	}{
		{
			name: "success",
			fields: fields{
				user: &UsersServiceMock{
					UpdateUserNameFunc: func(ctx context.Context, id, name string) error {
						assert.Equal(t, "1", id)
						assert.Equal(t, "name", name)
						return nil
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPatch, "/1", strings.NewReader(`{"name": "name"}`)),
			},
			wantCode: http.StatusOK,
			wantBody: []byte(`{"name":"name"}` + "\n"),
		},
		{
			name: "Internal Server Error",
			fields: fields{
				user: &UsersServiceMock{
					UpdateUserNameFunc: func(ctx context.Context, id, name string) error {
						return errors.New("error")
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPatch, "/1", strings.NewReader(`{"name": "name"}`)),
			},
			wantCode: http.StatusInternalServerError,
			wantBody: []byte(`Internal Server Error` + "\n"),
		},
		{
			name: "Not Found",
			fields: fields{
				user: &UsersServiceMock{
					UpdateUserNameFunc: func(ctx context.Context, id, name string) error {
						return models.NotFoundErr
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPatch, "/1", strings.NewReader(`{"name": "name"}`)),
			},
			wantCode: http.StatusNotFound,
			wantBody: []byte(`Not Found` + "\n"),
		},
		{
			name: "Bad JSON",
			fields: fields{
				user: &UsersServiceMock{
					UpdateUserNameFunc: func(ctx context.Context, id, name string) error {
						return errors.New("error")
					},
				},
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPatch, "/1", strings.NewReader(`{"name": "name"`)),
			},
			wantCode: http.StatusInternalServerError,
			wantBody: []byte(`Internal Server Error` + "\n"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUsers(tt.fields.user)
			u.Routes().ServeHTTP(tt.args.w, tt.args.r)

			assert.Equal(t, tt.wantCode, tt.args.w.Code)
			assert.Equal(t, tt.wantBody, tt.args.w.Body.Bytes(), "unxpected body")
		})
	}
}
