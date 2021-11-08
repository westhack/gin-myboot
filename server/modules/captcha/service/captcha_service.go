package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"gin-myboot/modules/captcha/model"
	"gin-myboot/modules/captcha/utils"
	"github.com/go-redis/redis/v8"
	"github.com/golang/freetype"
	uuid "github.com/satori/go.uuid"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"time"
)

const REDIS_CAPTCHA_KEY = "RUNNING:CAPTCHA:"
const REDIS_SECOND_CAPTCHA_KEY = "RUNNING:CAPTCHA:second-"
const RESOURCE_IMAGES_DIR = "modules/captcha/resource/defaultImages"
const RESOURCE_FONTS_DIR = "modules/captcha/resource/fonts"

func getImg(dir string) string {
	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var names []string
	for i := range fileInfoList {
		if !fileInfoList[i].IsDir() {
			names = append(names, fileInfoList[i].Name())
		}
	}

	n := rand.Intn(len(names))
	return dir + names[n]
}

func GetBlockPuzzle(client *redis.Client) (res model.RepData, err error) {
	original := getImg(RESOURCE_IMAGES_DIR + "/jigsaw/original/")
	slidingBlock := getImg(RESOURCE_IMAGES_DIR + "/jigsaw/slidingBlock/")

	ff, _ := ioutil.ReadFile(original) //读取文件
	bbb := bytes.NewBuffer(ff)
	img, _ := png.Decode(bbb)

	fff, _ := ioutil.ReadFile(slidingBlock) //读取文件
	bbbb := bytes.NewBuffer(fff)
	block, _ := png.Decode(bbbb)

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	blockBounds := block.Bounds()
	dx := blockBounds.Dx()
	dy := blockBounds.Dy()

	bgWidth := width / 2
	randNum := getRandOffset(bgWidth, dx)
	newBlock := image.NewRGBA(blockBounds) //new 一个新的图片

	xx := bgWidth + randNum

	bgHeight := height / 4
	randNum = rand.Intn(bgHeight - 1)
	yy := randNum - 10

	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := block.At(i, j)
			r, g, b, a := colorRgb.RGBA()

			r_uint8 := uint8(r >> 8) //转换为 255 值
			g_uint8 := uint8(g >> 8)
			b_uint8 := uint8(b >> 8)
			a_uint8 := uint8(a >> 8)

			newBlock.Set(i, j+yy, color.RGBA{r_uint8, g_uint8, b_uint8, a_uint8}) //设置像素点
		}
	}

	interfereX := xx - 100 - rand.Intn(50)
	interfereY := yy + 10 + rand.Intn(50)

	interfereXX := xx + 50 + rand.Intn(50)
	interfereYY := yy + 50 + rand.Intn(50)

	newImg := img.(draw.Image)
	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			colorRgb := newBlock.At(i, j)
			toBlockRgb := img.At(xx+i, j)

			_, _, _, a := colorRgb.RGBA()
			a_uint8 := uint8(a >> 8) //转换为 255 值

			if a_uint8 < 128 { // 透明
				//newRgba.SetRGBA(i, j, color.RGBA{r_uint8, g_uint8, b_uint8, a_uint8}) //设置像素点
			} else {
				r, g, b, a := toBlockRgb.RGBA()
				opacity := uint16(float64(a) * 0.5)
				v := img.ColorModel().Convert(color.NRGBA64{R: uint16(r), G: uint16(g), B: uint16(b), A: opacity})
				//Alpha = 0: Full transparent
				rr, gg, bb, aa := v.RGBA()
				newImg.Set(interfereXX+i, interfereYY+j, color.RGBA64{R: uint16(rr), G: uint16(gg), B: uint16(bb), A: uint16(aa)}) // 干扰
				newImg.Set(interfereX+i, interfereY+j, color.RGBA64{R: uint16(rr), G: uint16(gg), B: uint16(bb), A: uint16(aa)})   // 干扰
				newImg.Set(xx+i, j, color.RGBA64{R: uint16(rr), G: uint16(gg), B: uint16(bb), A: uint16(aa)})                      // 抠图
				newBlock.Set(i, j, toBlockRgb)                                                                                     // 抠图
			}

			if IsBoundary(newBlock, IsOpacity(a), i, j) {
				newImg.Set(interfereXX+i, interfereYY+j, color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(255)})
				newImg.Set(interfereX+i, interfereY+j, color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(255)})
				newImg.Set(xx+i, j, color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(255)})
				newBlock.Set(i, j, color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(255)})
			}
		}
	}

	emptyBuff := bytes.NewBuffer(nil) //开辟一个新的空buff
	err = png.Encode(emptyBuff, newImg)
	if err != nil {
		fmt.Println(err)
		return res, err
	}
	////img写入到buff
	img64 := base64.StdEncoding.EncodeToString(emptyBuff.Bytes()) //buff转成base64

	emptyBuff2 := bytes.NewBuffer(nil) //开辟一个新的空buff
	err = png.Encode(emptyBuff2, newBlock)
	if err != nil {
		fmt.Println(err)
		return res, err
	}
	////img写入到buff
	newBlockBase64 := base64.StdEncoding.EncodeToString(emptyBuff2.Bytes()) //buff转成base64

	uuidStr := uuid.NewV4().String()
	res.OriginalImageBase64 = img64
	res.JigsawImageBase64 = newBlockBase64
	res.Token = uuidStr
	res.SecretKey = RandChar(16)

	ctx := context.Background()

	ss, _ := json.Marshal(model.BlockPuzzleCheckInfo{SecretKey: res.SecretKey, Point: model.Point{float32(xx), float32(yy)}})
	timer := time.Duration(300) * time.Second
	client.Set(ctx, REDIS_CAPTCHA_KEY+res.Token, string(ss), timer).Result()

	return res, nil
}

