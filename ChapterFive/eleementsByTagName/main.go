package eleementsbytagname

import "golang.org/x/net/html"

func ElementsByName(h *html.Node, names ...string) []*html.Node {
	var nodes []*html.Node
	if h == nil {
		return nodes
	}
	var nameMap = make(map[string]struct{})
	for _, name := range names {
		nameMap[name] = struct{}{}
	}
	// for _, arg := range name {
	// 	if h.Type == html.ElementNode && h.Data == arg {
	// 		nodes = append(nodes, h)
	// 	}
	// }
	if h.Type == html.ElementNode {
		if _, found := nameMap[h.Data]; found {
			nodes = append(nodes, h)
		}
	}
	nodes = append(nodes, ElementsByName(h.FirstChild, names...)...)
	nodes = append(nodes, ElementsByName(h.NextSibling, names...)...)
	return nodes
}
