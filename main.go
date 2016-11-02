package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type flag struct {
	idiot  bool
	stupid bool
	dumb   bool
}

func main() {
	run(os.Args[1:])
}

func judge(input string) (flag, error) {
	// 3の倍数のとき "idiot"
	// 3のつく数のとき "stupid"
	// 3の倍数でかつ 3 のつく数のとき "dumb"
	// それぞれ判定し、フラグを立てる

	var nabeatsu flag

	num, err := strconv.Atoi(input)
	if err != nil {
		return nabeatsu, err
	}

	if num%3 == 0 {
		nabeatsu.idiot = true
	}

	stupidNum := regexp.MustCompile("[0-9]*3+[0-9]*")
	if stupidNum.MatchString(input) {
		nabeatsu.stupid = true
	}

	if nabeatsu.idiot && nabeatsu.stupid {
		nabeatsu.dumb = true
	}

	return nabeatsu, err
}

func run(args []string) {
	for _, v := range args {
		nabeatsu, err := judge(v)
		if err != nil {
			fmt.Println("invalid")
			continue
		}
		if nabeatsu.dumb {
			fmt.Println("dumb")
			continue
		}
		if nabeatsu.idiot {
			fmt.Println("idiot")
			continue
		}
		if nabeatsu.stupid {
			fmt.Println("stupid")
			continue
		}
		fmt.Println("smart")
	}
}
