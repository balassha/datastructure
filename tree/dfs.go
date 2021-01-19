func SearchElementByDFS(n *node, id string)(*node){
  if n.id == id {
    return n
  }
  if len(n.children)>0 {
    for _,child := range n.children{
      SearchElementByDFS(child,id)
    }
  }
  return nil
}
