package main

import (   // imports are only have file level scope and must be identified in each file where used
	"fmt"
	"github.com/Grayspawn/Goworkspace/ToddMc/04-scope/01-package_scope/02/viz"
)

func main ()  {

	fmt.Println(viz.MiName)
	viz.Printvarr()
}
