package addpassword

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// PasswordCredentialRequestBody 
type PasswordCredentialRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    passwordCredential *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PasswordCredential;
}
// NewPasswordCredentialRequestBody instantiates a new PasswordCredentialRequestBody and sets the default values.
func NewPasswordCredentialRequestBody()(*PasswordCredentialRequestBody) {
    m := &PasswordCredentialRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordCredentialRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetPasswordCredential gets the passwordCredential property value. 
func (m *PasswordCredentialRequestBody) GetPasswordCredential()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PasswordCredential) {
    if m == nil {
        return nil
    } else {
        return m.passwordCredential
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PasswordCredentialRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["passwordCredential"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPasswordCredential() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordCredential(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PasswordCredential))
        }
        return nil
    }
    return res
}
func (m *PasswordCredentialRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PasswordCredentialRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("passwordCredential", m.GetPasswordCredential())
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
func (m *PasswordCredentialRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetPasswordCredential sets the passwordCredential property value. 
func (m *PasswordCredentialRequestBody) SetPasswordCredential(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PasswordCredential)() {
    m.passwordCredential = value
}
