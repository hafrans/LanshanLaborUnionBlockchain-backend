package smsqueue

type Request interface {
	getDescriptor() string
}

type Authenticator struct {
	Account  string `schema:"account,required"`
	Password string `schema:"password,required"`
}

type SendRequest struct {
	Mobile  string `schema:"mobile,required"`
	Content string `schema:"content,required"`
}

func (req *SendRequest) getDescriptor() string {
	return "send"
}

type ReportRequest struct {
	TaskId string `schema:"taskid"`
}

func (req *ReportRequest) getDescriptor() string {
	return "report"
}

type UnicomSMSSender struct {
	Authenticator *Authenticator
}
