package crypto

import "testing"

func TestGetSha256String(t *testing.T) {
	type args struct {
		source string
		salt   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{source: "123456", salt: "mI19dU8K"}, "45151a66c94c3b6c336fabb156be97aaea4708e286e1c7dcd31d078fca904036"},
		{"2", args{source: "password", salt: "mI19dU8K"}, "699898fc8043375c1bd052d9aa1d73116dc4062409a5e4d5591d3b27bcc5bc5f"},
		{"3", args{source: "KbX#^CZydelzNzxh", salt: "8RikVvx32GXYlUQZ"}, "a71397d58c79d8dc07324bcfbe28e9c9d2513254781dc27a891e14456eaf55b9"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSha256String(tt.args.source, tt.args.salt); got != tt.want {
				t.Errorf("GetSha256String() = %v, want %v", got, tt.want)
			}
		})
	}
}
