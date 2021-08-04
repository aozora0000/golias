package main

import (
	"testing"
)

func TestCipher_Decrypt(t *testing.T) {
	type fields struct {
		key []byte
	}
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "正常に動作",
			fields:  fields{key: []byte("1234567890123456")},
			args:    args{t: "TxqLHl3f5gwJMkga8EsFITSsfxeSUnSYIkVop6LNlIVe8oh4rHyVjMt2UNozJ2V5OLfm5Q=="},
			want:    "rawstringrawstringrawstring",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cipher{
				key: tt.fields.key,
			}
			got, err := c.Decrypt(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCipher_Encrypt(t *testing.T) {
	type fields struct {
		key []byte
	}
	type args struct {
		text []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "正常に動作",
			fields:  fields{key: []byte("1234567890123456")},
			args:    args{text: []byte("rawstringrawstringrawstring")},
			want:    "TxqLHl3f5gwJMkga8EsFITSsfxeSUnSYIkVop6LNlIVe8oh4rHyVjMt2UNozJ2V5OLfm5Q==",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Cipher{
				key: tt.fields.key,
			}
			_, err := c.Encrypt(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
