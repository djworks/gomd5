package hash

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

// file md5
func FileMd5(filepath string) (string,error){
	file, err := os.Open(filepath)
	if err != nil {
		return "",errors.New(fmt.Sprintf("md5.go hash.FileMd5 os open file error: %v ",err))
	}

	h := md5.New()
	_, err = io.Copy(h,file)
	if err != nil {
		return "",errors.New(fmt.Sprintf("md5.go hash.FileMd5 io copy error: %v",err))
	}

	return hex.EncodeToString(h.Sum(nil)),nil
}

// string md5
func StringMd5(str string)string{
	_md5 := md5.New()
	_md5.Write([]byte(str))
	return hex.EncodeToString(_md5.Sum(nil))
}
