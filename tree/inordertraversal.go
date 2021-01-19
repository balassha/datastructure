type Node struct {
	val   int
	left  *Node
	right *Node
}

func InorderTraversal(root *node){
  if root == nil{
    return
  }
  InorderTraversal(root.left)
  fmt.Println(root.id)
  InorderTraversal(root.right)
}
