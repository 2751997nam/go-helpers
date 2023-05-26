package event

type RegisterFunc func()

type EventRegister struct {
	registers []RegisterFunc
}

func (r *EventRegister) NewRegister(rf RegisterFunc) {
	r.registers = append(r.registers, rf)
}

func (r *EventRegister) Register() {
	for _, r := range r.registers {
		r()
	}
}
