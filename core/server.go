package core

func ServerRun() {
	// router
	router := Routers()
	//运行
	router.Run(":8088")
}
