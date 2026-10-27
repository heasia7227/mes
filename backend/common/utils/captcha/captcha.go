package captcha

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type CaptchaItem struct {
	Code      string
	CreatedAt time.Time
}

// 验证码存储
var (
	captchaMu    sync.RWMutex
	captchaStore = make(map[string]CaptchaItem) // key: sessionID, value: CaptchaItem
)

// 清理过期的验证码
func CleanExpiredCaptchas() {
	captchaMu.Lock()
	defer captchaMu.Unlock()

	now := time.Now()
	// 验证码有效期3分钟
	expiration := 3 * time.Minute

	for key, item := range captchaStore {
		if now.Sub(item.CreatedAt) > expiration {
			delete(captchaStore, key)
		}
	}
}

// 生成随机验证码, 只包含数字
func GenerateCaptchaCode(length int) string {
	const charset = "0123456789"
	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

// 生成验证码图片 - 参考JS版本实现
func GenerateCaptchaImage(code string) image.Image {
	// 定义验证码尺寸
	width := 160
	height := 60
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	white := color.RGBA{255, 255, 255, 255}

	// 填充白色背景
	draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.Point{}, draw.Src)

	// 添加噪点 (50个噪点, 减少噪声)
	for i := 0; i < 50; i++ {
		x := rand.Intn(width)
		y := rand.Intn(height)
		// 随机颜色, 透明度0.3
		r := uint8(rand.Intn(255))
		g := uint8(rand.Intn(255))
		b := uint8(rand.Intn(255))
		// 使用较低的alpha值来模拟透明度
		alpha := uint8(77) // 0.3 * 255 ≈ 77
		noiseColor := color.RGBA{r, g, b, alpha}

		// 绘制小圆点
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if (dx*dx + dy*dy) <= 1 { // 半径为1的圆
					px := x + dx
					py := y + dy
					if px >= 0 && px < width && py >= 0 && py < height {
						img.Set(px, py, noiseColor)
					}
				}
			}
		}
	}

	// 添加干扰线 (3条线, 减少噪声)
	for i := 0; i < 3; i++ {
		x1 := rand.Intn(width)
		y1 := rand.Intn(height)
		x2 := rand.Intn(width)
		y2 := rand.Intn(height)

		// 随机颜色, 透明度0.5
		r := uint8(rand.Intn(100))
		g := uint8(rand.Intn(150))
		b := uint8(rand.Intn(255))
		alpha := uint8(128) // 0.5 * 255 = 128
		lineColor := color.RGBA{r, g, b, alpha}

		// 绘制线段
		drawLine(img, x1, y1, x2, y2, lineColor)
	}

	// 绘制验证码文本
	for i, char := range code {
		// 计算字符位置
		x := width / 5 * (i + 1)
		y := height / 2

		// 随机旋转角度 (-10度到10度)
		angle := (rand.Float64()*0.4 - 0.2) // 转换为弧度

		// 随机颜色
		r := uint8(rand.Intn(80))
		g := uint8(rand.Intn(120))
		b := uint8(rand.Intn(200))
		textColor := color.RGBA{r, g, b, 255}

		// 绘制旋转文字
		drawRotatedChar(img, x, y, char, textColor, angle)
	}

	return img
}

func GenerateImage(captchaKey string) image.Image {
	// 生成4位验证码
	code := GenerateCaptchaCode(4)

	// 存储验证码
	captchaMu.Lock()
	captchaStore[captchaKey] = CaptchaItem{
		Code:      code,
		CreatedAt: time.Now(),
	}
	captchaMu.Unlock()

	// 生成验证码图片
	return GenerateCaptchaImage(code)
}

