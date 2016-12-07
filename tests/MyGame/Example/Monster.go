// automatically generated by the FlatBuffers compiler, do not modify

package Example

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// an example documentation comment: monster object
type Monster struct {
	_tab flatbuffers.Table
}

func GetRootAsMonster(buf []byte, offset flatbuffers.UOffsetT) *Monster {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Monster{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Monster) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Monster) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Monster) Pos(obj *Vec3) *Vec3 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(Vec3)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Monster) Mana() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 150
}

func (rcv *Monster) MutateMana(n int16) bool {
	return rcv._tab.MutateInt16Slot(6, n)
}

func (rcv *Monster) Hp() int16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt16(o + rcv._tab.Pos)
	}
	return 100
}

func (rcv *Monster) MutateHp(n int16) bool {
	return rcv._tab.MutateInt16Slot(8, n)
}

func (rcv *Monster) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Monster) Inventory(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Monster) InventoryLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) InventoryBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Monster) Color() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 8
}

func (rcv *Monster) MutateColor(n int8) bool {
	return rcv._tab.MutateInt8Slot(16, n)
}

func (rcv *Monster) TestType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) MutateTestType(n byte) bool {
	return rcv._tab.MutateByteSlot(18, n)
}

func (rcv *Monster) Test(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func (rcv *Monster) Test4(obj *Test, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Monster) Test4Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) Testarrayofstring(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Monster) TestarrayofstringLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(24))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// an example documentation comment: this will end up in the generated code
/// multiline too
func (rcv *Monster) Testarrayoftables(obj *Monster, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Monster) TestarrayoftablesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(26))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

/// an example documentation comment: this will end up in the generated code
/// multiline too
func (rcv *Monster) Enemy(obj *Monster) *Monster {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(28))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Monster)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Monster) Testnestedflatbuffer(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Monster) TestnestedflatbufferLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) TestnestedflatbufferBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(30))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Monster) Testempty(obj *Stat) *Stat {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(32))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Stat)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *Monster) Testbool() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(34))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) MutateTestbool(n byte) bool {
	return rcv._tab.MutateByteSlot(34, n)
}

func (rcv *Monster) Testhashs32Fnv1() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(36))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) MutateTesthashs32Fnv1(n int32) bool {
	return rcv._tab.MutateInt32Slot(36, n)
}

func (rcv *Monster) Testhashu32Fnv1() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(38))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) MutateTesthashu32Fnv1(n uint32) bool {
	return rcv._tab.MutateUint32Slot(38, n)
}

func (rcv *Monster) Testhashs64Fnv1() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(40))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) MutateTesthashs64Fnv1(n int64) bool {
	return rcv._tab.MutateInt64Slot(40, n)
}

func (rcv *Monster) Testhashu64Fnv1() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(42))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) MutateTesthashu64Fnv1(n uint64) bool {
	return rcv._tab.MutateUint64Slot(42, n)
}

func (rcv *Monster) Testhashs32Fnv1a() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(44))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) MutateTesthashs32Fnv1a(n int32) bool {
	return rcv._tab.MutateInt32Slot(44, n)
}

func (rcv *Monster) Testhashu32Fnv1a() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(46))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) MutateTesthashu32Fnv1a(n uint32) bool {
	return rcv._tab.MutateUint32Slot(46, n)
}

func (rcv *Monster) Testhashs64Fnv1a() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(48))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) MutateTesthashs64Fnv1a(n int64) bool {
	return rcv._tab.MutateInt64Slot(48, n)
}

func (rcv *Monster) Testhashu64Fnv1a() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(50))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Monster) MutateTesthashu64Fnv1a(n uint64) bool {
	return rcv._tab.MutateUint64Slot(50, n)
}

func (rcv *Monster) Testarrayofbools(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *Monster) TestarrayofboolsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(52))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *Monster) Testf() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(54))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 3.14159
}

func (rcv *Monster) MutateTestf(n float32) bool {
	return rcv._tab.MutateFloat32Slot(54, n)
}

func (rcv *Monster) Testf2() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(56))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 3.0
}

func (rcv *Monster) MutateTestf2(n float32) bool {
	return rcv._tab.MutateFloat32Slot(56, n)
}

func (rcv *Monster) Testf3() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(58))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *Monster) MutateTestf3(n float32) bool {
	return rcv._tab.MutateFloat32Slot(58, n)
}

