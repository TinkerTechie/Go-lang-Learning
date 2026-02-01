## What Happens at Scale

At scale, the cache holds millions of entries. Each entry allocates multiple pointer-based objects, including:
- string headers and string data
- slice headers
- slice backing arrays

With millions of entries, this results in **tens of millions of pointers** allocated on the heap.

---

## Garbage Collection Behavior

Go uses a **concurrent mark-and-sweep garbage collector**. During each GC cycle:

- GC starts from root objects
- Traverses every reachable pointer
- Marks live objects
- Sweeps unreachable memory

The cost of garbage collection grows with the **number of pointers**, not with business logic complexity.

---

## The Core Problem

Even when an HTTP request does **not** access the cache at all, the garbage collector still scans all cache entries during each GC cycle. This causes:

- Long GC mark phases
- Increased CPU usage
- Latency spikes across the entire service

Unrelated endpoints slow down simply because the cache exists in memory.

---

## Observed Symptoms

Under production-like load, the following behavior was observed:

- Mean latency remained stable
- 99th percentile latency spiked to hundreds of milliseconds or seconds
- Maximum latency reached several seconds
- CPU usage increased during GC cycles

These effects occurred even for endpoints that did not access the cache.

---

## Why Eviction Does Not Fix the Issue

Time-based eviction does not solve the GC problem:

- Entries remain reachable until eviction
- GC scans them during every cycle
- Large heaps amplify GC scanning cost

Eviction limits memory growth but does not prevent GC overhead.

---

## Key Insight

The main performance bottleneck was **not cache logic**, but the garbage collector scanning millions of pointer-based cache entries.

At large scale, **memory layout matters more than algorithmic complexity**.

---

## Direction Toward the Solution

To meet latency requirements, the cache design needed to:

- Minimize pointer usage
- Reduce GC visibility of cache entries
- Avoid scanning millions of objects during GC

This led to a redesigned cache architecture using:
- Sharding
- Pointer-free maps
- Offset-based storage

---

## Conclusion

This case study demonstrates that a cache implementation can be correct and fast at small scale, yet fail under heavy load due to garbage collection behavior.

Designing **GC-aware data structures** is essential for building low-latency, high-throughput systems in Go.
