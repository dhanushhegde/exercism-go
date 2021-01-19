//Package contains function to calculate square of sums , sum of squares and their difference
//for first n number
package diffsquares

//brute force
func SquareOfSum(n int) int {
	// sum := 0
	// for i := 0; i <= n; i++ {
	// 	sum += i
	// }
	// squareOfSum := sum * sum
	// return squareOfSum
	return n * (n + 1) / 2 * n * (n + 1) / 2
}

func SumOfSquares(n int) int {
	// sumOfSquares := 0
	// for i := 0; i <= n; i++ {
	// 	sumOfSquares += i * i
	// }
	// return sumOfSquares
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	// sum := 0
	// sumOfSquares := 0
	// for i := 0; i <= n; i++ {
	// 	sum += i
	// 	sumOfSquares += i * i
	// }
	// squareOfSum := sum * sum
	// return squareOfSum - sumOfSquares

	return (n * (n + 1) / 2 * n * (n + 1) / 2) - (n * (n + 1) * (2*n + 1) / 6)
}

//sum of first n natural numbers = n*(n+1)/2
//sum of square of first n natural number = (n*(n+1)* (2n+1))/6
