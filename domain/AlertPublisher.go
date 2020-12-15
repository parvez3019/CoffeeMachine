package domain

type Publisher struct {
	subscribers []Subscriber
}

func (o *Publisher) AddSubscriber(obs Subscriber) {
	if o.subscribers == nil {
		o.subscribers = []Subscriber{obs}
	}
	o.subscribers = append(o.subscribers, obs)
}

func (o *Publisher) NotifyAll(ingredient Ingredient) {
	for _, ob := range o.subscribers {
		ob.notify(ingredient)
	}
}
