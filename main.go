package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	str, _ := rd.ReadString('\n')
	words := strings.Split(str, " ")
	fmt.Println(len(words))

}
