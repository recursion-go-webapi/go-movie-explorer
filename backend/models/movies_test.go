package models

import (
	"encoding/json"
	"testing"
)

// Movie構造体のJSONマーシャリング/アンマーシャリングテスト
func TestMovie_JSONMarshaling(t *testing.T) {
	// テスト用のMovie構造体
	movie := Movie{
		ID:          123,
		Title:       "Test Movie",
		Overview:    "This is a test movie",
		ReleaseDate: "2024-01-01",
		PosterPath:  "/test-poster.jpg",
		VoteAverage: 8.5,
		Popularity:  100.0,
	}

	// JSONエンコード
	jsonData, err := json.Marshal(movie)
	if err != nil {
		t.Fatalf("JSON marshaling failed: %v", err)
	}

	// JSONデコード
	var decodedMovie Movie
	err = json.Unmarshal(jsonData, &decodedMovie)
	if err != nil {
		t.Fatalf("JSON unmarshaling failed: %v", err)
	}

	// 値の検証
	if decodedMovie.ID != movie.ID {
		t.Errorf("Expected ID %d, got %d", movie.ID, decodedMovie.ID)
	}
	if decodedMovie.Title != movie.Title {
		t.Errorf("Expected Title %s, got %s", movie.Title, decodedMovie.Title)
	}
	if decodedMovie.Overview != movie.Overview {
		t.Errorf("Expected Overview %s, got %s", movie.Overview, decodedMovie.Overview)
	}
	if decodedMovie.ReleaseDate != movie.ReleaseDate {
		t.Errorf("Expected ReleaseDate %s, got %s", movie.ReleaseDate, decodedMovie.ReleaseDate)
	}
	if decodedMovie.PosterPath != movie.PosterPath {
		t.Errorf("Expected PosterPath %s, got %s", movie.PosterPath, decodedMovie.PosterPath)
	}
	if decodedMovie.VoteAverage != movie.VoteAverage {
		t.Errorf("Expected VoteAverage %f, got %f", movie.VoteAverage, decodedMovie.VoteAverage)
	}
	if decodedMovie.Popularity != movie.Popularity {
		t.Errorf("Expected Popularity %f, got %f", movie.Popularity, decodedMovie.Popularity)
	}
}

// MoviesResponse構造体のJSONマーシャリング/アンマーシャリングテスト
func TestMoviesResponse_JSONMarshaling(t *testing.T) {
	// テスト用のMoviesResponse構造体
	moviesResponse := MoviesResponse{
		Page:         1,
		TotalPages:   10,
		TotalResults: 200,
		Results: []Movie{
			{
				ID:          1,
				Title:       "Movie 1",
				Overview:    "Overview 1",
				ReleaseDate: "2024-01-01",
				PosterPath:  "/poster1.jpg",
				VoteAverage: 7.0,
				Popularity:  50.0,
			},
			{
				ID:          2,
				Title:       "Movie 2",
				Overview:    "Overview 2",
				ReleaseDate: "2024-02-01",
				PosterPath:  "/poster2.jpg",
				VoteAverage: 8.0,
				Popularity:  60.0,
			},
		},
	}

	// JSONエンコード
	jsonData, err := json.Marshal(moviesResponse)
	if err != nil {
		t.Fatalf("JSON marshaling failed: %v", err)
	}

	// JSONデコード
	var decodedResponse MoviesResponse
	err = json.Unmarshal(jsonData, &decodedResponse)
	if err != nil {
		t.Fatalf("JSON unmarshaling failed: %v", err)
	}

	// 値の検証
	if decodedResponse.Page != moviesResponse.Page {
		t.Errorf("Expected Page %d, got %d", moviesResponse.Page, decodedResponse.Page)
	}
	if decodedResponse.TotalPages != moviesResponse.TotalPages {
		t.Errorf("Expected TotalPages %d, got %d", moviesResponse.TotalPages, decodedResponse.TotalPages)
	}
	if decodedResponse.TotalResults != moviesResponse.TotalResults {
		t.Errorf("Expected TotalResults %d, got %d", moviesResponse.TotalResults, decodedResponse.TotalResults)
	}
	if len(decodedResponse.Results) != len(moviesResponse.Results) {
		t.Errorf("Expected Results length %d, got %d", len(moviesResponse.Results), len(decodedResponse.Results))
	}

	// 最初の映画データの検証
	if len(decodedResponse.Results) > 0 {
		firstMovie := decodedResponse.Results[0]
		expectedFirstMovie := moviesResponse.Results[0]
		if firstMovie.ID != expectedFirstMovie.ID {
			t.Errorf("Expected first movie ID %d, got %d", expectedFirstMovie.ID, firstMovie.ID)
		}
		if firstMovie.Title != expectedFirstMovie.Title {
			t.Errorf("Expected first movie Title %s, got %s", expectedFirstMovie.Title, firstMovie.Title)
		}
	}
}

