// Code generated by "stringer -type=execState"; DO NOT EDIT

package local

import "fmt"

const _execState_name = "execUnstartedexecInitexecCreatedexecRunningexecComplete"

var _execState_index = [...]uint8{0, 13, 21, 32, 43, 55}

func (i execState) String() string {
	if i < 0 || i >= execState(len(_execState_index)-1) {
		return fmt.Sprintf("execState(%d)", i)
	}
	return _execState_name[_execState_index[i]:_execState_index[i+1]]
}