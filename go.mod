module github.com/souliot/siot-util

go 1.12

require (
	github.com/boombuler/barcode v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/skip2/go-qrcode v0.0.0-20190110000554-dc11ecdae0a9
	golang.org/x/crypto v0.0.0-20181127143415-eb0de9b17e85
)

replace golang.org/x/crypto v0.0.0-20181127143415-eb0de9b17e85 => github.com/golang/crypto v0.0.0-20181127143415-eb0de9b17e85
