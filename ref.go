package spark

import (
	"context"
)

type V1Ref struct {
	ref Component
}

func Ref(ref Component, base Component) Reference {
	for key, val := range base.Props() {
		ref.Bind(key, func() string {
			return val
		})
	}
	return &V1Ref{
		ref: ref,
	}
}

func (r *V1Ref) FirstChild() Element {
	elements := r.ref.GetElements()
	if len(elements) > 0 {
		return elements[0]
	}
	return nil
}

func (r *V1Ref) AddChild(e Element) {
	r.ref.Push(e)
}

func (r *V1Ref) Attributes() []string {
	return r.FirstChild().Attributes()
}

func (r *V1Ref) Component() Component {
	return r.ref
}

func (r *V1Ref) Content(fn func() string) {
	r.FirstChild().Content(fn)
}

func (r *V1Ref) GetAttribute(key string) string {
	return r.FirstChild().GetAttribute(key)
}

func (r *V1Ref) SetAttribute(key string, value func() string) {
	r.FirstChild().SetAttribute(key, value)
}

func (r *V1Ref) RemoveAttribute(key string) {
	r.FirstChild().RemoveAttribute(key)
}

func (r *V1Ref) GetProp(s string) string {
	return r.ref.GetProp(s)
}

func (r *V1Ref) Render(ctx context.Context) string {
	return r.ref.Render(ctx)
}

func (r *V1Ref) Tag() (string, bool) {
	return r.FirstChild().Tag()
}
