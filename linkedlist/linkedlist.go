package main
import "fmt"

type node struct{
	data int 
	next *node
}

type linkedList struct {
	head *node
	length int
}

func (l *linkedList) prepend( n*node){
	second := l.head
	l.head = n
	l.head.next = second
	l.length++

}

func (l linkedList) printListData(){
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}
func (l *linkedList) deleteWithValue (value int){
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next  == nil {
			return 
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main(){
	myList := linkedList{}

	node1 := &node{data:48}
	node2 := &node{data:15}
	node3 := &node{data:22}
	node4 := &node{data:18}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	
	myList.printListData()
	myList.deleteWithValue(22)
	myList.printListData()


}