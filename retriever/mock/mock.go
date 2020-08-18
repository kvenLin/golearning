package mock

import "fmt"

type Retriever struct {
	Content string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever:{Content:%s}", r.Content)
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Content = form["content"]
	return "ok"
}

func (r *Retriever) Get(url string) string {
	return r.Content
}
