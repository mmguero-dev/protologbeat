// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

type AccountGateStatus string

// Enum values for AccountGateStatus
const (
	AccountGateStatusSucceeded AccountGateStatus = "SUCCEEDED"
	AccountGateStatusFailed    AccountGateStatus = "FAILED"
	AccountGateStatusSkipped   AccountGateStatus = "SKIPPED"
)

func (enum AccountGateStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AccountGateStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Capability string

// Enum values for Capability
const (
	CapabilityCapabilityIam        Capability = "CAPABILITY_IAM"
	CapabilityCapabilityNamedIam   Capability = "CAPABILITY_NAMED_IAM"
	CapabilityCapabilityAutoExpand Capability = "CAPABILITY_AUTO_EXPAND"
)

func (enum Capability) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Capability) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChangeAction string

// Enum values for ChangeAction
const (
	ChangeActionAdd    ChangeAction = "Add"
	ChangeActionModify ChangeAction = "Modify"
	ChangeActionRemove ChangeAction = "Remove"
)

func (enum ChangeAction) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChangeAction) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChangeSetStatus string

// Enum values for ChangeSetStatus
const (
	ChangeSetStatusCreatePending    ChangeSetStatus = "CREATE_PENDING"
	ChangeSetStatusCreateInProgress ChangeSetStatus = "CREATE_IN_PROGRESS"
	ChangeSetStatusCreateComplete   ChangeSetStatus = "CREATE_COMPLETE"
	ChangeSetStatusDeleteComplete   ChangeSetStatus = "DELETE_COMPLETE"
	ChangeSetStatusFailed           ChangeSetStatus = "FAILED"
)

