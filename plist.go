package main

import "encoding/xml"

type Plist struct {
	XMLName xml.Name `xml:"plist"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Dict    struct {
		Text  string   `xml:",chardata"`
		Key   []string `xml:"key"`
		Array struct {
			Text   string `xml:",chardata"`
			String string `xml:"string"`
		} `xml:"array"`
		Dict []struct {
			Text   string   `xml:",chardata"`
			Key    []string `xml:"key"`
			String []string `xml:"string"`
			False  string   `xml:"false"`
			True   string   `xml:"true"`
		} `xml:"dict"`
		String []string `xml:"string"`
		False  string   `xml:"false"`
		True   string   `xml:"true"`
	} `xml:"dict"`
}
