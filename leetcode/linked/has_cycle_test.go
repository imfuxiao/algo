package linked

import "testing"

func TestHasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	// 测试案例1
	// 1 ->2, 2 ->1
	n1 := &ListNode{
		Val: 1,
	}
	n2 := &ListNode{
		Val: 2,
	}
	n1.Next = n2
	n2.Next = n1

	// 测试案例2
	// 3 -> 2 -> 0 -> -4, 循环 -4 -> 2
	n21 := &ListNode{
		Val: 3,
	}
	n22 := &ListNode{
		Val: 2,
	}
	n23 := &ListNode{
		Val: 0,
	}
	n24 := &ListNode{
		Val: -4,
	}
	n21.Next = n22
	n22.Next = n23
	n23.Next = n24
	n24.Next = n22

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				head: n21,
			},
			want: true,
		},
		{
			args: args{
				head: n1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasCycle1(tt.args.head); got != tt.want {
				t.Errorf("HasCycle1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasCycle2(t *testing.T) {
	// 测试案例1
	// 1 ->2, 2 ->1
	n1 := &ListNode{
		Val: 1,
	}
	n2 := &ListNode{
		Val: 2,
	}
	n1.Next = n2
	n2.Next = n1

	// 测试案例2
	// 3 -> 2 -> 0 -> -4, 循环 -4 -> 2
	n21 := &ListNode{
		Val: 3,
	}
	n22 := &ListNode{
		Val: 2,
	}
	n23 := &ListNode{
		Val: 0,
	}
	n24 := &ListNode{
		Val: -4,
	}
	n21.Next = n22
	n22.Next = n23
	n23.Next = n24
	n24.Next = n22

	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				head: n21,
			},
			want: true,
		},
		{
			args: args{
				head: n1,
			},
			want: true,
		},
		{
			args: args{
				head: &ListNode{
					Val: 1,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasCycle2(tt.args.head); got != tt.want {
				t.Errorf("HasCycle2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkHasCycle1-12    	13355978	        81.0 ns/op
func BenchmarkHasCycle1(b *testing.B) {
	// 测试案例2
	// 3 -> 2 -> 0 -> -4, 循环 -4 -> 2
	n21 := &ListNode{
		Val: 3,
	}
	n22 := &ListNode{
		Val: 2,
	}
	n23 := &ListNode{
		Val: 0,
	}
	n24 := &ListNode{
		Val: -4,
	}
	n21.Next = n22
	n22.Next = n23
	n23.Next = n24
	n24.Next = n22
	for i := 0; i < b.N; i++ {
		HasCycle1(n21)
	}
}

// BenchmarkHasCycle2-12    	629073163	         1.69 ns/op
func BenchmarkHasCycle2(b *testing.B) {
	// 测试案例2
	// 3 -> 2 -> 0 -> -4, 循环 -4 -> 2
	n21 := &ListNode{
		Val: 3,
	}
	n22 := &ListNode{
		Val: 2,
	}
	n23 := &ListNode{
		Val: 0,
	}
	n24 := &ListNode{
		Val: -4,
	}
	n21.Next = n22
	n22.Next = n23
	n23.Next = n24
	n24.Next = n22
	for i := 0; i < b.N; i++ {
		HasCycle2(n21)
	}
}
