// Code generated by "stringer -type=IndexKind -output=index_kind_string_gen.go"; DO NOT EDIT.

package model

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[IndexKindInvalid-0]
	_ = x[IndexKindPrimaryKey-1]
	_ = x[IndexKindNormal-2]
	_ = x[IndexKindUnique-3]
	_ = x[IndexKindFullText-4]
	_ = x[IndexKindSpatial-5]
	_ = x[IndexKindForeignKey-6]
}

const _IndexKind_name = "IndexKindInvalidIndexKindPrimaryKeyIndexKindNormalIndexKindUniqueIndexKindFullTextIndexKindSpatialIndexKindForeignKey"

var _IndexKind_index = [...]uint8{0, 16, 35, 50, 65, 82, 98, 117}

func (i IndexKind) String() string {
	if i < 0 || i >= IndexKind(len(_IndexKind_index)-1) {
		return "IndexKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _IndexKind_name[_IndexKind_index[i]:_IndexKind_index[i+1]]
}
