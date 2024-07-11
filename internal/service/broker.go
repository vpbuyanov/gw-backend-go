package service

func RegisterBroker(f func()) {
	go f()
}
