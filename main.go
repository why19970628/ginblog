package main

import (
    "ginblog/model"
    "ginblog/routes"
    "ginblog/utils"
)

func main() {
    model.InitDb()
    r := routes.InitRouter()
    r.Run(utils.HttpPort)

}