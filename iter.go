// Package iter provides a syntantically different way to iterate over integers. That's it.
//
// Just github.com/bradfitz/iter with N renamed to Iter, so that you can dot import it
// with less chance of name conflicts. 
package iter

// Iter returns a slice of n 0-sized elements, suitable for ranging over.
// It is intended to be used with dot imports.
//
// For example:
//
//    import . "github.com/benkehoe/iter"
//
//    for i := range Iter(10) {
//        fmt.Println(i)
//    }
//
// ... will print 0 to 9, inclusive.
//
// It does not cause any allocations.
func Iter(n int) []struct{} {
	return make([]struct{}, n)
}
