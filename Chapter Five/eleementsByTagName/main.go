package eleementsbytagname

import "golang.org/x/net/html"

func ElementsByName(h *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	if h == nil {
		return nodes
	}
	for _, arg := range name {
		if h.Type == html.ElementNode && h.Data == arg {
			nodes = append(nodes, h)
		}
	}
	nodes = ElementsByName(h.FirstChild, name...)
	nodes2 := ElementsByName(h.NextSibling, name...)
	nodes = append(nodes, nodes2...)
	return nodes
}
