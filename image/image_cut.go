package image

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/fogleman/gg"
	"golang.org/x/image/bmp"
)

const (
	UNKNOWN = 0
	JPEG    = 0
	PNG     = 0
	PDF     = 1
	OFD     = 2
)

//坐标
type Coord struct {
	X int32 `protobuf:"zigzag32,1,opt,name=x,proto3" json:"x,omitempty"` //横坐标
	Y int32 `protobuf:"zigzag32,2,opt,name=y,proto3" json:"y,omitempty"` //纵坐标
}

//文本的坐标，以四个顶点坐标表示 注意：此字段可能返回 null，表示取不到有效值
type Polygon struct {
	LeftTop     *Coord `protobuf:"bytes,1,opt,name=left_top,json=leftTop,proto3" json:"left_top,omitempty"`             // 左上顶点坐标
	RightTop    *Coord `protobuf:"bytes,2,opt,name=right_top,json=rightTop,proto3" json:"right_top,omitempty"`          // 右上顶点坐标
	RightBottom *Coord `protobuf:"bytes,3,opt,name=right_bottom,json=rightBottom,proto3" json:"right_bottom,omitempty"` // 右下顶点坐标
	LeftBottom  *Coord `protobuf:"bytes,4,opt,name=left_bottom,json=leftBottom,proto3" json:"left_bottom,omitempty"`    // 左下顶点坐标
}

//矩形坐标
type Rect struct {
	X      int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`           //左上角x
	Y      int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`           //左上角y
	Width  int32 `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`   //宽度
	Height int32 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"` //高度
}


// 识别参数
type Image struct {
	LogId string
}

func NewImage(logId string) *Image {
	return &Image{
		LogId: logId,
	}
}

/*ClipByUrl 图片裁剪(支持URL和base64形式图片)

入参:图片地址,图片Base64,需裁剪图片矩形坐标,图片角度,精度

精度规则:精度范围:0-100,精度随数值增加而增加

返回:图片字节流输出、异常*/
func (i *Image) ClipImage(imageUrl, imageBase64 *string, rect *Rect, angle float64, quality int) (imageBytes []byte, err error) {

	if imageBase64 != nil && *imageBase64 != "" {

		newImageBase64 := i.CutImgBase64Hander(*imageBase64)

		imageBytes, err := base64.StdEncoding.DecodeString(newImageBase64)
		if err != nil {
			fmt.Println(i.LogId, "图片base64解码失败: ", err)
			return nil, err
		}

		imageBytes, err = i.clipImageCore(imageBytes, float64(rect.X), float64(rect.Y), float64(rect.Width), float64(rect.Height), angle, quality)
		if err != nil {
			fmt.Println(i.LogId, "发票图片裁剪失败:", err)
			return nil, err
		}
		fmt.Println(i.LogId, "发票图片裁剪成功:")
		return imageBytes, nil

	} else if imageUrl != nil && *imageUrl != "" {

		imageBytes, err = i.ClipByUrl(*imageUrl, float64(rect.X), float64(rect.Y), float64(rect.Width), float64(rect.Height), angle, quality)
		if err != nil {
			return nil, err
		} else {
			return imageBytes, nil
		}

	} else {
		err := errors.New("图片URL和图片BASE64均为空")
		fmt.Println(i.LogId, err)
		return nil, err
	}
}

