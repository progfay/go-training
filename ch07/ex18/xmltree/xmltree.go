package xmltree

import (
	"encoding/xml"
	"io"
)

type (
	Node     interface{}
	CharData string
	Element  struct {
		Type     xml.Name
		Attr     []xml.Attr
		Children []Node
	}
)

func newElement(s xml.StartElement) Element {
	e := Element{}
	e.Type = xml.Name{
		Space: s.Name.Space,
		Local: s.Name.Local,
	}
	e.Attr = make([]xml.Attr, 0)
	for _, attr := range s.Attr {
		e.Attr = append(e.Attr, xml.Attr{
			Name: xml.Name{
				Space: attr.Name.Space,
				Local: attr.Name.Local,
			},
			Value: attr.Value,
		})
	}
	e.Children = make([]Node, 0)
	return e
}

func New(r io.Reader) (Node, error) {
	var f func(parent *Element, decoder *xml.Decoder) error
	f = func(parent *Element, decoder *xml.Decoder) error {
		for {
			tok, err := decoder.Token()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			switch tok := tok.(type) {
			case xml.StartElement:
				c := newElement(tok)
				err := f(&c, decoder)
				if err != nil {
					return err
				}
				parent.Children = append(parent.Children, c)

			case xml.EndElement:
				return nil

			case xml.CharData:
				parent.Children = append(parent.Children, CharData(tok))
			}
		}
	}

	decoder := xml.NewDecoder(r)
	root := Element{
		Type: xml.Name{
			Space: "",
			Local: "ROOT",
		},
		Attr:     []xml.Attr{},
		Children: []Node{},
	}
	err := f(&root, decoder)
	if err != nil {
		return nil, err
	}
	return root, nil
}
