package main

import (
    "ginblog/model"
    "ginblog/routes"
    "ginblog/utils"
)

func main() {
    model.InitDb()
    oss.Initserver()
    r := routes.InitRouter()
    r.Run(utils.HttpPort)

}