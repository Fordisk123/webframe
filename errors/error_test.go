package errors

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestErrAs(t *testing.T) {

	e := NewBadRequestError("")
	fmt.Println(errors.As(NewBadRequestError(""), &e))

	fmt.Println(errors.As(NewBadRequestError(""), &BadRequestError{}))

}
