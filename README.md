# go-practice

### Basic

1. Hello goroutines — fire and forget (concept: go keyword, sync.WaitGroup, loop variable capture trap)
2. Ping-pong between two goroutines (concept: unbuffered channel as rendezvous point)
3. Job queue with buffered channel (concept: buffered vs unbuffered, backpressure when buffer full)
4. Parallel file processor — no channels (concept: WaitGroup + Mutex, two separate concerns)
5. Thread-safe counter struct (concept: sync.Mutex, pointer receivers)
6. Typed channel pipeline — directional channels (concept: chan<- T vs <-chan T)

### Intermediate

7. Fan-in: merge two channels with select (concept: select multiplexing, non-deterministic case selection)
8. Operation with deadline — time.After (concept: select + timeout, time.After vs time.NewTimer)
9. Cancellable goroutine tree (concept: context.WithCancel, propagation, no goroutine leaks)
10. Context with timeout + value passing (concept: WithTimeout, WithValue, typed keys)
11. Concurrent read-heavy cache (concept: sync.RWMutex, RLock vs Lock, when RWMutex is worse)
12. Singleton initializer (concept: sync.Once, why it beats double-checked locking)
13. Three-stage data pipeline (concept: channel chaining, close() for termination, range over channel)
14. Fixed-size worker pool (concept: bounded concurrency, jobs+results channels, sizing workers)

### Advanced

15. Semaphore via buffered channel (concept: acquire=send, release=receive, fairness)
16. Graceful shutdown with done channel (concept: close() broadcasts to ALL receivers)
17. Parallel tasks with first-error cancellation (concept: errgroup, automatic cancel on first error)
18. Token bucket rate limiter (concept: time.Ticker refill, buffered channel as token bucket)
19. Sharded concurrent map (concept: lock striping, hash-to-shard, reduce contention)
20. Distributed request tracing simulation (concept: context propagation across call boundaries)

### Expert

21. Lock-free stack with atomic CAS (concept: atomic.CompareAndSwap, ABA problem, when lock-free is slower)
22. Thread-safe LRU cache (concept: doubly linked list + hashmap, O(1) eviction under concurrency)
23. In-process pub/sub event bus (concept: per-subscriber buffered channel, non-blocking publish with select+default)
24. Circuit breaker with state machine (concept: atomic state transitions, Closed→Open→HalfOpen, time-based probing)

Here's a focused revision list — 10 problems, ~2 hours:

---

**Tier 1 — Must do, high interview probability (45 min)**

These test the patterns that come up most:

1. **#3 — Job queue** — buffered channel, workers, close() propagation
2. **#9 — Cancellable goroutine tree** — context.WithCancel, ticker+select
3. **#14 — Worker pool** — bounded concurrency, jobs+results, closer goroutine
4. **#16 — Graceful shutdown** — close() broadcasts, done channel

---

**Tier 2 — Likely to come up, tests depth (45 min)**

5. **#7 — Fan-in merge** — select multiplexing, termination with two channels
6. **#11 — RWMutex cache** — RLock vs Lock, thread-safe reads
7. **#17 — errgroup** — first-error cancellation, context propagation
8. **#13 — Three-stage pipeline** — close() propagation, range over channel

---

**Tier 3 — Good to refresh, tests polish (30 min)**

9. **#10 — Context with timeout+value** — typed keys, defer cancel(), ctx.Err()
10. **#19 — Sharded map** — lock striping, hash-to-shard, RWMutex per shard

---

**Before each problem, ask yourself:**

1. Who owns closing each channel?
2. What stops main from exiting?
3. Is every read of shared state locked?

---

**Skip for revision:** #1, #2, #4, #5, #6 (too basic), #15 (semaphore rarely asked), #18 (rate limiter — know concept, unlikely as coding question), #21–#24 (advanced, low probability).
