package pwnam

/*
#include <sys/types.h>
#include <pwd.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type Passwd struct {
	Uid uint32 ; Gid uint32 ; HomeDir string ; Shell string
}

func Getpwnam(name string) *Passwd {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cpw := C.getpwnam(cname)
	return &Passwd{ 
		Uid: uint32(cpw.pw_uid), Gid: uint32(cpw.pw_uid), 
		HomeDir: C.GoString(cpw.pw_dir), Shell: C.GoString(cpw.pw_shell) }
}
