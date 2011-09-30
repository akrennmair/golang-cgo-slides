package main

import (
	"fmt"
	"pwnam"
)

func main() {
	pw := pwnam.Getpwnam("root")
	fmt.Printf("uid = %v gid = %v homedir = %v shell = %v\n", 
		pw.Uid, pw.Gid, pw.HomeDir, pw.Shell)
}
