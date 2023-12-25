// Code generated by "stringer -type=Status"; DO NOT EDIT.

package alertconst

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[New-0]
	_ = x[Updated-1]
	_ = x[Ignored-2]
	_ = x[Imported-3]
}

const _Status_name = "NewUpdatedIgnoredImported"

var _Status_index = [...]uint8{0, 3, 10, 17, 25}

func (i Status) String() string {
	if i < 0 || i >= Status(len(_Status_index)-1) {
		return "Status(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}
