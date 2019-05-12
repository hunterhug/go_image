package go_image

import (
	"fmt"
	"testing"
)

//将某一图片文件进行缩放后存入另外的文件中
func TestImage(t *testing.T) {
	//打印当前文件夹位置
	fmt.Printf("本文件文件夹位置:%s\n", CurDir())

	//图像位置
	filename := "./testdata/gopher.png"

	//宽度,高度
	width := 500
	height := 800

	//保存位置
	save1 := "./testdata/gopher500.jpg"
	save2 := "./testdata/gopher500_800.png"

	//按照宽度进行等比例缩放
	err := ScaleF2F(filename, save1, width)
	if err != nil {
		fmt.Printf("生成按宽度缩放图失败：%s\n", err.Error())
	} else {
		fmt.Printf("生成按宽度缩放图：%s\n", save1)
	}

	//按照宽度和高度进行等比例缩放
	err = ThumbnailF2F(filename, save2, width, height)
	if err != nil {
		fmt.Printf("生成按宽度高度缩放图:%s\n", err.Error())
	} else {
		fmt.Printf("生成按宽度高度缩放图:%s\n", save2)
	}

	//查看图像文件的真正名字
	//如 ./testdata/gopher500.jpg其实是png类型,但是命名错误,需要纠正!
	realfilename, err := RealImageName(save1)
	if err != nil {
		fmt.Printf("真正的文件名: %s->? err:%s\n", save1, err.Error())
	} else {
		fmt.Printf("真正的文件名:%s->%s\n", save1, realfilename)
	}

	//文件改名,强制性
	err = ChangeImageName(save1, realfilename, true)
	if err != nil {
		fmt.Printf("文件改名失败:%s->%s,%s\n", save1, realfilename, err.Error())
	} else {
		fmt.Println("改名成功")
	}

	//文件改名,不强制性
	err = ChangeImageName(save1, realfilename, false)
	if err != nil {
		fmt.Printf("文件改名失败:%s->%s,%s\n", save1, realfilename, err.Error())
	}
}
