package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// ObjectMapping 
type ObjectMapping struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewObjectMapping instantiates a new objectMapping and sets the default values.
func NewObjectMapping()(*ObjectMapping) {
    m := &ObjectMapping{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateObjectMappingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateObjectMappingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewObjectMapping(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ObjectMapping) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAttributeMappings gets the attributeMappings property value. The attributeMappings property
func (m *ObjectMapping) GetAttributeMappings()([]AttributeMappingable) {
    val, err := m.GetBackingStore().Get("attributeMappings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]AttributeMappingable)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ObjectMapping) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetEnabled gets the enabled property value. The enabled property
func (m *ObjectMapping) GetEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("enabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ObjectMapping) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attributeMappings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttributeMappingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttributeMappingable, len(val))
            for i, v := range val {
                res[i] = v.(AttributeMappingable)
            }
            m.SetAttributeMappings(res)
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["flowTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseObjectFlowTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlowTypes(val.(*ObjectFlowTypes))
        }
        return nil
    }
    res["metadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateObjectMappingMetadataEntryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ObjectMappingMetadataEntryable, len(val))
            for i, v := range val {
                res[i] = v.(ObjectMappingMetadataEntryable)
            }
            m.SetMetadata(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(Filterable))
        }
        return nil
    }
    res["sourceObjectName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceObjectName(val)
        }
        return nil
    }
    res["targetObjectName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetObjectName(val)
        }
        return nil
    }
    return res
}
// GetFlowTypes gets the flowTypes property value. The flowTypes property
func (m *ObjectMapping) GetFlowTypes()(*ObjectFlowTypes) {
    val, err := m.GetBackingStore().Get("flowTypes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ObjectFlowTypes)
    }
    return nil
}
// GetMetadata gets the metadata property value. The metadata property
func (m *ObjectMapping) GetMetadata()([]ObjectMappingMetadataEntryable) {
    val, err := m.GetBackingStore().Get("metadata")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ObjectMappingMetadataEntryable)
    }
    return nil
}
// GetName gets the name property value. The name property
func (m *ObjectMapping) GetName()(*string) {
    val, err := m.GetBackingStore().Get("name")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ObjectMapping) GetOdataType()(*string) {
    val, err := m.GetBackingStore().Get("odataType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetScope gets the scope property value. The scope property
func (m *ObjectMapping) GetScope()(Filterable) {
    val, err := m.GetBackingStore().Get("scope")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(Filterable)
    }
    return nil
}
// GetSourceObjectName gets the sourceObjectName property value. The sourceObjectName property
func (m *ObjectMapping) GetSourceObjectName()(*string) {
    val, err := m.GetBackingStore().Get("sourceObjectName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTargetObjectName gets the targetObjectName property value. The targetObjectName property
func (m *ObjectMapping) GetTargetObjectName()(*string) {
    val, err := m.GetBackingStore().Get("targetObjectName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ObjectMapping) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttributeMappings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttributeMappings()))
        for i, v := range m.GetAttributeMappings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("attributeMappings", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetFlowTypes() != nil {
        cast := (*m.GetFlowTypes()).String()
        err := writer.WriteStringValue("flowTypes", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMetadata() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMetadata()))
        for i, v := range m.GetMetadata() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("metadata", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceObjectName", m.GetSourceObjectName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("targetObjectName", m.GetTargetObjectName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ObjectMapping) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAttributeMappings sets the attributeMappings property value. The attributeMappings property
func (m *ObjectMapping) SetAttributeMappings(value []AttributeMappingable)() {
    err := m.GetBackingStore().Set("attributeMappings", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ObjectMapping) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetEnabled sets the enabled property value. The enabled property
func (m *ObjectMapping) SetEnabled(value *bool)() {
    err := m.GetBackingStore().Set("enabled", value)
    if err != nil {
        panic(err)
    }
}
// SetFlowTypes sets the flowTypes property value. The flowTypes property
func (m *ObjectMapping) SetFlowTypes(value *ObjectFlowTypes)() {
    err := m.GetBackingStore().Set("flowTypes", value)
    if err != nil {
        panic(err)
    }
}
// SetMetadata sets the metadata property value. The metadata property
func (m *ObjectMapping) SetMetadata(value []ObjectMappingMetadataEntryable)() {
    err := m.GetBackingStore().Set("metadata", value)
    if err != nil {
        panic(err)
    }
}
// SetName sets the name property value. The name property
func (m *ObjectMapping) SetName(value *string)() {
    err := m.GetBackingStore().Set("name", value)
    if err != nil {
        panic(err)
    }
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ObjectMapping) SetOdataType(value *string)() {
    err := m.GetBackingStore().Set("odataType", value)
    if err != nil {
        panic(err)
    }
}
// SetScope sets the scope property value. The scope property
func (m *ObjectMapping) SetScope(value Filterable)() {
    err := m.GetBackingStore().Set("scope", value)
    if err != nil {
        panic(err)
    }
}
// SetSourceObjectName sets the sourceObjectName property value. The sourceObjectName property
func (m *ObjectMapping) SetSourceObjectName(value *string)() {
    err := m.GetBackingStore().Set("sourceObjectName", value)
    if err != nil {
        panic(err)
    }
}
// SetTargetObjectName sets the targetObjectName property value. The targetObjectName property
func (m *ObjectMapping) SetTargetObjectName(value *string)() {
    err := m.GetBackingStore().Set("targetObjectName", value)
    if err != nil {
        panic(err)
    }
}
// ObjectMappingable 
type ObjectMappingable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttributeMappings()([]AttributeMappingable)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetEnabled()(*bool)
    GetFlowTypes()(*ObjectFlowTypes)
    GetMetadata()([]ObjectMappingMetadataEntryable)
    GetName()(*string)
    GetOdataType()(*string)
    GetScope()(Filterable)
    GetSourceObjectName()(*string)
    GetTargetObjectName()(*string)
    SetAttributeMappings(value []AttributeMappingable)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetEnabled(value *bool)()
    SetFlowTypes(value *ObjectFlowTypes)()
    SetMetadata(value []ObjectMappingMetadataEntryable)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetScope(value Filterable)()
    SetSourceObjectName(value *string)()
    SetTargetObjectName(value *string)()
}
