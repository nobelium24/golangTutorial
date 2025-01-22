package main

func returnWithPanic() (result int) {
	defer func() {
		if r := recover(); r != nil {
			result = r.(int)
		}
	}()
	panic(42)
}