func GetClickWord(client *redis.Client) (res model.RepData, err error) {

	original := getImg(RESOURCE_IMAGES_DIR + "/pic-click/")

	ff, _ := ioutil.ReadFile(original) //读取文件
	b := bytes.NewBuffer(ff)
	img, _ := png.Decode(b)

	newImg := img.(draw.Image)

	//拷贝一个字体文件到运行目录
	fontBytes, err := ioutil.ReadFile(RESOURCE_FONTS_DIR + "/WenQuanZhengHei.ttf")
	if err != nil {
		log.Println(err)
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
	}

	f := freetype.NewContext()

	strs := RandChinaStr(3)
	var p []model.Point

	width := newImg.Bounds().Dx()
	height := newImg.Bounds().Dy()
	for _, str := range strs {
		f.SetDPI(72)
		f.SetFont(font)
		f.SetFontSize(20)
		f.SetClip(img.Bounds())
		f.SetDst(newImg)
		//f.SetSrc(image.NewUniform(color.RGBA{R: 255, G: 0, B: 0, A: 255}))
		f.SetSrc(image.NewUniform(getRandomColor()))

		xx := rand.Intn(width)
		yy := rand.Intn(height)

		if xx < 20 {
			xx = xx + 20
		}
		if yy < 20 {
			yy = yy + 20
		}
		if xx > (width - 20) {
			xx = xx - 20
		}
		if yy > (height - 20) {
			yy = yy - 20
		}

		pt := freetype.Pt(xx, yy)
		_, err = f.DrawString(str, pt)

		p = append(p, model.Point{X: float32(xx), Y: float32(yy)})
	}

	emptyBuff := bytes.NewBuffer(nil) //开辟一个新的空buff
	err = png.Encode(emptyBuff, newImg)
	if err != nil {
		fmt.Println(err)
		return res, err
	}
	////img写入到buff
	img64 := base64.StdEncoding.EncodeToString(emptyBuff.Bytes()) //buff转成base64

	uuidStr := uuid.NewV4().String()
	res.OriginalImageBase64 = img64
	res.JigsawImageBase64 = ""
	res.Token = uuidStr
	res.SecretKey = RandChar(16)
	res.WordList = strs

	ctx := context.Background()

	ss, _ := json.Marshal(model.ClickWordCheckInfo{SecretKey: res.SecretKey, Points: p})
	timer := time.Duration(300) * time.Second
	client.Set(ctx, REDIS_CAPTCHA_KEY+res.Token, string(ss), timer).Result()

	return res, nil
}

func getRandOffset(bgWidth int, tempWidth int) int {

	diffWidth := bgWidth - tempWidth

	return rand.Intn(diffWidth - 1)
}

func IsBoundary(img *image.RGBA, isOpacity bool, x int, y int) bool {

	b := img.Bounds()
	w := b.Dx()
	h := b.Dy()
	if x >= w-1 || y >= h-1 {
		return false
	}

	_, _, _, right := img.At(x+1, y).RGBA()
	_, _, _, down := img.At(x, y+1).RGBA()

	if isOpacity && !IsOpacity(right) ||
		!isOpacity && IsOpacity(right) ||
		isOpacity && !IsOpacity(down) ||
		!isOpacity && IsOpacity(down) {
		return true
	}

	return false
}

func IsOpacity(a uint32) bool {
	aa := uint8(a >> 8)
	if aa < 128 {
		return true
	}

	return false
}

const char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandChar(size int) string {
	rand.NewSource(time.Now().UnixNano()) // 产生随机种子
	var s bytes.Buffer
	for i := 0; i < size; i++ {
		s.WriteByte(char[rand.Int63()%int64(len(char))])
	}
	return s.String()
}

func RandInt(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}

func RandChinaStr(num int) []string {
	a := make([]string, num)
	for i := range a {
		a[i] = string(rune(RandInt(19968, 40869)))
	}
	return a
}

