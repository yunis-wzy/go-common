package status

// 51-200 文件服务错误
var ErrNotEnoughSpace = NewStatus(51, "NotEnoughSpace", "NotEnoughSpace")
var ErrBeenDelete = NewStatus(52, "ErrBeenDelete", "ErrBeenDelete")
var ErrDuplicateName = NewStatus(53, "ErrDuplicateName", "ErrDuplicateName")
var ErrBlockNotUpload = NewStatus(54, "BlockNotUpload", "BlockNotUpload")
var ErrNodeNotFound = NewStatus(55, "ErrNodeNotFound", "ErrNodeNotFound")
var ErrOperateExpand = NewStatus(56, "ErrOperateExpand", "ErrOperateExpand")
var ErrBlockIllegal = NewStatus(57, "ErrBlockIllegal", "ErrBlockIllegal")
var ErrBlockNotExist = NewStatus(58, "ErrBlockNotExist", "ErrBlockNotExist")
var ErrShareLimitExceed = NewStatus(59, "ErrShareLimitExceed", "ErrShareLimitExceed")
var ErrUserNotInShareGroup = NewStatus(60, "ErrUserNotInShareGroup", "ErrUserNotInShareGroup")
var ErrUserHasNoPermit = NewStatus(61, "ErrUserHasNoPermit", "ErrUserHasNoPermit")
var ErrUpdatedInCloud = NewStatus(62, "ErrUpdatedInCloud", "ErrUpdatedInCloud")
var ErrDuplicateAppID = NewStatus(63, "ErrDuplicateAppID", "ErrDuplicateAppID")

// 201-300 文件集合错误
var ErrSetKeyExist = NewStatus(201, "ErrSetKeyExist", "setKey已存在")
var ErrSetKeyNotExist = NewStatus(202, "ErrSetKeyNotExist", "setKey不存在")
var ErrSetLevel = NewStatus(203, "ErrSetLevel", "无权限的setLevel")
var ErrFileNotExist = NewStatus(204, "ErrFileNotExist", "部分文件不存在")
