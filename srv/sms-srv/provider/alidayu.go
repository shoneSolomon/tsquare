package provider

/**
阿里大鱼 短信通道
http://www.ihuyi.cn/
**/

type Alidayu struct {
	sms *SMS
}

func (ali *Alidayu) Send(sms *SMS) error {

	return nil
}
