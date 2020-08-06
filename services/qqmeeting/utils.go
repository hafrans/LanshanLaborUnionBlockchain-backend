package qqmeeting

import "fmt"

// 默认URL Parameter处理方式
func defaultFillPlaceholder(req MeetingRequest, args ...interface{}) string {
	return fmt.Sprintf(req.getDescriptor().Url, args...)
}
