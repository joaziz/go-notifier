package local

import "sync"

type channelList struct {
	items map[string]map[string]chan string
	lock  *sync.Mutex
}

func (l *channelList) RegisterToList(id string, session string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if _, ok := (l.items)[id]; !ok {
		(l.items)[id] = make(map[string]chan string)
	}
	(l.items)[id][session] = make(chan string)
}

func (l *channelList) DeregisterFromList(id string, session string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if _, ok := (l.items)[id]; ok {
		if _, ok := (l.items)[id][session]; ok {
			close((l.items)[id][session])
			delete((l.items)[id], session)
		}
		if len((l.items)[id]) == 0 {
			delete((l.items), id)
		}
	}
}
