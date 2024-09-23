package spark

import (
	"context"
	"net/http"
)

type Component interface {
	// Unique ID

	GetID() string

	// Props & Attributes

	// Bind binds a new prop or rewrites the current
	Bind(key string, value func() string)
	Props() map[string]string
	GetProp(s string) string

	Content(content string)
	GetContent() string

	// Elements

	GetElements() []Element
	Push(e Element)

	// Generators

	Render(ctx context.Context) string
	Response(w http.ResponseWriter, r *http.Request) error
}

type Element interface {
	// Component returns the element's component
	Component() Component

	// Tag returns the html tag for the element and a boolean if the tag is self closing or not
	Tag() (string, bool)

	Attributes() []string
	SetAttribute(key string, value func() string)
	RemoveAttribute(key string)
	GetAttribute(key string) string

	GetProp(s string) string

	Content(func() string)

	AddChild(e Element)

	Render(ctx context.Context) string
}

type Reference interface {
	Element
	FirstChild() Element
}
