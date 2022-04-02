package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// User Structure 의 JSON형식, `json key값 설정`
type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreateAt  time.Time `json:"create_at"`
}

type fooHandler struct{}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Golang World!!")
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	// request body를 인자로 넣어줌  r.body는 ioReader객체 타입을 갖고있어서 NewDecoder의 ioReader인자로 넣을 수 있다.
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	// 현재 시간값을 넣어줌
	user.CreateAt = time.Now()

	// 데이터를 json형태로 변환시켜줌, return data, err
	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	// ?name=주원 이라는 쿼리스트링 생성
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	// fmt.Fprint(w, "Hello bar")
	fmt.Fprintf(w, "Hello %s!", name)
}

// mux를 만들어내는 함수

func NewHttpHandler() http.Handler {
	// mux 라우터 생성
	mux := http.NewServeMux()

	// func으로 handler를 직접 구현
	mux.HandleFunc("/", indexHandler)

	// mux : 경로 별로 handler를 만드는것 => mux
	mux.HandleFunc("/bar", barHandler)

	// 인스턴스 형태로 handler를 불러오기
	mux.Handle("/foo", &fooHandler{})
	return mux
}
