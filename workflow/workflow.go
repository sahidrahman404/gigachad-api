package workflow

import (
	"github.com/cschleiden/go-workflows/workflow"
	"github.com/sahidrahman404/gigachad-api/activity"
)

func SendEmail(ctx workflow.Context, data activity.MailDetails) error {
	var a *activity.Mailer
	mailError, err := workflow.ExecuteActivity[error](ctx, workflow.DefaultActivityOptions, a.Send, data).Get(ctx)
	if err != nil {
		return err
	}

	if mailError != nil {
		return mailError
	}

	return nil
}
