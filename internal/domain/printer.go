package domain

import "fmt"

type MarsReport interface {
	Print()
}

type Reporter struct {
	Explorer *MarsExplorer
}

func (r Reporter) Print() {
	for _, r := range r.Explorer.Robots {
		fmt.Println(r.ToString())
	}
}
