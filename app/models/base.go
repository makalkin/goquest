package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Base struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`

	// We want this to default to false without any work. So this will be the opposite of isNew. We want it to be new unless set to existing
	exists bool
}

// Satisfy the new tracker interface
func (d *Base) SetIsNew(isNew bool) {
	d.exists = !isNew
}

func (d *Base) IsNew() bool {
	return !d.exists
}

// Satisfy the document interface
func (d *Base) GetId() bson.ObjectId {
	return d.Id
}

func (d *Base) SetId(id bson.ObjectId) {
	d.Id = id
}
