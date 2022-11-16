package state

// [-0, +1, –]
func (self *luaState) PushNil() {
	self.stack.push(nil)
}

// [-0, +1, –]
func (self *luaState) PushBoolean(b bool) {
	self.stack.push(b)
}

// [-0, +1, –]
func (self *luaState) PushInteger(n int64) {
	self.stack.push(n)
}

// [-0, +1, –]
func (self *luaState) PushNumber(n float64) {
	self.stack.push(n)
}

// [-0, +1, m]
func (self *luaState) PushString(s string) {
	self.stack.push(s)
}
