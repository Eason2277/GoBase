package main

import "fmt"

// OrderNode 表示待发货队列中的一个订单节点。
// Next 保存下一个节点的地址，尾节点的 Next 为 nil。
type OrderNode struct {
	OrderID string
	Next    *OrderNode
}

func main() {
	first := &OrderNode{
		OrderID: "A1001",
	}

	second := &OrderNode{
		OrderID: "A1002",
	}

	third := &OrderNode{
		OrderID: "A1003",
	}

	// 按顺序连接节点：A1001 -> A1002 -> A1003 -> nil。
	first.Next = second
	second.Next = third

	// 如果要删除 A1003，就让它的前驱节点直接指向它的后继节点：
	// second.Next = third.Next

	// 删除头节点时，只需要把链表入口移动到第二个节点。
	// A1001 节点仍然存在，但已经无法从 first 访问到它。
	first = first.Next

	// 从当前头节点开始遍历，遇到 nil 时说明已经到达链表末尾。
	current := first

	for current != nil {
		fmt.Println(current.OrderID, "当前订单的编号")

		// 移动遍历游标，但不修改节点之间的连接关系。
		current = current.Next
	}
}
