package validate

import idvalidator "github.com/guanguans/id-validator"

func IdCardVerify(idCard string) bool {
	return idvalidator.IsValid(idCard, false)
}
