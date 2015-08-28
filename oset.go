
// Ordered Set of strings

package oset

// duplicated entries in a map for fast lookup
type Oset struct {
    m map[string]struct{}
    slice []string
}


func NewOset() *Oset {
    return &Oset{m: make(map[string]struct{}), slice: make([]string, 0)}
}

func (o *Oset) Add(s string) {
    if _, prs := o.m[s]; prs {
        // do nothing
    } else {
        o.m[s] = struct{}{}
        o.slice = append(o.slice, s)
    }

}

func (o *Oset) ToSlice() []string {
    return o.slice
}

