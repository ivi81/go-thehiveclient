// Code generated by "stringer -type=ResolutionStatus"; DO NOT EDIT.

package caseconst

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Indeterminate-0]
	_ = x[FalsePositive-1]
	_ = x[TruePositive-2]
	_ = x[Other-3]
	_ = x[Duplicated-4]
}

const _ResolutionStatus_name = "IndeterminateFalsePositiveTruePositiveOtherDuplicated"

var _ResolutionStatus_index = [...]uint8{0, 13, 26, 38, 43, 53}

func (i ResolutionStatus) String() string {
	if i < 0 || i >= ResolutionStatus(len(_ResolutionStatus_index)-1) {
		return "ResolutionStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ResolutionStatus_name[_ResolutionStatus_index[i]:_ResolutionStatus_index[i+1]]
}