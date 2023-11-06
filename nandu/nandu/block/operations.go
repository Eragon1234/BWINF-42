package block

func PassThrough(i1, _ bool) bool {
	return i1
}

func BlueBlock(i1, _ bool) bool {
	return i1
}

func RedBlockNonSensor(_, i2 bool) bool {
	return !i2
}

func RedBlockSensor(i1, _ bool) bool {
	return !i1
}

func WhiteBlock(i1, i2 bool) bool {
	o := !(i1 && i2)
	return o
}

func NoBlock(_, _ bool) bool {
	return false
}
