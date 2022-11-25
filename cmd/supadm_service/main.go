package main

/* func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := tools.SetupLogger()
	config, err := tools.LoadDammnAppConfig(logger)
	if err != nil {
		logger.Log().Fatal().AnErr("load env configs err", err)
	}

	db, err := tools.BeginDbInstance(config, logger)
	if err != nil {
		logger.Log().Fatal().AnErr("db connection err", err)
	}

	cache, err := tools.NewExternalCache(ctx, config, logger)
	if err != nil {
		logger.Log().Fatal().AnErr("connect to external cache err", err)
	}

	// mock
	logger.Log().Debug().Interface("a", db)
	logger.Log().Debug().Interface("a", cache)

	repo := repository.NewAuthCoreRepository(ctx)

	serivce := auth.NewAuthService(ctx, repo)

	proxyServer := transport.NewAuthProxy(ctx, serivce)

	go func() {
		if err := proxyServer.Run(); err != nil {
			//
		}
	}()

	proxyServer.GracefullNotify()
	if err := proxyServer.Shotdown(); err != nil {
		// log

	}

	// closeDeps

} */
