package main

// ここでconfigがimportされた時点でconfig.goのinit関数が実行される
import (
	"fmt"
	"log"
	"todo_app/config"
)

// ここではConfigListじゃなくてConfigでいい。なぜならＣｏｎｆｉｇ変数ががインスタンスとして働いているから
func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")
}
