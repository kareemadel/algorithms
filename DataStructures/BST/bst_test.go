package bst

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var r *node
	r = r.insert(&item{50})
	r.insert(&item{25})
	r.insert(&item{25})
	r.insert(&item{25})
	r.insert(&item{25})
	r.insert(&item{75})
	r.insert(&item{12})
	r.insert(&item{37})
	r.insert(&item{62})
	r.insert(&item{87})
	r.insert(&item{57})
	r.insert(&item{68})
	r.insert(&item{82})
	// m := r.search(&item{37})
	// fmt.Println(r.leftSize)
	// r.delete(1)
	// fmt.Println(r.right.leftSize)
	fmt.Println(r.selectOrder(7).data)
	fmt.Println(r.getRank(&item{62}))
	Print(r)
}
