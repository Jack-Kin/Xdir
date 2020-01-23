package msg6030

// 无符号 large int
func ConvertToInt (a uint32) (b uint64) {
	switch a >> 30 {
	case 0:
		return uint64(a)
	case 1:
		return uint64(a << 2) << 2
	case 2:
		return uint64(a << 2) << 6
	case 3:
		return uint64(a << 2) << 10 //最大 4,398,046,507,008‬
	default:
		return 0
	}
}

