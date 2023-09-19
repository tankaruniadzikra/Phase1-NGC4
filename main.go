// NG Challenges 4 : function 2
package main

import (
	"fmt"
	"strings"
)

func AlayGen(kata ...string) string {
	//untuk mengubah kata-kata menjadi "alay" sesuai dengan ketentuan yang sudah didefinisikan.
	alayRules := map[rune]string{
		//pasangan konversi yang akan diterapkan pada karakter tertentu dalam fungsi AlayGen.
		'a': "4",
		'e': "3",
		'i': "!",
		'l': "1",
		'n': "N",
		's': "5",
		'x': "*",
	}

	result := ""
	for _, word := range kata {
		for _, char := range word {
			if alay, ok := alayRules[char]; ok {
				result += alay
			} else {
				result += string(char)
			}
		}
		result += " "
	}

	return strings.TrimSpace(result)
}

func main() {
	kalimat := AlayGen("hello", "world", "zzz", "senin")
	fmt.Println(kalimat)
}

// NG Challenges 4 : function 3
package main

import "fmt"

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	fmt.Println(Fibonacci(10))
}
