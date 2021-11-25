package aqueue

type Aqueue struct {
	total   int
	tasks   []func()
	tickets chan int
	running bool
}

func NewAqueue(total int) *Aqueue {
	return &Aqueue{
		total:   total,
		tickets: make(chan int, total),
	}
}

func (a *Aqueue) Add(task func()) {
	if len(a.tasks) < a.total {
		a.tasks = append(a.tasks, task)
	}
	a.run()
}

func (a *Aqueue) Wait() {
	for i := 0; i < a.total; i++ {
		<-a.tickets
	}
}

func (a *Aqueue) run() {
	if a.running {
		return
	}
	if len(a.tasks) == 0 {
		go a.run()
		return
	}
	task := a.tasks[0]
	a.tasks = a.tasks[1:]
	a.running = true
	go func(t func(), a *Aqueue) {
		t()

		a.tickets <- 1
		a.running = false
		a.run()
	}(task, a)
}
