package pkga

import "basic/pkgb"

type User struct {
	Name string
	Age  uint8
}

func init() {
	pkgb.Log("pkgb init")
}
