package main

// #include "libABC.h"
// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lABC
import "C"

import (
    "fmt"
)

func main() {
    var pu_count C.int
    var do_count C.int
    C.count_PU_DO_Locations(C.CString("out.bin"), &pu_count, &do_count);
    fmt.Println(pu_count)
    fmt.Println(do_count)
}
