package linked

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	// 示例 1
	// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
	// 输出：[1,1,2,3,4,4,5,6]
	l1 := make([]*ListNode, 3)
	l1[0] = NewListNode(1).Add(4).Add(5)
	l1[1] = NewListNode(1).Add(3).Add(4)
	l1[2] = NewListNode(2).Add(6)

	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{lists: l1},
			want: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := print(mergeKLists(tt.args.lists)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
