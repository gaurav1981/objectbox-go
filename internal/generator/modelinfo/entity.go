package modelinfo

import (
	"fmt"
	"strings"
)

type Entity struct {
	Id             IdUid       `json:"id"`
	Name           string      `json:"name"`
	LastPropertyId IdUid       `json:"lastPropertyId"`
	Properties     []*Property `json:"properties"`

	model *ModelInfo
}

func CreateEntity(model *ModelInfo, id id, uid uid) *Entity {
	return &Entity{
		model:      model,
		Id:         CreateIdUid(id, uid),
		Properties: make([]*Property, 0),
	}
}

// performs initial validation of loaded data so that it doesn't have to be checked in each function
func (entity *Entity) Validate() (err error) {
	if err = entity.Id.Validate(); err != nil {
		return err
	}

	if len(entity.Name) == 0 {
		return fmt.Errorf("name is undefined")
	}

	if len(entity.Properties) > 0 {
		if err = entity.LastPropertyId.Validate(); err != nil {
			return fmt.Errorf("lastPropertyId: %s", err)
		}

		var lastId = entity.LastPropertyId.getIdSafe()
		var lastUid = entity.LastPropertyId.getUidSafe()

		var found = false
		for _, property := range entity.Properties {
			if lastId == property.Id.getIdSafe() {
				if lastUid != property.Id.getUidSafe() {
					return fmt.Errorf("lastPropertyId %s doesn't match property %s %s",
						entity.LastPropertyId, property.Name, property.Id)
				}
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("lastPropertyId %s doesn't match any property", entity.LastPropertyId)
		}
	}

	if entity.Properties == nil {
		return fmt.Errorf("properties are not defined or not an array")
	}

	for _, property := range entity.Properties {
		err = property.Validate()
		if err != nil {
			return fmt.Errorf("property %s %s is invalid: %s", entity.Name, string(entity.Id), err)
		}
	}

	return nil
}

func (entity *Entity) FindPropertyByUid(uid uid) (*Property, error) {
	for _, property := range entity.Properties {
		propertyUid, _ := property.Id.GetUid()
		if propertyUid == uid {
			return property, nil
		}
	}

	return nil, fmt.Errorf("property with Uid %d not found", uid)
}

func (entity *Entity) FindPropertyByName(name string) (*Property, error) {
	for _, property := range entity.Properties {
		if strings.ToLower(property.Name) == strings.ToLower(name) {
			return property, nil
		}
	}

	return nil, fmt.Errorf("property with Name %s not found", name)
}

func (entity *Entity) CreateProperty() (*Property, error) {
	var id id = 1
	if len(entity.Properties) > 0 {
		id = entity.LastPropertyId.getIdSafe() + 1
	}

	// generate a unique UID
	uniqueUid, err := entity.model.generateUid()

	if err != nil {
		return nil, err
	}

	var property = CreateProperty(id, uniqueUid)

	entity.Properties = append(entity.Properties, property)
	entity.LastPropertyId = property.Id

	return property, nil
}

// recursively checks whether given UID is present in the model
func (entity *Entity) containsUid(searched uid) bool {
	if entity.Id.getUidSafe() == searched {
		return true
	}

	if entity.LastPropertyId.getUidSafe() == searched {
		return true
	}

	for _, property := range entity.Properties {
		if property.containsUid(searched) {
			return true
		}
	}

	return false
}
