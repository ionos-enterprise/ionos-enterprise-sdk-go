package profitbricks

import "testing"

func Test_cache_Add(t *testing.T) {
	type fields struct {
		typeCache map[string]map[string]cachable
	}
	type args struct {
		obj cachable
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cache{
				typeCache: tt.fields.typeCache,
			}
		})
	}
}