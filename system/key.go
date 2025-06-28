package system

type Key byte

func (k Key) ToKeySet() KeySet {
	return 1 << k
}

type KeySet uint16

func (ks KeySet) FirstKey() (Key, bool) {
	if ks == 0 {
		return 0, false
	}

	var result Key

	for ks != 1 {
		ks >>= 1
		result++
	}

	return result, true
}

func (ks KeySet) KeysNotIn(old KeySet) KeySet {
	return ^old & ks
}
