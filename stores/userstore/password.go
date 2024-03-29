package userstore

import "golang.org/x/crypto/bcrypt"

func (u *userImpl) hashPassword(password string) ([]byte, error) {
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return []byte{}, err
	}

	return hashedPassword, nil
}
