package common

const (
	// Inactive is the state representing an inactive process
	STATE_INACTIVE = 0
	// Ready is the state representing a process that is ready to be started
	STATE_READY = 1
	// Active is the state representing a process that is currently running
	STATE_ACTIVE = 2
	// Completing is the state representing a process that is finishing
	STATE_COMPLETING = 3
	// Failing is the state representing a process that is failing
	STATE_FAILING = 4
	// Terminating is the state representing a process that is terminating
	STATE_TERMINATING = 5
	// Compensating is the state representing a process that is compensating
	STATE_COMPENSATING = 6
	// Terminated is the state representing a process that has terminated
	STATE_TERMINATED = 7
	// Compensated is the state representing a process that has been compensated
	STATE_COMPENSATED = 8
	// Failed is the state representing a process that has failed
	STATE_FAILED = 9
	// Completed is the state representing a process that has completed
	STATE_COMPLETED = 10
)
