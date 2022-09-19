package ytd

import "testing"

func TestUnit(t *testing.T) {
	goroutine := "2"
	fp := "/Users/zen/Github/youtube-dl-bat/ytd/list.txt"
	addr := "127.0.0.1"
	port := "8889"
	target := "/Users/zen/Github/youtube-dl-bat/ytd"
	isproxy := "true"
	Master(fp, addr, port, target, goroutine, isproxy)
}
