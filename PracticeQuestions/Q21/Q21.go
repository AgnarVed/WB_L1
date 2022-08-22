package main

// Реализовать паттерн «адаптер» на любом примере.

// Есть компьютер, у которого порт для монитора - VGA. Имеется два кабеля: VGA и DVI. Не зависимо от того, какой провод
// возьмёт пользователь, картинка на монитор должна выводиться

func main() {
	user := &user{}         // пользователь
	vga := &vga{}           // провод VGA
	user.InsertVGAinPC(vga) // пользователь подключает провод VGA к компьютеру
	dvi := &DVI{}           // провод DVI
	adapt := &AdapterVGA{   // условно, подключаем провод DVI к переходнику(адаптеру)
		CableDVI: dvi,
	}
	user.InsertVGAinPC(adapt) // используем провод DVI с переходником в качестве провода VGA
}
