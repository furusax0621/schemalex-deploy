// Code generated by "stringer -type=ReferenceMatch -output=reference_match_string_gen.go"; DO NOT EDIT.

package model

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ReferenceMatchNone-0]
	_ = x[ReferenceMatchFull-1]
	_ = x[ReferenceMatchPartial-2]
	_ = x[ReferenceMatchSimple-3]
}

const _ReferenceMatch_name = "ReferenceMatchNoneReferenceMatchFullReferenceMatchPartialReferenceMatchSimple"

var _ReferenceMatch_index = [...]uint8{0, 18, 36, 57, 77}

func (i ReferenceMatch) String() string {
	if i < 0 || i >= ReferenceMatch(len(_ReferenceMatch_index)-1) {
		return "ReferenceMatch(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ReferenceMatch_name[_ReferenceMatch_index[i]:_ReferenceMatch_index[i+1]]
}
