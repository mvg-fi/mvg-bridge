package users

/*
func TestNewUser(t *testing.T) {
	ctx := context.Background()
	usr, _ := user.Current()
	cp := "../config/config1.toml"
	bp := filepath.Join(usr.HomeDir, "/mvg/bridge")
	conf, err := config.ReadConfiguration(cp)
	if err != nil {
		panic(err)
	}
	db, err := store.OpenBadger(ctx, bp)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	p := NewProxy(ctx, conf.MTG, conf.Proxy)
	user, err := p.NewUser(ctx, db)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	fmt.Printf("%+v", user)
}
*/