// 绘制线段
func drawLine(img *image.RGBA, x1, y1, x2, y2 int, c color.Color) {
	// 使用Bresenham算法绘制线段
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	sx := 1
	sy := 1

	if x1 > x2 {
		sx = -1
	}
	if y1 > y2 {
		sy = -1
	}

	err := dx - dy

	for {
		// 设置像素颜色
		if x1 >= 0 && x1 < img.Bounds().Dx() && y1 >= 0 && y1 < img.Bounds().Dy() {
			img.Set(x1, y1, c)
		}

		if x1 == x2 && y1 == y2 {
			break
		}

		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

// 绘制旋转的字符
func drawRotatedChar(img *image.RGBA, x, y int, char rune, c color.Color, angle float64) {
	// 增大临时图像尺寸, 使字符更清晰
	charImg := image.NewRGBA(image.Rect(0, 0, 60, 60))

	// 简单绘制字符的点阵图案
	drawCharPattern(charImg, 30, 30, char, c)

	// 计算旋转后的坐标
	for i := 0; i < 60; i++ {
		for j := 0; j < 60; j++ {
			// 只处理非透明像素
			pixelColor := charImg.At(i, j)
			if r, g, b, a := pixelColor.RGBA(); a > 0 {
				// 计算相对于中心点的坐标
				dx := float64(i - 30)
				dy := float64(j - 30)

				// 应用旋转
				newDx := dx*math.Cos(angle) - dy*math.Sin(angle)
				newDy := dx*math.Sin(angle) + dy*math.Cos(angle)

				// 转换回绝对坐标
				px := x + int(newDx)
				py := y + int(newDy)

				// 确保在图像范围内
				if px >= 0 && px < img.Bounds().Dx() && py >= 0 && py < img.Bounds().Dy() {
					img.Set(px, py, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
				}
			}
		}
	}

	// 简化平滑效果, 只在需要的地方添加
	for i := 0; i < 60; i++ {
		for j := 0; j < 60; j++ {
			// 检查是否为主要字符的边缘像素
			pixelColor := charImg.At(i, j)
			r, g, b, a := pixelColor.RGBA()
			if a > 0 {
				// 检查周围8个方向的像素, 如果有透明像素, 则添加平滑效果
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						if dx == 0 && dy == 0 {
							continue // 跳过自己
						}
						ni, nj := i+dx, j+dy
						if ni >= 0 && ni < 60 && nj >= 0 && nj < 60 {
							neighborColor := charImg.At(ni, nj)
							_, _, _, neighborA := neighborColor.RGBA()
							if neighborA == 0 {
								// 边缘像素, 添加平滑效果
								centerDx := float64(i - 30)
								centerDy := float64(j - 30)
								newDx := centerDx*math.Cos(angle) - centerDy*math.Sin(angle)
								newDy := centerDx*math.Sin(angle) + centerDy*math.Cos(angle)
								px := x + int(newDx+float64(dx))
								py := y + int(newDy+float64(dy))
								if px >= 0 && px < img.Bounds().Dx() && py >= 0 && py < img.Bounds().Dy() {
									// 调整透明度, 使边缘更柔和但不模糊
									img.Set(px, py, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 96})
								}
							}
						}
					}
				}
			}
		}
	}
}

