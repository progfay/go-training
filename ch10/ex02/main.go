package main

import (
	"log"

	"github.com/progfay/go-training/ch10/ex02/archive"

	_ "github.com/progfay/go-training/ch10/ex02/archive/tar"
	_ "github.com/progfay/go-training/ch10/ex02/archive/zip"
)

func main() {
	log.Println(archive.Load("archive/data/archive.zip"))
	log.Println(archive.Load("archive/data/archive.tar"))
}
