package lastwood

type NodeId string

type Node struct {
	NodeId NodeId
	connect bool
	address string
}
