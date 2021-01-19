func GetElementByDFS(root *node, id string)(*node){
  q := make([]*node,0)
  q = append(q,root)
    for len(q) > 0 {
      el := q[0]
      q = q[1:]
        if el.id == id {
          return el
        }
        if len(el.children) > 0 {
          for _, child := range el.children {
            q = append(q,child)
          }
        }
    }
    return nil
}