func (enum ChangeSetStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChangeSetStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChangeSetType string

// Enum values for ChangeSetType
const (
	ChangeSetTypeCreate ChangeSetType = "CREATE"
	ChangeSetTypeUpdate ChangeSetType = "UPDATE"
)

func (enum ChangeSetType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChangeSetType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChangeSource string

// Enum values for ChangeSource
const (
	ChangeSourceResourceReference  ChangeSource = "ResourceReference"
	ChangeSourceParameterReference ChangeSource = "ParameterReference"
	ChangeSourceResourceAttribute  ChangeSource = "ResourceAttribute"
	ChangeSourceDirectModification ChangeSource = "DirectModification"
	ChangeSourceAutomatic          ChangeSource = "Automatic"
)

func (enum ChangeSource) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChangeSource) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ChangeType string

// Enum values for ChangeType
const (
	ChangeTypeResource ChangeType = "Resource"
)

func (enum ChangeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ChangeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DifferenceType string

// Enum values for DifferenceType
const (
	DifferenceTypeAdd      DifferenceType = "ADD"
	DifferenceTypeRemove   DifferenceType = "REMOVE"
	DifferenceTypeNotEqual DifferenceType = "NOT_EQUAL"
)

func (enum DifferenceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DifferenceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EvaluationType string

// Enum values for EvaluationType
const (
	EvaluationTypeStatic  EvaluationType = "Static"
	EvaluationTypeDynamic EvaluationType = "Dynamic"
)

func (enum EvaluationType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EvaluationType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ExecutionStatus string

// Enum values for ExecutionStatus
const (
	ExecutionStatusUnavailable       ExecutionStatus = "UNAVAILABLE"
	ExecutionStatusAvailable         ExecutionStatus = "AVAILABLE"
	ExecutionStatusExecuteInProgress ExecutionStatus = "EXECUTE_IN_PROGRESS"
	ExecutionStatusExecuteComplete   ExecutionStatus = "EXECUTE_COMPLETE"
	ExecutionStatusExecuteFailed     ExecutionStatus = "EXECUTE_FAILED"
	ExecutionStatusObsolete          ExecutionStatus = "OBSOLETE"
)

func (enum ExecutionStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ExecutionStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OnFailure string

// Enum values for OnFailure
const (
	OnFailureDoNothing OnFailure = "DO_NOTHING"
	OnFailureRollback  OnFailure = "ROLLBACK"
	OnFailureDelete    OnFailure = "DELETE"
)

func (enum OnFailure) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OnFailure) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Replacement string

// Enum values for Replacement
const (
	ReplacementTrue        Replacement = "True"
	ReplacementFalse       Replacement = "False"
	ReplacementConditional Replacement = "Conditional"
)

func (enum Replacement) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Replacement) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RequiresRecreation string

// Enum values for RequiresRecreation
const (
	RequiresRecreationNever         RequiresRecreation = "Never"
	RequiresRecreationConditionally RequiresRecreation = "Conditionally"
	RequiresRecreationAlways        RequiresRecreation = "Always"
)

func (enum RequiresRecreation) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RequiresRecreation) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceAttribute string

// Enum values for ResourceAttribute
const (
	ResourceAttributeProperties     ResourceAttribute = "Properties"
	ResourceAttributeMetadata       ResourceAttribute = "Metadata"
	ResourceAttributeCreationPolicy ResourceAttribute = "CreationPolicy"
	ResourceAttributeUpdatePolicy   ResourceAttribute = "UpdatePolicy"
	ResourceAttributeDeletionPolicy ResourceAttribute = "DeletionPolicy"
	ResourceAttributeTags           ResourceAttribute = "Tags"
)

func (enum ResourceAttribute) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceAttribute) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceSignalStatus string

// Enum values for ResourceSignalStatus
const (
	ResourceSignalStatusSuccess ResourceSignalStatus = "SUCCESS"
	ResourceSignalStatusFailure ResourceSignalStatus = "FAILURE"
)

func (enum ResourceSignalStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceSignalStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceStatus string

// Enum values for ResourceStatus
const (
	ResourceStatusCreateInProgress ResourceStatus = "CREATE_IN_PROGRESS"
	ResourceStatusCreateFailed     ResourceStatus = "CREATE_FAILED"
	ResourceStatusCreateComplete   ResourceStatus = "CREATE_COMPLETE"
	ResourceStatusDeleteInProgress ResourceStatus = "DELETE_IN_PROGRESS"
	ResourceStatusDeleteFailed     ResourceStatus = "DELETE_FAILED"
	ResourceStatusDeleteComplete   ResourceStatus = "DELETE_COMPLETE"
	ResourceStatusDeleteSkipped    ResourceStatus = "DELETE_SKIPPED"
	ResourceStatusUpdateInProgress ResourceStatus = "UPDATE_IN_PROGRESS"
	ResourceStatusUpdateFailed     ResourceStatus = "UPDATE_FAILED"
	ResourceStatusUpdateComplete   ResourceStatus = "UPDATE_COMPLETE"
)

func (enum ResourceStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StackDriftDetectionStatus string

// Enum values for StackDriftDetectionStatus
const (
	StackDriftDetectionStatusDetectionInProgress StackDriftDetectionStatus = "DETECTION_IN_PROGRESS"
	StackDriftDetectionStatusDetectionFailed     StackDriftDetectionStatus = "DETECTION_FAILED"
	StackDriftDetectionStatusDetectionComplete   StackDriftDetectionStatus = "DETECTION_COMPLETE"
)

func (enum StackDriftDetectionStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StackDriftDetectionStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StackDriftStatus string

// Enum values for StackDriftStatus
const (
	StackDriftStatusDrifted    StackDriftStatus = "DRIFTED"
	StackDriftStatusInSync     StackDriftStatus = "IN_SYNC"
	StackDriftStatusUnknown    StackDriftStatus = "UNKNOWN"
	StackDriftStatusNotChecked StackDriftStatus = "NOT_CHECKED"
)

func (enum StackDriftStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StackDriftStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StackInstanceStatus string

// Enum values for StackInstanceStatus
const (
	StackInstanceStatusCurrent    StackInstanceStatus = "CURRENT"
	StackInstanceStatusOutdated   StackInstanceStatus = "OUTDATED"
	StackInstanceStatusInoperable StackInstanceStatus = "INOPERABLE"
)

func (enum StackInstanceStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StackInstanceStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StackResourceDriftStatus string

// Enum values for StackResourceDriftStatus
const (
	StackResourceDriftStatusInSync     StackResourceDriftStatus = "IN_SYNC"
	StackResourceDriftStatusModified   StackResourceDriftStatus = "MODIFIED"
	StackResourceDriftStatusDeleted    StackResourceDriftStatus = "DELETED"
	StackResourceDriftStatusNotChecked StackResourceDriftStatus = "NOT_CHECKED"
)

func (enum StackResourceDriftStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StackResourceDriftStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StackSetOperationAction string

// Enum values for StackSetOperationAction
const (
	StackSetOperationActionCreate StackSetOperationAction = "CREATE"
	StackSetOperationActionUpdate StackSetOperationAction = "UPDATE"
	StackSetOperationActionDelete StackSetOperationAction = "DELETE"
)

func (enum StackSetOperationAction) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StackSetOperationAction) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StackSetOperationResultStatus string

// Enum values for StackSetOperationResultStatus
const (
	StackSetOperationResultStatusPending   StackSetOperationResultStatus = "PENDING"
	StackSetOperationResultStatusRunning   StackSetOperationResultStatus = "RUNNING"
	StackSetOperationResultStatusSucceeded StackSetOperationResultStatus = "SUCCEEDED"
	StackSetOperationResultStatusFailed    StackSetOperationResultStatus = "FAILED"
	StackSetOperationResultStatusCancelled StackSetOperationResultStatus = "CANCELLED"
)

func (enum StackSetOperationResultStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StackSetOperationResultStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StackSetOperationStatus string

// Enum values for StackSetOperationStatus
const (
	StackSetOperationStatusRunning   StackSetOperationStatus = "RUNNING"
	StackSetOperationStatusSucceeded StackSetOperationStatus = "SUCCEEDED"
	StackSetOperationStatusFailed    StackSetOperationStatus = "FAILED"
	StackSetOperationStatusStopping  StackSetOperationStatus = "STOPPING"
	StackSetOperationStatusStopped   StackSetOperationStatus = "STOPPED"
)

func (enum StackSetOperationStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StackSetOperationStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StackSetStatus string

// Enum values for StackSetStatus
const (
	StackSetStatusActive  StackSetStatus = "ACTIVE"
	StackSetStatusDeleted StackSetStatus = "DELETED"
)

func (enum StackSetStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StackSetStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StackStatus string

// Enum values for StackStatus
const (
	StackStatusCreateInProgress                        StackStatus = "CREATE_IN_PROGRESS"
	StackStatusCreateFailed                            StackStatus = "CREATE_FAILED"
	StackStatusCreateComplete                          StackStatus = "CREATE_COMPLETE"
	StackStatusRollbackInProgress                      StackStatus = "ROLLBACK_IN_PROGRESS"
	StackStatusRollbackFailed                          StackStatus = "ROLLBACK_FAILED"
	StackStatusRollbackComplete                        StackStatus = "ROLLBACK_COMPLETE"
	StackStatusDeleteInProgress                        StackStatus = "DELETE_IN_PROGRESS"
	StackStatusDeleteFailed                            StackStatus = "DELETE_FAILED"
	StackStatusDeleteComplete                          StackStatus = "DELETE_COMPLETE"
	StackStatusUpdateInProgress                        StackStatus = "UPDATE_IN_PROGRESS"
	StackStatusUpdateCompleteCleanupInProgress         StackStatus = "UPDATE_COMPLETE_CLEANUP_IN_PROGRESS"
	StackStatusUpdateComplete                          StackStatus = "UPDATE_COMPLETE"
	StackStatusUpdateRollbackInProgress                StackStatus = "UPDATE_ROLLBACK_IN_PROGRESS"
	StackStatusUpdateRollbackFailed                    StackStatus = "UPDATE_ROLLBACK_FAILED"
	StackStatusUpdateRollbackCompleteCleanupInProgress StackStatus = "UPDATE_ROLLBACK_COMPLETE_CLEANUP_IN_PROGRESS"
	StackStatusUpdateRollbackComplete                  StackStatus = "UPDATE_ROLLBACK_COMPLETE"
	StackStatusReviewInProgress                        StackStatus = "REVIEW_IN_PROGRESS"
)

func (enum StackStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StackStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TemplateStage string

// Enum values for TemplateStage
const (
	TemplateStageOriginal  TemplateStage = "Original"
	TemplateStageProcessed TemplateStage = "Processed"
)

func (enum TemplateStage) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TemplateStage) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
