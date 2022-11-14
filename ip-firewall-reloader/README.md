# IP Firewall Reloader / Multi-Reader Single-Writer Performance Benchmarks
This example also provides the performance benchmarks for different mechanisms to notify the readers about update of a variables storing some data (atomic pointers for the variables, atomic pointers for the version numbers, contexts with cancelations, mutexes).

The *primary* reason for this example is to benchmark the performance difference of the approaches considered from the reader's perspective (and not the writer/reloader/updater perspective).

### Run Tests
In order to run the test, run the following:
```sh
docker-compose up --build

```

## Approaches considered

For our benchmarks we consider multiple approaches as follows:

**Atomic Updates of Certificate Variable:** Using an pointer to the certificate variable and atomically updating it. Reads are also atomic.

**Logical Clock/Atomic Updates of Version Numbers:** Using a logical clock to track the version number of the certificate and atomically updating it from the writer. In this approach, once the version number changes/increases, the certificate variable is updated (using atomic instructions).

**Multi-Reader Single-Writer Mutexes:** Using `sync.RWMutex` in golang.

**Mutexes:** Using `sync.Mutex` in golang.

**Channels:** Using a channel of the pointer. This approach requires the write to know exactly the number of readers and is prone to race conditions with multiple readers (so it's only feasible for a single reader).

**Contexts:** Using a context to update the readers about the certifcate variable being changed.

### Results
