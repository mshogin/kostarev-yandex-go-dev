package utils

import "testing"

func TestRandomString(t *testing.T) {
	str := RandomString()

	tests := []struct {
		name         string
		randomString string
		want         string
	}{
		{
			name:         "good test",
			randomString: str,
			want:         str,
		},
		//{ TODO не знаю как обработать bad test
		//	name:         "bad test",
		//	randomString: "zxcvbnmk",
		//	want:         str,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.randomString != tt.want {
				t.Error("random string is not want")
			}
		})
	}
}
