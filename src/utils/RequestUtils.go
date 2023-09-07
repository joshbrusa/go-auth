package utils

import (
	"net/http"
)

func CheckRequestBody(r *http.Request, expectedItems []string) bool {
	
	println(r.Body)


	// err := json.Unmarshal(r.Body, v)

	// if err != nil {
	// 	panic(err.Error())
	// }


	return true
}