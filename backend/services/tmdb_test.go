package services

import (
	"context"
	"net/http"
	"os"
	"sync"
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

// TestHTTPClientSingleton - シングルトンHTTPクライアントのテスト
func TestHTTPClientSingleton(t *testing.T) {
	// 同じインスタンスが返されることを確認
	client1 := getHTTPClient()
	client2 := getHTTPClient()
	
	if client1 != client2 {
		t.Error("Expected same HTTP client instance, but got different instances")
	}
	
	// クライアントの設定確認
	if client1.Timeout != 10*time.Second {
		t.Errorf("Expected timeout 10s, got %v", client1.Timeout)
	}
	
	// Transportが設定されていることを確認
	if client1.Transport == nil {
		t.Error("Expected Transport to be set")
	}
}

// TestPingHTTPClient - Ping用HTTPクライアントのテスト
func TestPingHTTPClient(t *testing.T) {
	pingClient := getPingHTTPClient()
	
	// タイムアウトが5秒に設定されていることを確認
	if pingClient.Timeout != 5*time.Second {
		t.Errorf("Expected ping client timeout 5s, got %v", pingClient.Timeout)
	}
	
	// 通常クライアントとTransportが共有されていることを確認
	normalClient := getHTTPClient()
	if pingClient.Transport != normalClient.Transport {
		t.Error("Expected ping client to share Transport with normal client")
	}
}

// TestHTTPClientConcurrency - 並行アクセスでのシングルトンテスト
func TestHTTPClientConcurrency(t *testing.T) {
	const goroutines = 100
	clients := make([]*http.Client, goroutines)
	
	// 並行してHTTPクライアントを取得
	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			clients[index] = getHTTPClient()
		}(i)
	}
	wg.Wait()
	
	// 全て同じインスタンスであることを確認
	firstClient := clients[0]
	for i := 1; i < goroutines; i++ {
		if clients[i] != firstClient {
			t.Errorf("Client %d is different from first client", i)
		}
	}
}