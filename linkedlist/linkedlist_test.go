package linkedlist

import (
	"reflect"
	"testing"
)

func TestList_Add(t *testing.T) {
	type fields struct {
		list *List
	}
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			fields: fields{
				list: New(1, 2, 3, 4),
			},
			args: args{
				v: []interface{}{5, 6, 7},
			},
			want: "[1 2 3 4 5 6 7]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.list.Add(tt.args.v...)
			if got := tt.fields.list.String(); got != tt.want {
				t.Errorf("List.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_String(t *testing.T) {
	type fields struct {
		list *List
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{
				list: New(1, 2, 3, 4),
			},
			want: "[1 2 3 4]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.list.String(); got != tt.want {
				t.Errorf("List.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Insert(t *testing.T) {
	type fields struct {
		list *List
	}
	type args struct {
		index int
		v     interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    string
	}{
		{
			fields: fields{
				list: New(1, 2, 3, 4),
			},
			args: args{
				index: -1,
				v:     0,
			},
			wantErr: true,
		},
		{
			fields: fields{
				list: New(1, 2, 3, 4),
			},
			args: args{
				index: 0,
				v:     0,
			},
			wantErr: false,
			want:    "[0 1 2 3 4]",
		},
		{
			fields: fields{
				list: New(1, 2, 3, 4),
			},
			args: args{
				index: 2,
				v:     6,
			},
			wantErr: false,
			want:    "[1 2 6 3 4]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.list.Insert(tt.args.index, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("List.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestList_Prepend(t *testing.T) {
	type fields struct {
		list *List
	}
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			fields: fields{
				list: New(1, 2, 3),
			},
			args: args{
				v: []interface{}{4, 5},
			},
			want: "[4 5 1 2 3]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.list.Prepend(tt.args.v...)
			if got := tt.fields.list.String(); got != tt.want {
				t.Errorf("List.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Get(t *testing.T) {
	type fields struct {
		list *List
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  bool
	}{
		{
			fields: fields{
				list: New(1, 2, 3),
			},
			args: args{
				index: 0,
			},
			want:  1,
			want1: true,
		},
		{
			fields: fields{
				list: New(1, 2, 3),
			},
			args: args{
				index: 3,
			},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.fields.list.Get(tt.args.index)
			if got1 != tt.want1 {
				t.Errorf("List.Get() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Remove(t *testing.T) {
	type fields struct {
		list *List
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    string
	}{
		{
			fields: fields{
				list: New(1, 2, 3),
			},
			args: args{
				index: 0,
			},
			wantErr: false,
			want:    "[2 3]",
		},
		{
			fields: fields{
				list: New(1, 2, 3),
			},
			args: args{
				index: 3,
			},
			wantErr: true,
			want:    "[1 2 3]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.list.Remove(tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("List.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := tt.fields.list.String(); got != tt.want {
				t.Errorf("List.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Contain(t *testing.T) {
	type fields struct {
		list *List
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			fields: fields{
				list: New(1, 2, 3),
			},
			args: args{
				v: 2,
			},
			want: true,
		},
		{
			fields: fields{
				list: New(1, 2, 3),
			},
			args: args{
				v: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.list.Contain(tt.args.v); got != tt.want {
				t.Errorf("List.Contain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_IndexOf(t *testing.T) {
	type fields struct {
		list *List
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			fields: fields{
				list: New(1, 2, 3),
			},
			args: args{
				v: 3,
			},
			want: 2,
		},
		{
			fields: fields{
				list: New(1, 2, 3),
			},
			args: args{
				v: 4,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.list.IndexOf(tt.args.v); got != tt.want {
				t.Errorf("List.IndexOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Size(t *testing.T) {
	type fields struct {
		list *List
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			fields: fields{
				list: New(1, 2, 3),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.list.Size(); got != tt.want {
				t.Errorf("List.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Clean(t *testing.T) {
	type fields struct {
		list *List
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{
				list: New(1, 2, 3),
			},
			want: "[]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.list.Clean()
		})
		if got := tt.fields.list.String(); got != tt.want {
			t.Errorf("List.String() = %v, want %v", got, tt.want)
		}
	}
}
