# CPP + Golang Shared library on Ubuntu (WSL2)

To set this up:

1. When creating the shared library (*.so) file, if it is a C++ application (in this case, it is), then ensure that the header is surrounded with an `extern "C"`.
2. Create the shared library (running `make` in the `cpp/` directory)
2. Copy the resulting shared library (in this case, called `exe.so`) into the same directory as the Go file, but prefix the filename with `lib` (i.e. the new name should be `libexe.so`).
4. Copy the header file (in this case, `main.h`) into the same directory as the Go file, but remove the `extern "C"`.
5. Add the following lines to the top of the `main.go` file (but without the blank lines between each line):

> // #include "exe.h"
> 
> // #cgo CFLAGS: -I/home/user/repos/go_cpp_shared_library
> 
> // #cgo LDFLAGS: -L/home/user/repos/go_cpp_shared_library  -lexe
> 
> import "C"


...where both CFLAGS and LFLAGS point to the directory of the shared library `libexe.so`, and the `-l` part points to the shared library name without the `lib` nor the `.so` extension.

6. Build the Go app, and run it!

> go build main.go
> ./main.go

7. Tada! You've linked a C++ shared library to a Go application!
