// Code generated by "github.com/hedhyw/semerr"; DO NOT EDIT.

package semerr_test

import (
	"errors"
	"testing"

	"github.com/hedhyw/semerr/pkg/v1/semerr"
)

func TestWrappedErrors(t *testing.T) {
	t.Parallel()

	const err semerr.Error = "error"

	testCases := []struct {
		Name      string
		Create    func(err error) error
		Permanent bool
	}{
		{
			Name:      "BadRequestError",
			Create:    semerr.NewBadRequestError,
			Permanent: true,
		},
		{
			Name:      "ForbiddenError",
			Create:    semerr.NewForbiddenError,
			Permanent: true,
		},
		{
			Name:      "InternalServerError",
			Create:    semerr.NewInternalServerError,
			Permanent: true,
		},
		{
			Name:      "NotFoundError",
			Create:    semerr.NewNotFoundError,
			Permanent: true,
		},
		{
			Name:      "ServiceUnavailableError",
			Create:    semerr.NewServiceUnavailableError,
			Permanent: false,
		},
		{
			Name:      "UnauthorizedError",
			Create:    semerr.NewUnauthorizedError,
			Permanent: true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			errWrapped := tc.Create(err)
			switch {
			case tc.Create(nil) != nil:
				t.Fatal()
			case errWrapped.Error() != err.Error():
				t.Fatal("exp", err.Error(), "got", errWrapped.Error())
			case !errors.Is(errWrapped, err):
				t.Fatal()
			case semerr.IsTemporaryError(errWrapped) == tc.Permanent:
				t.Fatal()
			}
		})
	}
}
