package model

import (
	"container/list"
	"sort"
)

func (graph *BPMNGraph) GetExecutableNodes() []*Node {
	// Filter nodes by type
	// Executable nodes are Task, IntermediateCatchEvent, IntermediateThrowEvent
	var executableNodes []*Node
	for _, node := range graph.Nodes {
		switch node.Type {
		case "Task", "IntermediateCatchEvent", "IntermediateThrowEvent":
			executableNodes = append(executableNodes, node)
		}
	}
	layers := layerize(graph)

	// Reorder nodes by layer
	sort.Slice(executableNodes, func(i, j int) bool {
		if layers[executableNodes[i].ID] == layers[executableNodes[j].ID] {
			return executableNodes[i].ID < executableNodes[j].ID
		}
		return layers[executableNodes[i].ID] < layers[executableNodes[j].ID]
	})

	return executableNodes
}

func layerize(graph *BPMNGraph) map[string]int {
	startNodes := graph.GetStartNodes()
	result := make(map[string]int)

	for _, startNode := range startNodes {
		// Initialize a queue (FIFO) for BFS and a visited map
		queue := list.New()
		nextLayer := list.New()
		layer := 0
		visited := make(map[string]bool)

		// Enqueue the start node
		queue.PushBack(startNode.ID)
		visited[startNode.ID] = true

		// Traverse the graph
		for queue.Len() > 0 {
			// Dequeue the next node
			element := queue.Front()
			nodeID := element.Value.(string)
			queue.Remove(element)

			// Get the node from the graph
			node := graph.Nodes[nodeID]

			// Enqueue all outgoing edges' target nodes (if not visited)
			for _, edge := range node.Outgoing {
				targetID := edge.TargetRef.ID
				if !visited[targetID] {
					result[targetID] = layer
					nextLayer.PushBack(targetID)
					visited[targetID] = true
				}
			}

			if queue.Len() == 0 {
				queue.PushBackList(nextLayer)
				nextLayer.Init()
				layer++
			}
		}
	}

	return result
}
