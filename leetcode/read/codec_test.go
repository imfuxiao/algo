package read

import "testing"

func TestConstructorCodec(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
	}
	ser := ConstructorCodec()
	data := ser.serialize(root)
	if want := "[1,2,3,null,null,4,5]"; want != data {
		t.Fatalf("want: %s, got: %s", want, data)
	}
	ans := ser.deserialize(data)
	t.Logf("%+v", ans)
}

func TestConstructorCodec2(t *testing.T) {
	coder := ConstructorCodec()
	node := coder.deserialize("[]")
	t.Log(node)
	//node := coder.deserialize("[1,2,3,null,null,4,5,6,7]")
	//t.Log(node)
}
