package middleware

import (
	"fmt"
	"net/http"
)

// GlobalMiddle : with jwt on the verification, no jwt on the verification
type GlobalMiddle struct {
}

func NewCommonJwtAuthMiddleware() *GlobalMiddle {
	return &GlobalMiddle{}
}

func (m *GlobalMiddle) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("global before")
		next(w, r)
		fmt.Println("global end")
	}
}
