package main

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"github.com/go-co-op/gocron"
	"gocv.io/x/gocv"
)

func createProgressBar() {
	// 現在の日付を取得
	now := time.Now()
	startOfYear := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
	endOfYear := time.Date(now.Year(), time.December, 31, 23, 59, 59, 0, now.Location())
	yearProgress := float64(now.Sub(startOfYear)) / float64(endOfYear.Sub(startOfYear))

	// 画像のサイズを設定
	width := 1000
	height := 200

	// 画像を作成
	img := gocv.NewMatWithSize(height, width, gocv.MatTypeCV8UC3)
	defer img.Close()

	// 背景を塗りつぶし
	bgColor := color.RGBA{0, 0, 50, 255}
	gocv.Rectangle(&img, image.Rect(0, 0, width, height), bgColor, -1)

	// プログレスバーの枠を描画
	borderColor := color.RGBA{50, 100, 255, 255}
	gocv.Rectangle(&img, image.Rect(20, 20, width-20, height-20), borderColor, 3)

	// プログレスバーを描画
	progressBarColor := color.RGBA{50, 100, 255, 255}
	progressWidth := int(float64(width-40) * yearProgress)
	gocv.Rectangle(&img, image.Rect(20, 20, 20+progressWidth, height-20), progressBarColor, -1)

	// タイトルを描画
	title := fmt.Sprintf("%d Progress Bar", now.Year())
	gocv.PutText(&img, title, image.Point{width/2 - 100, 50}, gocv.FontHersheySimplex, 1.0, borderColor, 2)

	// 画像を保存
	filename := fmt.Sprintf("progress_bar_%d_%02d_%02d.png", now.Year(), now.Month(), now.Day())
	gocv.IMWrite(filename, img)
	fmt.Printf("Progress bar image saved as %s\n", filename)
}

func main() {
	s := gocron.NewScheduler(time.UTC)

	// 毎日0時にcreateProgressBar関数を実行
	s.Every(1).Day().At("00:00").Do(createProgressBar)

	// スケジューラを開始
	s.StartBlocking()
}
