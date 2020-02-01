package gamewin

//CheckWin checks to see if the win condition has been met.
func CheckWin() bool {
	for i := 0; i < len(Win); i++ {
		wConditions()
		if Win[i] {
			return true
		}
	}
	return false
}
