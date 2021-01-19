type Node struct {
	val   int
	left  *Node
	right *Node
}

func PostorderTraversal(root *node){
  if root == nil{
    return
  }
  PostorderTraversal(root.left)
  PostorderTraversal(root.right)
  fmt.Println(root.id)
}
