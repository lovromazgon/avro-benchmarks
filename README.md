![Logo](http://svg.wiersma.co.za/github/project?lang=go&title=avro-benchmarks)

This fork includes the following changes:

* All libraries are updated to the latest version (on 18.05.2023)
  * Consequently, the models for gogen-avro are regenerated
* It adds benchmarks `BenchmarkHambaDecodeMap` and `BenchmarkHambaEncodeMap`,
  which test the speed of encoding and decoding a generic map using hamba, as
  opposed to using a Go struct as the existing benchmarks. This is relevant for
  use cases where your implementation does not know the Go struct in advance.

#### Libraries

* [github.com/hamba/avro/v2](https://github.com/hamba/avro)
* [github.com/go-avro/avro](https://github.com/go-avro/avro) (unmaintained)
* [github.com/linkedin/goavro/v2](https://github.com/linkedin/goavro)
* [github.com/actgardner/gogen-avro/v10](https://github.com/actgardner/gogen-avro)

#### Results

Go: 1.20.4

```
goos: darwin
goarch: amd64
pkg: github.com/nrwiersma/avro-benchmarks
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkGoAvroDecode-16       	  431052	      2341 ns/op	     418 B/op	      27 allocs/op
BenchmarkGoAvroEncode-16       	  352249	      2966 ns/op	     973 B/op	      63 allocs/op
BenchmarkGoGenAvroDecode-16    	  865196	      1169 ns/op	     320 B/op	      11 allocs/op
BenchmarkGoGenAvroEncode-16    	 2053659	       586.9 ns/op	     240 B/op	       3 allocs/op
BenchmarkHambaDecode-16        	 3026668	       388.5 ns/op	      64 B/op	       4 allocs/op
BenchmarkHambaEncode-16        	 4245242	       273.2 ns/op	     112 B/op	       1 allocs/op
BenchmarkHambaDecodeMap-16     	  533870	      2216 ns/op	    1440 B/op	      36 allocs/op
BenchmarkHambaEncodeMap-16     	  591298	      1967 ns/op	     112 B/op	       1 allocs/op
BenchmarkLinkedinDecode-16     	  801688	      1531 ns/op	    1688 B/op	      35 allocs/op
BenchmarkLinkedinEncode-16     	 2291268	       519.0 ns/op	     248 B/op	       5 allocs/op
PASS
ok  	github.com/nrwiersma/avro-benchmarks	16.426s
```