package advance16_xml_json

import (
	"encoding/xml"
	"testing"
)

type Social struct {
	XMLName  xml.Name `xml:"social"`
	Facebook string   `xml:"facebook"`
	Twitter  string   `xml:"twitter"`
	Youtube  string   `xml:"youtube,omitempty"`
}

func TestXML(t *testing.T) {
	soc := &Social{
		Facebook: "fb",
		Twitter:  "tw",
	}
	if bytes, err := xml.Marshal(soc); err != nil {
		t.Errorf("err=%+v", err)
	} else {
		t.Logf("soc = %s", string(bytes))
	}
}
