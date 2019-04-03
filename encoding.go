package feutil

import (
	"bytes"
	"encoding/binary"
	"math"
	"strconv"

	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"strings"
)

const hextable = "0123456789abcdef"

func EncodedLen(n int) int {
	return n * 2
}

func Encode(dst, src []byte) int {
	for i, v := range src {
		dst[i*2] = hextable[v>>4]
		dst[i*2+1] = hextable[v&0x0f]
	}
	return len(src) * 2
}

//生成GUID编码
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func GetSerialNumber() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return strings.ToUpper(GetMd5String(base64.URLEncoding.EncodeToString(b)))
}

func GetGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func BytesReverse(src []byte) []byte {
	for i, j := 0, len(src)-1; i < j; i, j = i+1, j-1 {
		src[i], src[j] = src[j], src[i]
	}
	return src
}

func EncodeByteArrayToString(src []byte) string {
	dst := make([]byte, EncodedLen(len(src)))
	Encode(dst, src)
	return string(dst)
}

//用于UTC时间戳转换
func IntToBytes(n int64) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

//字节数组转整数
func BytesToInt(b []byte) int64 {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int64(tmp)
}

///////////////////////////////////////////////////
func Bin(i int, prefix bool) string {
	i64 := int64(i)

	if prefix {
		return "0b" + strconv.FormatInt(i64, 2) // base 2 for binary
	} else {
		return strconv.FormatInt(i64, 2) // base 2 for binary
	}
}

func Bin2int(binStr string) int {

	// base 2 for binary
	result, _ := strconv.ParseInt(binStr, 2, 64)
	return int(result)
}

func Oct(i int, prefix bool) string {
	i64 := int64(i)

	if prefix {
		return "0o" + strconv.FormatInt(i64, 8) // base 8 for octal
	} else {
		return strconv.FormatInt(i64, 8) // base 8 for octal
	}
}

func Oct2int(octStr string) int {
	// base 8 for octal
	result, _ := strconv.ParseInt(octStr, 8, 64)
	return int(result)
}

func Int2hex(i int, prefix bool) string {
	i64 := int64(i)

	tmpStr := strconv.FormatInt(i64, 16)

	if len(tmpStr)%2 != 0 {
		tmpStr = "0" + tmpStr
	}

	if prefix {
		return "0x" + tmpStr // base 16 for hexadecimal
	} else {
		return tmpStr // base 16 for hexadecimal
	}
}

func Int642hex(i int64, prefix bool) string {
	tmpStr := strconv.FormatInt(i, 16)

	tmpStr = FormatHex(tmpStr, 16)

	if prefix {
		return "0x" + tmpStr // base 16 for hexadecimal
	} else {
		return tmpStr // base 16 for hexadecimal
	}
}
func Int322hex(i int64, prefix bool) string {
	tmpStr := strconv.FormatInt(i, 16)

	tmpStr = FormatHex(tmpStr, 8)

	if prefix {
		return "0x" + tmpStr // base 16 for hexadecimal
	} else {
		return tmpStr // base 16 for hexadecimal
	}
}
func Int162hex(i int64, prefix bool) string {
	tmpStr := strconv.FormatInt(i, 16)

	tmpStr = FormatHex(tmpStr, 4)

	if prefix {
		return "0x" + tmpStr // base 16 for hexadecimal
	} else {
		return tmpStr // base 16 for hexadecimal
	}
}

func Str2hex(str string, length int) (hexStr string) {
	for _, char := range []rune(str) {
		hexStr = hexStr + Int2hex(int(char), false)
	}
	hexStr = FormatHex(hexStr, length)
	return
}

func FormatStr(str string, length int) string {
	tmpStr := ""
	if len(str) < length {
		for i := 0; i < length-len(str); i++ {
			tmpStr = tmpStr + "0"
		}
		tmpStr = tmpStr + str
		return tmpStr
	} else {
		return str
	}
}

func FormatHex(hexStr string, length int) string {
	tmpStr := ""
	if len(hexStr) < length {
		for i := 0; i < length-len(hexStr); i++ {
			tmpStr = tmpStr + "0"
		}
		tmpStr = tmpStr + hexStr
		return tmpStr
	} else {
		return hexStr
	}
}

func Hex2int64(hexStr string) int64 {
	// base 16 for hexadecimal
	result, _ := strconv.ParseInt(hexStr, 16, 64)
	return result
}

func Hex2int(hexStr string) int {
	// base 16 for hexadecimal
	result, _ := strconv.ParseInt(hexStr, 16, 64)
	return int(result)
}

func Hex2int16(hexStr string) int {
	// base 16 for hexadecimal
	result, _ := strconv.ParseInt(hexStr, 16, 16)
	return int(result)
}

func HexStr2byte(hexStr string) []byte {
	var lengthByte int
	if len(hexStr)%2 != 0 {
		hexStr = "00" + hexStr
		lengthByte = len(hexStr)/2 + 1
	} else {
		lengthByte = len(hexStr) / 2
	}

	result := make([]byte, lengthByte)

	for i := 0; i < lengthByte; i++ {
		strTmp := SubStr(hexStr, i*2, 2)
		intTmp, _ := strconv.ParseInt(strTmp, 16, 64)
		result[i] = byte(int(intTmp))
	}
	return result
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}

func HexStr2float32(hexStr string) float32 {
	return ByteToFloat32(HexStr2byte(hexStr))
}

func HexStr_and(hexStr1 string, hexStr2 string) string {
	return Int2hex((Hex2int(hexStr1) & Hex2int(hexStr2)), false)
}

func SubStr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}
