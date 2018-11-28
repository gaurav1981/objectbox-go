// This file was automatically generated by ObjectBox, do not modify

package iot

import (
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type event_ struct {
	Id  objectbox.TypeId
	Uid uint64
}

var EventBinding = event_{
	Id:  1,
	Uid: 1468539308767086854,
}

func (event_) AddToModel(model *objectbox.Model) {
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

func (event_) GetId(entity interface{}) (uint64, error) {
	return entity.(*Event).Id, nil
}

func (event_) Flatten(entity interface{}, fbb *flatbuffers.Builder, id uint64) {
	ent := entity.(*Event)
	var offsetUid = fbutils.CreateStringOffset(fbb, ent.Uid)
	var offsetDevice = fbutils.CreateStringOffset(fbb, ent.Device)
	var offsetPicture = fbutils.CreateByteVectorOffset(fbb, ent.Picture)

	// build the FlatBuffers object
	fbb.StartObject(5)
	fbb.PrependUint64Slot(0, id, 0)
	fbb.PrependUOffsetTSlot(3, offsetUid, 0)
	fbb.PrependUOffsetTSlot(1, offsetDevice, 0)
	fbb.PrependInt64Slot(2, ent.Date, 0)
	fbb.PrependUOffsetTSlot(4, offsetPicture, 0)
}

func (event_) ToObject(bytes []byte) interface{} {
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

func (event_) MakeSlice(capacity int) interface{} {
	return make([]*Event, 0, capacity)
}

func (event_) AppendToSlice(slice interface{}, entity interface{}) interface{} {
	return append(slice.([]*Event), entity.(*Event))
}

type EventBox struct {
	*objectbox.Box
}

func BoxForEvent(ob *objectbox.ObjectBox) *EventBox {
	return &EventBox{
		Box: ob.Box(1),
	}
}

func (box *EventBox) Get(id uint64) (*Event, error) {
	entity, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if entity == nil {
		return nil, nil
	}
	return entity.(*Event), nil
}

func (box *EventBox) GetAll() ([]*Event, error) {
	entities, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return entities.([]*Event), nil
}

func (box *EventBox) Remove(entity *Event) (err error) {
	return box.Box.Remove(entity.Id)
}

type reading_ struct {
	Id  objectbox.TypeId
	Uid uint64
}

var ReadingBinding = reading_{
	Id:  2,
	Uid: 5284076134434938613,
}

func (reading_) AddToModel(model *objectbox.Model) {
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
	model.EntityLastPropertyId(7, 7102253623343671118)
}

func (reading_) GetId(entity interface{}) (uint64, error) {
	return entity.(*Reading).Id, nil
}

func (reading_) Flatten(entity interface{}, fbb *flatbuffers.Builder, id uint64) {
	ent := entity.(*Reading)
	var offsetValueName = fbutils.CreateStringOffset(fbb, ent.ValueName)
	var offsetValueString = fbutils.CreateStringOffset(fbb, ent.ValueString)

	// build the FlatBuffers object
	fbb.StartObject(7)
	fbb.PrependUint64Slot(0, id, 0)
	fbb.PrependInt64Slot(1, ent.Date, 0)
	fbb.PrependUint64Slot(2, ent.EventId, 0)
	fbb.PrependUOffsetTSlot(3, offsetValueName, 0)
	fbb.PrependUOffsetTSlot(4, offsetValueString, 0)
	fbb.PrependInt64Slot(5, ent.ValueInteger, 0)
	fbb.PrependFloat64Slot(6, ent.ValueFloating, 0)
}

func (reading_) ToObject(bytes []byte) interface{} {
	table := &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	return &Reading{
		Id:            table.GetUint64Slot(4, 0),
		Date:          table.GetInt64Slot(6, 0),
		EventId:       table.GetUint64Slot(8, 0),
		ValueName:     fbutils.GetStringSlot(table, 10),
		ValueString:   fbutils.GetStringSlot(table, 12),
		ValueInteger:  table.GetInt64Slot(14, 0),
		ValueFloating: table.GetFloat64Slot(16, 0),
	}
}

func (reading_) MakeSlice(capacity int) interface{} {
	return make([]*Reading, 0, capacity)
}

func (reading_) AppendToSlice(slice interface{}, entity interface{}) interface{} {
	return append(slice.([]*Reading), entity.(*Reading))
}

type ReadingBox struct {
	*objectbox.Box
}

func BoxForReading(ob *objectbox.ObjectBox) *ReadingBox {
	return &ReadingBox{
		Box: ob.Box(2),
	}
}

func (box *ReadingBox) Get(id uint64) (*Reading, error) {
	entity, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if entity == nil {
		return nil, nil
	}
	return entity.(*Reading), nil
}

func (box *ReadingBox) GetAll() ([]*Reading, error) {
	entities, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return entities.([]*Reading), nil
}

func (box *ReadingBox) Remove(entity *Reading) (err error) {
	return box.Box.Remove(entity.Id)
}
