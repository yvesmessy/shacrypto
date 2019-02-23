package main

import "crypto/sha1"
import "fmt" 

func main() {

	s := "I swear on Gryffindor that I am up to no good!"


	h := sha1.New()



	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)

	fmt.Println("This is the SHA1 coded version of your phrase") 

	fmt.Printf("%x\n", bs)

}
