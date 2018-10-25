package main

import (
	"fmt"

	majBranch "github.com/go-modules-by-example-staging/goinfo-maj-branch/designers"
	majBranch2 "github.com/go-modules-by-example-staging/goinfo-maj-branch/v2/designers"

	majSubdir "github.com/go-modules-by-example-staging/goinfo-maj-subdir/designers"
	majSubdir2 "github.com/go-modules-by-example-staging/goinfo-maj-subdir/v2/designers"

	majLossy "github.com/go-modules-by-example-staging/goinfo-maj-lossy/designers"
	majLossy2 "github.com/go-modules-by-example-staging/goinfo-maj-lossy/v2/designers"
)

func main() {
	fmt.Println(majBranch.Names())
	fmt.Println(majBranch2.FullNames())
	fmt.Println(majSubdir.Names())
	fmt.Println(majSubdir2.FullNames())
	fmt.Println(majLossy.Names())
	fmt.Println(majLossy2.FullNames())
}
