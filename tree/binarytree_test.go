package tree

import (
	"reflect"
	"testing"
)

func TestBinarySearchTree_Insert(t *testing.T) {
	type fields struct {
		b *BinarySearchTree
	}
	type args struct {
		args []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []interface{}
	}{
		{
			fields: fields{
				b: NewBinarySearchTree(NewIntNode(13)),
			},
			args: args{
				args: []int{16, 18, 10, 8, 6, 20},
			},
			want: []interface{}{6, 8, 10, 13, 16, 18, 20},
		},
		{
			fields: fields{
				b: NewBinarySearchTree(NewIntNode(15)),
			},
			args: args{
				args: []int{7, 11, 3, 8, 18, 20},
			},
			want: []interface{}{3, 7, 8, 11, 15, 18, 20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.args.args {
				tt.fields.b.Insert(NewIntNode(v))
			}
			if v := tt.fields.b.Values(); !reflect.DeepEqual(v, tt.want) {
				t.Fatalf("you want %+v, real value %+v", tt.want, v)
			}
		})
	}
}

func TestBinarySearchTree_Search(t *testing.T) {
	type fields struct {
		b *BinarySearchTree
	}
	type args struct {
		values []int
		n      *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			fields: fields{
				b: NewBinarySearchTree(NewIntNode(13)),
			},
			args: args{
				values: []int{90, 18, 30, 1, 4, 3},
				n:      NewIntNode(3),
			},
			want: 3,
		},
		{
			fields: fields{
				b: NewBinarySearchTree(NewIntNode(13)),
			},
			args: args{
				values: []int{90, 18, 30, 1, 4, 3},
				n:      NewIntNode(999),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.args.values {
				tt.fields.b.Insert(NewIntNode(v))
			}
			if got := tt.fields.b.Search(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinarySearchTree.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchTree_Delete(t *testing.T) {
	type fields struct {
		b *BinarySearchTree
	}
	type args struct {
		n    *Node
		args []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []interface{}
	}{
		{
			fields: fields{
				b: NewBinarySearchTree(NewIntNode(300)),
			},
			args: args{
				n:    NewIntNode(1000),
				args: []int{100, 500, 0, 200, 400, 1000, 700, 1500, 600, 900, 1250, 2000, 1400},
			},
			want: []interface{}{0, 100, 200, 300, 400, 500, 600, 700, 900, 1250, 1400, 1500, 2000},
		},
		{
			fields: fields{
				b: NewBinarySearchTree(NewIntNode(300)),
			},
			args: args{
				n:    NewIntNode(100),
				args: []int{100, 500, 0, 200, 400, 1000, 700, 1500, 600, 900, 1250, 2000, 1400},
			},
			want: []interface{}{0, 200, 300, 400, 500, 600, 700, 900, 1000, 1250, 1400, 1500, 2000},
		},
		{
			fields: fields{
				b: NewBinarySearchTree(NewIntNode(300)),
			},
			args: args{
				n:    NewIntNode(0),
				args: []int{100, 500, 0, 200, 400, 1000, 700, 1500, 600, 900, 1250, 2000, 1400},
			},
			want: []interface{}{100, 200, 300, 400, 500, 600, 700, 900, 1000, 1250, 1400, 1500, 2000},
		},
		{
			fields: fields{
				b: NewBinarySearchTree(NewIntNode(300)),
			},
			args: args{
				n:    NewIntNode(1250),
				args: []int{100, 500, 0, 200, 400, 1000, 700, 1500, 600, 900, 1250, 2000, 1400},
			},
			want: []interface{}{0, 100, 200, 300, 400, 500, 600, 700, 900, 1000, 1400, 1500, 2000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.args.args {
				tt.fields.b.Insert(NewIntNode(v))
			}
			tt.fields.b.Delete(tt.args.n)
			if v := tt.fields.b.Values(); !reflect.DeepEqual(v, tt.want) {
				t.Fatalf("you want value: %+v, real value: %+v", tt.want, v)
			}
		})
	}
}

func TestBinaryTree_InNode(t *testing.T) {
	type fields struct {
		rootNode *Node
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			fields: fields{rootNode: &Node{
				value:     intNode(1),
				leftNode:  &Node{
					value:     intNode(3),
					leftNode:  &Node{value: intNode(5)},
					rightNode: &Node{
						value:     intNode(4),
						rightNode: &Node{
							value:     intNode(6),
						},
					},
				},
				rightNode: &Node{
					value:     intNode(2),
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &BinaryTree{
				rootNode: tt.fields.rootNode,
			}
			//b.InNode()
			//b.PostNode()
			b.PreNode()
		})
	}
}