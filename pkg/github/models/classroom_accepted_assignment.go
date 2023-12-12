package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClassroomAcceptedAssignment a GitHub Classroom accepted assignment
type ClassroomAcceptedAssignment struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A GitHub Classroom assignment
    assignment SimpleClassroomAssignmentable
    // Count of student commits.
    commit_count *int32
    // Most recent grade.
    grade *string
    // Unique identifier of the repository.
    id *int32
    // Whether a submission passed.
    passing *bool
    // A GitHub repository view for Classroom
    repository SimpleClassroomRepositoryable
    // The students property
    students []SimpleClassroomUserable
    // Whether an accepted assignment has been submitted.
    submitted *bool
}
// NewClassroomAcceptedAssignment instantiates a new classroomAcceptedAssignment and sets the default values.
func NewClassroomAcceptedAssignment()(*ClassroomAcceptedAssignment) {
    m := &ClassroomAcceptedAssignment{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateClassroomAcceptedAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassroomAcceptedAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClassroomAcceptedAssignment(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ClassroomAcceptedAssignment) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignment gets the assignment property value. A GitHub Classroom assignment
func (m *ClassroomAcceptedAssignment) GetAssignment()(SimpleClassroomAssignmentable) {
    return m.assignment
}
// GetCommitCount gets the commit_count property value. Count of student commits.
func (m *ClassroomAcceptedAssignment) GetCommitCount()(*int32) {
    return m.commit_count
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClassroomAcceptedAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleClassroomAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignment(val.(SimpleClassroomAssignmentable))
        }
        return nil
    }
    res["commit_count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommitCount(val)
        }
        return nil
    }
    res["grade"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrade(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["passing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassing(val)
        }
        return nil
    }
    res["repository"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimpleClassroomRepositoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRepository(val.(SimpleClassroomRepositoryable))
        }
        return nil
    }
    res["students"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSimpleClassroomUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SimpleClassroomUserable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SimpleClassroomUserable)
                }
            }
            m.SetStudents(res)
        }
        return nil
    }
    res["submitted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubmitted(val)
        }
        return nil
    }
    return res
}
// GetGrade gets the grade property value. Most recent grade.
func (m *ClassroomAcceptedAssignment) GetGrade()(*string) {
    return m.grade
}
// GetId gets the id property value. Unique identifier of the repository.
func (m *ClassroomAcceptedAssignment) GetId()(*int32) {
    return m.id
}
// GetPassing gets the passing property value. Whether a submission passed.
func (m *ClassroomAcceptedAssignment) GetPassing()(*bool) {
    return m.passing
}
// GetRepository gets the repository property value. A GitHub repository view for Classroom
func (m *ClassroomAcceptedAssignment) GetRepository()(SimpleClassroomRepositoryable) {
    return m.repository
}
// GetStudents gets the students property value. The students property
func (m *ClassroomAcceptedAssignment) GetStudents()([]SimpleClassroomUserable) {
    return m.students
}
// GetSubmitted gets the submitted property value. Whether an accepted assignment has been submitted.
func (m *ClassroomAcceptedAssignment) GetSubmitted()(*bool) {
    return m.submitted
}
// Serialize serializes information the current object
func (m *ClassroomAcceptedAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("assignment", m.GetAssignment())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("commit_count", m.GetCommitCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("grade", m.GetGrade())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("passing", m.GetPassing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("repository", m.GetRepository())
        if err != nil {
            return err
        }
    }
    if m.GetStudents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStudents()))
        for i, v := range m.GetStudents() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("students", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("submitted", m.GetSubmitted())
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
func (m *ClassroomAcceptedAssignment) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignment sets the assignment property value. A GitHub Classroom assignment
func (m *ClassroomAcceptedAssignment) SetAssignment(value SimpleClassroomAssignmentable)() {
    m.assignment = value
}
// SetCommitCount sets the commit_count property value. Count of student commits.
func (m *ClassroomAcceptedAssignment) SetCommitCount(value *int32)() {
    m.commit_count = value
}
// SetGrade sets the grade property value. Most recent grade.
func (m *ClassroomAcceptedAssignment) SetGrade(value *string)() {
    m.grade = value
}
// SetId sets the id property value. Unique identifier of the repository.
func (m *ClassroomAcceptedAssignment) SetId(value *int32)() {
    m.id = value
}
// SetPassing sets the passing property value. Whether a submission passed.
func (m *ClassroomAcceptedAssignment) SetPassing(value *bool)() {
    m.passing = value
}
// SetRepository sets the repository property value. A GitHub repository view for Classroom
func (m *ClassroomAcceptedAssignment) SetRepository(value SimpleClassroomRepositoryable)() {
    m.repository = value
}
// SetStudents sets the students property value. The students property
func (m *ClassroomAcceptedAssignment) SetStudents(value []SimpleClassroomUserable)() {
    m.students = value
}
// SetSubmitted sets the submitted property value. Whether an accepted assignment has been submitted.
func (m *ClassroomAcceptedAssignment) SetSubmitted(value *bool)() {
    m.submitted = value
}
// ClassroomAcceptedAssignmentable 
type ClassroomAcceptedAssignmentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignment()(SimpleClassroomAssignmentable)
    GetCommitCount()(*int32)
    GetGrade()(*string)
    GetId()(*int32)
    GetPassing()(*bool)
    GetRepository()(SimpleClassroomRepositoryable)
    GetStudents()([]SimpleClassroomUserable)
    GetSubmitted()(*bool)
    SetAssignment(value SimpleClassroomAssignmentable)()
    SetCommitCount(value *int32)()
    SetGrade(value *string)()
    SetId(value *int32)()
    SetPassing(value *bool)()
    SetRepository(value SimpleClassroomRepositoryable)()
    SetStudents(value []SimpleClassroomUserable)()
    SetSubmitted(value *bool)()
}
