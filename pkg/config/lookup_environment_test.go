package config

import "testing"

func TestEnvironmentLookup_Lookup(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name     string
		lookuper *EnvironmentLookup
		args     args
		want     string
		want1    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.lookuper.Lookup(tt.args.key)
			if got != tt.want {
				t.Errorf("EnvironmentLookup.Lookup() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("EnvironmentLookup.Lookup() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
