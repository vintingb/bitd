package main

import "encoding/xml"

type success struct {
	XMLName xml.Name `xml:"feed"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	M       string   `xml:"m,attr"`
	D       string   `xml:"d,attr"`
	Base    string   `xml:"base,attr"`
	ID      struct {
		Text string `xml:",chardata"`
	} `xml:"id"`
	Title struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
	} `xml:"title"`
	Updated struct {
		Text string `xml:",chardata"`
	} `xml:"updated"`
	Author struct {
		Text string `xml:",chardata"`
		Name struct {
			Text string `xml:",chardata"`
		} `xml:"name"`
	} `xml:"author"`
	Link struct {
		Text  string `xml:",chardata"`
		Href  string `xml:"href,attr"`
		Rel   string `xml:"rel,attr"`
		Title string `xml:"title,attr"`
	} `xml:"link"`
	Entry struct {
		Text string `xml:",chardata"`
		ID   struct {
			Text string `xml:",chardata"`
		} `xml:"id"`
		Title struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"title"`
		Updated struct {
			Text string `xml:",chardata"`
		} `xml:"updated"`
		Category struct {
			Text   string `xml:",chardata"`
			Term   string `xml:"term,attr"`
			Scheme string `xml:"scheme,attr"`
		} `xml:"category"`
		Link []struct {
			Text  string `xml:",chardata"`
			Href  string `xml:"href,attr"`
			Rel   string `xml:"rel,attr"`
			Title string `xml:"title,attr"`
			Type  string `xml:"type,attr"`
		} `xml:"link"`
		Content struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			Src  string `xml:"src,attr"`
		} `xml:"content"`
		Properties struct {
			Text string `xml:",chardata"`
			ID   struct {
				Text string `xml:",chardata"`
			} `xml:"Id"`
			Name struct {
				Text string `xml:",chardata"`
			} `xml:"Name"`
			ContentType struct {
				Text string `xml:",chardata"`
			} `xml:"ContentType"`
			ContentLength struct {
				Text string `xml:",chardata"`
			} `xml:"ContentLength"`
			ChildrenNumber struct {
				Text string `xml:",chardata"`
			} `xml:"ChildrenNumber"`
			Value struct {
				Text string `xml:",chardata"`
				Null string `xml:"null,attr"`
			} `xml:"Value"`
			CreationDate struct {
				Text string `xml:",chardata"`
			} `xml:"CreationDate"`
			IngestionDate struct {
				Text string `xml:",chardata"`
			} `xml:"IngestionDate"`
			ModificationDate struct {
				Text string `xml:",chardata"`
			} `xml:"ModificationDate"`
			EvictionDate struct {
				Text string `xml:",chardata"`
				Null string `xml:"null,attr"`
			} `xml:"EvictionDate"`
			Online struct {
				Text string `xml:",chardata"`
			} `xml:"Online"`
			OnDemand struct {
				Text string `xml:",chardata"`
			} `xml:"OnDemand"`
			ContentDate struct {
				Text  string `xml:",chardata"`
				Type  string `xml:"type,attr"`
				Start struct {
					Text string `xml:",chardata"`
				} `xml:"Start"`
				End struct {
					Text string `xml:",chardata"`
				} `xml:"End"`
			} `xml:"ContentDate"`
			Checksum struct {
				Text      string `xml:",chardata"`
				Type      string `xml:"type,attr"`
				Algorithm struct {
					Text string `xml:",chardata"`
				} `xml:"Algorithm"`
				Value struct {
					Text string `xml:",chardata"`
				} `xml:"Value"`
			} `xml:"Checksum"`
			ContentGeometry struct {
				Text string `xml:",chardata"`
				Null string `xml:"null,attr"`
			} `xml:"ContentGeometry"`
		} `xml:"properties"`
	} `xml:"entry"`
}
