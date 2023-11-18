package nandu

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

func Passthrough(i1, _ bool) bool {
	return i1
}

func AlwaysFalse(_, _ bool) bool {
	return false
}
