package provider

/*
阿里云-云通信
*/

type Yuntongxun struct {
	sms *SMS
}

func (y *Yuntongxun) Send(sms *SMS) error {
	return nil
}
