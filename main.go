package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func CUT(word string) (string, bool) {
	ans := ""
	ok := false
	for i := 0; i < len(word); i++ {
		if word[i] == '.' {
			ok = true
			break
		} else {
			ans += string(word[i])
		}
	}
	return ans, ok
}
func gen(ch string, num int) string {
	ans := ""
	for num > 0 {
		ans += ch
		num--
	}
	return ans
}
func main() {
	var s string
	fmt.Scan(&s)
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	it := bufio.NewScanner(file)
	it.Split(bufio.ScanWords)
	table := make(map[string]bool)
	for it.Scan() {
		table[strings.ToLower(it.Text())] = true
	}
	err = it.Err()
	if err != nil {
		log.Fatal(err)
	}
	It := bufio.NewReader(os.Stdin)
	s, _ = It.ReadString('\n')
	s = strings.TrimSpace(s)
	for s != "exit" {
		if len(s) == 0 {
			s, _ = It.ReadString('\n')
			s = strings.TrimSpace(s)
			continue
		}
		SL := strings.Split(s, " ")
		for auto := range SL {
			tmp, ok := CUT(SL[auto])
			if table[strings.ToLower(tmp)] {
				tmp = gen(string('*'), len(tmp))
			}
			fmt.Print(tmp)
			if ok {
				fmt.Print(".\n")
			} else {
				fmt.Print(" ")
			}
		}
		s, _ = It.ReadString('\n')
		s = strings.TrimSpace(s)
	}
	fmt.Println("Bye!")
}