func (rcv *Monster) Testarrayofstring2(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(60))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *Monster) Testarrayofstring2Length() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(60))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func MonsterStart(builder *flatbuffers.Builder) {
	builder.StartObject(29)
}
func MonsterAddPos(builder *flatbuffers.Builder, pos flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(pos), 0)
}
func MonsterAddMana(builder *flatbuffers.Builder, mana int16) {
	builder.PrependInt16Slot(1, mana, 150)
}
func MonsterAddHp(builder *flatbuffers.Builder, hp int16) {
	builder.PrependInt16Slot(2, hp, 100)
}
func MonsterAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(name), 0)
}
func MonsterAddInventory(builder *flatbuffers.Builder, inventory flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(inventory), 0)
}
func MonsterStartInventoryVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func MonsterAddColor(builder *flatbuffers.Builder, color int8) {
	builder.PrependInt8Slot(6, color, 8)
}
func MonsterAddTestType(builder *flatbuffers.Builder, testType byte) {
	builder.PrependByteSlot(7, testType, 0)
}
func MonsterAddTest(builder *flatbuffers.Builder, test flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(test), 0)
}
func MonsterAddTest4(builder *flatbuffers.Builder, test4 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(test4), 0)
}
func MonsterStartTest4Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 2)
}
func MonsterAddTestarrayofstring(builder *flatbuffers.Builder, testarrayofstring flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(10, flatbuffers.UOffsetT(testarrayofstring), 0)
}
func MonsterStartTestarrayofstringVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MonsterAddTestarrayoftables(builder *flatbuffers.Builder, testarrayoftables flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(11, flatbuffers.UOffsetT(testarrayoftables), 0)
}
func MonsterStartTestarrayoftablesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MonsterAddEnemy(builder *flatbuffers.Builder, enemy flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(12, flatbuffers.UOffsetT(enemy), 0)
}
func MonsterAddTestnestedflatbuffer(builder *flatbuffers.Builder, testnestedflatbuffer flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(13, flatbuffers.UOffsetT(testnestedflatbuffer), 0)
}
func MonsterStartTestnestedflatbufferVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func MonsterAddTestempty(builder *flatbuffers.Builder, testempty flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(14, flatbuffers.UOffsetT(testempty), 0)
}
func MonsterAddTestbool(builder *flatbuffers.Builder, testbool byte) {
	builder.PrependByteSlot(15, testbool, 0)
}
func MonsterAddTesthashs32Fnv1(builder *flatbuffers.Builder, testhashs32Fnv1 int32) {
	builder.PrependInt32Slot(16, testhashs32Fnv1, 0)
}
func MonsterAddTesthashu32Fnv1(builder *flatbuffers.Builder, testhashu32Fnv1 uint32) {
	builder.PrependUint32Slot(17, testhashu32Fnv1, 0)
}
func MonsterAddTesthashs64Fnv1(builder *flatbuffers.Builder, testhashs64Fnv1 int64) {
	builder.PrependInt64Slot(18, testhashs64Fnv1, 0)
}
func MonsterAddTesthashu64Fnv1(builder *flatbuffers.Builder, testhashu64Fnv1 uint64) {
	builder.PrependUint64Slot(19, testhashu64Fnv1, 0)
}
func MonsterAddTesthashs32Fnv1a(builder *flatbuffers.Builder, testhashs32Fnv1a int32) {
	builder.PrependInt32Slot(20, testhashs32Fnv1a, 0)
}
func MonsterAddTesthashu32Fnv1a(builder *flatbuffers.Builder, testhashu32Fnv1a uint32) {
	builder.PrependUint32Slot(21, testhashu32Fnv1a, 0)
}
func MonsterAddTesthashs64Fnv1a(builder *flatbuffers.Builder, testhashs64Fnv1a int64) {
	builder.PrependInt64Slot(22, testhashs64Fnv1a, 0)
}
func MonsterAddTesthashu64Fnv1a(builder *flatbuffers.Builder, testhashu64Fnv1a uint64) {
	builder.PrependUint64Slot(23, testhashu64Fnv1a, 0)
}
func MonsterAddTestarrayofbools(builder *flatbuffers.Builder, testarrayofbools flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(24, flatbuffers.UOffsetT(testarrayofbools), 0)
}
func MonsterStartTestarrayofboolsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func MonsterAddTestf(builder *flatbuffers.Builder, testf float32) {
	builder.PrependFloat32Slot(25, testf, 3.14159)
}
func MonsterAddTestf2(builder *flatbuffers.Builder, testf2 float32) {
	builder.PrependFloat32Slot(26, testf2, 3.0)
}
func MonsterAddTestf3(builder *flatbuffers.Builder, testf3 float32) {
	builder.PrependFloat32Slot(27, testf3, 0.0)
}
func MonsterAddTestarrayofstring2(builder *flatbuffers.Builder, testarrayofstring2 flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(28, flatbuffers.UOffsetT(testarrayofstring2), 0)
}
func MonsterStartTestarrayofstring2Vector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func MonsterEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
