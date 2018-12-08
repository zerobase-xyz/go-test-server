package main

import (
	"fmt"
	"net/http"
)

type String string

func factorize(number int) {
	var a []int
	for i := 1; i < number+1; i++ {
		if number%i == 0 {
			a = append(a, i)
		}
	}
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	numbers := 214748363
	factorize(numbers)
	fmt.Fprint(w, s)
}

func main() {
	http.Handle("/", String("Success!"))
	http.ListenAndServe(":8000", nil)
}
