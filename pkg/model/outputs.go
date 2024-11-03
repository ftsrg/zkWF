package model

import "fmt"

func (node *Node) GetPairs() []*Node {
	if node.Type != "Task" {
		return []*Node{}
	}
	perv := node.Outgoing[0].TargetRef.GetPreviousNodes()
	output := []*Node{}
	for _, n := range perv {
		if n.ID == node.ID {
			continue
		}
		if n.Type == "Task" {
			output = append(output, n)
		}
	}

	return output
}

// GetPreviousNodes returns a slice of tasks (or events) before a node, handling incoming edges.
func (node *Node) GetPreviousNodes() []*Node {
	var previousNodes []*Node
	visited := make(map[string]bool) // To avoid cycles or infinite recursion

	// Helper function to recursively traverse the graph
	var traverse func(n *Node)
	traverse = func(n *Node) {
		// If this node has already been visited, avoid reprocessing it
		if visited[n.ID] {
			return
		}
		visited[n.ID] = true

		// Iterate through the incoming edges to get the previous nodes
		for _, edge := range n.Incoming {
			prevNode := edge.SourceRef

			// If the previous node is a task or event, add it to the result
			if prevNode.Type == "Task" || prevNode.Type == "StartEvent" || prevNode.Type == "IntermediateCatchEvent" || prevNode.Type == "IntermediateThrowEvent" {
				previousNodes = append(previousNodes, prevNode)
			} else if isGateway(prevNode.Type) {
				// If it's a gateway, recursively follow all its incoming edges
				traverse(prevNode)
			}
		}
	}

	// Start traversal from the current node (which would be a gateway, typically)
	traverse(node)
	return previousNodes
}

// GetNextNodes returns a slice of the next tasks (or events) recursively, handling gateways.
func (node *Node) GetNextNodes() []*Node {
	var nextNodes []*Node
	visited := make(map[string]bool) // To avoid cycles or infinite recursion

	// Helper function to recursively traverse the graph
	var traverse func(n *Node)
	traverse = func(n *Node) {
		// If this node has already been visited, avoid reprocessing it
		if visited[n.ID] {
			return
		}
		visited[n.ID] = true

		// Iterate through the outgoing edges to get the next nodes
		for _, edge := range n.Outgoing {
			nextNode := edge.TargetRef

			// If the next node is a task or event, add it to the result
			if nextNode.Type == "Task" || nextNode.Type == "EndEvent" || nextNode.Type == "IntermediateCatchEvent" || nextNode.Type == "IntermediateThrowEvent" {
				nextNodes = append(nextNodes, nextNode)
			} else if isGateway(nextNode.Type) {
				// If it's a gateway, recursively follow all its outgoing edges
				traverse(nextNode)
			}
		}
	}

	// Start traversal from the current node
	traverse(node)
	return nextNodes
}

// Helper function to determine if a node is a gateway
func isGateway(nodeType string) bool {
	return nodeType == "ExclusiveGateway" || nodeType == "ParallelGateway" || nodeType == "InclusiveGateway"
}

// CheckCompletion returns a recursive structure representing the groups of nodes
// that need to be completed before the current node can proceed.
func (node *Node) GetCompletionGroups() []interface{} {
	var result []interface{}
	visited := make(map[string]bool) // To prevent infinite recursion

	// Helper function to recursively traverse and group nodes
	var traverse func(n *Node) []interface{}
	traverse = func(n *Node) []interface{} {
		// If we have already visited this node, return nil to avoid cycles
		if visited[n.ID] {
			return nil
		}
		visited[n.ID] = true

		// Base case: If the node is a task or event, return it as a single entry
		if n.Type == "Task" || n.Type == "StartEvent" || n.Type == "IntermediateCatchEvent" || n.Type == "IntermediateThrowEvent" {
			return []interface{}{n}
		}

		// If the node is a gateway, handle differently based on the type
		if isExclusiveGateway(n.Type) {
			// For exclusive gateways, only one of the incoming paths needs to be completed
			var exclusiveGroup []interface{}
			for _, edge := range n.Incoming {
				prevNode := edge.SourceRef
				// Recursively check the previous nodes and append them
				exclusiveGroup = append(exclusiveGroup, traverse(prevNode)...)
			}
			// Return a single group for the exclusive gateway (only one path needs to be completed)
			return []interface{}{exclusiveGroup}
		} else if isParallelGateway(n.Type) {
			// For parallel gateways, all incoming paths must be completed
			var parallelGroups []interface{}
			for _, edge := range n.Incoming {
				prevNode := edge.SourceRef
				// Recursively check the previous nodes and append them
				parallelGroups = append(parallelGroups, traverse(prevNode))
			}
			// Return all groups for the parallel gateway (all paths must be completed)
			return parallelGroups
		}

		// For other node types (events, etc.), return an empty group
		return nil
	}

	// Start traversing from the current node
	result = traverse(node)
	return result
}

// Helper function to determine if a node is an Exclusive Gateway
func isExclusiveGateway(nodeType string) bool {
	return nodeType == "ExclusiveGateway"
}

// Helper function to determine if a node is a Parallel Gateway
func isParallelGateway(nodeType string) bool {
	return nodeType == "ParallelGateway"
}

func PrintCompletionGroups(groups []interface{}) {
	fmt.Print('[')
	for _, group := range groups {
		fmt.Print('[')

		if node, ok := group.(*Node); ok {
			fmt.Print(",", node.ID, " ", node.Type)
		} else if subgroups, ok := group.([]interface{}); ok {
			// Recursively print subgroups
			PrintCompletionGroups(subgroups)
		}
		fmt.Print(']')
	}
	fmt.Print(']')
}
