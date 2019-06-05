package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}
	node.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}


/**---------------Tranverse by Channel---------------------**/


//Node内部函数 遍历 返回 带有Node引用的通道
func (myNode *Node)  TraverseByChannel() chan *Node{

	c:=make(chan *Node)

	go func() {
		//每次遍历将对应的node作为采纳数
		myNode.TraverseFunc(func(node *Node) { //左中右
			c<-node
		})
		close(c)
	}()

	return c

}