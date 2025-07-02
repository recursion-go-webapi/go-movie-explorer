package services

import (
	"context"
	"os"
	"testing"
	"time"
)

// TestGetMoviesFromTMDB_NoAPIKey - APIキーなしのテスト
func TestGetMoviesFromTMDB_NoAPIKey(t *testing.T) {
	// APIキーを一時的に削除
	originalKey := os.Getenv("TMDB_API_KEY")
	os.Unsetenv("TMDB_API_KEY")
	defer func() {
		if originalKey != "" {
			os.Setenv("TMDB_API_KEY", originalKey)
		}
	}()

	_, err := GetMoviesFromTMDB(1)
	if err == nil {
		t.Error("Expected error when TMDB_API_KEY is not set")
	}
	expectedMsg := "TMDB_API_KEYが設定されていません"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

// TestGetMovieDetailFromTMDB_NoAPIKey - 映画詳細取得のAPIキーなしテスト
func TestGetMovieDetailFromTMDB_NoAPIKey(t *testing.T) {
	// APIキーを一時的に削除
	originalKey := os.Getenv("TMDB_API_KEY")
	os.Unsetenv("TMDB_API_KEY")
	defer func() {
		if originalKey != "" {
			os.Setenv("TMDB_API_KEY", originalKey)
		}
	}()

	_, err := GetMovieDetailFromTMDB(123)
	if err == nil {
		t.Error("Expected error when TMDB_API_KEY is not set")
	}
	expectedMsg := "TMDB_API_KEYが設定されていません"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

// TestSearchMoviesFromTMDB_EmptyQuery - 映画検索の空クエリテスト
func TestSearchMoviesFromTMDB_EmptyQuery(t *testing.T) {
	// APIキーを設定
	os.Setenv("TMDB_API_KEY", "test-key")
	defer os.Unsetenv("TMDB_API_KEY")

	_, err := SearchMoviesFromTMDB("", 1)
	if err == nil {
		t.Error("Expected error when query is empty")
	}
	expectedMsg := "検索クエリが指定されていません"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

// TestSearchMoviesFromTMDB_NoAPIKey - 映画検索のAPIキーなしテスト
func TestSearchMoviesFromTMDB_NoAPIKey(t *testing.T) {
	// APIキーを一時的に削除
	originalKey := os.Getenv("TMDB_API_KEY")
	os.Unsetenv("TMDB_API_KEY")
	defer func() {
		if originalKey != "" {
			os.Setenv("TMDB_API_KEY", originalKey)
		}
	}()

	_, err := SearchMoviesFromTMDB("test", 1)
	if err == nil {
		t.Error("Expected error when TMDB_API_KEY is not set")
	}
	expectedMsg := "TMDB_API_KEYが設定されていません"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}
}

// TestTmdbPinger - TMDBヘルスチェックのテスト
func TestTmdbPinger(t *testing.T) {
	pinger := &TmdbPinger{}

	// Name() メソッドのテスト
	if pinger.Name() != "TMDB" {
		t.Errorf("Expected name 'TMDB', got '%s'", pinger.Name())
	}

	// Ping() メソッドのテスト（APIキーなし）
	originalKey := os.Getenv("TMDB_API_KEY")
	os.Unsetenv("TMDB_API_KEY")
	defer func() {
		if originalKey != "" {
			os.Setenv("TMDB_API_KEY", originalKey)
		}
	}()

	ctx := context.Background()
	err := pinger.Ping(ctx)
	if err == nil {
		t.Error("Expected error when TMDB_API_KEY is not set")
	}
	expectedMsg := "TMDB_API_KEYが設定されていません"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
	}

	// タイムアウトのテスト
	os.Setenv("TMDB_API_KEY", "test-key")
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	err = pinger.Ping(ctxWithTimeout)
	if err == nil {
		t.Error("Expected timeout error")
	}
}

// TestGetTMDBApiKey - APIキー取得のテスト
func TestGetTMDBApiKey(t *testing.T) {
	// APIキーが設定されていない場合
	originalKey := os.Getenv("TMDB_API_KEY")
	os.Unsetenv("TMDB_API_KEY")
	defer func() {
		if originalKey != "" {
			os.Setenv("TMDB_API_KEY", originalKey)
		}
	}()

	key := GetTMDBApiKey()
	if key != "" {
		t.Errorf("Expected empty string when TMDB_API_KEY is not set, got '%s'", key)
	}

	// APIキーが設定されている場合
	expectedKey := "test-api-key-12345"
	os.Setenv("TMDB_API_KEY", expectedKey)
	
	key = GetTMDBApiKey()
	if key != expectedKey {
		t.Errorf("Expected '%s', got '%s'", expectedKey, key)
	}
}