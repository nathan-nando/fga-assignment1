package main

import (
	"fmt"
	"main/helper"
	"os"
	"strconv"
)

func main() {
	args, err := strconv.Atoi(os.Args[1])
	if err != nil {
		args = 0
		panic(err)
	}
	if helper.Data[args] == (helper.Biodata{}) {
		fmt.Println("belum ada data")
	} else {
		fmt.Println(helper.Data[args].GetBiodata())
	}
}
