package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {

	var wg sync.WaitGroup
	var mu sync.Mutex

	ch := make(chan struct{}, pool)

	var i int64

	for i = 0; i < n; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(i int64) {
			usr := getOne(i)
			<-ch
			mu.Lock()
			res = append(res, usr)
			mu.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()

	return
}
