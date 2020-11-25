package stego

func setSuffix(x, suffix, suffixLength byte) byte {
	var suffixMask byte
	suffixMask = ^(1<<suffixLength - 1)
	x &= suffixMask
	x |= suffix
	return x
}
