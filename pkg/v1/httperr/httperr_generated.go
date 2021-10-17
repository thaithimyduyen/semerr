// Code generated by "github.com/hedhyw/semerr"; DO NOT EDIT.

package httperr

import (
	"errors"
	"net/http"

	"github.com/hedhyw/semerr/pkg/v1/semerr"
)

// Code returns http status code for err.
func Code(err error) int {
	switch err.(type) {
	case nil:
		return http.StatusOK
	case semerr.BadRequestError:
		return 400
	case semerr.ForbiddenError:
		return 403
	case semerr.InternalServerError:
		return 500
	case semerr.NotFoundError:
		return 404
	case semerr.ServiceUnavailableError:
		return 503
	case semerr.UnauthorizedError:
		return 401
	}

	if err = errors.Unwrap(err); err == nil {
		return http.StatusInternalServerError
	}

	return Code(err)
}