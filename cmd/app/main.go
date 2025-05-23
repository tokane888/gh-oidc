package main

import (
	"errors"
	"log"

	"github.com/tokane888/gh-oidc/configs"
	common "github.com/tokane888/go_common_module/v2"
	"go.uber.org/zap"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Println("failed to load config:", err)
		return
	}
	logger, err := common.NewLogger(cfg.Logger)
	if err != nil {
		// zap loggerの初期化に失敗した場合のエラーハンドリング
		// zapを使用できないため、標準のlogパッケージを使用
		log.Println("failed to initialize logger:", err)
		return
	}
	defer logger.Sync()

	logger.Info("sample info")
	logger.Info("additional field sample", zap.String("key", "value"))
	logger.Warn("sample warn")
	logger.Error("sample error")
	err = errors.New("errorのサンプル")
	logger.Error("DB Connection failed", zap.Error(err))
}
