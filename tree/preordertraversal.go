type Node struct {
	val   int
	left  *Node
	right *Node
}

func PreorderTraversal(root *node){
  if root == nil{
    return
  }
  fmt.Println(root.id)
  PreorderTraversal(root.left)
  PreorderTraversal(root.right)
}
