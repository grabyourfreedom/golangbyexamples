package main

import "fmt"

func isPrime(number uint) bool {
	prime := true
	var i uint = 0

	if number == 1 {
		return false
	}

	for i = 2; i < number; i++ {
		if number % i == 0 {
			prime = false
			break
		}
	}

	return prime
}

/*func getPrimes(start, end uint) uint[] {
	var primes [1]uint
	TODO: After comfortable with Arrays and slices
	return primes

}*/

func getNextPrime(number uint) uint {
	var i = number
	for ; true; i++ {
		if isPrime(i) {
			break
		}
	}

	return i
}

func main()  {
	fmt.Println(isPrime(1))
	fmt.Println(isPrime(2))
	fmt.Println(isPrime(3))
	fmt.Println(isPrime(25))
	fmt.Println(isPrime(31))
	fmt.Println(isPrime(113))
	fmt.Println(getNextPrime(111))
}