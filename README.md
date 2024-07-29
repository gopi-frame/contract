# Contract

Package contract provides interfaces and types for Gopi frame.

## Installation

```shell
go get -u -v github.com/gopi-frame/contract
```

## Import
```go
import "github.com/gopi-frame/contract"
```

## Interfaces

### Arrayable

```go
type Numbers struct {
	numbers []int
}

func (n *Numbers) ToArray() []int {
	return n.numbers
}
```

### Comparator

```go
type User struct {
	ID   int
	Name string
}

func (u *User) Compare(other *User) int {
	if u.ID > other.ID {
		return 1
	}
	if u.ID < other.ID {
		return -1
	}
	if u.Name > other.Name {
		return 1
	}
	if u.Name < other.Name {
		return -1
	}
	return 0
}
```

### Countable

```go
type Numbers struct {
	numbers []int
}

func (n *Numbers) Count() int64 {
	return int64(len(n.numbers))
}
```

### Delayable

```go
type Delay[int] struct {
	value int
	until time.Time
}

func (d *Delay[int]) Value() int { return d.value }

func (d *Delay[int]) Until() time.Time { return d.until }

func (d *Delay[int]) MarshalJSON(until time.Time) error {
    var jsonObject struct {
        Value int `json:"value"`
		Until time.Time `json:"until"`
	}
	jsonObject.Value = d.value
	jsonObject.Until = d.until
	return json.Marshal(jsonObject)
}

func (d *Delay[int]) UnmarshalJSON(b []byte) error {
	var jsonObject struct {
		Value int `json:"value"`
		Until time.Time `json:"until"`
    }
	if err := json.Unmarshal(b, &jsonObject); err!= nil {
		return err
	}
	d.value = jsonObject.Value
	d.until = jsonObject.Until
	return nil
}
```

### Jsonable

```go
type User struct {
	ID   int
	Name string
}

func (u *User) ToJSON() ([]byte, error) {
    return json.Marshal(u)
}
```

### Mappable

```go
type User struct {
	ID   int
    Name string
}

func (u *User) Map() map[string]interface{} {
	return map[string]interface{}{
		"id":   u.ID,
		"name": u.Name,
    }
}
```


### List

See [collection.List](https://github.com/gopi-frame/collection/tree/main/list) for more information.

### Map

See [collection.Map](https://github.com/gopi-frame/collection/tree/main/kv) for more information.

### Queue

See [collection.Queue](https://github.com/gopi-frame/collection/tree/main/queue) for more information.

### BlockingQueue

See [collection.Queue](https://github.com/gopi-frame/collection/tree/main/queue) for more information.

### DelayedQueue

See [collection.Queue](https://github.com/gopi-frame/collection/tree/main/queue) for more information.

### Set

See [collection.Set](https://github.com/gopi-frame/collection/tree/main/set) for more information.

### Traversable

```go
type Numbers struct {
	numbers []int
}

func (n *Numbers) Each(fn func(index int, value int) bool) {
	for i, v := range n.numbers {
		if !fn(i, v) {
			break
		}
	}
}
```

### Sortable

```go
type Numbers struct {
    numbers []int
}

func (n *Numbers) Sort(fn func(a, b int) int) {
	for i := 0; i < len(n.numbers); i++ {
		for j := i + 1; j < len(n.numbers); j++ {
			if fn(n.numbers[i], n.numbers[j]) > 0 {
				n.numbers[i], n.numbers[j] = n.numbers[j], n.numbers[i]
			}
		}
	}
}
```

### Stringable

```go
type User struct {
	ID   int
    Name string
}

func (u *User) String() string {
    return fmt.Sprintf("%v %v", u.ID, u.Name)
}
```

### Reversible

```go
type Numbers struct {
	numbers []int
}

func (n *Numbers) Reverse() {
	for i, j := 0, len(n.numbers)-1; i < j; i, j = i+1, j-1 {
		n.numbers[i], n.numbers[j] = n.numbers[j], n.numbers[i]
	}
}
```

## Sub-packages
