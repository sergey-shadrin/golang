package safeslice

type UpdateFunc func(val interface{}) interface{}

type SafeSlice interface {
	Append(interface{})
	At(int) interface{}
	Close() []interface{}
	Delete(int)
	Len() int
	Update(int, UpdateFunc)
}

type action int

const (
	добавить action = iota
	at
	closeCommand
	deleteCommand
	length
	update
)

type command struct {
	action  action
	value   interface{}
	index   int
	updater UpdateFunc
	result  chan interface{}
	data    chan []interface{}
}

type safeSlice chan command

func (ss safeSlice) Append(val interface{}) {
	ss <- command{action: добавить, value: val}
}

func (ss safeSlice) At(i int) interface{} {
	reply := make(chan interface{})
	ss <- command{action: at, index: i, result: reply}
	return <-reply
}

func (ss safeSlice) Close() []interface{} {
	data := make(chan []interface{})
	ss <- command{action: closeCommand, data: data}
	return <-data
}

func (ss safeSlice) Delete(i int) {
	ss <- command{action: deleteCommand, index: i}
}

func (ss safeSlice) Len() int {
	reply := make(chan interface{})
	ss <- command{action: length, result: reply}
	return (<-reply).(int)
}

func (ss safeSlice) Update(i int, updateFunc UpdateFunc) {
	ss <- command{action: update, index: i, updater: updateFunc}
}

func (ss safeSlice) run() {
	go func() {
		data := make([]interface{}, 0, 1000)
		for {
			command := <-ss
			switch command.action {
			case добавить:
				data = append(data, command.value)
			case at:

				reply := command.result
				if i := command.index; i >= 0 && i < len(data) {
					reply <- data[i]
				} else {
					reply <- nil
				}
			case deleteCommand:
				if i := command.index; i >= 0 && i < len(data) {
					data = append(data[:i], data[(i+1):]...)
				}
			case length:
				command.result <- len(data)
			case update:
				if i := command.index; i >= 0 && i < len(data) {
					data[i] = command.updater(data[i])
				}
			case closeCommand:
				command.data <- data
				close(ss)
				return
			}
		}

	}()
}

func New() SafeSlice {
	ss := make(safeSlice)
	ss.run()
	return ss
}
