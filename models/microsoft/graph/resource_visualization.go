package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// resourceVisualization 
type ResourceVisualization struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A string describing where the item is stored. For example, the name of a SharePoint site or the user name identifying the owner of the OneDrive storing the item.
    containerDisplayName *string;
    // Can be used for filtering by the type of container in which the file is stored. Such as Site or OneDriveBusiness.
    containerType *string;
    // A path leading to the folder in which the item is stored.
    containerWebUrl *string;
    // The item's media type. Can be used for filtering for a specific type of file based on supported IANA Media Mime Types. Note that not all Media Mime Types are supported.
    mediaType *string;
    // A URL leading to the preview image for the item.
    previewImageUrl *string;
    // A preview text for the item.
    previewText *string;
    // The item's title text.
    title *string;
    // The item's media type. Can be used for filtering for a specific file based on a specific type. See below for supported types.
    type_escaped *string;
}
// NewResourceVisualization instantiates a new resourceVisualization and sets the default values.
func NewResourceVisualization()(*ResourceVisualization) {
    m := &ResourceVisualization{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResourceVisualization) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetContainerDisplayName gets the containerDisplayName property value. A string describing where the item is stored. For example, the name of a SharePoint site or the user name identifying the owner of the OneDrive storing the item.
func (m *ResourceVisualization) GetContainerDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.containerDisplayName
    }
}
// GetContainerType gets the containerType property value. Can be used for filtering by the type of container in which the file is stored. Such as Site or OneDriveBusiness.
func (m *ResourceVisualization) GetContainerType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.containerType
    }
}
// GetContainerWebUrl gets the containerWebUrl property value. A path leading to the folder in which the item is stored.
func (m *ResourceVisualization) GetContainerWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.containerWebUrl
    }
}
// GetMediaType gets the mediaType property value. The item's media type. Can be used for filtering for a specific type of file based on supported IANA Media Mime Types. Note that not all Media Mime Types are supported.
func (m *ResourceVisualization) GetMediaType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mediaType
    }
}
// GetPreviewImageUrl gets the previewImageUrl property value. A URL leading to the preview image for the item.
func (m *ResourceVisualization) GetPreviewImageUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.previewImageUrl
    }
}
// GetPreviewText gets the previewText property value. A preview text for the item.
func (m *ResourceVisualization) GetPreviewText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.previewText
    }
}
// GetTitle gets the title property value. The item's title text.
func (m *ResourceVisualization) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// GetType_escaped gets the type_escaped property value. The item's media type. Can be used for filtering for a specific file based on a specific type. See below for supported types.
func (m *ResourceVisualization) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResourceVisualization) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["containerDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainerDisplayName(val)
        }
        return nil
    }
    res["containerType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainerType(val)
        }
        return nil
    }
    res["containerWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContainerWebUrl(val)
        }
        return nil
    }
    res["mediaType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaType(val)
        }
        return nil
    }
    res["previewImageUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewImageUrl(val)
        }
        return nil
    }
    res["previewText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewText(val)
        }
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType_escaped(val)
        }
        return nil
    }
    return res
}
func (m *ResourceVisualization) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ResourceVisualization) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("containerDisplayName", m.GetContainerDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("containerType", m.GetContainerType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("containerWebUrl", m.GetContainerWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaType", m.GetMediaType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("previewImageUrl", m.GetPreviewImageUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("previewText", m.GetPreviewText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type_escaped", m.GetType_escaped())
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResourceVisualization) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContainerDisplayName sets the containerDisplayName property value. A string describing where the item is stored. For example, the name of a SharePoint site or the user name identifying the owner of the OneDrive storing the item.
func (m *ResourceVisualization) SetContainerDisplayName(value *string)() {
    m.containerDisplayName = value
}
// SetContainerType sets the containerType property value. Can be used for filtering by the type of container in which the file is stored. Such as Site or OneDriveBusiness.
func (m *ResourceVisualization) SetContainerType(value *string)() {
    m.containerType = value
}
// SetContainerWebUrl sets the containerWebUrl property value. A path leading to the folder in which the item is stored.
func (m *ResourceVisualization) SetContainerWebUrl(value *string)() {
    m.containerWebUrl = value
}
// SetMediaType sets the mediaType property value. The item's media type. Can be used for filtering for a specific type of file based on supported IANA Media Mime Types. Note that not all Media Mime Types are supported.
func (m *ResourceVisualization) SetMediaType(value *string)() {
    m.mediaType = value
}
// SetPreviewImageUrl sets the previewImageUrl property value. A URL leading to the preview image for the item.
func (m *ResourceVisualization) SetPreviewImageUrl(value *string)() {
    m.previewImageUrl = value
}
// SetPreviewText sets the previewText property value. A preview text for the item.
func (m *ResourceVisualization) SetPreviewText(value *string)() {
    m.previewText = value
}
// SetTitle sets the title property value. The item's title text.
func (m *ResourceVisualization) SetTitle(value *string)() {
    m.title = value
}
// SetType_escaped sets the type_escaped property value. The item's media type. Can be used for filtering for a specific file based on a specific type. See below for supported types.
func (m *ResourceVisualization) SetType_escaped(value *string)() {
    m.type_escaped = value
}
