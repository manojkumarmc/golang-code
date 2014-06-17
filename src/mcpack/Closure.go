package mcpack

func Adder(i int) (func(j int) int){
	return func (k int) int {
		i = i + k
		return i
	}
}

func Multi(i int) (func(j int) int) {
	return func(k int) int {
		i = i * k
		return i
	}
}

