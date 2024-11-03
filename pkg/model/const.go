package model

const (
	// BPMN types
	START_EVENT              = "StartEvent"
	END_EVENT                = "EndEvent"
	TASK                     = "Task"
	INTERMEDIATE_CATCH_EVENT = "IntermediateCatchEvent"
	INTERMEDIATE_THROW_EVENT = "IntermediateThrowEvent"
	EXCLUSIVE_GATEWAY        = "ExclusiveGateway"
	PARALLEL_GATEWAY         = "ParallelGateway"
	INCLUSIVE_GATEWAY        = "InclusiveGateway"
	PAYMENT_TASK             = "PaymentTask"
)
