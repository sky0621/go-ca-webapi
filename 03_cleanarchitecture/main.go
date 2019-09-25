package main

import (
	"go-ca-webapi/02_cleanarchitecture/adapter"
	"go-ca-webapi/02_cleanarchitecture/application"
	"log"
)

// 注意: プロダクション品質ではありません。
func main() {

	// FIXME: DIは、ここで？

	env := application.NewEnv()
	router := adapter.NewRouter()
	webServer := adapter.NewWebServer(env, router)

	err := application.NewApp(env, webServer).Start()
	if err != nil {
		log.Fatal(err)
	}
}
