package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"io"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
	Xml     string   `xml:",innerxml"`
}

type Author struct {
	Id   string      `xml:"id,attr"`
	Name string      `xml:",chardata"`
}

func main() {
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := io.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}
	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)

	post_xml := Post{
		Id : "1",
		Content: "Hello World!",
		Author: Author{
			Id: "2",
			Name: "Sau Sheong",
		},
	}
	//output, err := xml.Marshal(&post_xml)
	output, err := xml.MarshalIndent(&post_xml, "", "\t")
	if err != nil {
		fmt.Println("Error marshaling to XML:", err)
		return
	}
	err = os.WriteFile("post_xml.xml", []byte(xml.Header + string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}
}