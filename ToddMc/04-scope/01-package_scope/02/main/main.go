package main

import (
	"fmt"
	"github.com/GoesToEleven/GolangTraining/04_scope/01_package-scope/02_visibility/vis"
	"github.com/Grayspawn/Goworkspace/ToddMc/04-scope/01-package_scope/02/viz"
)

func main ()  {

	fmt.Println(vis.MiName)
	vis.PrintVar()
}
