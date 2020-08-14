package command

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
	"unicode/utf8"
)

/**
当前工作路径
*/
func CurrentDir() string {
	path, _ := os.Getwd()
	fmt.Println("当前工作路径:", path)
	return path
}

/**
字符转换
*/
func ConvertEncode(out []byte) []byte {
	var duplicate []byte
	deepCopy(out, duplicate)
	if utf8.Valid(duplicate) {
		fmt.Println("字节切片不是UTF-8编码")
		//GBK转UTF-8
		ret, _ := simplifiedchinese.GBK.NewDecoder().Bytes(out)
		return ret
	}
	return out
}

/**
判断是否是UTF-8,注意类型 只能本包访问,注意会进行修改字节切片
不是返回 true
*/
func validUTF8(buf []byte) bool {
	nBytes := 0
	for i := 0; i < len(buf); i++ {
		if nBytes == 0 {
			if (buf[i] & 0x80) != 0 { //与操作之后不为0，说明首位为1
				for (buf[i] & 0x80) != 0 {
					buf[i] <<= 1 //左移一位
					nBytes++     //记录字符共占几个字节
				}

				if nBytes < 2 || nBytes > 6 { //因为UTF8编码单字符最多不超过6个字节
					return false
				}

				nBytes-- //减掉首字节的一个计数
			}
		} else { //处理多字节字符
			if buf[i]&0xc0 != 0x80 { //判断多字节后面的字节是否是10开头
				return false
			}
			nBytes--
		}
	}
	return nBytes == 0
}

/**
深拷贝
*/
func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
