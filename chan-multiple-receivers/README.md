Lesson learnt:

1. it is okay to have many writers and many readers for one channel.
2. always check `val, isOpen := <-chan`. If `isOpen` == `false`, then stop using the channel.
3. `sync.RWMutex.RLock()` does not guarantee other goroutine from making changes to the map. So, becareful. This is how it should be done.
   ```
    func (mux *MapMux) AddInt(i int) {
        mux.RLock()
        if _, isExist := mux.m[i]; isExist {
            log.Fatal("race condition")
        }
        mux.RUnlock()
        mux.Lock()
        if _, isExist := mux.m[i]; isExist {
            log.Fatal("race condition")
        }
        mux.m[i] = true
        mux.Unlock()
        mux.Done()
    }
    ```