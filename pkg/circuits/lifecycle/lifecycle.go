package lifecycle

import (
	"github.com/consensys/gnark/frontend"
	"github.com/ftsrg/zkWF/pkg/circuits/utils"
	"github.com/ftsrg/zkWF/pkg/common"
)

type Circuit struct {
	State_curr frontend.Variable
	State_new  frontend.Variable
}

func LifecycleCheck(api frontend.API, State_curr, State_new frontend.Variable) frontend.Variable {
	sum := make([]frontend.Variable, 11)

	inactiveSelector := utils.IsEqual(api, State_curr, common.STATE_INACTIVE)
	inactiveToReady := utils.IsEqual(api, State_new, 1)
	sum[0] = api.Select(inactiveSelector, inactiveToReady, 0)

	readySelector := utils.IsEqual(api, State_curr, common.STATE_READY)
	readyToActive := utils.IsEqual(api, State_new, common.STATE_ACTIVE)
	readyToFailing := utils.IsEqual(api, State_new, common.STATE_FAILING)
	readyToTerminating := utils.IsEqual(api, State_new, common.STATE_TERMINATING)
	readySum := make([]frontend.Variable, 3)
	readySum[0] = api.Select(readySelector, readyToActive, 0)
	readySum[1] = api.Add(api.Select(readySelector, readyToFailing, 0), readySum[0])
	readySum[2] = api.Add(api.Select(readySelector, readyToTerminating, 0), readySum[1])
	sum[1] = api.Add(sum[0], readySum[2])

	activeSelector := utils.IsEqual(api, State_curr, common.STATE_ACTIVE)
	activeToCompleting := utils.IsEqual(api, State_new, common.STATE_COMPLETING)
	activeToFailing := utils.IsEqual(api, State_new, common.STATE_FAILING)
	activeToTerminating := utils.IsEqual(api, State_new, common.STATE_TERMINATING)
	activeSum := make([]frontend.Variable, 3)
	activeSum[0] = api.Select(activeSelector, activeToCompleting, 0)
	activeSum[1] = api.Add(api.Select(activeSelector, activeToFailing, 0), activeSum[0])
	activeSum[2] = api.Add(api.Select(activeSelector, activeToTerminating, 0), activeSum[1])
	sum[2] = api.Add(sum[1], activeSum[2])

	completingSelector := utils.IsEqual(api, State_curr, common.STATE_COMPLETING)
	completingToCompleted := utils.IsEqual(api, State_new, common.STATE_COMPLETED)
	completingToFailing := utils.IsEqual(api, State_new, common.STATE_FAILING)
	completingToTerminating := utils.IsEqual(api, State_new, common.STATE_TERMINATING)
	completingSum := make([]frontend.Variable, 3)
	completingSum[0] = api.Select(completingSelector, completingToCompleted, 0)
	completingSum[1] = api.Add(api.Select(completingSelector, completingToFailing, 0), completingSum[0])
	completingSum[2] = api.Add(api.Select(completingSelector, completingToTerminating, 0), completingSum[1])
	sum[3] = api.Add(sum[2], completingSum[2])

	failingSelector := utils.IsEqual(api, State_curr, common.STATE_FAILING)
	failingToFailed := utils.IsEqual(api, State_new, common.STATE_FAILED)
	sum[4] = api.Add(sum[3], api.Select(failingSelector, failingToFailed, 0))

	terminatingSelector := utils.IsEqual(api, State_curr, common.STATE_TERMINATING)
	terminatingToTerminated := utils.IsEqual(api, State_new, common.STATE_TERMINATED)
	terminatingToFailed := utils.IsEqual(api, State_new, common.STATE_FAILED)
	terminatingSum := make([]frontend.Variable, 2)
	terminatingSum[0] = api.Select(terminatingSelector, terminatingToTerminated, 0)
	terminatingSum[1] = api.Add(api.Select(terminatingSelector, terminatingToFailed, 0), terminatingSum[0])
	sum[5] = api.Add(sum[4], terminatingSum[1])

	compensatingSelector := utils.IsEqual(api, State_curr, common.STATE_COMPENSATING)
	compensatingToCompensated := utils.IsEqual(api, State_new, common.STATE_COMPENSATED)
	compensatingToFailed := utils.IsEqual(api, State_new, common.STATE_FAILED)
	compensatingToTerminated := utils.IsEqual(api, State_new, common.STATE_TERMINATED)
	compensatingSum := make([]frontend.Variable, 3)
	compensatingSum[0] = api.Select(compensatingSelector, compensatingToCompensated, 0)
	compensatingSum[1] = api.Add(api.Select(compensatingSelector, compensatingToFailed, 0), compensatingSum[0])
	compensatingSum[2] = api.Add(api.Select(compensatingSelector, compensatingToTerminated, 0), compensatingSum[1])
	sum[6] = api.Add(sum[5], compensatingSum[2])

	terminatedSelector := utils.IsEqual(api, State_curr, common.STATE_TERMINATED)
	terminatedToInactive := utils.IsEqual(api, State_new, common.STATE_INACTIVE)
	sum[7] = api.Add(sum[6], api.Select(terminatedSelector, terminatedToInactive, 0))

	compensatedSelector := utils.IsEqual(api, State_curr, common.STATE_COMPENSATED)
	compensatedToInactive := utils.IsEqual(api, State_new, common.STATE_INACTIVE)
	sum[8] = api.Add(sum[7], api.Select(compensatedSelector, compensatedToInactive, 0))

	failedSelector := utils.IsEqual(api, State_curr, common.STATE_FAILED)
	failedToInactive := utils.IsEqual(api, State_new, common.STATE_INACTIVE)
	sum[9] = api.Add(sum[8], api.Select(failedSelector, failedToInactive, 0))

	completedSelector := utils.IsEqual(api, State_curr, common.STATE_COMPLETED)
	completedToInactive := utils.IsEqual(api, State_new, common.STATE_INACTIVE)
	completedToCompensating := utils.IsEqual(api, State_new, common.STATE_COMPENSATING)
	completedSum := make([]frontend.Variable, 2)
	completedSum[0] = api.Select(completedSelector, completedToInactive, 0)
	completedSum[1] = api.Add(api.Select(completedSelector, completedToCompensating, 0), completedSum[0])
	sum[10] = api.Add(sum[9], completedSum[1])

	// justOne <== IsEqual()([sum[10], 1]);

	return utils.IsEqual(api, sum[10], 1)
}
