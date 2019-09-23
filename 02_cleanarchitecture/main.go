package main

import "github.com/labstack/echo"

// 注意: プロダクション品質ではありません。
func main() {

	// Webサーバセットアップ
	e := echo.New()

	e.Logger.Fatal(e.Start(":8080"))
}
