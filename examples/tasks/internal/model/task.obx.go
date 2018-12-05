// Code generated by ObjectBox; DO NOT EDIT.

package model

import (
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type task_EntityInfo struct {
	Id  objectbox.TypeId
	Uid uint64
}

var TaskBinding = task_EntityInfo{
	Id:  1,
	Uid: 1306759095002958910,
}

// Task_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Task_ = struct {
	Id           *objectbox.PropertyUint64
	Text         *objectbox.PropertyString
	DateCreated  *objectbox.PropertyInt64
	DateFinished *objectbox.PropertyInt64
}{
	Id: &objectbox.PropertyUint64{
		Property: &objectbox.Property{
			Id: 1,
		},
	},
	Text: &objectbox.PropertyString{
		Property: &objectbox.Property{
			Id: 2,
		},
	},
	DateCreated: &objectbox.PropertyInt64{
		Property: &objectbox.Property{
			Id: 3,
		},
	},
	DateFinished: &objectbox.PropertyInt64{
		Property: &objectbox.Property{
			Id: 4,
		},
	},
}

// GeneratorVersion is called by the ObjectBox to verify the compatibility of the generator used to generate this code
func (task_EntityInfo) GeneratorVersion() int {
	return 1
}

// AddToModel is called by the ObjectBox during model build
func (task_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Task", 1, 1306759095002958910)
	model.Property("Id", objectbox.PropertyType_Long, 1, 2193439623591184445)
	model.PropertyFlags(objectbox.PropertyFlags_ID)
	model.Property("Text", objectbox.PropertyType_String, 2, 6177929178231325611)
	model.Property("DateCreated", objectbox.PropertyType_Long, 3, 9141374017424160113)
	model.Property("DateFinished", objectbox.PropertyType_Long, 4, 8083684673086871702)
	model.EntityLastPropertyId(4, 8083684673086871702)
}

// GetId is called by the ObjectBox during Put operations to check for existing ID on an object
func (task_EntityInfo) GetId(object interface{}) (uint64, error) {
	return object.(*Task).Id, nil
}

// SetId is called by the ObjectBox during Put to update an ID on an object that has just been inserted
func (task_EntityInfo) SetId(object interface{}, id uint64) error {
	object.(*Task).Id = id
	return nil
}

// Flatten is called by the ObjectBox to transform an object to a FlatBuffer
func (task_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) {
	obj := object.(*Task)
	var offsetText = fbutils.CreateStringOffset(fbb, obj.Text)

	// build the FlatBuffers object
	fbb.StartObject(4)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetUOffsetTSlot(fbb, 1, offsetText)
	fbutils.SetInt64Slot(fbb, 2, obj.DateCreated)
	fbutils.SetInt64Slot(fbb, 3, obj.DateFinished)
}

// ToObject is called by the ObjectBox to load an object from a FlatBuffer
func (task_EntityInfo) ToObject(bytes []byte) interface{} {
	table := &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	return &Task{
		Id:           table.GetUint64Slot(4, 0),
		Text:         fbutils.GetStringSlot(table, 6),
		DateCreated:  table.GetInt64Slot(8, 0),
		DateFinished: table.GetInt64Slot(10, 0),
	}
}

// MakeSlice is called by the ObjectBox to construct a new slice to hold the read objects
func (task_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]*Task, 0, capacity)
}

// AppendToSlice is called by the ObjectBox to fill the slice of the read objects
func (task_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]*Task), object.(*Task))
}

// Box provides CRUD access to Task objects
type TaskBox struct {
	*objectbox.Box
}

// BoxForTask opens a box of Task objects
func BoxForTask(ob *objectbox.ObjectBox) *TaskBox {
	return &TaskBox{
		Box: ob.InternalBox(1),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Task.Id property on the passed object will be assigned the new ID as well.
func (box *TaskBox) Put(object *Task) (uint64, error) {
	return box.Box.Put(object)
}

// PutAsync asynchronously inserts/updates a single object.
// When inserting, the Task.Id property on the passed object will be assigned the new ID as well.
//
// It's executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "Put & Forget:" you gain faster puts as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
//
// In situations with (extremely) high async load, this method may be throttled (~1ms) or delayed (<1s).
// In the unlikely event that the object could not be enqueued after delaying, an error will be returned.
//
// Note that this method does not give you hard durability guarantees like the synchronous Put provides.
// There is a small time window (typically 3 ms) in which the data may not have been committed durably yet.
func (box *TaskBox) PutAsync(object *Task) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutAll inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Task.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Task.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *TaskBox) PutAll(objects []*Task) ([]uint64, error) {
	return box.Box.PutAll(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *TaskBox) Get(id uint64) (*Task, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Task), nil
}

// Get reads all stored objects
func (box *TaskBox) GetAll() ([]*Task, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]*Task), nil
}

// Remove deletes a single object
func (box *TaskBox) Remove(object *Task) (err error) {
	return box.Box.Remove(object.Id)
}

// Creates a query with the given conditions. Use the fields of the Task_ struct to create conditions.
// Keep the *TaskQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *TaskBox) Query(conditions ...objectbox.Condition) *TaskQuery {
	return &TaskQuery{
		box.Box.Query(conditions...),
	}
}

// Query provides a way to search stored objects
//
// For example, you can find all Task which Id is lower than 100:
// 		box.Query(Task_.Id.LessThan(100)).Find()
type TaskQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *TaskQuery) Find() ([]*Task, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]*Task), nil
}
