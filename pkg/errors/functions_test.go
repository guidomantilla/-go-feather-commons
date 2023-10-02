package errors

import "testing"

func TestErrJoin(t *testing.T) {
	type args struct {
		errs []error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ErrJoin(tt.args.errs...); (err != nil) != tt.wantErr {
				t.Errorf("ErrJoin() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
