package workflow

// func SendEmail(ctx workflow.Context, data activity.MailDetails) error {
// 	var a *activity.Mailer
// 	mailError, err := workflow.ExecuteActivity[error](ctx, workflow.DefaultActivityOptions, a.Send, data).Get(ctx)
// 	if err != nil {
// 		return err
// 	}
//
// 	if mailError != nil {
// 		return mailError
// 	}
//
// 	return nil
// }
