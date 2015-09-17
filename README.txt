Based on github.com/bradfitz/iter, but with N() renamed to Iter(), 
for use with dot imports, just to save some typing:

import . "github.com/benkehoe/iter"

//later:

N := 10

for i := range Iter(N) {
    //stuff
}

See http://godoc.org/github.com/benkehoe/iter
