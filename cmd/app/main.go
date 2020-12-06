package main

func main() {
	injector, _, err := CreateApp()
	if err != nil {
		panic(err)
	}

	err = injector.Engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}
