# CPP + Golang Shared library on Ubuntu (WSL2)

To set this up:

1. Create the shared library (running `make` in the `cpp/` directory)
2. Move the resulting shared library into your `/usr/lib` folder.
3. Copy the header file (in this case, `libABC.h`) into the same directory as your Go application, but remove the `extern "C"` from the file.
4. Add the following lines to the top of the `main.go` file.

> // #include "libABC.h"<br/>
> // #cgo CFLAGS: -I.<br/>
> // #cgo LDFLAGS: -L. -lABC<br/>
> import "C"

...where the `-l` part is the name of shared library name without the `lib` prefix nor the `.so` suffix.

5. Build the Go app, and run it!

> go build main.go<br/>
> ./main.go

6. Tada! You've linked a C++ shared library to a Go application!
