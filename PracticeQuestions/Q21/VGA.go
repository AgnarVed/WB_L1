package main

import "fmt"

type vga struct {
}

func (v *vga) InsertVGA() {
	fmt.Println("Inserted VGA")
}
