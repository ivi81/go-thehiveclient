// Code generated by "stringer -type=SortConst"; DO NOT EDIT.

package v1const

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[desc-0]
	_ = x[asc-1]
	_ = x[_fields-2]
}

const _SortConst_name = "descasc_fields"

var _SortConst_index = [...]uint8{0, 4, 7, 14}

func (i SortConst) String() string {
	if i < 0 || i >= SortConst(len(_SortConst_index)-1) {
		return "SortConst(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SortConst_name[_SortConst_index[i]:_SortConst_index[i+1]]
}