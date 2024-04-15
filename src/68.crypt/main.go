package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "ABCDE")

	//b := []byte{12, 34, 55, 3}

	fmt.Println(h.Sum(nil))
	//fmt.Println(h.Sum(b))

	//fmt.Println(b)

	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s)

}

//?2
// import (
// 	"fmt"
// 	"io"

// 	"crypto/sha1"
// 	"crypto/sha256"
// 	"crypto/sha512"
// )

// func main() {
// 	s1 := sha1.New()
// 	io.WriteString(s1, "ABCDE")
// 	fmt.Printf("%x\n", s1.Sum(nil))

// 	s256 := sha256.New()
// 	io.WriteString(s256, "ABCDE")
// 	fmt.Printf("%x\n", s256.Sum(nil))

// 	s512 := sha512.New()
// 	io.WriteString(s512, "ABCDE")
// 	fmt.Printf("%x\n", s512.Sum(nil))
// }

