package out

type MailPort interface {
	Send(templateFile, username, email string, data any, isSandbox bool) (int, error)
}
