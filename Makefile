test: regen
	go test -v -race -cover

benchmark: regen
	go test -bench=. -benchmem
	wc -l parser_re.go parser_rl.go

regen: parser.re parser.rl
	re2go parser.re -o parser_re.go
	ragel -Z -G2 -o parser_rl.go parser.rl
.PHONY: regen
