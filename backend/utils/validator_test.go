package utils

import (
	"errors"
	"testing"
)

func TestValidatePage(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		hasError bool
	}{
		{"空文字列はデフォルト値1", "", 1, false},
		{"正常な値", "5", 5, false},
		{"最大値", "1000", 1000, false},
		{"最小値", "1", 1, false},
		{"ゼロは無効", "0", 0, true},
		{"負の値は無効", "-1", 0, true},
		{"上限超過は無効", "1001", 0, true},
		{"文字列は無効", "abc", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ValidatePage(tt.input)
			
			if tt.hasError {
				if err == nil {
					t.Errorf("エラーが期待されていましたが、エラーがありませんでした")
				}
			} else {
				if err != nil {
					t.Errorf("エラーが期待されていませんでしたが、エラーが発生しました: %v", err)
				}
				if result != tt.expected {
					t.Errorf("期待値 %d, 実際の値 %d", tt.expected, result)
				}
			}
		})
	}
}

func TestValidateMovieID(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		hasError bool
	}{
		{"正常な値", "12345", 12345, false},
		{"最小値", "1", 1, false},
		{"最大値", "999999", 999999, false},
		{"空文字列は無効", "", 0, true},
		{"ゼロは無効", "0", 0, true},
		{"負の値は無効", "-1", 0, true},
		{"上限超過は無効", "1000000", 0, true},
		{"スラッシュ含む", "123/456", 0, true},
		{"文字列は無効", "abc", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ValidateMovieID(tt.input)
			
			if tt.hasError {
				if err == nil {
					t.Errorf("エラーが期待されていましたが、エラーがありませんでした")
				}
			} else {
				if err != nil {
					t.Errorf("エラーが期待されていませんでしたが、エラーが発生しました: %v", err)
				}
				if result != tt.expected {
					t.Errorf("期待値 %d, 実際の値 %d", tt.expected, result)
				}
			}
		})
	}
}

func TestValidateSearchQuery(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		hasError bool
	}{
		{"正常な検索クエリ", "batman", false},
		{"日本語の検索クエリ", "バットマン", false},
		{"数字を含む検索クエリ", "movie 2023", false},
		{"空文字列は無効", "", true},
		{"長すぎるクエリ", string(make([]rune, 101)), true},
		{"スクリプトタグ", "<script>alert('xss')</script>", true},
		{"JavaScriptプロトコル", "javascript:alert('xss')", true},
		{"onloadイベント", "onload=alert('xss')", true},
		{"大文字小文字混在のスクリプト", "<SCRIPT>alert('xss')</SCRIPT>", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateSearchQuery(tt.input)
			
			if tt.hasError {
				if err == nil {
					t.Errorf("エラーが期待されていましたが、エラーがありませんでした")
				}
			} else {
				if err != nil {
					t.Errorf("エラーが期待されていませんでしたが、エラーが発生しました: %v", err)
				}
			}
		})
	}
}

func TestValidateGenreID(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		hasError bool
	}{
		{"正常な値", "28", 28, false},
		{"最小値", "1", 1, false},
		{"最大値", "999", 999, false},
		{"空文字列は無効", "", 0, true},
		{"ゼロは無効", "0", 0, true},
		{"負の値は無効", "-1", 0, true},
		{"上限超過は無効", "1000", 0, true},
		{"文字列は無効", "abc", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ValidateGenreID(tt.input)
			
			if tt.hasError {
				if err == nil {
					t.Errorf("エラーが期待されていましたが、エラーがありませんでした")
				}
			} else {
				if err != nil {
					t.Errorf("エラーが期待されていませんでしたが、エラーが発生しました: %v", err)
				}
				if result != tt.expected {
					t.Errorf("期待値 %d, 実際の値 %d", tt.expected, result)
				}
			}
		})
	}
}

func TestIsValidationError(t *testing.T) {
	validationErr := &ValidationError{Field: "test", Message: "test error"}
	normalErr := errors.New("normal error")

	if !IsValidationError(validationErr) {
		t.Errorf("ValidationErrorが正しく識別されませんでした")
	}

	if IsValidationError(normalErr) {
		t.Errorf("通常のエラーがValidationErrorとして誤識別されました")
	}
}