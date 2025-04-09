package bsql

import (
	"log"
	"testing"
	"unsafe"

	"golang.org/x/crypto/bcrypt"
)

func TestMasterAdminAdd(t *testing.T) {
	generate := func(password string, cost int) ([]byte, error) {
		p, err := bcrypt.GenerateFromPassword([]byte(password), cost)
		if err != nil {
			return nil, err
		}

		return p, nil
	}

	pw, err := generate("12345", bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(ZeroCopyByteToString(pw), len(pw))

}

/*
byte to string with zero copy

@warning
전달되는 b의 데이터가 수정이 되는 경우 좋지 않음.
*/
func ZeroCopyByteToString(b []byte) string {
	return *((*string)(unsafe.Pointer(&b)))
}
