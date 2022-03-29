package fileos

type BaseApplication struct {

}
func (ab *BaseApplication) add(a int, b int) int {
	return a + b
}

func (ab *BaseApplication) del(a int, b int) int {
	return a - b
}

func (ab *BaseApplication) mul(a int, b int) int {
	return a * b
}

func (ab *BaseApplication) shang(a int, b int) int {
	return a / b
}