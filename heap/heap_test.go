package heap

import (
	"reflect"
	"testing"
)

func TestHeap_Add(t *testing.T) {
	type fields struct {
		h *Heap
	}
	type args struct {
		element []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			fields: fields{
				h: NewHeap(10),
			},
			args: args{
				element: []int{1, 2, 3, 4, 5, 6, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.args.element {
				if err := tt.fields.h.Add(v); (err != nil) != tt.wantErr {
					t.Errorf("Heap.Add() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
			t.Logf("%s", tt.fields.h)
		})
	}
}

func TestHeap_RemoveHeapTop(t *testing.T) {
	type fields struct {
		h *Heap
	}
	type args struct {
		elements    []int
		removeCount int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			fields: fields{
				h: NewHeap(10),
			},
			args: args{
				elements:    []int{1, 2, 3, 4, 5, 6, 7},
				removeCount: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.args.elements {
				_ = tt.fields.h.Add(v)
			}
			t.Logf("%s", tt.fields.h.String())
			for i := 0; i < tt.args.removeCount; i++ {
				tt.fields.h.RemoveHeapTop()
				t.Logf("%s", tt.fields.h.String())
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	type args struct {
		elements []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				elements: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.elements)
			if !reflect.DeepEqual(tt.args.elements, tt.want) {
				t.Errorf("you want: %+v, result: %+v", tt.want, tt.args.elements)
			}
		})
	}
}
