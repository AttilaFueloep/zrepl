// Code generated by "stringer -type=Status"; DO NOT EDIT.

package rpc

import "fmt"

const _Status_name = "StatusOKStatusRequestErrorStatusServerErrorStatusError"

var _Status_index = [...]uint8{0, 8, 26, 43, 54}

func (i Status) String() string {
	i -= 1
	if i >= Status(len(_Status_index)-1) {
		return fmt.Sprintf("Status(%d)", i+1)
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}
