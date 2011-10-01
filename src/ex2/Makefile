include $(GOROOT)/src/Make.inc

TARG=foo
CGOFILES=foo.go
CLEANFILES+=ex2

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O

