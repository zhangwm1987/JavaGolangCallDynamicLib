package main

/*
#cgo LDFLAGS: -ldl

#include <dlfcn.h>
void do_test_so_func(void);

void do_test_so_func()
{
    void* handle;
    typedef int (*FPTR)();

    handle = dlopen("./test.so", 1);
    FPTR fptr = (FPTR)dlsym(handle, "HelloWorld");

    (*fptr)();
    return;
}

void do_test_client_so()
{
    void* handle;
    typedef int (*FPTR)(int);

    handle = dlopen("./client.so", 1);
    FPTR fptr = (FPTR)dlsym(handle, "SayHello");

    (*fptr)(5);
    return;
}


*/
import "C"
import "fmt"

func main() {
    C.do_test_so_func()
    fmt.Printf("test client below\n")
    C.do_test_client_so()
}
