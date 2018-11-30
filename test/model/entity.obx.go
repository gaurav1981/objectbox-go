// Code generated by ObjectBox; DO NOT EDIT.

package model

import (
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type entity_EntityInfo struct {
	Id  objectbox.TypeId
	Uid uint64
}

var EntityBinding = entity_EntityInfo{
	Id:  1,
	Uid: 3022148985475790732,
}

var Entity_ = struct {
	Id         *objectbox.PropertyUint64
	Int        *objectbox.PropertyInt
	Int8       *objectbox.PropertyInt8
	Int16      *objectbox.PropertyInt16
	Int32      *objectbox.PropertyInt32
	Int64      *objectbox.PropertyInt64
	Uint       *objectbox.PropertyUint
	Uint8      *objectbox.PropertyUint8
	Uint16     *objectbox.PropertyUint16
	Uint32     *objectbox.PropertyUint32
	Uint64     *objectbox.PropertyUint64
	Bool       *objectbox.PropertyBool
	String     *objectbox.PropertyString
	Byte       *objectbox.PropertyByte
	ByteVector *objectbox.PropertyByteVector
	Rune       *objectbox.PropertyRune
	Float32    *objectbox.PropertyFloat32
	Float64    *objectbox.PropertyFloat64
	Date       *objectbox.PropertyInt64
}{
	Id: &objectbox.PropertyUint64{
		Property: &objectbox.Property{
			Id: 1,
		},
	},
	Int: &objectbox.PropertyInt{
		Property: &objectbox.Property{
			Id: 2,
		},
	},
	Int8: &objectbox.PropertyInt8{
		Property: &objectbox.Property{
			Id: 3,
		},
	},
	Int16: &objectbox.PropertyInt16{
		Property: &objectbox.Property{
			Id: 4,
		},
	},
	Int32: &objectbox.PropertyInt32{
		Property: &objectbox.Property{
			Id: 5,
		},
	},
	Int64: &objectbox.PropertyInt64{
		Property: &objectbox.Property{
			Id: 6,
		},
	},
	Uint: &objectbox.PropertyUint{
		Property: &objectbox.Property{
			Id: 7,
		},
	},
	Uint8: &objectbox.PropertyUint8{
		Property: &objectbox.Property{
			Id: 8,
		},
	},
	Uint16: &objectbox.PropertyUint16{
		Property: &objectbox.Property{
			Id: 9,
		},
	},
	Uint32: &objectbox.PropertyUint32{
		Property: &objectbox.Property{
			Id: 10,
		},
	},
	Uint64: &objectbox.PropertyUint64{
		Property: &objectbox.Property{
			Id: 11,
		},
	},
	Bool: &objectbox.PropertyBool{
		Property: &objectbox.Property{
			Id: 12,
		},
	},
	String: &objectbox.PropertyString{
		Property: &objectbox.Property{
			Id: 13,
		},
	},
	Byte: &objectbox.PropertyByte{
		Property: &objectbox.Property{
			Id: 14,
		},
	},
	ByteVector: &objectbox.PropertyByteVector{
		Property: &objectbox.Property{
			Id: 15,
		},
	},
	Rune: &objectbox.PropertyRune{
		Property: &objectbox.Property{
			Id: 16,
		},
	},
	Float32: &objectbox.PropertyFloat32{
		Property: &objectbox.Property{
			Id: 17,
		},
	},
	Float64: &objectbox.PropertyFloat64{
		Property: &objectbox.Property{
			Id: 18,
		},
	},
	Date: &objectbox.PropertyInt64{
		Property: &objectbox.Property{
			Id: 19,
		},
	},
}

func (entity_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Entity", 1, 3022148985475790732)
	model.Property("Id", objectbox.PropertyType_Long, 1, 1213346202559552829)
	model.PropertyFlags(objectbox.PropertyFlags_ID)
	model.Property("Int", objectbox.PropertyType_Long, 2, 6609825840127351046)
	model.Property("Int8", objectbox.PropertyType_Byte, 3, 741904540265547276)
	model.Property("Int16", objectbox.PropertyType_Short, 4, 2102961483425256790)
	model.Property("Int32", objectbox.PropertyType_Int, 5, 5086065890743931723)
	model.Property("Int64", objectbox.PropertyType_Long, 6, 7993850425898343586)
	model.Property("Uint", objectbox.PropertyType_Long, 7, 993618338451248101)
	model.Property("Uint8", objectbox.PropertyType_Byte, 8, 8794162191867215541)
	model.Property("Uint16", objectbox.PropertyType_Short, 9, 8820932096807845950)
	model.Property("Uint32", objectbox.PropertyType_Int, 10, 6060435196424462358)
	model.Property("Uint64", objectbox.PropertyType_Long, 11, 6159011311237949479)
	model.Property("Bool", objectbox.PropertyType_Bool, 12, 13717766026420552)
	model.Property("String", objectbox.PropertyType_String, 13, 3525810560076343996)
	model.Property("Byte", objectbox.PropertyType_Byte, 14, 4035373893984224671)
	model.Property("ByteVector", objectbox.PropertyType_ByteVector, 15, 1294888641203478533)
	model.Property("Rune", objectbox.PropertyType_Int, 16, 445652208596094853)
	model.Property("Float32", objectbox.PropertyType_Float, 17, 2321055489159952634)
	model.Property("Float64", objectbox.PropertyType_Double, 18, 681625187526498317)
	model.Property("Date", objectbox.PropertyType_Date, 19, 2927532418453906842)
	model.EntityLastPropertyId(19, 2927532418453906842)
}

func (entity_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*Entity).Id, nil
}

