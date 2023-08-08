package validate

import idvalidator "github.com/guanguans/id-validator"

// IdCardVerify
//  @Description: 身份证号码的合法性校验
//  @param idCard 身份证号
//  @return bool 是否合法

func IdCardVerify(idCard string) bool {
	return idvalidator.IsValid(idCard, false)
}
