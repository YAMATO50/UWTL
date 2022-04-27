package docxio

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

func GetEmptyXMLClass() Wxml {
	return Wxml{}
}

func ParseXML(data []byte, v *Wxml) error {
	err := xml.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}

func ParseXMLFromFile(path string, v *Wxml, removeLines int) error {
	//Open file
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//remove lines by removeLines
	lines = lines[removeLines:]

	var fileContent string

	for _, line := range lines {
		fileContent = fileContent + strings.TrimSuffix(line, "\n")
	}

	//call ParseXML
	err = ParseXML([]byte(fileContent), v)

	if err != nil {
		return err
	}

	return nil
}

//DEBUGGING FUNCTION
func PrintStruct(v Wxml) {
	fmt.Println(v.Ignoreable)
	fmt.Println(v.Document.Body.SectPr.PageSize.Heigth)
	fmt.Println(v.Document.Body.ParagraphID[0])
	fmt.Println(v.Document.Body.Paragraph[0])
}

type Wxml struct {
	Document   Document `xml:"w:document"`
	Ignoreable string   `xml:"mc:Ignorable,attr"`
}

type Document struct {
	Body Body `xml:"w:body"`
}

type Body struct {
	Paragraph    []Paragraph `xml:"w:p"`
	ParagraphID  []string    `xml:"w14:paraID,attr"`
	TextID       []string    `xml:"w14:textID,attr"`
	RsidR        []string    `xml:"w:rsidR,attr"`
	RsidRDefault []string    `xml:"w:rsidRDefault,attr"`
	RsidRPr      []string    `xml:"w:rsidRPr,attr"`

	SectPr SectPr `xml:"w:sectPr"`
}

//w:sectPr w:rsidR="00A63C32" w:rsidRPr="00A63C32"

type SectPr struct {
	PageSize PageSize `xml:"w:pgSz"`
	PageMar  PageMar  `xml:"w:pgMar"`
	Cols     Cols     `xml:"w:cols"`
	DocGrid  DocGrid  `xml:"w:docGrid"`
}

type PageSize struct {
	Width  int `xml:"w:w,attr"`
	Heigth int `xml:"w:h,attr"`
}

type PageMar struct {
	Top    int `xml:"w:top,attr"`
	Right  int `xml:"w:right,attr"`
	Bottom int `xml:"w:bottom,attr"`
	Left   int `xml:"w:left,attr"`
	Header int `xml:"w:header,attr"`
	Footer int `xml:"w:footer,attr"`
	Gutter int `xml:"w:gutter,attr"`
}

type Cols struct {
	Space int `xml:"w:space,attr"`
}

type DocGrid struct {
	LinePitch int `xml:"w:linePitch,attr"`
}

type Paragraph struct {
}

//w:document = Wurzelelement
//<eintrag>Eintrag</eintrag> = Element mit start und endeintrag
//<eintrag /> = Element mit Leereintrag
//
//<Parent>
//	<Child1>
//		<ChildChild/>
//	</Child1>
//	<Child2>
//		<isSiblingOfChild1/>
//	</Child2>
//</Parent>
//
//<hasAttribute attribute="attribute"/>
