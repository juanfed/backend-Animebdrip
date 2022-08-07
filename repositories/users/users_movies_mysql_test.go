package repositories

import (
	"database/sql"
	"reflect"
	"testing"
	"v1/models"
)

func TestUserRepositories_Get(t *testing.T) {
	type fields struct {
		database *sql.DB
	}
	type args struct {
		id int
	}
	type test struct {
		name    string
		fields  fields
		args    args
		want    models.Movie
		wantErr bool
	}

	tests := []test{
		func() test {
			return test{}
		}(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sq := &UserMovieRepositories{
				database: tt.fields.database,
			}
			got, err := sq.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
