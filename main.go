package main

import (
	"log"

	"github.com/YAMATO50/UWTL/internal/docxio"
	"github.com/YAMATO50/UWTL/internal/wxml"
)

func main() {
	//Flag checking and shit
	err := docxio.Unzip("testfile.docx", "testing")

	if err != nil {
		log.Fatal(err)
	}

	//var v wxml.Wxml
	v := make(wxml.Wxml)

	wxml.ParseXMLFromFile("documentTemplate1.xml", &v, 1)
	wxml.PrintStruct(v)
}
