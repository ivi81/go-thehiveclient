// Code generated by "stringer -type=Status"; DO NOT EDIT.

package caseconst

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Open-0]
	_ = x[Resolved-1]
	_ = x[Deleted-2]
}

const _Status_name = "OpenResolvedDeleted"

var _Status_index = [...]uint8{0, 4, 12, 19}

func (i Status) String() string {
	if i < 0 || i >= Status(len(_Status_index)-1) {
		return "Status(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}