func (entity_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Entity).Id = id
	return nil
}

func (entity_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) {
	obj := object.(*Entity)
	var offsetString = fbutils.CreateStringOffset(fbb, obj.String)
	var offsetByteVector = fbutils.CreateByteVectorOffset(fbb, obj.ByteVector)

	// build the FlatBuffers object
	fbb.StartObject(19)
	fbb.PrependUint64Slot(0, id, 0)
	fbb.PrependInt64Slot(1, int64(obj.Int), 0)
	fbb.PrependInt8Slot(2, obj.Int8, 0)
	fbb.PrependInt16Slot(3, obj.Int16, 0)
	fbb.PrependInt32Slot(4, obj.Int32, 0)
	fbb.PrependInt64Slot(5, obj.Int64, 0)
	fbb.PrependUint64Slot(6, uint64(obj.Uint), 0)
	fbb.PrependUint8Slot(7, obj.Uint8, 0)
	fbb.PrependUint16Slot(8, obj.Uint16, 0)
	fbb.PrependUint32Slot(9, obj.Uint32, 0)
	fbb.PrependUint64Slot(10, obj.Uint64, 0)
	fbb.PrependBoolSlot(11, obj.Bool, false)
	fbb.PrependUOffsetTSlot(12, offsetString, 0)
	fbb.PrependByteSlot(13, obj.Byte, 0)
	fbb.PrependUOffsetTSlot(14, offsetByteVector, 0)
	fbb.PrependInt32Slot(15, obj.Rune, 0)
	fbb.PrependFloat32Slot(16, obj.Float32, 0)
	fbb.PrependFloat64Slot(17, obj.Float64, 0)
	fbb.PrependInt64Slot(18, obj.Date, 0)
}

func (entity_EntityInfo) ToObject(bytes []byte) interface{} {
	table := &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	return &Entity{
		Id:         table.GetUint64Slot(4, 0),
		Int:        int(table.GetUint64Slot(6, 0)),
		Int8:       table.GetInt8Slot(8, 0),
		Int16:      table.GetInt16Slot(10, 0),
		Int32:      table.GetInt32Slot(12, 0),
		Int64:      table.GetInt64Slot(14, 0),
		Uint:       uint(table.GetUint64Slot(16, 0)),
		Uint8:      table.GetUint8Slot(18, 0),
		Uint16:     table.GetUint16Slot(20, 0),
		Uint32:     table.GetUint32Slot(22, 0),
		Uint64:     table.GetUint64Slot(24, 0),
		Bool:       table.GetBoolSlot(26, false),
		String:     fbutils.GetStringSlot(table, 28),
		Byte:       table.GetByteSlot(30, 0),
		ByteVector: fbutils.GetByteVectorSlot(table, 32),
		Rune:       rune(table.GetInt32Slot(34, 0)),
		Float32:    table.GetFloat32Slot(36, 0),
		Float64:    table.GetFloat64Slot(38, 0),
		Date:       table.GetInt64Slot(40, 0),
	}
}

func (entity_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Entity, 0, capacity)
}

func (entity_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]*Entity), object.(*Entity))
}

type EntityBox struct {
	*objectbox.Box
}

func BoxForEntity(ob *objectbox.ObjectBox) *EntityBox {
	return &EntityBox{
		Box: ob.Box(1),
	}
}

func (box *EntityBox) Put(object *Entity) (uint64, error) {
	return box.Box.Put(object)
}

func (box *EntityBox) PutAsync(object *Entity) (uint64, error) {
	return box.Box.PutAsync(object)
}

func (box *EntityBox) PutAll(objects []*Entity) ([]uint64, error) {
	return box.Box.PutAll(objects)
}

func (box *EntityBox) Get(id uint64) (*Entity, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Entity), nil
}

func (box *EntityBox) GetAll() ([]*Entity, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Entity), nil
}

func (box *EntityBox) Remove(object *Entity) (err error) {
	return box.Box.Remove(object.Id)
}

func (box *EntityBox) Query(conditions ...objectbox.Condition) *EntityQuery {
	return &EntityQuery{
		box.Box.Query(conditions...),
	}
}

type EntityQuery struct {
	*objectbox.Query
}

func (query *EntityQuery) Find() ([]*Entity, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Entity), nil
}
