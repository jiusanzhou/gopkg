package gopkg

import "testing"

func TestPackage_Init(t *testing.T) {
	type args struct {
		gopkg *Gopkg
	}
	tests := []struct {
		name    string
		pkg     *Package
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if err := tt.pkg.Init(tt.args.gopkg); (err != nil) != tt.wantErr {
			t.Errorf("%q. Package.Init() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
