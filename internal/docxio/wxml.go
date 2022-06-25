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
	fmt.Println(v)
	fmt.Println("Ignorable:", v.Ignoreable)
	fmt.Println("Height:", v.Body.SectPr.PageSize.Heigth)
	fmt.Println("ParagraphID:", v.Body.Paragraph[0].ParaID)
	fmt.Println("Paragraph:", v.Body.Paragraph)
	fmt.Println("Text:", v.Body.Paragraph[0].R.Text)
	fmt.Println("Properties:", v.Body.Paragraph[3].ParagraphProperties.RunProperties.Properties)
	fmt.Println("NewPage:", v.Body.Paragraph[30].R.LastRenderedPageBreak.XMLName.Local)
}

type Wxml struct {
	XMLName    xml.Name `xml:"document"`
	Body       Body     `xml:"body"`
	Ignoreable string   `xml:"Ignorable,attr"`
}

type Body struct {
	Paragraph []Paragraph `xml:"p"`
	SectPr    SectPr      `xml:"sectPr"`
}

type SectPr struct {
	PageSize PageSize `xml:"pgSz"`
	PageMar  PageMar  `xml:"pgMar"`
	Cols     Cols     `xml:"cols"`
	DocGrid  DocGrid  `xml:"docGrid"`
}

type PageSize struct {
	Width  string `xml:"w,attr"`
	Heigth string `xml:"h,attr"`
}

type PageMar struct {
	Top    string `xml:"top,attr"`
	Right  string `xml:"right,attr"`
	Bottom string `xml:"bottom,attr"`
	Left   string `xml:"left,attr"`
	Header string `xml:"header,attr"`
	Footer string `xml:"footer,attr"`
	Gutter string `xml:"gutter,attr"`
}

type Cols struct {
	Space string `xml:"space,attr"`
}

type DocGrid struct {
	LinePitch string `xml:"linePitch,attr"`
}

type Paragraph struct {
	R                   R      `xml:"r"`
	ParaID              string `xml:"paraId,attr"`
	TextID              string `xml:"textId,attr"`
	RsIDR               string `xml:"rsidR,attr"`
	RsidRPr             string `xml:"rsidRPr,attr"`
	RsIDRDefault        string `xml:"rsidRDefault,attr""`
	ParagraphProperties PPR    `xml:"pPr"`
}

type R struct {
	Text                  string `xml:"t"`
	Space                 string `xml:"xml:space"`
	RsidRPr               string `xml:"rsidRPr,attr"`
	RunProperties         RPR    `xml:"rPr"`
	LastRenderedPageBreak LRPB   `xml:"lastRenderedPageBreak"`
}

type LRPB struct {
	XMLName xml.Name
}

type PPR struct {
	RunProperties RPR `xml:"rPr"`
}

type RPR struct {
	Properties string `xml:",innerxml"`
}

//document = Wurzelelement
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
