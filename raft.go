package lastwood

import "fmt"

type State uint64

const (
	Follower State = iota
	Candidate
	Leader
)

type Raft struct {
	state         State
	nodes         map[NodeId]*Node
	currentTerm   uint64
	votedFor      NodeId
	voteCount     uint64
	heartBeatChan chan bool
	toLeaderChan  chan bool
}

func newRaft() *Raft {
	return &Raft{
		state:         Follower,
		nodes:         make(map[NodeId]*Node),
		currentTerm:   0,
		votedFor:      "",
		voteCount:     0,
		heartBeatChan: make(chan bool),
		toLeaderChan:  make(chan bool),
	}
}

func (r *Raft) String() string {
	return fmt.Sprintf("{ state: %d, nodes: %+v, currentTerm: %d, votedFor: %s }", r.state, r.nodes, r.currentTerm, r.votedFor)
}

func (r *Raft) start() {
}
