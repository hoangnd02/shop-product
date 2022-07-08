package params

type Error struct {
	Code    int
	Message string
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func (e *Error) Error() string {
	return e.Message
}

var (
	FailedToParseBody = NewError(500, "failed to parse")

	ServerInternalError = NewError(500, "Server internal error")

	FailedConnectDataInDatabase = NewError(500, "Failed to connect data in database")

	FailedToParseJWT = NewError(500, "Failed to parse jwt")

	FailedConnectToSessions = NewError(500, "Failed to connect to sessions")

	ServerInvalidQueryErr = NewError(500, "server.invalid_query")

	AuthzInvalidPermissionErr = NewError(500, "authz.invalid_permission")

	AuthzCsrfTokenMismatchErr = NewError(500, "authz.csrf_token_mismatch")

	AuthzMissingCsrfTokenErr = NewError(500, "authz.missing_csrf_token")

	AuthzClientSessionMismatchErr = NewError(500, "authz.client_session_mismatch")

	AuthzUserNotActiveErr = NewError(500, "authz.user_not_active")

	AuthzUserNotExistErr = NewError(500, "authz.user_not_exist")

	AuthzInvalidSessionErr = NewError(500, "authz.invalid_session")

	AuthZPermissionDeniedErr = NewError(500, "authz.permission_denied")

	RecordNotFoundErr = NewError(500, "record.not_found")

	UserNotFound = NewError(500, "user.not_found")

	OrderIsNotPending = NewError(500, "order.is_not_pending")
)
