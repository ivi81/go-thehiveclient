package cons

import "net/http"

const (
	HIVEReqSuccessfully = http.StatusOK
	HIVEReqCreated      = http.StatusCreated
	HIVEReqDeleted      = http.StatusNoContent

	HIVEReqErrAttrChecking   = http.StatusBadRequest
	HIVEReqErrUthorization   = http.StatusUnauthorized
	HIVEReqErrAuthentication = http.StatusForbidden
	HIVEReqErrNotFound       = http.StatusNotFound
)
