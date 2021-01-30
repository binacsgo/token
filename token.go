package token

import (
	"time"
)

var TokenSalt = "Binacs_Token_Salt"

// genToken generate token
func genToken(id string) string {
	t := time.Now().Unix()
	buffer := NewBuffer().Append(t).Append(id).Append(TokenSalt)
	md5 := Md5(buffer.String())
	buffer = NewBuffer().Append(t).Append(md5).Append("id").Append(id)
	return buffer.String()
}

// GenTokenWithRefresh generate token
func GenTokenWithRefresh(id, refresh string) string {
	return genToken(id) + "F" + refresh
}

// GenTokenAndRefresh generate token
func GenTokenAndRefresh(id string) (string, string) {
	token, refresh := genToken(id), genToken(id)
	return token + "F" + refresh, refresh
}
