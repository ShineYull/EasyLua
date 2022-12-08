// @Title	api_push.go
// @Description	External call api for lua push operation.
// @Author	ShineYu 2022/12/08 23:52:00
// @Update	ShineYu 2022/12/08 23:52:00
package state

// @title	PushNil()
// @description	push nil type value in to the stack. [-0, +1, –]
// @auth	ShineYu	2022/12/08 23:52:00
func (self *luaState) PushNil() {
	self.stack.push(nil)
}

// @title	PushBoolean(b bool)
// @description	push bool type value in to the stack. [-0, +1, –]
// @auth	ShineYu	2022/12/08 23:52:00
func (self *luaState) PushBoolean(b bool) {
	self.stack.push(b)
}

// @title	PushInteger(n int64)
// @description	push int64 type value in to the stack. [-0, +1, –]
// @auth	ShineYu	2022/12/08 23:52:00
func (self *luaState) PushInteger(n int64) {
	self.stack.push(n)
}

// @title	PushNumber(n float64)
// @description	push float64 type value in to the stack. [-0, +1, –]
// @auth	ShineYu	2022/12/08 23:52:00
func (self *luaState) PushNumber(n float64) {
	self.stack.push(n)
}

// @title	PushString(s string)
// @description	push string type value in to the stack. [-0, +1, m]
// @auth	ShineYu	2022/12/08 23:52:00
func (self *luaState) PushString(s string) {
	self.stack.push(s)
}
