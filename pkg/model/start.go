package model

func (graph *BPMNGraph) GetStartNodes() []*Node {
	// Filter nodes by type
	// Start nodes are StartEvent
	var startNodes []*Node
	for _, node := range graph.Nodes {
		if node.Type == "StartEvent" {
			startNodes = append(startNodes, node)
		}
	}

	return startNodes
}
