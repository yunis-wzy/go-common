package status

// 1-50 通用错误
var (
	CommonSuccess       = NewStatus(0, "success", "success")
	CommonUnknown       = NewStatus(999, "unknown", "unknown")
	CommonInternalError = NewStatus(1, "internalError", "internal error")

	// version.
	CommonVersionNotSupported = NewStatus(2, "VersionNotSupported", "version not supported")
	// user.
	CommonUserNoLogin       = NewStatus(3, "noLogin", "user no login")
	CommonUserNotExist      = NewStatus(4, "UserNotExist", "user not exist")
	CommonUserAuthorityDeny = NewStatus(5, "UserAuthorityDeny", "user authority deny")
	CommonUserPunished      = NewStatus(6, "Punished", "user punished")
	// param.
	CommonLackRequiredParams = NewStatus(7, "LackRequiredParams", "lack required params")
	CommonInvalidParams      = NewStatus(8, "InvalidParams", "invalid params")
	CommonInvalidResult      = NewStatus(10, "InvalidResult", "invalid result")
	// sign.
	CommonInvalidSign = NewStatus(9, "InvalidSign", "invalid sign")
	// verify
	CommonVerifyFailed = NewStatus(11, "VerifyFailed", "verify failed")
)
