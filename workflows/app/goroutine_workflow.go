package app

import (
	"fmt"
	"math/rand"
	"time"

	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/workflow"
)

/**
* selfnote: Modified from go sample goroutine + schedules
* This sample workflow demonstrates how to use multiple Temporal goroutines (instead of native goroutine) to process a
* a sequence of activities in parallel.
* In Temporal workflow, you should create goroutines using workflow.Go method.
 */

// SampleScheduledGoroutineWorkflow workflow definition
func SampleScheduledGoroutineWorkflow(ctx workflow.Context, parallelism int) (results []string, err error) {
	// Extract schedule metadata
	info := workflow.GetInfo(ctx)

	// Workflow Executions started by a Schedule have the following additional properties appended to their search attributes
	//lint:ignore SA1019 - this is a sample
	scheduledByIDPayload := info.SearchAttributes.IndexedFields["TemporalScheduledById"]
	var scheduledByID string
	err = converter.GetDefaultDataConverter().FromPayload(scheduledByIDPayload, &scheduledByID)
	if err != nil {
		return []string{}, err
	}
	//lint:ignore SA1019 - this is a sample
	startTimePayload := info.SearchAttributes.IndexedFields["TemporalScheduledStartTime"]
	var startTime time.Time
	err = converter.GetDefaultDataConverter().FromPayload(startTimePayload, &startTime)
	if err != nil {
		return []string{}, err
	}

	// Set activity options on ctx1 (like schedule example)
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx1 := workflow.WithActivityOptions(ctx, ao)

	for i := 0; i < parallelism; i++ {
		input := fmt.Sprintf("job-%d|scheduleID:%s|startTime:%s", i, scheduledByID, startTime.Format(time.RFC3339)) // Should be outside lambda to be captured correctly
		// Start a goroutine in a workflow safe way
		workflow.Go(ctx1, func(gCtx workflow.Context) {
			var result1 string
			err = workflow.ExecuteActivity(gCtx, Step1, input).Get(gCtx, &result1)
			if err != nil {
				// Very naive error handling. Only the last error will be returned by the workflow
				return
			}
			var result2 string
			err = workflow.ExecuteActivity(gCtx, Step2, result1).Get(gCtx, &result2)
			if err != nil {
				return
			}
			results = append(results, result2)
		})
	}

	_ = workflow.Await(ctx, func() bool {
		return err != nil || len(results) == parallelism
	})
	return
}

func Step1(input string) (output string, err error) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return input + ", Step1", nil
}

func Step2(input string) (output string, err error) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return input + ", Step2", nil
}
