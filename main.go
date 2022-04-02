package main

import (
	"net/http"

	"github.com/Juwon-Yun/goRest/app"
)

func main() {
	http.ListenAndServe(":4000", app.NewHttpHandler())
}

// 보낸 데이터
// {
// "first_name" : "juwon",
// "last_name" : "yun",
// "email" : "juwon@naver.com"
// }

// 받은 데이터 = {"first_name":"juwon","last_name":"yun","email":"juwon@naver.com","create_at":"2022-03-31T20:15:16.910464+09:00"}

// header를 json으로 한 응답 값
// {
//     "first_name": "juwon",
//     "last_name": "yun",
//     "email": "juwon@naver.com",
//     "create_at": "2022-03-31T20:17:56.427465+09:00"
// }
