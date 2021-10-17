// Code generated by "github.com/hedhyw/semerr"; DO NOT EDIT.

package httperr_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hedhyw/semerr/pkg/v1/httperr"
	"github.com/hedhyw/semerr/pkg/v1/semerr"
)

func TestCode(t *testing.T) {
	t.Parallel()

	const err = semerr.Error("some error")

	testCases := []struct {
		Err  error
		Code int
	}{
		{
			Err:  nil,
			Code: http.StatusOK,
		},
		{
			Err:  err,
			Code: http.StatusInternalServerError,
		},
		{
			Err:  semerr.NewBadRequestError(err),
			Code: 400,
		},
		{
			Err:  semerr.NewForbiddenError(err),
			Code: 403,
		},
		{
			Err:  semerr.NewInternalServerError(err),
			Code: 500,
		},
		{
			Err:  semerr.NewNotFoundError(err),
			Code: 404,
		},
		{
			Err:  semerr.NewServiceUnavailableError(err),
			Code: 503,
		},
		{
			Err:  semerr.NewUnauthorizedError(err),
			Code: 401,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(fmt.Sprint(tc.Err), func(t *testing.T) {
			t.Parallel()

			gotCode := httperr.Code(tc.Err)
			if tc.Code != gotCode {
				t.Fatal("exp", tc.Code, "got", gotCode)
			}
		})
	}
}