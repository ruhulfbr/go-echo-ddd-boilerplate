package bootstrap

func Run() error {
	cfg, err := LoadEnv()
	if err != nil {
		return err
	}

	loggerCleanup, err := InitLogger(cfg)
	if err != nil {
		return err
	}
	defer loggerCleanup()

	gormDB, err := InitDatabase(cfg)
	if err != nil {
		return err
	}
	defer CloseDatabase(gormDB)

	// Wire Repos → Services → Handlers
	app, err := InitServer(cfg, gormDB)
	if err != nil {
		return err
	}

	return StartServer(app, cfg)
}
