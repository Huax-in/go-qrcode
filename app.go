package main

import (
	"image"
	"image/png"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func writePng(filename string, img image.Image) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(file, img)
	// err = jpeg.Encode(file, img, &jpeg.Options{100})      //图像质量值为100，是最好的图像显示
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	log.Println(file.Name())
}

func main() {
	// 获取命令参数
	args := os.Args
	// 没有参数提示并返回
	if len(args) < 2 {
		log.Println("ERROR: no args!")
		return
	}
	// 第一个参数当作二维码内容
	base64 := args[1]
	// 文件默认名称（当前时间的纳秒的36进制字符串）
	fileName := strconv.FormatInt(time.Now().Unix(), 36)
	// 如果有第二个参数 第二个参数当作文件名称
	if len(args) >= 3 {
		fileName = string(args[2])
	}
	log.Println("Original data:", base64)
	code, err := qr.Encode(base64, qr.L, qr.Unicode)
	// code, err := code39.Encode(base64)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Encoded data: ", code.Content())

	if base64 != code.Content() {
		log.Fatal("data differs")
	}

	code, err = barcode.Scale(code, 300, 300)
	if err != nil {
		log.Fatal(err)
	}

	writePng(fileName+".png", code)
}
