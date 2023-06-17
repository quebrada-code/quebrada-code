package config

import (
	"os"
	"quebrada_api/internal/config"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	type env struct {
		dbUser     string
		dbName     string
		dbPassword string
		dbHost     string
	}

	type args struct {
		path string
		env  env
	}

	setEnv := func(env env) {
		_ = os.Setenv("DB_USER", env.dbUser)
		_ = os.Setenv("DB_PASSWORD", env.dbPassword)
		_ = os.Setenv("DB_NAME", env.dbName)
		_ = os.Setenv("DB_HOST", env.dbHost)
		_ = os.Setenv("APP_ENV", "local")
	}

	tests := []struct {
		name    string
		args    args
		want    *config.Config
		wantErr bool
	}{
		{
			name: "test config",
			args: args{
				path: "fixtures",
				env: env{
					dbUser:     "postgres",
					dbPassword: "postgres",
					dbName:     "postgres",
					dbHost:     "localhost",
				},
			},
			want: &config.Config{
				Environment: "local",
				Database: config.DatabaseConfig{
					ConnectionString: "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(tt.args.env)

			got, err := config.Init(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() got = %v, want %v", got, tt.want)
			}
		})
	}
}
