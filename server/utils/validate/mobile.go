package validate

import "regexp"

// MobileVerify
//  @Description: 手机号合法性校验
//  @param mobileNum 手机号
//  @return bool 是否校验通过

func MobileVerify(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
