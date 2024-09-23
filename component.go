package spark

import (
	"context"
	"net/http"
)

type V1Component struct {
	genID    string
	props    []string
	binds    map[string]func() string
	content  string
	elements []Element
}

func NewV1Component(props []string) Component {
	return &V1Component{
		props: props,
		binds: make(map[string]func() string),
	}
}

func (c *V1Component) Bind(name string, value func() string) {
	c.binds[name] = value
}

func (c *V1Component) Props() map[string]string {
	var data = make(map[string]string)
	for key, value := range c.binds {
		data[key] = value()
	}
	return data
}

func (c *V1Component) GetProp(s string) string {
	fn, ok := c.binds[s]
	if !ok {
		return ""
	}
	return fn()
}

func (c *V1Component) Content(s string) {
	c.content = s
}

func (c *V1Component) GetContent() string {
	return c.content
}

func (c *V1Component) GetElements() []Element {
	return c.elements
}

func (c *V1Component) Push(e Element) {
	c.elements = append(c.elements, e)
}

func (c *V1Component) Render(ctx context.Context) string {
	var generated string

	for _, e := range c.elements {
		generated += e.Render(ctx)
	}

	return generated
}

func (c *V1Component) Response(w http.ResponseWriter, r *http.Request) error {
	content := c.Render(r.Context())
	_, err := w.Write([]byte(content))
	return err
}

func (c *V1Component) GetID() string {
	if c.genID != "" {
		return c.genID
	}
	return randomString(5)
}
