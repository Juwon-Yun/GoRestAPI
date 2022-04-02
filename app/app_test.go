package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// test 함수 선언, 양식이 정해져 있다.
func TestIndexPathHandler(t *testing.T) {
	// 가짜 http 리더 생성
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	indexHandler(res, req)

	if res.Code != http.StatusBadRequest {
		// 실패하면 프로그램 종료
		t.Fatal("Failed! ", res.Code)
	}
}
