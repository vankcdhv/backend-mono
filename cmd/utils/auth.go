package utils

import (
	"backend-mono/cmd/utils/constant"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
)

func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
func ComparePasswords(hashedPwd string, originPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(originPwd))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
func ValidateUsername(username string) (bool, string) {
	userNameMinLength := viper.GetInt("account.username_min_length")
	if len(username) < userNameMinLength {
		return false, constant.UsernameToShort
	}
	return true, ""
}
func ValidateEmail(email string) (bool, string) {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(email) {
		return false, constant.EmailNotValid
	}
	return true, ""

}
func ValidatePhoneNumber(phoneNumber string) (bool, string) {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if !re.MatchString(phoneNumber) {
		return false, constant.PhoneNotValid
	}
	return true, ""
}
func ValidatePassword(pwd string) (bool, string) {
	passwordMinLength := viper.GetInt("account.password_min_length")
	if len(pwd) < passwordMinLength {
		return false, constant.PasswordToShort
	}
	return true, ""
}
