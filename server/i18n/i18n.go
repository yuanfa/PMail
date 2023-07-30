package i18n

var (
	cn = map[string]string{
		"all_email":   "全部邮件数据",
		"inbox":       "收件箱",
		"outbox":      "发件箱",
		"sketch":      "草稿箱",
		"aperror":     "账号或密码错误",
		"unknowError": "未知错误",
		"succ":        "成功",
		"send_fail":   "发送失败",
		"att_err":     "附件解码错误",
	}
	en = map[string]string{
		"all_email":   "All Email",
		"inbox":       "Inbox",
		"outbox":      "Outbox",
		"sketch":      "Sketch",
		"aperror":     "Incorrect account number or password",
		"unknowError": "Unknow Error",
		"succ":        "Success",
		"send_fail":   "Send Failure",
		"att_err":     "Attachment decoding error",
	}
)

func GetText(lang, key string) string {
	if lang == "zhCn" {
		text, exist := cn[key]
		if !exist {
			return ""
		}
		return text
	}
	text, exist := en[key]
	if !exist {
		return ""
	}
	return text
}
