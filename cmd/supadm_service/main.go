package main

import (
	"context"
	"log"

	"github.com/damnn/tulahack/your-supadmin-service/tools"
	"github.com/damnn/tulahack/your-supadmin-service/transport"
	_ "github.com/lib/pq"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := tools.SetupLogger()
	if err != nil {
		log.Fatal(err)
	}

	config, err := tools.LoadDammnAppConfig(logger)
	if err != nil {
		logger.Log().Fatal().AnErr("load env configs err", err)
	}

	logger.Log().Printf("[%#v]", config)

	err = tools.BeginDbInstance(config, logger)
	if err != nil {
		logger.Log().Printf("[%v]", err)
	}

	err = tools.NewExternalCache(ctx, config, logger)
	if err != nil {
		logger.Log().Printf("[%v]", err)
	}

	proxyServer := transport.NewAuthProxy(ctx, logger, config)

	go func() {
		if err := proxyServer.Run(); err != nil {
			logger.Log().Fatal().Err(err)
		}
	}()

	proxyServer.GracefullNotify()

	// closeDeps

}
