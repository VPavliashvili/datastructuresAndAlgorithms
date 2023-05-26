package main

func main() {
	println(foo(5))
}

func foo(n int) int {
    //base case
	if n == 0 {
		return 0
	}

    it := n + foo(n - 1)
    println(n)
	return it
}
