package utils

import "golang.org/x/crypto/bcrypt"

// BcryptHash
//
//	@Description: 密码加密
//	@param password 密码
//	@return string 加密后的密码
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck
//  @Description: 密码校验
//  @param password 未加密密码
//  @param hash 加密后的密码(数据库)
//  @return bool 是否校验成功

func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