/*ClipByUrl URL图片裁剪

入参:图片地址、需裁剪图片左上角X、Y轴,裁剪图宽,裁剪图高,图片角度,精度

精度规则:精度范围:0-100,精度随数值增加而增加

返回:图片字节流输出、异常*/
func (i *Image) ClipByUrl(imageUrl string, x0, y0, width, height, angle float64, quality int) (outImg []byte, err error) {
	resp, err := http.Get(imageUrl)
	if err != nil {
		err := errors.New("URL图片裁剪错误:图片获取失败" + err.Error())
		fmt.Println(i.LogId, err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || resp.StatusCode != 200 {
		err := errors.New("URL图片裁剪错误:图片获取失败 " + err.Error())
		fmt.Println(i.LogId, err)
		return nil, err
	}

	outImg, err = i.clipImageCore(body, x0, y0, width, height, angle, quality)
	if err != nil {
		err := errors.New("URL图片裁剪错误:图片获取失败 " + err.Error())
		fmt.Println(i.LogId, err)
		return nil, err
	}
	return outImg, nil
}

/*clipImageCore 图片裁剪核心

入参:图片字节流输入、图片左上角X、Y轴,图宽,图高,图片角度,精度

精度规则:精度范围:0-100,精度随数值增加而增加

返回:图片字节流输出、异常*/
func (i *Image) clipImageCore(inImg []byte, x0, y0, width, height, angle float64, quality int) (outImg []byte, err error) {

	var (
		imgBuffer = new(bytes.Buffer)
		imgReader = bytes.NewReader(inImg)
		rotateX   = x0 + width/2
		rotateY   = y0 + height/2
		x1 = x0 + width
		y1 = y0 + height
	)

	err = errors.New("图片读取失败")
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	originImage, fm, err := image.Decode(imgReader)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	canvas := i.RotateImage(originImage, rotateX, rotateY, angle, Counterclockwise)
	switch fm {
	case "jpeg":
		switch canvas.(type) {
		case *image.YCbCr:
			img := canvas.(*image.YCbCr)
			subImg := img.SubImage(image.Rect(int(x0), int(y0), int(x1), int(y1))).(*image.YCbCr)
			err = jpeg.Encode(imgBuffer, subImg, &jpeg.Options{Quality: quality})
		case *image.RGBA:
			img := canvas.(*image.RGBA)
			subImg := img.SubImage(image.Rect(int(x0), int(y0), int(x1), int(y1))).(*image.RGBA)
			err = png.Encode(imgBuffer, subImg)
		}
	case "png":
		switch canvas.(type) {
		case *image.NRGBA:
			img := canvas.(*image.NRGBA)
			subImg := img.SubImage(image.Rect(int(x0), int(y0), int(x1), int(y1))).(*image.NRGBA)
			err = png.Encode(imgBuffer, subImg)
		case *image.RGBA:
			img := canvas.(*image.RGBA)
			subImg := img.SubImage(image.Rect(int(x0), int(y0), int(x1), int(y1))).(*image.RGBA)
			err = png.Encode(imgBuffer, subImg)
		}
	case "gif":
		img := canvas.(*image.Paletted)
		subImg := img.SubImage(image.Rect(int(x0), int(y0), int(x1), int(y1))).(*image.Paletted)
		err = gif.Encode(imgBuffer, subImg, &gif.Options{})
	case "bmp":
		img := canvas.(*image.RGBA)
		subImg := img.SubImage(image.Rect(int(x0), int(y0), int(x1), int(y1))).(*image.RGBA)
		err = bmp.Encode(imgBuffer, subImg)
	default:
		err = errors.New("无效的图片")
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return imgBuffer.Bytes(), nil
}

func (i *Image) CutImgBase64Hander(imageBase64 string) string {
	if len(imageBase64) > 21 {
		if strings.Contains(imageBase64[:22], "data:image/jpg;base64,") { //JPG、JPEG
			return imageBase64[22:]
		} else if strings.Contains(imageBase64[:23], "data:image/jpeg;base64,") { //JPG、JPEG
			return imageBase64[23:]
		} else if strings.Contains(imageBase64[:22], "data:image/png;base64,") { //PNG
			return imageBase64[22:]
		} else if strings.Contains(imageBase64[:28], "data:application/pdf;base64,") { //PDF
			return imageBase64[28:]
		} else if strings.Contains(imageBase64[:37], "data:application/octet-stream;base64,") { //OFD
			return imageBase64[37:]
		}
	}
	return imageBase64
}

type Direction bool //旋转方向

const (
	Counterclockwise Direction = true  //逆时针方向
	Clockwise        Direction = false //顺时针方向
)

//图片旋转(不改变原有图像尺寸)
//x,y:旋转中心坐标点
//angle角度与direction旋转方向配合使用,没有传入旋转方向则默认旋转角度为顺时针
func (i *Image) RotateImage(image image.Image, x, y, angle float64, direction ...Direction) image.Image {

	canvas := gg.NewContext(int(image.Bounds().Size().X), int(image.Bounds().Size().Y))
	if len(direction) > 0 && direction[0] {
		angle = 360.00 - angle //逆时针角度
	}

	canvas.RotateAbout(gg.Radians(360-angle), x, y)
	canvas.DrawImage(image, 0, 0)
	// canvas.SavePNG("test.png")//保存文件方便查看
	return canvas.Image()
}
