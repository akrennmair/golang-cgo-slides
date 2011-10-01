package foo

/*

extern void myprint(int i);

void dofoo(void) {
	int i;
	for (i=0;i<10;i++) {
		myprint(i);
	}
}
*/
import "C"
import "fmt"

//export myprint
func myprint(i C.int) {
	fmt.Printf("i = %v\n", uint32(i))
}

func DoFoo() {
	C.dofoo()
}
