package main

import "github.com/johnsudaar/acp/config"

func main() {
	viper := config.GetViperConfig()
	err := viper.WriteConfigAs("./acp.yml")
	if err != nil {
		panic(err)
	}
}
