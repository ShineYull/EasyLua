// @Title	api_access.go
// @Description
//
//	The external call API for the acquisition
//	operation of the lua stack value.
//
// @Author	ShineYu 2022/12/08 23:52:00
// @Update	ShineYu 2022/12/08 23:52:00
package state

import (
	"fmt"
	. "luago/api"
)

// @title	TypeName(tp LuaType)
// @description	Lua type is converted to the corresponding string representation. [-0, +0, –]
// @auth	ShineYu	2022/12/08 23:52:00
// @param	tp	LuaType	The lua Type.
// @return		string	String of type.
func (self *luaState) TypeName(tp LuaType) string {
	switch tp {
	case LUA_TNONE:
		return "no value"
	case LUA_TNIL:
		return "nil"
	case LUA_TBOOLEAN:
		return "boolean"
	case LUA_TNUMBER:
		return "number"
	case LUA_TSTRING:
		return "string"
	case LUA_TTABLE:
		return "table"
	case LUA_TFUNCTION:
		return "function"
	case LUA_TTHREAD:
		return "thread"
	default:
		return "userdata"
	}
}

// @title	Type(idx int)
// @description	Returns the type of the value according to the index. [-0, +0, –]
// @auth	ShineYu	2022/12/08 23:52:00
// @param	idx	int		The index of stack.
// @return		LuaType	The type of vale.
func (self *luaState) Type(idx int) LuaType {
	if self.stack.isValid(idx) {
		val := self.stack.get(idx)
		return typeOf(val)
	}
	return LUA_TNONE
}

// [-0, +0, –]
func (self *luaState) IsNone(idx int) bool {
	return self.Type(idx) == LUA_TNONE
}

// [-0, +0, –]
func (self *luaState) IsNil(idx int) bool {
	return self.Type(idx) == LUA_TNIL
}

// [-0, +0, –]
func (self *luaState) IsNoneOrNil(idx int) bool {
	return self.Type(idx) <= LUA_TNIL
}

// [-0, +0, –]
func (self *luaState) IsBoolean(idx int) bool {
	return self.Type(idx) == LUA_TBOOLEAN
}

// [-0, +0, –]
func (self *luaState) IsTable(idx int) bool {
	return self.Type(idx) == LUA_TTABLE
}

// [-0, +0, –]
func (self *luaState) IsFunction(idx int) bool {
	return self.Type(idx) == LUA_TFUNCTION
}

// [-0, +0, –]
func (self *luaState) IsThread(idx int) bool {
	return self.Type(idx) == LUA_TTHREAD
}

// [-0, +0, –]
func (self *luaState) IsString(idx int) bool {
	t := self.Type(idx)
	return t == LUA_TSTRING || t == LUA_TNUMBER
}

// [-0, +0, –]
func (self *luaState) IsNumber(idx int) bool {
	_, ok := self.ToNumberX(idx)
	return ok
}

// [-0, +0, –]
func (self *luaState) IsInteger(idx int) bool {
	val := self.stack.get(idx)
	_, ok := val.(int64)
	return ok
}

// [-0, +0, –]
func (self *luaState) ToBoolean(idx int) bool {
	val := self.stack.get(idx)
	return convertToBoolean(val)
}

// [-0, +0, –]
func (self *luaState) ToInteger(idx int) int64 {
	i, _ := self.ToIntegerX(idx)
	return i
}

// [-0, +0, –]
func (self *luaState) ToIntegerX(idx int) (int64, bool) {
	val := self.stack.get(idx)
	i, ok := val.(int64)
	return i, ok
}

// [-0, +0, –]
func (self *luaState) ToNumber(idx int) float64 {
	n, _ := self.ToNumberX(idx)
	return n
}

// [-0, +0, –]
func (self *luaState) ToNumberX(idx int) (float64, bool) {
	val := self.stack.get(idx)
	switch x := val.(type) {
	case float64:
		return x, true
	case int64:
		return float64(x), true
	default:
		return 0, false
	}
}

// [-0, +0, m]
func (self *luaState) ToString(idx int) string {
	s, _ := self.ToStringX(idx)
	return s
}

func (self *luaState) ToStringX(idx int) (string, bool) {
	val := self.stack.get(idx)

	switch x := val.(type) {
	case string:
		return x, true
	case int64, float64:
		s := fmt.Sprintf("%v", x) // todo
		self.stack.set(idx, s)
		return s, true
	default:
		return "", false
	}
}
