package bpmn

import "encoding/xml"

// BPMN Definitions struct
type Definitions struct {
	XMLName xml.Name      `xml:"definitions"`
	Collab  Collaboration `xml:"collaboration"`
	Process []Process     `xml:"process"`
}

// Collaboration struct
type Collaboration struct {
	ID          string        `xml:"id,attr"`
	Participant []Participant `xml:"participant"`
	MessageFlow []MessageFlow `xml:"messageFlow"`
}

// Participant struct
type Participant struct {
	ID         string `xml:"id,attr"`
	Name       string `xml:"name,attr"`
	ProcessRef string `xml:"processRef,attr"`
	PublicKey  string `xml:"publicKey,attr"`
}

// Process struct
type Process struct {
	ID                     string                   `xml:"id,attr"`
	IsExecutable           bool                     `xml:"isExecutable,attr"`
	StartEvent             []StartEvent             `xml:"startEvent"`
	Tasks                  []Task                   `xml:"task"`
	IntermediateCatchEvent []IntermediateCatchEvent `xml:"intermediateCatchEvent"`
	IntermediateThrowEvent []IntermediateThrowEvent `xml:"intermediateThrowEvent"`
	SequenceFlows          []SequenceFlow           `xml:"sequenceFlow"`
	EndEvent               []EndEvent               `xml:"endEvent"`
	ExclusiveGateways      []ExclusiveGateway       `xml:"exclusiveGateway"`
	ParallelGateways       []ParallelGateway        `xml:"parallelGateway"`
	InclusiveGateways      []InclusiveGateway       `xml:"inclusiveGateway"`
}

// StartEvent struct
type StartEvent struct {
	ID       string `xml:"id,attr"`
	Outgoing string `xml:"outgoing"`
}

// Task struct
type Task struct {
	ID       string `xml:"id,attr"`
	Name     string `xml:"name,attr"`
	Incoming string `xml:"incoming"`
	Outgoing string `xml:"outgoing"`

	// Custom fields
	Variables   string `xml:"variables,attr"`
	Type        string `xml:"type,attr"`
	Participant string `xml:"participant,attr"`
	Amount      string `xml:"amount,attr"`
}

// IntermediateCatchEvent struct
type IntermediateCatchEvent struct {
	ID       string `xml:"id,attr"`
	Incoming string `xml:"incoming"`
	Outgoing string `xml:"outgoing"`
}

// IntermediateThrowEvent struct
type IntermediateThrowEvent struct {
	ID       string `xml:"id,attr"`
	Incoming string `xml:"incoming"`
	Outgoing string `xml:"outgoing"`
}

// SequenceFlow struct
type SequenceFlow struct {
	ID        string `xml:"id,attr"`
	Name      string `xml:"name,attr"`
	SourceRef string `xml:"sourceRef,attr"`
	TargetRef string `xml:"targetRef,attr"`
}

// EndEvent struct
type EndEvent struct {
	ID       string `xml:"id,attr"`
	Incoming string `xml:"incoming"`
}

/**************************************************
 *                     GATES                      *
 **************************************************/

type ExclusiveGateway struct {
	ID       string   `xml:"id,attr"`
	Name     string   `xml:"name,attr"`
	Default  string   `xml:"default,attr"`
	Incoming []string `xml:"incoming"`
	Outgoing []string `xml:"outgoing"`
}

type ParallelGateway struct {
	ID       string   `xml:"id,attr"`
	Name     string   `xml:"name,attr"`
	Incoming []string `xml:"incoming"`
	Outgoing []string `xml:"outgoing"`
}

type InclusiveGateway struct {
	ID       string   `xml:"id,attr"`
	Name     string   `xml:"name,attr"`
	Incoming []string `xml:"incoming"`
	Outgoing []string `xml:"outgoing"`
}

type Lane struct {
	ID       string   `xml:"id,attr"`
	FlowRefs []string `xml:"flowNodeRef"`
}

type LaneSet struct {
	ID    string `xml:"id,attr"`
	Lanes []Lane `xml:"lane"`
}

type Message struct {
	Id string `xml:"id,attr"`
}

type MessageFlow struct {
	ID        string `xml:"id,attr"`
	SourceRef string `xml:"sourceRef,attr"`
	TargetRef string `xml:"targetRef,attr"`
}
