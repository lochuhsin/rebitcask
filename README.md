# Re-Bitcask

#### A redemption journey from a backend software engineer

---
### Description:
This is a redemption journey of a backend engineer.<br> 
Trying to re-implement
bitcask key-value storage database mechanism to understand how
database underlying works.<br>
In short, this is a project re-implementing Key/Value storage with log style
storage mechanism.

---
### Feature
1. Basic Key/Value Storage mechanism
   1. Get
   2. Set
   3. Delete
2. The database could be tuned using environment variable
    ```text
    LOG_FOLDER_PATH="./log/"
   
    TOMBSTONE="abcdefghijklmnopqrstuvwxyz"
   
    MEMORY_KEY_COUNT_LIMIT=50000
   
    FILE_BYTE_LIMIT=30000
   
    SEGMENT_FILE_COUNT_LIMIT=3
    ```
3. Currently, using AVL-Tree as memory table (Could be change to Binary Search Tree). There will more types
   of tree based indexing mechanism.

4. Could be used as standalone server.

5. Could be used as library.

---
### Future Panning
- [x]  Basic Get, Set, Delete mechanism  
- [x]  Implement vanilla hashtable key value storage
- [x]  Implement Segment storage,
- [x]  Implement Seek file header method (file.sync or writer.flush)
- [x]  Implement Binary Search Tree memory style
- [x]  Re-implement Get, Set storage method
- [x]  Implement SSTable (**Sorted String Table**) (Last one is implement compress function) （Finished)
- [x]  Implement AVL Tree
- [ ]  Supporting backend compression periodically
- [ ]  Implement range based key query
- [ ]  Optimize compression mechanism
- [ ]  Implement Red-Black Tree
- [ ]  Implement BloomFilter and cache mechanism for Read
- [ ]  Implement B-Tree
- [ ]  Support more generic types 

---
### How to ....use?
1. Used as library:

`git clone project`

`go mod tidy`

```go
package foo
import "rebitcask/src"

src.Get(key)

src.Set(key, value)

src.Delete(key)
```

2. Used as Server:
```bash
cd rebitcask/cli

go mod tidy

create rebitcask.env or using default

go build .

go run ./cli
```

