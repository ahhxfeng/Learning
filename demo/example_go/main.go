package main

func main() {
	list := LinkList{}

	list.append("1")
	list.append("20")
	list.append("33")
	list.append("99")
	list.append("77")
	list.append("end")

	list.prepend("0")
	list.prepend("-1")
	list.display()
	list.deleteWithValue("end")
	list.display()

}
