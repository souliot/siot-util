module github.com/souliot/siot-util

go 1.12

require (
	github.com/astaxie/beego v1.12.1
	github.com/boombuler/barcode v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/robfig/cron/v3 v3.0.1
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/skip2/go-qrcode v0.0.0-20190110000554-dc11ecdae0a9
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
)

replace golang.org/x/crypto v0.0.0-20181127143415-eb0de9b17e85 => github.com/golang/crypto v0.0.0-20181127143415-eb0de9b17e85
