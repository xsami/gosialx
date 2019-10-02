package mongodb

import mgo "gopkg.in/mgo.v2"

type Service struct {
	baseSession *mgo.Session
	queue       chan int
	URL         string
	Open        int
}

var service Service

func (s *Service) New() error {
	var err error
	s.queue = make(chan int, MaxPool)
	for i := 0; i < MaxPool; i = i + 1 {
		s.queue <- 1
	}
	s.Open = 0
	s.baseSession, err = mgo.Dial(s.URL)
	return err
}

func (s *Service) Session() *mgo.Session {
	if len(s.queue) < MaxPool {
		s.queue <- 1
	}
	s.Open++
	return s.baseSession.Copy()
}

func (s *Service) Close(c *Collection) {
	c.db.s.Close()
	if len(s.queue) > 0 {
		<-s.queue
	}
	s.Open--
}