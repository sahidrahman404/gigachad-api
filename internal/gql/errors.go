package gql

import (
	"runtime/debug"

	"github.com/sahidrahman404/gigachad-api/internal/validator"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *Resolver) reportError(err error) {
	trace := debug.Stack()

	r.logger.Error(err, trace)
}

func (r *Resolver) errorMessage(v *validator.Validator) *gqlerror.List {
	errList := gqlerror.List{}
	for _, v := range v.FieldErrors {
		errList = append(errList, gqlerror.Errorf(v))
	}
	return &errList
}

func (r *Resolver) serverError(err error) *gqlerror.Error {
	r.reportError(err)

	message := "The server encountered a problem and could not process your request"

	return gqlerror.Errorf(message)
}

func (r *Resolver) notFound() *gqlerror.Error {
	message := "The requested resource could not be found"
	return gqlerror.Errorf(message)
}

func (r *Resolver) editConflict() *gqlerror.Error {
	message := "unable to update the record due to an edit conflict, please try again"
	return gqlerror.Errorf(message)
}

func (r *Resolver) invalidCredentials() *gqlerror.Error {
	message := "invalid authentication credentials"
	return gqlerror.Errorf(message)
}

func (r *Resolver) authenticationRequired() *gqlerror.Error {
	message := "you must be authenticated to access this resource"
	return gqlerror.Errorf(message)
}

func (r *Resolver) inactiveAccount() *gqlerror.Error {
	message := "your user account must be activated to access this resource"
	return gqlerror.Errorf(message)
}

func (r *Resolver) notPremittedError() *gqlerror.Error {
	message := "your user account doesn't have the necessary permissions to access this resource"
	return gqlerror.Errorf(message)
}

func (r *Resolver) invalidEmail() *gqlerror.Error {
	message := "no matching email address found"
	return gqlerror.Errorf(message)
}