// 绘制字符的点阵图案 - 使用更平滑的模式
func drawCharPattern(img *image.RGBA, x, y int, char rune, c color.Color) {
	// 增大字符点的大小, 使字符更清晰
	size := 6 // 增大尺寸

	// 定义平滑的字符图案
	var pattern [][]bool

	switch char {
	case 'A':
		pattern = [][]bool{
			{false, true, true, false},
			{true, false, false, true},
			{true, false, false, true},
			{true, true, true, true},
			{true, false, false, true},
			{true, false, false, true},
		}
	case 'B':
		pattern = [][]bool{
			{true, true, true, false},
			{true, false, false, true},
			{true, false, false, true},
			{true, true, true, false},
			{true, false, false, true},
			{true, true, true, false},
		}
	case 'C':
		pattern = [][]bool{
			{false, true, true, true},
			{true, false, false, false},
			{true, false, false, false},
			{true, false, false, false},
			{false, true, true, true},
		}
	case 'D':
		pattern = [][]bool{
			{true, true, false, false},
			{true, false, true, false},
			{true, false, true, false},
			{true, false, true, false},
			{true, true, false, false},
		}
	case 'E':
		pattern = [][]bool{
			{true, true, true, true},
			{true, false, false, false},
			{true, true, true, false},
			{true, false, false, false},
			{true, true, true, true},
		}
	case 'F':
		pattern = [][]bool{
			{true, true, true, true},
			{true, false, false, false},
			{true, true, true, false},
			{true, false, false, false},
			{true, false, false, false},
		}
	case 'G':
		pattern = [][]bool{
			{false, true, true, true},
			{true, false, false, false},
			{true, false, false, false},
			{true, false, true, true},
			{false, true, true, true},
		}
	case 'H':
		pattern = [][]bool{
			{true, false, false, true},
			{true, false, false, true},
			{true, true, true, true},
			{true, false, false, true},
			{true, false, false, true},
		}
	case 'I':
		pattern = [][]bool{
			{true, true, true},
			{false, true, false},
			{false, true, false},
			{false, true, false},
			{true, true, true},
		}
	case 'J':
		pattern = [][]bool{
			{true, true, true},
			{false, false, true},
			{false, false, true},
			{false, false, true},
			{true, true, true},
		}
	case 'K':
		pattern = [][]bool{
			{true, false, false, true},
			{true, false, true, false},
			{true, true, false, false},
			{true, false, true, false},
			{true, false, false, true},
		}
	case 'L':
		pattern = [][]bool{
			{true, false, false},
			{true, false, false},
			{true, false, false},
			{true, false, false},
			{true, true, true},
		}
	case 'M':
		pattern = [][]bool{
			{true, false, false, false, true},
			{true, true, false, true, true},
			{true, false, true, false, true},
			{true, false, false, false, true},
		}
	case 'N':
		pattern = [][]bool{
			{true, false, false, true},
			{true, true, false, true},
			{true, false, true, true},
			{true, false, false, true},
		}
	case 'O':
		pattern = [][]bool{
			{false, true, true, false},
			{true, false, false, true},
			{true, false, false, true},
			{true, false, false, true},
			{false, true, true, false},
		}
	case 'P':
		pattern = [][]bool{
			{true, true, true, false},
			{true, false, false, true},
			{true, true, true, false},
			{true, false, false, false},
			{true, false, false, false},
		}
	case 'Q':
		pattern = [][]bool{
			{false, true, true, false},
			{true, false, false, true},
			{true, false, false, true},
			{true, false, true, true},
			{false, true, true, false},
		}
	case 'R':
		pattern = [][]bool{
			{true, true, true, false},
			{true, false, false, true},
			{true, true, true, false},
			{true, false, false, true},
			{true, false, false, true},
		}
	case 'S':
		pattern = [][]bool{
			{true, true, true, true},
			{true, false, false, false},
			{true, true, true, false},
			{false, false, false, true},
			{true, true, true, true},
		}
	case 'T':
		pattern = [][]bool{
			{true, true, true, true},
			{false, false, true, false},
			{false, false, true, false},
			{false, false, true, false},
			{false, false, true, false},
		}
	case 'U':
		pattern = [][]bool{
			{true, false, false, true},
			{true, false, false, true},
			{true, false, false, true},
			{true, false, false, true},
			{false, true, true, false},
		}
	case 'V':
		pattern = [][]bool{
			{true, false, false, true},
			{true, false, false, true},
			{true, false, false, true},
			{false, true, true, false},
		}
	case 'W':
		pattern = [][]bool{
			{true, false, false, false, true},
			{true, false, false, false, true},
			{true, false, true, false, true},
			{true, true, false, true, true},
		}
	case 'X':
		pattern = [][]bool{
			{true, false, false, true},
			{false, true, true, false},
			{false, true, true, false},
			{true, false, false, true},
		}
	case 'Y':
		pattern = [][]bool{
			{true, false, false, true},
			{true, false, false, true},
			{false, true, true, false},
			{false, false, true, false},
			{false, false, true, false},
		}
	case 'Z':
		pattern = [][]bool{
			{true, true, true, true},
			{false, false, false, true},
			{false, true, true, false},
			{true, false, false, false},
			{true, true, true, true},
		}
	case '0':
		pattern = [][]bool{
			{false, true, true, false},
			{true, false, false, true},
			{true, false, false, true},
			{true, false, false, true},
			{false, true, true, false},
		}
	case '1':
		pattern = [][]bool{
			{false, false, true, false},
			{false, true, true, false},
			{false, false, true, false},
			{false, false, true, false},
			{false, true, true, true},
		}
	case '2':
		pattern = [][]bool{
			{false, true, true, true},
			{true, false, false, true},
			{false, false, true, false},
			{false, true, false, false},
			{true, true, true, true},
		}
	case '3':
		pattern = [][]bool{
			{false, true, true, true},
			{true, false, false, true},
			{false, false, true, true},
			{true, false, false, true},
			{false, true, true, true},
		}
	case '4':
		pattern = [][]bool{
			{true, false, false, true},
			{true, false, false, true},
			{true, true, true, true},
			{false, false, false, true},
			{false, false, false, true},
		}
	case '5':
		pattern = [][]bool{
			{true, true, true, true},
			{true, false, false, false},
			{true, true, true, true},
			{false, false, false, true},
			{true, true, true, true},
		}
	case '6':
		pattern = [][]bool{
			{true, true, true, true},
			{true, false, false, false},
			{true, true, true, true},
			{true, false, false, true},
			{true, true, true, true},
		}
	case '7':
		pattern = [][]bool{
			{true, true, true, true},
			{false, false, false, true},
			{false, false, true, false},
			{false, true, false, false},
			{true, false, false, false},
		}
	case '8':
		pattern = [][]bool{
			{true, true, true, true},
			{true, false, false, true},
			{true, true, true, true},
			{true, false, false, true},
			{true, true, true, true},
		}
	case '9':
		pattern = [][]bool{
			{true, true, true, true},
			{true, false, false, true},
			{true, true, true, true},
			{false, false, false, true},
			{true, true, true, true},
		}
	default:
		// 默认字符, 简单的方块
		pattern = [][]bool{
			{true, true, true},
			{true, true, true},
			{true, true, true},
			{true, true, true},
		}
	}

	// 计算字符的起始位置
	height := len(pattern)
	width := len(pattern[0])
	offsetX := -width * size / 2
	offsetY := -height * size / 2

	// 绘制字符, 使用更大的点并改进平滑效果
	for i, row := range pattern {
		for j, pixel := range row {
			if pixel {
				// 绘制主方块, 使用更大的尺寸
				for dx := 0; dx < size; dx++ {
					for dy := 0; dy < size; dy++ {
						px := x + offsetX + j*size + dx
						py := y + offsetY + i*size + dy
						if px >= 0 && px < img.Bounds().Dx() && py >= 0 && py < img.Bounds().Dy() {
							img.Set(px, py, c)
						}
					}
				}

				// 简化平滑效果, 只添加必要的边缘平滑
				// 检查右侧相邻像素是否为false (边缘)
				hasRight := j+1 < width && row[j+1]
				// 检查底部相邻像素是否为false (边缘)
				hasBottom := i+1 < height && len(pattern[i+1]) > j && pattern[i+1][j]

				// 如果右侧是边缘, 添加右侧平滑像素
				if !hasRight {
					for dy := 1; dy < size-1; dy++ {
						px := x + offsetX + j*size + size
						py := y + offsetY + i*size + dy
						if px >= 0 && px < img.Bounds().Dx() && py >= 0 && py < img.Bounds().Dy() {
							// 使用更高的透明度, 使字符更清晰
							if r, g, b, _ := c.RGBA(); r > 0 || g > 0 || b > 0 {
								img.Set(px, py, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 192})
							}
						}
					}
				}

				// 如果底部是边缘, 添加底部平滑像素
				if !hasBottom {
					for dx := 1; dx < size-1; dx++ {
						px := x + offsetX + j*size + dx
						py := y + offsetY + i*size + size
						if px >= 0 && px < img.Bounds().Dx() && py >= 0 && py < img.Bounds().Dy() {
							// 使用更高的透明度, 使字符更清晰
							if r, g, b, _ := c.RGBA(); r > 0 || g > 0 || b > 0 {
								img.Set(px, py, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 192})
							}
						}
					}
				}
			}
		}
	}
}

// abs 返回整数的绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 验证码校验
func ValidateCaptcha(captchaKey, code string) bool {
	captchaMu.RLock()
	item, exists := captchaStore[captchaKey]
	captchaMu.RUnlock()

	if !exists {
		return false
	}

	// 无论验证是否成功, 都删除验证码 (一次性使用, 防止暴力破解和撞库)
	captchaMu.Lock()
	delete(captchaStore, captchaKey)
	captchaMu.Unlock()

	// 检查是否过期 (3分钟)
	if time.Since(item.CreatedAt) > 3*time.Minute {
		return false
	}

	// 验证码不区分大小写
	result := strings.EqualFold(item.Code, code)

	return result
}
