package domain

import "fmt"

type MarsReport interface {
	Print()
}

type Reporter struct {
	Explorer *MarsExplorer
}

// Print range over the explorer robots and print their status over the standard output
func (r Reporter) Print() {
	for _, r := range r.Explorer.Robots {
		fmt.Println(r.ToString())
	}
}
