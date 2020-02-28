//
//  Practicing ELK
//
//  Copyright © 2020. All rights reserved.
//

package main

import (
	conf "github.com/moemoe89/practicing-elk-golang/config"
	"github.com/moemoe89/practicing-elk-golang/routers"

	"fmt"
)

func main() {
	log := conf.InitLog()

	app := routers.GetRouter(log)
	err := app.Start(":" + conf.Configuration.Port)
	if err != nil {
		panic(fmt.Sprintf("Can't start the app: %s", err.Error()))
	}
}
