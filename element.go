package spark

import (
	"context"
	"fmt"
	"strings"
)

type V1Element struct {
	genID       string
	attributes  map[string]func() string
	element     string
	selfClosing bool
	base        Component
	children    []Element
	content     func() string
}

func NewV1Element(element string, selfClosing bool, base Component) Element {
	return &V1Element{
		element:     element,
		selfClosing: selfClosing,
		base:        base,
		attributes:  make(map[string]func() string),
	}
}

func (e *V1Element) Component() Component {
	return e.base
}

func (e *V1Element) AddChild(el Element) {
	e.children = append(e.children, el)
}

func (e *V1Element) Tag() (string, bool) {
	return e.element, e.selfClosing
}

func (e *V1Element) Attributes() []string {
	var attributes []string

	for key := range e.attributes {
		attributes = append(attributes, key)
	}

	return attributes
}

func (e *V1Element) SetAttribute(key string, value func() string) {
	e.attributes[key] = value
}

func (e *V1Element) RemoveAttribute(key string) {
	delete(e.attributes, key)
}

func (e *V1Element) GetAttribute(key string) string {
	attr, ok := e.attributes[key]

	if !ok {
		return ""
	}

	return strings.ReplaceAll(attr(), "$(id)", e.Component().GetID())
}

func (e *V1Element) GetProp(s string) string {
	return e.base.GetProp(s)
}

func (e *V1Element) Content(fn func() string) {
	e.content = fn
}

func (e *V1Element) Render(c context.Context) string {
	var out string

	out += fmt.Sprintf("<%s", e.element)
	for _, name := range e.Attributes() {
		out += fmt.Sprintf(" %s=\"%s\"", name, e.attrEscape(e.GetAttribute(name)))
	}

	if !e.selfClosing {
		out += ">" + e.content()
	} else {
		out += "/>"
		return out
	}

	for _, child := range e.children {
		out += child.Render(c)
	}

	if !e.selfClosing {
		out += fmt.Sprintf("</%s>", e.element)
	}

	return out
}

func (e *V1Element) attrEscape(s string) string {
	return s
}
