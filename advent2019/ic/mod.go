package ic

type mod struct {
	Pos int
	Val int
}

// Mods encode changes to the program to be made at runtime.
func Mod(pos, val int) mod {
	return mod{
		Pos: pos,
		Val: val,
	}
}
