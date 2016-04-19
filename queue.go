package sop

import (
	"math/rand"
	"time"
)

type Object struct {
	Ptr     *interface{}
	taskNum int
}

type Job *func()

func (j *Job) Do(...params) {

}

type Pool struct {
	Objects map[int]Object
	Queue   []Job
}

func (p *Pool) Add(ptr *interface{}) {
	rnd := rand.NewSource(rand.Seed(time.Now().UnixNano()))
	id := rnd.Int63()
	p[id] = Object{Ptr: ptr}
}

func (p *Pool) remove(id int) {
	delete(p[id])
}

func (p *Pool) Remove(id int) {
	if p[id].taskNum > 0 {
		p.Do()
	} else {
		p.remove(id)
	}
}
