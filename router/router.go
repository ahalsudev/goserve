package router

import (
	"math"
	"net/http"
	"strconv"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ok"))
	})

	mux.HandleFunc("/calcPrimes/", func(w http.ResponseWriter, r *http.Request) {
		primes := calcPrimes(10, 10000)
		w.WriteHeader(http.StatusOK)
		bs := []byte(strconv.Itoa(primes))
		w.Write(bs)
	})

	return mux
}

func calcPrimes(s int, e int) int {
	sum := 0
	for i := s; i <= e; i++ {
		prime := isPrime(i)
		if prime {
			sum++
		}
	}
	return sum
}

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
