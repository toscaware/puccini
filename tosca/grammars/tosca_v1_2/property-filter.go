package tosca_v1_2

import (
	"github.com/tliron/puccini/tosca"
	"github.com/tliron/puccini/tosca/normal"
)

//
// PropertyFilter
//
// [TOSCA-Simple-Profile-YAML-v1.2] @ 3.6.4
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.5.3
//

type PropertyFilter struct {
	*Entity `name:"property filter"`
	Name    string

	ConstraintClauses ConstraintClauses
}

func NewPropertyFilter(context *tosca.Context) *PropertyFilter {
	return &PropertyFilter{
		Entity: NewEntity(context),
		Name:   context.Name,
	}
}

// tosca.Reader signature
func ReadPropertyFilter(context *tosca.Context) interface{} {
	self := NewPropertyFilter(context)

	context.ReadListItems(ReadConstraintClause, func(item interface{}) {
		self.ConstraintClauses = append(self.ConstraintClauses, item.(*ConstraintClause))
	})

	return self
}

// tosca.Mappable interface
func (self *PropertyFilter) GetKey() string {
	return self.Name
}

func (self *PropertyFilter) Normalize(functionsMap normal.FunctionsMap) normal.Functions {
	if len(self.ConstraintClauses) == 0 {
		return nil
	}

	functions := self.ConstraintClauses.Normalize(self.Context)
	functionsMap[self.Name] = functions
	return functions
}

//
// PropertyFilters
//

type PropertyFilters map[string]*PropertyFilter

func (self PropertyFilters) Normalize(functionsMap normal.FunctionsMap) {
	for _, propertyFilter := range self {
		propertyFilter.Normalize(functionsMap)
	}
}
