// Code generated by ObjectBox; DO NOT EDIT.

package iot

import (
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type event_EntityInfo struct {
	Id  objectbox.TypeId
	Uid uint64
}

var EventBinding = event_EntityInfo{
	Id:  1,
	Uid: 1468539308767086854,
}

var Event_ = struct {
	Id      *objectbox.PropertyUint64
	Uid     *objectbox.PropertyString
	Device  *objectbox.PropertyString
	Date    *objectbox.PropertyInt64
	Picture *objectbox.PropertyByteVector
}{
	Id: &objectbox.PropertyUint64{
		Property: &objectbox.Property{
			Id: 1,
		},
	},
	Uid: &objectbox.PropertyString{
		Property: &objectbox.Property{
			Id: 4,
		},
	},
	Device: &objectbox.PropertyString{
		Property: &objectbox.Property{
			Id: 2,
		},
	},
	Date: &objectbox.PropertyInt64{
		Property: &objectbox.Property{
			Id: 3,
		},
	},
	Picture: &objectbox.PropertyByteVector{
		Property: &objectbox.Property{
			Id: 5,
		},
	},
}

func (event_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Event", 1, 1468539308767086854)
	model.Property("Id", objectbox.PropertyType_Long, 1, 3098166604415018001)
	model.PropertyFlags(objectbox.PropertyFlags_ID)
	model.Property("Uid", objectbox.PropertyType_String, 4, 472416569173577818)
	model.PropertyFlags(objectbox.PropertyFlags_UNIQUE)
	model.PropertyIndex(1, 3297791712577314158)
	model.Property("Device", objectbox.PropertyType_String, 2, 1213411729427304641)
	model.Property("Date", objectbox.PropertyType_Date, 3, 5907655274386702697)
	model.Property("Picture", objectbox.PropertyType_ByteVector, 5, 6024563395733984005)
	model.EntityLastPropertyId(5, 6024563395733984005)
}

func (event_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*Event).Id, nil
}

func (event_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Event).Id = id
	return nil
}

func (event_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) {
	obj := object.(*Event)
	var offsetUid = fbutils.CreateStringOffset(fbb, obj.Uid)
	var offsetDevice = fbutils.CreateStringOffset(fbb, obj.Device)
	var offsetPicture = fbutils.CreateByteVectorOffset(fbb, obj.Picture)

	// build the FlatBuffers object
	fbb.StartObject(5)
	fbb.PrependUint64Slot(0, id, 0)
	fbb.PrependUOffsetTSlot(3, offsetUid, 0)
	fbb.PrependUOffsetTSlot(1, offsetDevice, 0)
	fbb.PrependInt64Slot(2, obj.Date, 0)
	fbb.PrependUOffsetTSlot(4, offsetPicture, 0)
}

func (event_EntityInfo) ToObject(bytes []byte) interface{} {
	table := &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	return &Event{
		Id:      table.GetUint64Slot(4, 0),
		Uid:     fbutils.GetStringSlot(table, 10),
		Device:  fbutils.GetStringSlot(table, 6),
		Date:    table.GetInt64Slot(8, 0),
		Picture: fbutils.GetByteVectorSlot(table, 12),
	}
}

func (event_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Event, 0, capacity)
}

func (event_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]*Event), object.(*Event))
}

type EventBox struct {
	*objectbox.Box
}

func BoxForEvent(ob *objectbox.ObjectBox) *EventBox {
	return &EventBox{
		Box: ob.Box(1),
	}
}

func (box *EventBox) Put(object *Event) (uint64, error) {
	return box.Box.Put(object)
}

func (box *EventBox) PutAsync(object *Event) (uint64, error) {
	return box.Box.PutAsync(object)
}

func (box *EventBox) PutAll(objects []*Event) ([]uint64, error) {
	return box.Box.PutAll(objects)
}

func (box *EventBox) Get(id uint64) (*Event, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Event), nil
}

func (box *EventBox) GetAll() ([]*Event, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Event), nil
}

func (box *EventBox) Remove(object *Event) (err error) {
	return box.Box.Remove(object.Id)
}

func (box *EventBox) Query(conditions ...objectbox.Condition) *EventQuery {
	return &EventQuery{
		box.Box.Query(conditions...),
	}
}

type EventQuery struct {
	*objectbox.Query
}

func (query *EventQuery) Find() ([]*Event, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Event), nil
}

type reading_EntityInfo struct {
	Id  objectbox.TypeId
	Uid uint64
}

var ReadingBinding = reading_EntityInfo{
	Id:  2,
	Uid: 5284076134434938613,
}

