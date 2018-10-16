all:
	c-for-go -ccdefs -out .. vma.yml

clean:
	rm -f doc.go types.go const.go
	rm -f cgo_helpers.go cgo_helpers.c cgo_helpers.h
	rm -f vma.go

test:
	go build
