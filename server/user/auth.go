package user

import (
	"fmt"
	"projects/resellbay/server/cfg"
	"projects/resellbay/server/util"
	"projects/sample-tutorials/server/pkg/crypt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/dgrijalva/jwt-go"
)

//Login logs user into the app
func (u *User) Login(input *LoginInput) *LoginOutput {
	item := u.GetByEmail(input.Useremail)
	if *item.Count > 0 && crypt.ComparePasswords(input.Password, *item.Items[0]["password"].S) {
		token := createToken(*item.Items[0]["userID"].S, input)
		if u.isAdmin(*item.Items[0]["userID"].S) {
			return &LoginOutput{true, true, token, *item.Items[0]["userID"].S, "Login successful"}
		}
		return &LoginOutput{true, false, token, *item.Items[0]["userID"].S, "Login successful"}
	}
	return &LoginOutput{Valid: false, Msg: "Login failed: Wrong Email or Password"}
}

func signupValid(input *AddInput) *Output {
	if !govalidator.IsEmail(input.Useremail) {
		return &Output{false, "Email Invalid", "Email"}
	}
	if !passwordValid(input.Password) {
		return &Output{false, "Password must be 8 characters or more", "Password"}
	}

	if !(input.Password == input.PasswordRepeat) {
		return &Output{false, "Passwords don't match", "Password"}
	}

	return &Output{true, "", ""}
}

func passwordValid(password string) bool {
	if len(password) >= 8 {
		return true
	}
	return false
}

func (u *User) changePasswordValid(input ChangePasswordInput, userID string) *Output {
	if !(passwordValid(input.New) && passwordValid(input.Old)) {
		return &Output{false, "Please Provide all Passwords", ""}
	}
	if !(input.New == input.NewRepeat) {
		return &Output{false, "Password's dont match", ""}
	}
	user := u.Get(userID)
	if !crypt.ComparePasswords(input.Old, *user.Item["password"].S) {
		return &Output{false, "Old Password invalid", "Old"}
	}
	return u.ChangePassword(userID, input.New)
}

func (u *User) isAdmin(userID string) bool {
	user := u.Get(userID)
	if user.Item["role"] != nil && *user.Item["role"].S == "admin" {
		return true
	}
	return false
}

func createToken(userID string, input *LoginInput) string {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	token.Claims["userID"] = userID
	token.Claims["useremail"] = input.Useremail
	token.Claims["password"] = input.Password
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(cfg.SigningKey)
	util.CheckErr(err)
	return tokenString
}

func verifyToken(myToken string, userID string) bool {
	token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		if token.Claims["userID"] == userID {
			return cfg.SigningKey, nil
		}

		return "fail", nil

	})
	util.CheckErr(err)

	if err == nil && token.Valid {
		return true
	}
	return false

}
