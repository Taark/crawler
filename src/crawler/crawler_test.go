package crawler

import (
	"github.com/taark/crawler/src/models"
	"reflect"
	"testing"
)

func TestScan(t *testing.T) {


	type args struct {
		urls []string
	}
	tests := []struct {
		name string
		args args
		want []*models.Target
	}{
		{
			name: "success",
			args: args{
				urls: []string{"https://gmail.com", "https://google.com"},
			},
			want: []*models.Target{
				{
					Url:   "https://gmail.com",
					Title: "Gmail",
					Err:   "",
				},
				{
					Url:   "https://google.com",
					Title: "Google",
					Err:   "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Scan(tt.args.urls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getContent(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "simple test for yandex api",
			args:    args{
				url: "https://api.directory.yandex.net/v6/departments/",
			},
			want:    `{"message": "Our API requires authentication", "code": "authentication-error"}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getContent(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("getContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getContent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTitle(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "for example.com",
			args: args{url: "http://example.com/"},
			want: "Example Domain",
			wantErr: false,
		},
		{
			name:    "fail http://testtest.ci/",
			args:    args{url: "http://testtest.ci/"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTitle(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTitle() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "correct: https://gmail.com ",
			args:    args{
				"https://gmail.com",
			},
			wantErr: false,
		},
		{
			name:    "wrong: gmail.com ",
			args:    args{
				"gmail.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateUrl(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("validateUrl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}