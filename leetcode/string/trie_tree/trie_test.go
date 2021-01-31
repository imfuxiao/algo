package trie_tree

import "testing"

func TestTrieNode_Insert(t *testing.T) {
	type fields struct {
		trie *TrieNode
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			fields: fields{
				trie: NewTrieNode(),
			},
			args: args{str: "hello"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.trie.Insert(tt.args.str); got != tt.want {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrieNode_Find(t *testing.T) {
	type fields struct {
		trie *TrieNode
	}
	type args struct {
		init []string
		str  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			fields: fields{trie: NewTrieNode()},
			args: args{
				init: []string{
					"hi",
					"how",
					"her",
					"hello",
					"so",
					"see",
				},
				str: "hell",
			},
			want: false,
		},
		{
			fields: fields{trie: NewTrieNode()},
			args: args{
				init: []string{
					"hi",
					"how",
					"her",
					"hello",
					"so",
					"see",
				},
				str: "what",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, str := range tt.args.init {
				tt.fields.trie.Insert(str)
			}
			if got := tt.fields.trie.Find(tt.args.str); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
