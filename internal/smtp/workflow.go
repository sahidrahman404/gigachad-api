package smtp

// func SendDailyWorkoutReminderWorkflow(ctx workflow.Context, input DailyWorkoutReminderInput) error {
// 	var mailer *Mailer
//
// 	// RetryPolicy specifies how to automatically handle retries if an Activity fails.
// 	retrypolicy := &temporal.RetryPolicy{
// 		InitialInterval:    time.Second,
// 		BackoffCoefficient: 2.0,
// 		MaximumInterval:    100 * time.Second,
// 		MaximumAttempts:    1, // unlimited retries
// 	}
//
// 	options := workflow.ActivityOptions{
// 		// Timeout options specify when to automatically timeout Activity functions.
// 		StartToCloseTimeout: time.Minute,
// 		// Optionally provide a customized RetryPolicy.
// 		// Temporal retries failed Activities by default.
// 		RetryPolicy: retrypolicy,
// 	}
//
// 	// Apply the options.
// 	ctx = workflow.WithActivityOptions(ctx, options)
//
// 	return workflow.ExecuteActivity(ctx, mailer.SendDailyWorkoutReminder, input).Get(ctx, nil)
// }
