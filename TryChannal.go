package main

import (
	"fmt"
	"strings"
)

//read a string, return different words with letter counts
// naive approach (copy-paste)
func main() {
	input := "I am someone La Honda RC100 et ses dérivées RC101 et RC101B sont des prototypes de monoplace de Formule 1 conçus, entre 1991 et 1996, par les jeunes ingénieurs du département recherche et développement de Honda ; elles sont destinées à offrir un nouveau défi technique aux ingénieurs du constructeur japonais qui s'investissent dans ce projet en dehors de leur temps de travail.La Honda RC100 et ses dérivées RC101 et RC101B sont des prototypes de monoplace de Formule 1 conçus, entre 1991 et 1996, par les jeunes ingénieurs du département recherche et développement de Honda ; elles sont destinées à offrir un nouveau défi technique aux ingénieurs du constructeur japonais qui s'investissent dans ce projet en dehors de leur temps de travail.La Honda RC100 et ses dérivées RC101 et RC101B sont des prototypes de monoplace de Formule 1 conçus, entre 1991 et 1996, par les jeunes ingénieurs du département recherche et développement de Honda ; elles sont destinées à offrir un nouveau défi technique aux ingénieurs du constructeur japonais qui s'investissent dans ce projet en dehors de leur temps de travail.La Honda RC100 et ses dérivées RC101 et RC101B sont des prototypes de monoplace de Formule 1 conçus, entre 1991 et 1996, par les jeunes ingénieurs du département recherche et développement de Honda ; elles sont destinées à offrir un nouveau défi technique aux ingénieurs du constructeur japonais qui s'investissent dans ce projet en dehors de leur temps de travail.La Honda RC100 et ses dérivées RC101 et RC101B sont des prototypes de monoplace de Formule 1 conçus, entre 1991 et 1996, par les jeunes ingénieurs du département recherche et développement de Honda ; elles sont destinées à offrir un nouveau défi technique aux ingénieurs du constructeur japonais qui s'investissent dans ce projet en dehors de leur temps de travail.La Honda RC100 et ses dérivées RC101 et RC101B sont des prototypes de monoplace de Formule 1 conçus, entre 1991 et 1996, par les jeunes ingénieurs du département recherche et développement de Honda ; elles sont destinées à offrir un nouveau défi technique aux ingénieurs du constructeur japonais qui s'investissent dans ce projet en dehors de leur temps de travail.La Honda RC100 et ses dérivées RC101 et RC101B sont des prototypes de monoplace de Formule 1 conçus, entre 1991 et 1996, par les jeunes ingénieurs du département recherche et développement de Honda ; elles sont destinées à offrir un nouveau défi technique aux ingénieurs du constructeur japonais qui s'investissent dans ce projet en dehors de leur temps de travail.La Honda RC100 et ses dérivées RC101 et RC101B sont des prototypes de monoplace de Formule 1 conçus, entre 1991 et 1996, par les jeunes ingénieurs du département recherche et développement de Honda ; elles sont destinées à offrir un nouveau défi technique aux ingénieurs du constructeur japonais qui s'investissent dans ce projet en dehors de leur temps de travail.La Honda RC100 et ses dérivées RC101 et RC101B sont des prototypes de monoplace de Formule 1 conçus, entre 1991 et 1996, par les jeunes ingénieurs du département recherche et développement de Honda ; elles sont destinées à offrir un nouveau défi technique aux ingénieurs du constructeur japonais qui s'investissent dans ce projet en dehors de leur temps de travail.La Honda RC100 et ses dérivées RC101 et RC101B sont des prototypes de monoplace de Formule 1 conçus, entre 1991 et 1996, par les jeunes ingénieurs du département recherche et développement de Honda ; elles sont destinées à offrir un nouveau défi technique aux ingénieurs du constructeur japonais qui s'investissent dans ce projet en dehors de leur temps de travail.La Honda RC100 et ses dérivées RC101 et RC101B sont des prototypes de monoplace de Formule 1 conçus, entre 1991 et 1996, par les jeunes ingénieurs du département recherche et développement de Honda ; elles sont destinées à offrir un nouveau défi technique aux ingénieurs du constructeur japonais qui s'investissent dans ce projet en dehors de leur temps de travail.La Honda RC100 et ses dérivées RC101 et RC101B sont des prototypes de monoplace de Formule 1 conçus, entre 1991 et 1996, par les jeunes ingénieurs du département recherche et développement de Honda ; elles sont destinées à offrir un nouveau défi technique aux ingénieurs du constructeur japonais qui s'investissent dans ce projet en dehors de leur temps de travail.La Honda RC100 et ses dérivées RC101 et RC101B sont des prototypes de monoplace de Formule 1 conçus, entre 1991 et 1996, par les jeunes ingénieurs du département recherche et développement de Honda ; elles sont destinées à offrir un nouveau défi technique aux ingénieurs du constructeur japonais qui s'investissent dans ce projet en dehors de leur temps de travail.La Honda RC100 et ses dérivées RC101 et RC101B sont des prototypes de monoplace de Formule 1 conçus, entre 1991 et 1996, par les jeunes ingénieurs du département recherche et développement de Honda ; elles sont destinées à offrir un nouveau défi technique aux ingénieurs du constructeur japonais qui s'investissent dans ce projet en dehors de leur temps de travail.La Honda RC100 et ses dérivées RC101 et RC101B sont des prototypes de monoplace de Formule 1 conçus, entre 1991 et 1996, par les jeunes ingénieurs du département recherche et développement de Honda ; elles sont destinées à offrir un nouveau défi technique aux ingénieurs du constructeur japonais qui s'investissent dans ce projet en dehors de leur temps de travail."
	s := strings.Split(input, " ")
	lettersCount := 0

	counts := make(chan int)
	words := make(chan string)

	for i:=0; i<4; i++{
		go wordCounter(words, i, counts)
	}

	for _, w := range s{
		words <- w
	}
	close(words)

	for i:=0; i<4; i++{
		count := <- counts
		lettersCount = lettersCount + count
	}
	close(counts)

	fmt.Printf("input has %d letters\n", lettersCount)
}

func wordCounter(words chan string, num int, counts chan int){
	fmt.Printf("hi I am goroutine %d\n", num)
	count := 0
	for{
		w, ok := <- words
		if !ok{
			println("exit")
			break
		}
		count = count + len(w)
	}
	counts <- count
}


