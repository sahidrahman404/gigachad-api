package workflow

import (
	"time"

	"github.com/sahidrahman404/gigachad-api/internal/activity"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func NewUser(ctx workflow.Context, input activity.MailDetails) error {
	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    100 * time.Second,
		MaximumAttempts:    0, // unlimited retries
	}

	options := workflow.ActivityOptions{
		// Timeout options specify when to automatically timeout Activity functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failed Activities by default.
		RetryPolicy: retrypolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, options)
	var a *activity.Mailer

	err := workflow.ExecuteActivity(ctx, a.Send, input).Get(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}
