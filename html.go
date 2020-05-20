package main

import "encoding/xml"

type Html struct {
	XMLName xml.Name `xml:"html"`
	Text    string   `xml:",chardata"`
	Head    struct {
		Text string `xml:",chardata"`
		Meta struct {
			Text      string `xml:",chardata"`
			HTTPEquiv string `xml:"http-equiv,attr"`
			Content   string `xml:"content,attr"`
			Title     string `xml:"title"`
		} `xml:"meta"`
	} `xml:"head"`
	Body struct {
		Text string `xml:",chardata"`
		H2   string `xml:"h2"`
		Hr   struct {
			Text    string `xml:",chardata"`
			Noshade string `xml:"noshade,attr"`
			Size    string `xml:"size,attr"`
		} `xml:"hr"`
		P []struct {
			Text string `xml:",chardata"`
			A    struct {
				Text string `xml:",chardata"`
				Href string `xml:"href,attr"`
			} `xml:"a"`
			Br struct {
				Text  string `xml:",chardata"`
				Clear string `xml:"clear,attr"`
			} `xml:"br"`
			Hr struct {
				Text    string `xml:",chardata"`
				Noshade string `xml:"noshade,attr"`
				Size    string `xml:"size,attr"`
			} `xml:"hr"`
			Address string `xml:"address"`
		} `xml:"p"`
		Ul struct {
			Text string `xml:",chardata"`
			Li   struct {
				Text   string `xml:",chardata"`
				Strong string `xml:"strong"`
			} `xml:"li"`
		} `xml:"ul"`
	} `xml:"body"`
}
