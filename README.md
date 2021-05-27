# Parsing SIP URI
Different technics to parse SIP URI:
- re2go
- ragel
- dummy parsing with simple string split
- handmade lexer
- and regexp

Lexer re2go (re2c), ragel and lexer are following RFC3261 specs. Others have just basic implementation.

Benchmarks:
```
$ go test -bench=. -benchmem                                                                                                             19:30:57
goos: linux
goarch: amd64
pkg: uri
cpu: Intel(R) Core(TM) i5-5300U CPU @ 2.30GHz
BenchmarkLexerParse-4    	   94401	     13018 ns/op	    1882 B/op	      34 allocs/op
BenchmarkNetURLParse-4   	 1713309	       685.6 ns/op	     192 B/op	       2 allocs/op
BenchmarkDummyParser-4   	 9100580	       130.8 ns/op	      80 B/op	       1 allocs/op
BenchmarkRegexParse-4    	   49527	     24573 ns/op	   16462 B/op	     111 allocs/op
BenchmarkRe2GoParse-4    	 2830198	       393.9 ns/op	      80 B/op	       1 allocs/op
BenchmarkRagelParse-4    	 4594177	       270.4 ns/op	      80 B/op	       1 allocs/op
PASS
ok  	uri	9.104s
```
