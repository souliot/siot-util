package sutil

import (
	qrcode "github.com/skip2/go-qrcode"

	"bytes"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func GetQrByte(chd string, chs int, qrType int) (file []byte) {
	if qrType == 1 {
		file, _ = qrcode.Encode(chd, qrcode.Medium, chs)
	} else if qrType == 2 {
		code, err := qr.Encode(chd, qr.L, qr.Unicode)
		if err != nil {
			// beego.Info(err)
			return
		}
		// beego.Info("Encoded data: ", code.Content())

		if chd != code.Content() {
			// beego.Info("data differs")
			return
		}

		code, err = barcode.Scale(code, chs, chs)
		if err != nil {
			// beego.Info(err)
			return
		}

		encoder := png.Encoder{CompressionLevel: png.BestCompression}

		var b bytes.Buffer
		err = encoder.Encode(&b, code)

		if err != nil {
			// beego.Info(err)
			return
		}

		file = b.Bytes()
	}
	return
}
