package internal

import (
	"errors"
	"strconv"
	"sync"
)

type Notifier struct {
	lock *sync.Mutex
	list *channelList
	name string
}

func (s *Notifier) Name() string {
	return s.name
}

func (s *Notifier) Register(id string, session string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.list.RegisterToList(id, session)

}

func (s *Notifier) RegisterAndGet(id string, session string) (map[string]chan string, error) {
	s.Register(id, session)
	return s.Get(id)
}

func (s *Notifier) Send(id string, msg string) {

	for _, channel := range s.list.items[id] {
		channel <- msg
	}

}

func (s *Notifier) Get(id string) (map[string]chan string, error) {

	if _, ok := s.list.items[id]; ok {
		return s.list.items[id], nil
	}

	return nil, errors.New("not found")
}

func (s *Notifier) Deregister(id string, session string) {
	// remove from map
	s.lock.Lock()
	defer s.lock.Unlock()
	s.list.DeregisterFromList(id, session)

}

func (s *Notifier) IsExist(id string) bool {
	_, err := s.Get(id)
	return !(err != nil)
}

func (s *Notifier) GetList() []string {
	// retrun list keys
	s.lock.Lock()
	defer s.lock.Unlock()

	var list []string
	for k, v := range s.list.items {
		list = append(list, k+"-"+strconv.Itoa(len(v)))
	}

	return list
}

func New(name string) *Notifier {
	return &Notifier{
		name: name,
		lock: &sync.Mutex{},
		list: &channelList{
			items: make(map[string]map[string]chan string),
			lock:  &sync.Mutex{},
		},
	}
}
