package single

import (
	"reflect"
	"testing"
)

func TestLinkedNode_Add(t *testing.T) {
	type fields struct {
		linkedNode *LinkedNode
	}
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []interface{}
	}{
		{
			fields: fields{linkedNode: NewLinkedNode()},
			args:   args{values: []interface{}{1, 2, 3, 4, 5, 6}},
			want:   []interface{}{1, 2, 3, 4, 5, 6},
		},
		{
			fields: fields{linkedNode: NewLinkedNode()},
			args:   args{values: []interface{}{}},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			linkedNode := tt.fields.linkedNode
			linkedNode.Add(tt.args.values...)
			if got := print(linkedNode); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got = %+v, want = %+v", got, tt.want)
			}
		})
	}
}

func TestLinkedNode_Find(t *testing.T) {
	type fields struct {
		linkedNode *LinkedNode
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		{
			fields: fields{linkedNode: NewLinkedNode().Add(1, 2, 3, 4, 5)},
			args:   args{value: 3},
			want: &Node{
				val: 3,
			},
		},
		{
			fields: fields{linkedNode: NewLinkedNode().Add(1, 2, 3, 4, 5)},
			args:   args{value: 6},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.linkedNode.Find(tt.args.value); got != nil && tt.want != nil && !reflect.DeepEqual(got.Value(), tt.want.Value()) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedNode_Delete(t *testing.T) {
	type fields struct {
		linkedNode *LinkedNode
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		wantValue []interface{}
	}{
		{
			fields:    fields{linkedNode: NewLinkedNode().Add(1, 2, 3, 4, 5)},
			args:      args{value: 3},
			wantErr:   false,
			wantValue: []interface{}{1, 2, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			node := tt.fields.linkedNode.Find(tt.args.value)
			if node == nil {
				t.Error("Find() node is nil")
			}
			if err := tt.fields.linkedNode.Delete(node); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := print(tt.fields.linkedNode); !tt.wantErr && !reflect.DeepEqual(got, tt.wantValue) {
				t.Errorf("got = %+v, wantValue %+v", got, tt.wantValue)
			}
		})
	}
}

func TestLinkedNode_Invert(t *testing.T) {
	type fields struct {
		linkedNode *LinkedNode
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		{
			fields: fields{
				linkedNode: NewLinkedNode().Add(1, 2, 3, 4, 5),
			},
			want: []interface{}{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.linkedNode.Invert(); !reflect.DeepEqual(print(got), tt.want) {
				t.Errorf("Invert() = %v, want %v", got, tt.want)
			}
		})
	}
}