var Reading_ = struct {
	Id              *objectbox.PropertyUint64
	Date            *objectbox.PropertyInt64
	EventId         *objectbox.PropertyUint64
	ValueName       *objectbox.PropertyString
	ValueString     *objectbox.PropertyString
	ValueInteger    *objectbox.PropertyInt64
	ValueFloating   *objectbox.PropertyFloat64
	ValueInt32      *objectbox.PropertyInt32
	ValueFloating32 *objectbox.PropertyFloat32
}{
	Id: &objectbox.PropertyUint64{
		Property: &objectbox.Property{
			Id: 1,
		},
	},
	Date: &objectbox.PropertyInt64{
		Property: &objectbox.Property{
			Id: 2,
		},
	},
	EventId: &objectbox.PropertyUint64{
		Property: &objectbox.Property{
			Id: 3,
		},
	},
	ValueName: &objectbox.PropertyString{
		Property: &objectbox.Property{
			Id: 4,
		},
	},
	ValueString: &objectbox.PropertyString{
		Property: &objectbox.Property{
			Id: 5,
		},
	},
	ValueInteger: &objectbox.PropertyInt64{
		Property: &objectbox.Property{
			Id: 6,
		},
	},
	ValueFloating: &objectbox.PropertyFloat64{
		Property: &objectbox.Property{
			Id: 7,
		},
	},
	ValueInt32: &objectbox.PropertyInt32{
		Property: &objectbox.Property{
			Id: 8,
		},
	},
	ValueFloating32: &objectbox.PropertyFloat32{
		Property: &objectbox.Property{
			Id: 9,
		},
	},
}

func (reading_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Reading", 2, 5284076134434938613)
	model.Property("Id", objectbox.PropertyType_Long, 1, 3968063745680890327)
	model.PropertyFlags(objectbox.PropertyFlags_ID)
	model.Property("Date", objectbox.PropertyType_Date, 2, 4852407661923085028)
	model.Property("EventId", objectbox.PropertyType_Relation, 3, 1403806151574554320)
	model.PropertyRelation("Event", 2, 2642563953244304959)
	model.Property("ValueName", objectbox.PropertyType_String, 4, 5626221656121286670)
	model.Property("ValueString", objectbox.PropertyType_String, 5, 7303099924122013060)
	model.Property("ValueInteger", objectbox.PropertyType_Long, 6, 1404333021836291657)
	model.Property("ValueFloating", objectbox.PropertyType_Double, 7, 7102253623343671118)
	model.Property("ValueInt32", objectbox.PropertyType_Int, 8, 7566830186276557216)
	model.Property("ValueFloating32", objectbox.PropertyType_Float, 9, 6040892611651481730)
	model.EntityLastPropertyId(9, 6040892611651481730)
}

func (reading_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*Reading).Id, nil
}

func (reading_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Reading).Id = id
	return nil
}

func (reading_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) {
	obj := object.(*Reading)
	var offsetValueName = fbutils.CreateStringOffset(fbb, obj.ValueName)
	var offsetValueString = fbutils.CreateStringOffset(fbb, obj.ValueString)

	// build the FlatBuffers object
	fbb.StartObject(9)
	fbb.PrependUint64Slot(0, id, 0)
	fbb.PrependInt64Slot(1, obj.Date, 0)
	fbb.PrependUint64Slot(2, obj.EventId, 0)
	fbb.PrependUOffsetTSlot(3, offsetValueName, 0)
	fbb.PrependUOffsetTSlot(4, offsetValueString, 0)
	fbb.PrependInt64Slot(5, obj.ValueInteger, 0)
	fbb.PrependFloat64Slot(6, obj.ValueFloating, 0)
	fbb.PrependInt32Slot(7, obj.ValueInt32, 0)
	fbb.PrependFloat32Slot(8, obj.ValueFloating32, 0)
}

func (reading_EntityInfo) ToObject(bytes []byte) interface{} {
	table := &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	return &Reading{
		Id:              table.GetUint64Slot(4, 0),
		Date:            table.GetInt64Slot(6, 0),
		EventId:         table.GetUint64Slot(8, 0),
		ValueName:       fbutils.GetStringSlot(table, 10),
		ValueString:     fbutils.GetStringSlot(table, 12),
		ValueInteger:    table.GetInt64Slot(14, 0),
		ValueFloating:   table.GetFloat64Slot(16, 0),
		ValueInt32:      table.GetInt32Slot(18, 0),
		ValueFloating32: table.GetFloat32Slot(20, 0),
	}
}

func (reading_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Reading, 0, capacity)
}

func (reading_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]*Reading), object.(*Reading))
}

type ReadingBox struct {
	*objectbox.Box
}

func BoxForReading(ob *objectbox.ObjectBox) *ReadingBox {
	return &ReadingBox{
		Box: ob.Box(2),
	}
}

func (box *ReadingBox) Put(object *Reading) (uint64, error) {
	return box.Box.Put(object)
}

func (box *ReadingBox) PutAsync(object *Reading) (uint64, error) {
	return box.Box.PutAsync(object)
}

func (box *ReadingBox) PutAll(objects []*Reading) ([]uint64, error) {
	return box.Box.PutAll(objects)
}

func (box *ReadingBox) Get(id uint64) (*Reading, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Reading), nil
}

func (box *ReadingBox) GetAll() ([]*Reading, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Reading), nil
}

func (box *ReadingBox) Remove(object *Reading) (err error) {
	return box.Box.Remove(object.Id)
}

func (box *ReadingBox) Query(conditions ...objectbox.Condition) *ReadingQuery {
	return &ReadingQuery{
		box.Box.Query(conditions...),
	}
}

type ReadingQuery struct {
	*objectbox.Query
}

func (query *ReadingQuery) Find() ([]*Reading, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Reading), nil
}
