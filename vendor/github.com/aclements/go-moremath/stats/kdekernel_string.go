// generated by stringer -type=KDEKernel; DO NOT EDIT

package stats

import "fmt"

const _KDEKernel_name = "GaussianKernelDeltaKernel"

var _KDEKernel_index = [...]uint8{0, 14, 25}

func (i KDEKernel) String() string {
	if i < 0 || i+1 >= KDEKernel(len(_KDEKernel_index)) {
		return fmt.Sprintf("KDEKernel(%d)", i)
	}
	return _KDEKernel_name[_KDEKernel_index[i]:_KDEKernel_index[i+1]]
}
