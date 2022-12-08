// @Title	api_stack.go
// @Description	Basic stack manipulation method.
// @Author	ShineYu 2022/12/08 23:52:00
// @Update	ShineYu 2022/12/08 23:52:00
package state

// @title	GetTop()
// @description	return stack top index. [-0, +0, –]
// @auth	ShineYu	2022/12/08 23:52:00
// @return	int	stack top index.
func (self *luaState) GetTop() int {
	return self.stack.top
}

// @title	AbsIndex(idx int)
// @description	Convert index to absolute index. [-0, +0, –]
// @auth	ShineYu	2022/12/08 23:52:00
// @return	int	absolute index.
func (self *luaState) AbsIndex(idx int) int {
	return self.stack.absIndex(idx)
}

// @title	CheckStack(n int)
// @description	The capacity of the lua stack will not increase automatically,
//				and the api user must call the CheckStack() method to check the
//				remaining space of the stack when necessary. [-0, +0, –]
// @auth	ShineYu	2022/12/08 23:52:00
// @return	bool	stack status.
func (self *luaState) CheckStack(n int) bool {
	self.stack.check(n)
	return true // never fails
}

// @title	Pop(n int)
// @description	Pop n values from the top of the stack. [-n, +0, –]
// @auth	ShineYu	2022/12/08 23:52:00
func (self *luaState) Pop(n int) {
	for i := 0; i < n; i++ {
		self.stack.pop()
	}
}

// @title	Copy(fromIdx, toIdx int)
// @description	Copies a value from one location to another. [-0, +0, –]
// @auth	ShineYu	2022/12/08 23:52:00
func (self *luaState) Copy(fromIdx, toIdx int) {
	val := self.stack.get(fromIdx)
	self.stack.set(toIdx, val)
}

// @title	PushValue(idx int)
// @description	Push the value at the specified index onto
//				the top of the stack. [-0, +1, –]
// @auth	ShineYu	2022/12/08 23:52:00
func (self *luaState) PushValue(idx int) {
	val := self.stack.get(idx)
	self.stack.push(val)
}

// @title	Replace(idx int)
// @description	Pop the top value off the stack and write it
//				to the specified location. [-1, +0, –]
// @auth	ShineYu	2022/12/08 23:52:00
func (self *luaState) Replace(idx int) {
	val := self.stack.pop()
	self.stack.set(idx, val)
}

// @title	Insert(idx int)
// @description	Pop the top value off the stack and insert it
//				into the specified position. [-1, +1, –]
// @auth	ShineYu	2022/12/08 23:52:00
func (self *luaState) Insert(idx int) {
	self.Rotate(idx, 1)
}

// @title	Remove(idx int)
// @description	Deletes the value at the specified index and
//				shifts all values above it down one position. [-1, +0, –]
// @auth	ShineYu	2022/12/08 23:52:00
func (self *luaState) Remove(idx int) {
	self.Rotate(idx, -1)
	self.Pop(1)
}

// @title	Rotate(idx, n int)
// @description	Rotate the value in the [idx, top] index interval
//				towards the top of the stack by n positions. If n is
//				negative, the actual effect is to rotate towards the
//				bottom of the stack. [-0, +0, –]
// @auth	ShineYu	2022/12/08 23:52:00
func (self *luaState) Rotate(idx, n int) {
	t := self.stack.top - 1           /* end of stack segment being rotated */
	p := self.stack.absIndex(idx) - 1 /* start of segment */
	var m int                         /* end of prefix */
	if n >= 0 {
		m = t - n
	} else {
		m = p - n - 1
	}
	self.stack.reverse(p, m)   /* reverse the prefix with length 'n' */
	self.stack.reverse(m+1, t) /* reverse the suffix */
	self.stack.reverse(p, t)   /* reverse the entire segment */
}

// @title	Rotate(idx, n int)
// @description	Set the top index of the stack to the specified value.
//				If the specified value is less than the current stack top index,
//				the effect is equivalent to a pop operation (a specified value
//				of 0 is equivalent to clearing the stack). [-?, +?, –]
// @auth	ShineYu	2022/12/08 23:52:00
func (self *luaState) SetTop(idx int) {
	newTop := self.stack.absIndex(idx)
	if newTop < 0 {
		panic("stack underflow!")
	}

	n := self.stack.top - newTop
	if n > 0 {
		for i := 0; i < n; i++ {
			self.stack.pop()
		}
	} else if n < 0 {
		for i := 0; i > n; i-- {
			self.stack.push(nil)
		}
	}
}
