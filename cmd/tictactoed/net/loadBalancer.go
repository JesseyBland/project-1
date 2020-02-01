package main

import (
	"log"
	"net/url"
)

func main() {
	tsources := []string{"http://localhost:7777", "http://localhost:8888", "http://localhost:9999"}
	for {
		for _, ts := range tsources {
			ttt, err := url.Parse(ts)
			if err != nil {
				log.Fatal(err)
			}

		}
	}
}

/*	var a, b, c int
	if a == b && a == c {

		a++
		fmt.Println("On Server :8888  Server 1 connection")
		ticURL, err := url.Parse("http://localhost:8888")

		if err != nil {
			log.Fatal(err)
		}
		return ticURL

	} else if a > b && b == c {
		b++
		//ConnSignal <- "On Server :9999  Server 2 connection"
		fmt.Println("On Server :9999   b=", b)
		ticURL, err1 := url.Parse("http://localhost:9999")
		if err1 != nil {
			log.Fatal(err1)
		}
		return ticURL

	} else if a > c && b > c {
		c++
		//ConnSignal <- "On Server :7777  Server 3 connection"
		fmt.Println("On Server :7777   c=", c)
		ticURL, err2 := url.Parse("http://localhost:7777")
		if err2 != nil {
			log.Fatal(err2)
		}
		return ticURL
	}

	return nil

}*/
