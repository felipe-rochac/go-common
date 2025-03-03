package crypto

import (
	"testing"
)

func TestHmac(t *testing.T) {
	tests := []struct {
		name   string
		hmacFn func(string, string) string
		want   string
	}{
		{
			name:   "HmacSha256",
			hmacFn: HmacSha256,
			want:   "9307b3b915efb5171ff14d8cb55fbcc798c6c0ef1456d66ded1a6aa723a58b7b",
		},
		{
			name:   "HmacSha1",
			hmacFn: HmacSha1,
			want:   "b34ceac4516ff23a143e61d79d0fa7a4fbe5f266",
		},
		{
			name:   "HmacSha512",
			hmacFn: HmacSha512,
			want:   "ff06ab36757777815c008d32c8e14a705b4e7bf310351a06a23b612dc4c7433e7757d20525a5593b71020ea2ee162d2311b247e9855862b270122419652c0c92",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.hmacFn("key", "hello")

			if got != tt.want {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.want)
			}

		})
	}
}

func TestMD5(t *testing.T) {
	md5 := MD5("hello")
	expect := "5d41402abc4b2a76b9719d911017c592"
	if md5 != expect {
		t.Error("MD5 Error")
	}
}

func TestSha(t *testing.T) {
	hashTests := []struct {
		name   string
		hashFn func(string) string
		input  string
		expect string
	}{
		{
			name:   "Sha1",
			hashFn: Sha1,
			input:  "hello",
			expect: "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d",
		},
		{
			name:   "Sha256",
			hashFn: Sha256,
			input:  "hello",
			expect: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
		{
			name:   "Sha512",
			hashFn: Sha512,
			input:  "hello",
			expect: "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043",
		},
		{
			name:   "Sha3",
			hashFn: Sha3,
			input:  "hello",
			expect: "3338be694f50c5f338814986cdf0686453a888b84f424d792af4b9202398f392",
		},
	}

	for _, tt := range hashTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.hashFn(tt.input)
			if got != tt.expect {
				t.Errorf("%s() = %v, want %v", tt.name, got, tt.expect)
			}
		})
	}
}
