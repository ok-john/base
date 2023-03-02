package main

import (
	"math/big"
	"os"
	"io"
	"strconv"
	"fmt"
)

func usage() {
	fmt.Println("[USAGE]")
	fmt.Println("base, by ok-john (github.com/ok-john/base)")
	fmt.Println("base converts input from STDIN to a provided BASE.")
	fmt.Println("the resulting output is written to STDOUT.")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("for example:")
	fmt.Println("")
	fmt.Println("		echo ok-john|base 32")
	fmt.Println("")
	fmt.Println("or")
	fmt.Println("")
	fmt.Println("		echo ok-john|base 16")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("you can provide any base between 2-63.")
	fmt.Println("the default base is base 10")
}

func main() {
	var base int64 = 10
	var err error
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" {
			usage()
			os.Exit(1)
		}
		base, err = strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			panic(err)	
		}
	}

	if (base < 2 || base > 62) {
		fmt.Println("must use a base between 2-62")
		os.Exit(1)
	}

	buff, err :=io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	
	fmt.Println((&big.Int{}).SetBytes(buff).Text(int(base)))


}
