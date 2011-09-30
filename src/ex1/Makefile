include $(GOROOT)/src/Make.inc

TARG=pwnam
CGOFILES=pwnam.go
CLEANFILES+=ex1

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O

