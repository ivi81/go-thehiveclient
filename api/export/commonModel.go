package export

import "gitlab.cloud.gcm/i.ippolitov/go-thehiveclient/api/cons/common"

//HiveCommonAttr атрибуты общие для всех типов сущиностей (Формируются и возваращаются theHive):
//- createdBy (text) : login of the user who created the entity
//- createdAt (date) : date and time of the creation
//- updatedBy (text) : login of the user who last updated the entity
//- updatedAt (date) : date and time of the last update
//- user (text) : same value as createdBy (this field is deprecated) These attributes are handled by the back-end and can't be directly updated.
type HiveCommonAttr struct {
	Id        string            `json:"id"`
	CreatedBy string            `json:"createdBy"`
	CreatedAt int               `json:"createdAt"` //CustomTime       `json:"createdAt"` //time.Time        `json:"createdAt"`
	UpdatedBy string            `json:"updatedBy"`
	UpdatedAt int               `json:"updatedAt"` //CustomTime       `json:"updatedAt,omitempty"` //time.Time        `json:"updatedAt,omitempty"`
	Type      common.ObjectType `json:"type"`
}

/*type ReqValidator interface {
	Validate() bool
}*/
