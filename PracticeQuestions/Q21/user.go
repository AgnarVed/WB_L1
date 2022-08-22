package main

type user struct {
}

func (u *user) InsertVGAinPC(pc computer) {
	pc.InsertVGA()
}
