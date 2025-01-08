package typecast

func BoolToBoolPtr(b bool) *bool {
	return &b
}

func BoolPtrToBool(b *bool) bool {
	if b == nil {
		return false
	}
	return *b
}
