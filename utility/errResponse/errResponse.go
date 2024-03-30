package errResponse

import (
	"errors"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func DbOperationError(s string) error {
	return gerror.WrapCode(gcode.New(
		gcode.CodeDbOperationError.Code(),
		"",
		nil),
		errors.New(s),
		gcode.CodeDbOperationError.Message(),
	)
}

func DbOperationErrorDefault() error {
	return gerror.WrapCode(gcode.New(
		gcode.CodeDbOperationError.Code(),
		"",
		nil),
		errors.New("查询失败"),
		gcode.CodeDbOperationError.Message(),
	)
}

func OperationFailed(s string) error {
	return gerror.WrapCode(gcode.New(
		gcode.CodeOperationFailed.Code(),
		"",
		nil),
		errors.New(s),
		gcode.CodeOperationFailed.Message(),
	)
}

func NotAuthorized(s string) error {
	return gerror.WrapCode(gcode.New(
		gcode.CodeNotAuthorized.Code(),
		"",
		nil),
		errors.New(s),
		gcode.CodeNotAuthorized.Message(),
	)
}

func Unknown(s string) error {
	return gerror.WrapCode(gcode.New(
		gcode.CodeUnknown.Code(),
		"",
		nil),
		errors.New(s),
		gcode.CodeUnknown.Message(),
	)
}

func CodeInvalidParameter(s string) error {
	return gerror.WrapCode(gcode.New(
		gcode.CodeInvalidParameter.Code(),
		"",
		nil),
		errors.New(s),
		gcode.CodeInvalidParameter.Message(),
	)
}
