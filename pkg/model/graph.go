package model

import "math/big"

// BPMNGraph struct
type BPMNGraph struct {
	Nodes           map[string]*Node // Map of element IDs to Nodes
	Edges           map[string]*Edge // Map of flow IDs to Edges
	Variables       []string
	MessageCount    int
	MessageMap      map[string]string
	ParticpnatCount int
}

// Node struct representing a BPMN element
type Node struct {
	ID        string
	Name      string
	Type      string  // Type of BPMN element (e.g., Task, StartEvent, EndEvent, etc.)
	Incoming  []*Edge // Edges leading to this node
	Outgoing  []*Edge // Edges leading from this node
	Variables []string
	Owner     Participant
	Payment   Payment
}

type Participant struct {
	ID        string
	Name      string
	PublicKey [2]big.Int
}

type Payment struct {
	Receiver string
	Amount   string
}

// Edge struct representing a sequence flow
type Edge struct {
	ID        string
	Name      string
	SourceRef *Node // Source node (where the flow starts)
	TargetRef *Node // Target node (where the flow ends)
}

// NewBPMNGraph initializes an empty BPMN graph
func NewBPMNGraph() *BPMNGraph {
	return &BPMNGraph{
		Nodes:        make(map[string]*Node),
		Edges:        make(map[string]*Edge),
		Variables:    make([]string, 0),
		MessageCount: 0,
		MessageMap:   make(map[string]string),
	}
}
