package main

import (
	"testing"
)

func Test_pickWelcomeMessage(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pickWelcomeMessage()
			if (err != nil) != tt.wantErr {
				t.Errorf("pickWelcomeMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("pickWelcomeMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pickKeywordMessage(t *testing.T) {
	t.SkipNow()
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"", args{""}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pickKeywordMessage(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("pickKeywordMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("pickKeywordMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
