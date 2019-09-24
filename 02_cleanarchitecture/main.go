package main

import (
	"go-ca-webapi/02_cleanarchitecture/adapter"
	"go-ca-webapi/02_cleanarchitecture/application"
	"log"
)

func init() {
	// TODO: 環境変数からシステム起動モード（デフォルト＝本番。開発・テストモード等を追加。）を取得して別パッケージのグローバル変数に保持するロジックを実装！
}

// 注意: プロダクション品質ではありません。
func main() {
	// TODO: Panic捕捉ロジック実装！

	env := application.NewEnv()
	err := application.NewApp(env, adapter.NewWebServer(env)).Start()
	if err != nil {
		log.Fatal(err)
	}
}
