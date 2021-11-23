package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// sharePointIdentity 
type SharePointIdentity struct {
    Identity
    // 
    loginName *string;
}
// NewSharePointIdentity instantiates a new sharePointIdentity and sets the default values.
func NewSharePointIdentity()(*SharePointIdentity) {
    m := &SharePointIdentity{
        Identity: *NewIdentity(),
    }
    return m
}
// GetLoginName gets the loginName property value. 
func (m *SharePointIdentity) GetLoginName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.loginName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharePointIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["loginName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoginName(val)
        }
        return nil
    }
    return res
}
func (m *SharePointIdentity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SharePointIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("loginName", m.GetLoginName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLoginName sets the loginName property value. 
func (m *SharePointIdentity) SetLoginName(value *string)() {
    m.loginName = value
}
