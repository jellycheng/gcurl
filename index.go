package gcurl

import (
	"encoding/xml"
	"net/http"
	"strings"
)

func NewClient() *Request {
	req := &Request{}
	req.options = NewOptions()
	return req
}

func Get(uri string, opts ...Options) (*Response, error) {
	r := NewClient()
	return r.Request(http.MethodGet, uri, opts...)
}

func Post(uri string, opts ...Options) (*Response, error) {
	r := NewClient()
	return r.Request(http.MethodPost, uri, opts...)
}

func Put(uri string, opts ...Options) (*Response, error) {
	r := NewClient()
	return r.Request(http.MethodPut, uri, opts...)
}

func Patch(uri string, opts ...Options) (*Response, error) {
	r := NewClient()
	return r.Request(http.MethodPatch, uri, opts...)
}

func Delete(uri string, opts ...Options) (*Response, error) {
	r := NewClient()
	return r.Request(http.MethodDelete, uri, opts...)
}

func Map2XML(m map[string]string, rootName ...string) ([]byte, error) {
	rootTag := "xml"
	if len(rootName) > 0 {
		rootTag = rootName[0]
	}
	if rootTag == "" {
		rootTag = "xml"
	}
	d := MyXmldata{XMLName:xml.Name{Local: rootTag}, Data: m}
	bt, err := xml.Marshal(d)
	if err != nil {
		return nil, err
	}
	return bt, nil
}

type MyXmldata struct {
	XMLName xml.Name
	Data    map[string]string
}

func (m MyXmldata) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m.Data) == 0 {
		return nil
	}
	start.Name.Local = m.XMLName.Local
	err := e.EncodeToken(start)
	if err != nil {
		return err
	}
	for k, v := range m.Data {
		if strings.HasPrefix(v, "cdata:") {
			v = strings.Replace(v, "cdata:", "", 1)
			xs := struct {
				XMLName xml.Name
				Value   interface{} `xml:",cdata"`
			}{xml.Name{Local: k}, v}
			e.Encode(xs)
		} else {
			xs := struct {
				XMLName xml.Name
				Value   interface{} `xml:",chardata"`
			}{xml.Name{Local: k}, v}
			e.Encode(xs)
		}
	}

	return e.EncodeToken(start.End())
}
