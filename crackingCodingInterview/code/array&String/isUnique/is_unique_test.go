package isUnique

import "testing"

func TestIsUnique(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "normal",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUnique(); got != tt.want {
				t.Errorf("IsUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
