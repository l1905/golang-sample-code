package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

func main() {

	e, err := casbin.NewEnforcer("/Users/smzdm/codesrc/go/ltexample/demo8/rbac_with_resource/model.conf", "/Users/smzdm/codesrc/go/ltexample/demo8/rbac_with_resource/policy.csv")
	if err != nil {
		fmt.Println("===")
		fmt.Println(err.Error())
	}

	// 普通用户
	sub := "alice" // the user that wants to access a resource.
	obj := "data2" // the resource that is going to be accessed.
	act := "write" // the operation that the user performs on the resource.


	ok, err := e.Enforce(sub, obj, act)
	fmt.Println(e.GetAllRoles())

	if err != nil {
		// handle err
		fmt.Println(err.Error())
	}

	if ok == true {
		// permit alice to read data1
		fmt.Println("ok")
	} else {
		// deny the request, show an error
		fmt.Println("deny")
	}



}
