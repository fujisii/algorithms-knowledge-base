package main

import "fmt"

func main() {
	uf := NewUnionFind(5)
	uf.Union(0, 1)
	uf.Union(2, 3)
	fmt.Println("connected(0,1):", uf.Connected(0, 1))
	fmt.Println("connected(0,2):", uf.Connected(0, 2))
	fmt.Println("find(1):", uf.Find(1))
}
