package users

import (
	"context"
	"reflect"
	"testing"

	"github.com/mikenai/gowork/internal/models"
)

func TestService_Create(t *testing.T) {
	type fields struct {
		repo Repositry
	}
	type args struct {
		name        string
		phoneNumber string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.User
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				repo: &RepositryMock{
					CreateFunc: func(ctx context.Context, name string, phoneNumber string) (models.User, error) {
						return models.User{ID: "1", Name: name, PhoneNumber: "123"}, nil
					},
				},
			},
			args: args{
				name:        "name",
				phoneNumber: "123",
			},
			want:    models.User{ID: "1", Name: "name", PhoneNumber: "123"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}
			got, err := s.Create(context.Background(), tt.args.name, tt.args.phoneNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
