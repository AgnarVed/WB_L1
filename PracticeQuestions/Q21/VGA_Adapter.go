package main

type AdapterVGA struct {
	CableDVI *DVI
}

func (a *AdapterVGA) InsertVGA() {
	a.CableDVI.InsertDVI()
}
