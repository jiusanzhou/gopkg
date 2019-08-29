package gopkg

import (
	"reflect"
	"testing"
)

func TestFile_Write(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name    string
		fl      *File
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.fl.Write(tt.args.target); (err != nil) != tt.wantErr {
			t.Errorf("%q. File.Write() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestTemplate_Render(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		tpl     *Template
		args    args
		want    *File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		got, err := tt.tpl.Render(tt.args.data)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Template.Render() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Template.Render() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