func getRandomColor() (c color.RGBA) {

	r := RandInt(0, 255)
	g := RandInt(0, 255)
	b := RandInt(0, 255)
	c.A = 255
	c.R = uint8(r)
	c.G = uint8(g)
	c.B = uint8(b)
	//color.RGBA{R: 255, G: 0, B: 0, A: 255}
	return c
}

func BlockPuzzleCheck(client *redis.Client, captchaCheckReq model.CaptchaCheckRequest, del bool) (b bool, err error) {
	ctx := context.Background()
	key := REDIS_CAPTCHA_KEY + captchaCheckReq.Token
	jsonStr, _ := client.Get(ctx, key).Result()
	var checkInfo model.BlockPuzzleCheckInfo
	err = json.Unmarshal([]byte(jsonStr), &checkInfo)
	if err != nil {
		return false, errors.New("验证码已失效，请重新获取")
	}

	if checkInfo.Point.X == 0 {
		return false, errors.New("验证码已失效，请重新获取")
	}

	fmt.Printf("=======> %s \n", checkInfo.SecretKey)
	fmt.Printf("=======> %v \n", captchaCheckReq.PointJSON)

	decodeString, err := base64.StdEncoding.DecodeString(captchaCheckReq.PointJSON)
	if err != nil {
		return
	}
	decrypt := utils.EcbDecrypt(decodeString, []byte(checkInfo.SecretKey))

	userPoint := model.Point{}
	err = json.Unmarshal(decrypt, &userPoint)
	if err != nil {
		fmt.Printf("=======> %s \n", err.Error())
		fmt.Printf("=======> %s \n", string(decrypt))
		return
	}

	var xx = checkInfo.Point.X - 2
	if math.Abs(float64(userPoint.X-xx)) < float64(2) || math.Abs(float64(userPoint.X-checkInfo.Point.X)) < float64(2) {

		if del {
			client.Del(ctx, key).Result()
		}

		base64str := base64.StdEncoding.EncodeToString(utils.EcbEncrypt([]byte(captchaCheckReq.Token+"---"+string(decrypt)), []byte(checkInfo.SecretKey)))
		key = REDIS_SECOND_CAPTCHA_KEY + base64str
		timer := time.Duration(600) * time.Second
		client.Set(ctx, key, captchaCheckReq.PointJSON, timer).Result()
		fmt.Printf("=======> %s \n", base64str)
		return true, nil
	}

	return false, errors.New("验证码已失效，请重新获取")
}

func ClickWordCheck(client *redis.Client, captchaCheckReq model.CaptchaCheckRequest, del bool) (b bool, err error) {
	ctx := context.Background()
	key := REDIS_CAPTCHA_KEY + captchaCheckReq.Token
	jsonStr, _ := client.Get(ctx, key).Result()

	var checkInfo model.ClickWordCheckInfo
	err = json.Unmarshal([]byte(jsonStr), &checkInfo)
	if err != nil {
		return false, errors.New("验证码已失效，请重新获取")
	}

	decodeString, err := base64.StdEncoding.DecodeString(captchaCheckReq.PointJSON)
	if err != nil {
		return
	}
	decrypt := utils.EcbDecrypt(decodeString, []byte(checkInfo.SecretKey))

	var userPoints []model.Point
	err = json.Unmarshal(decrypt, &userPoints)
	if err != nil {
		fmt.Printf("=======> %s \n", err.Error())
		fmt.Printf("=======> %s \n", string(decrypt))
		return
	}

	for i, userPoint := range userPoints {
		xx := checkInfo.Points[i].X
		yy := checkInfo.Points[i].Y
		if !(math.Abs(float64(userPoint.X-xx)) < float64(20) &&
			math.Abs(float64(userPoint.Y-yy)) < float64(20)) {
			return false, errors.New("验证码已失效，请重新获取")
		}
	}

	if del {
		client.Del(ctx, key).Result()
	}

	base64str := base64.StdEncoding.EncodeToString(utils.EcbEncrypt([]byte(captchaCheckReq.Token+"---"+string(decrypt)), []byte(checkInfo.SecretKey)))
	key = REDIS_SECOND_CAPTCHA_KEY + base64str
	timer := time.Duration(600) * time.Second
	client.Set(ctx, key, captchaCheckReq.PointJSON, timer).Result()

	return true, nil
}

func Verification(client *redis.Client, req model.CaptchaVerificationRequest) (b bool, err error) {
	ctx := context.Background()

	key := REDIS_SECOND_CAPTCHA_KEY + req.CaptchaVerification
	result, err := client.Get(ctx, key).Result()
	if err != nil {
		return false, errors.New("验证失败")
	}

	fmt.Printf("=======> Verification %s \n", result)

	client.Del(ctx, key).Result()

	return true, err
}
