package foo

import "fmt"

//在  v1.0.0 版本有效
//在  v1.3.0 版本fix
func Foo()  {
	fmt.Println("foo foo foo ...")
	fmt.Println("fixed ...")
}

//在 v1.1.0 版本有效
func Foo1()  {
	fmt.Println("foo1 foo1 foo1 ...")
}

//在 v1.2.0 版本有效
func Foo2()  {
	fmt.Println("foo2 foo2 foo2 ...")
}