package hash

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T){
	const file_md5 = "3dfcbc696f70c39a6932a7872ccd8cf0"
	_file_md5,err := FileMd5("./md5.go")
	if err != nil{
		panic(err)
	}

	fmt.Printf("file md5 === %s\n",_file_md5)

	if file_md5 != _file_md5 {
		t.Errorf("expect file md5 is %s, but had %s\n",file_md5,_file_md5)
	}

	const str = "golang mod study"
	const str_md5 = "f3db0e3fda35fc0a051ae2caf7f60107"
	_str_md5 := StringMd5(str)
	if str_md5 != _str_md5 {
		t.Errorf("expect str md5 is %s, but had %s \n",str_md5,_str_md5)
	}
}
