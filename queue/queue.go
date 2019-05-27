package queue


type  Queue []interface{} //队列本质为任何类型的数组


func (q *Queue) Push (v interface{}){
	*q = append(*q,v)
}


func (q *Queue) Pop() interface{}{
	head:=(*q)[0]
	*q= (*q)[1:]
	return head
}



//如何把任何类型强制转为其他类型


func (q *Queue) PopInt() int{
	head:=(*q)[0]
	*q= (*q)[1:]
	return head.(int) //将interface 转为特定类型
}

func (q *Queue) IsEmpty() bool{
	return len(*q)==0
}