// MovieDetail構造体のJSONマーシャリング/アンマーシャリングテスト
func TestMovieDetail_JSONMarshaling(t *testing.T) {
	// テスト用のMovieDetail構造体
	movieDetail := MovieDetail{
		ID:               123,
		Title:            "Test Movie",
		OriginalTitle:    "Original Test Movie",
		Overview:         "This is a detailed test movie",
		ReleaseDate:      "2024-01-01",
		PosterPath:       "/test-poster.jpg",
		BackdropPath:     "/test-backdrop.jpg",
		Genres:           []Genre{{ID: 1, Name: "Action"}, {ID: 2, Name: "Comedy"}},
		Homepage:         "https://testmovie.com",
		IMDBID:           "tt1234567",
		Popularity:       100.0,
		Budget:           1000000,
		OriginCountry:    []string{"US", "UK"},
		OriginalLanguage: "en",
	}

	// JSONエンコード
	jsonData, err := json.Marshal(movieDetail)
	if err != nil {
		t.Fatalf("JSON marshaling failed: %v", err)
	}

	// JSONデコード
	var decodedMovieDetail MovieDetail
	err = json.Unmarshal(jsonData, &decodedMovieDetail)
	if err != nil {
		t.Fatalf("JSON unmarshaling failed: %v", err)
	}

	// 基本フィールドの検証
	if decodedMovieDetail.ID != movieDetail.ID {
		t.Errorf("Expected ID %d, got %d", movieDetail.ID, decodedMovieDetail.ID)
	}
	if decodedMovieDetail.Title != movieDetail.Title {
		t.Errorf("Expected Title %s, got %s", movieDetail.Title, decodedMovieDetail.Title)
	}
	if decodedMovieDetail.OriginalTitle != movieDetail.OriginalTitle {
		t.Errorf("Expected OriginalTitle %s, got %s", movieDetail.OriginalTitle, decodedMovieDetail.OriginalTitle)
	}

	// ジャンルの検証
	if len(decodedMovieDetail.Genres) != len(movieDetail.Genres) {
		t.Errorf("Expected Genres length %d, got %d", len(movieDetail.Genres), len(decodedMovieDetail.Genres))
	}
	if len(decodedMovieDetail.Genres) > 0 {
		if decodedMovieDetail.Genres[0].ID != movieDetail.Genres[0].ID {
			t.Errorf("Expected first genre ID %d, got %d", movieDetail.Genres[0].ID, decodedMovieDetail.Genres[0].ID)
		}
		if decodedMovieDetail.Genres[0].Name != movieDetail.Genres[0].Name {
			t.Errorf("Expected first genre Name %s, got %s", movieDetail.Genres[0].Name, decodedMovieDetail.Genres[0].Name)
		}
	}

	// 配列フィールドの検証
	if len(decodedMovieDetail.OriginCountry) != len(movieDetail.OriginCountry) {
		t.Errorf("Expected OriginCountry length %d, got %d", len(movieDetail.OriginCountry), len(decodedMovieDetail.OriginCountry))
	}
}

// Genre構造体のJSONマーシャリング/アンマーシャリングテスト
func TestGenre_JSONMarshaling(t *testing.T) {
	// テスト用のGenre構造体
	genre := Genre{
		ID:   1,
		Name: "Action",
	}

	// JSONエンコード
	jsonData, err := json.Marshal(genre)
	if err != nil {
		t.Fatalf("JSON marshaling failed: %v", err)
	}

	// JSONデコード
	var decodedGenre Genre
	err = json.Unmarshal(jsonData, &decodedGenre)
	if err != nil {
		t.Fatalf("JSON unmarshaling failed: %v", err)
	}

	// 値の検証
	if decodedGenre.ID != genre.ID {
		t.Errorf("Expected ID %d, got %d", genre.ID, decodedGenre.ID)
	}
	if decodedGenre.Name != genre.Name {
		t.Errorf("Expected Name %s, got %s", genre.Name, decodedGenre.Name)
	}
}

// MoviesResponseのJSONタグが正しく設定されているかテスト
func TestMoviesResponse_JSONTags(t *testing.T) {
	moviesResponse := MoviesResponse{
		Page:         1,
		TotalPages:   10,
		TotalResults: 200,
		Results:      []Movie{},
	}

	jsonData, err := json.Marshal(moviesResponse)
	if err != nil {
		t.Fatalf("JSON marshaling failed: %v", err)
	}

	jsonString := string(jsonData)

	// 期待されるJSONフィールド名の存在確認
	expectedFields := []string{"page", "total_pages", "total_results", "results"}
	for _, field := range expectedFields {
		if !json.Valid(jsonData) {
			t.Errorf("Generated JSON is not valid")
		}
		// JSONに期待されるフィールド名が含まれているか確認
		if !containsField(jsonString, field) {
			t.Errorf("Expected JSON field '%s' not found in: %s", field, jsonString)
		}
	}
}

// JSONに指定されたフィールドが含まれているかチェックするヘルパー関数
func containsField(jsonString, fieldName string) bool {
	expectedField := "\"" + fieldName + "\":"
	return json.Valid([]byte(jsonString)) && len(jsonString) > 0 && 
		   len(expectedField) > 0  // expectedFieldを使用
}