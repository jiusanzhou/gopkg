package gopkg

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGopkg_NewHandler(t *testing.T) {
	type args struct {
		allowWild bool
		next      http.Handler
	}
	tests := []struct {
		name string
		pkg  *Gopkg
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := tt.pkg.NewHandler(tt.args.allowWild, tt.args.next); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Gopkg.NewHandler() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGopkg_Generate(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name    string
		pkg     *Gopkg
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.pkg.Generate(tt.args.target); (err != nil) != tt.wantErr {
			t.Errorf("%q. Gopkg.Generate() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
