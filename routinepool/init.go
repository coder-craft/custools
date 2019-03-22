package routinepool

var (
	MaxWorker = 100
	MaxQueue  = 10000
)

type Delivery struct {
	workerPool chan chan work
	jobQueue   chan work
	maxWorker  int
}

var routinePool *Delivery

func NewDelivery(maxWorker int) *Delivery {
	return &Delivery{maxWorker: maxWorker, workerPool: make(chan chan work, maxWorker)}
}
func (d *Delivery) run() {
	for i := 0; i < d.maxWorker; i++ {
		worker := NewWorker(d.workerPool)
		worker.Working()
	}
	go d.delivery()
}

func (d *Delivery) delivery() {
	for {
		select {
		case newJob := <-d.jobQueue:
			go func(j work) {
				jobChannel := <-d.workerPool
				jobChannel <- j
			}(newJob)
		}
	}
}

func init() {
	routinePool = NewDelivery(MaxWorker)
	routinePool.run()
}
