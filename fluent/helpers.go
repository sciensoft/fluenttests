package fluent

func GetMessage(original string, messagesf ...string) string {
	if len(messagesf) > 0 {
		original = messagesf[0]
	}
	return original
}
