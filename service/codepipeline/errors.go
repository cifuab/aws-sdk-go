// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codepipeline

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeActionNotFoundException for service response error code
	// "ActionNotFoundException".
	//
	// The specified action cannot be found.
	ErrCodeActionNotFoundException = "ActionNotFoundException"

	// ErrCodeActionTypeNotFoundException for service response error code
	// "ActionTypeNotFoundException".
	//
	// The specified action type cannot be found.
	ErrCodeActionTypeNotFoundException = "ActionTypeNotFoundException"

	// ErrCodeApprovalAlreadyCompletedException for service response error code
	// "ApprovalAlreadyCompletedException".
	//
	// The approval action has already been approved or rejected.
	ErrCodeApprovalAlreadyCompletedException = "ApprovalAlreadyCompletedException"

	// ErrCodeConcurrentModificationException for service response error code
	// "ConcurrentModificationException".
	//
	// Unable to modify the tag due to a simultaneous update request.
	ErrCodeConcurrentModificationException = "ConcurrentModificationException"

	// ErrCodeConcurrentPipelineExecutionsLimitExceededException for service response error code
	// "ConcurrentPipelineExecutionsLimitExceededException".
	//
	// The pipeline has reached the limit for concurrent pipeline executions.
	ErrCodeConcurrentPipelineExecutionsLimitExceededException = "ConcurrentPipelineExecutionsLimitExceededException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// Your request cannot be handled because the pipeline is busy handling ongoing
	// activities. Try again later.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeDuplicatedStopRequestException for service response error code
	// "DuplicatedStopRequestException".
	//
	// The pipeline execution is already in a Stopping state. If you already chose
	// to stop and wait, you cannot make that request again. You can choose to stop
	// and abandon now, but be aware that this option can lead to failed tasks or
	// out of sequence tasks. If you already chose to stop and abandon, you cannot
	// make that request again.
	ErrCodeDuplicatedStopRequestException = "DuplicatedStopRequestException"

	// ErrCodeInvalidActionDeclarationException for service response error code
	// "InvalidActionDeclarationException".
	//
	// The action declaration was specified in an invalid format.
	ErrCodeInvalidActionDeclarationException = "InvalidActionDeclarationException"

	// ErrCodeInvalidApprovalTokenException for service response error code
	// "InvalidApprovalTokenException".
	//
	// The approval request already received a response or has expired.
	ErrCodeInvalidApprovalTokenException = "InvalidApprovalTokenException"

	// ErrCodeInvalidArnException for service response error code
	// "InvalidArnException".
	//
	// The specified resource ARN is invalid.
	ErrCodeInvalidArnException = "InvalidArnException"

	// ErrCodeInvalidBlockerDeclarationException for service response error code
	// "InvalidBlockerDeclarationException".
	//
	// Reserved for future use.
	ErrCodeInvalidBlockerDeclarationException = "InvalidBlockerDeclarationException"

	// ErrCodeInvalidClientTokenException for service response error code
	// "InvalidClientTokenException".
	//
	// The client token was specified in an invalid format
	ErrCodeInvalidClientTokenException = "InvalidClientTokenException"

	// ErrCodeInvalidJobException for service response error code
	// "InvalidJobException".
	//
	// The job was specified in an invalid format or cannot be found.
	ErrCodeInvalidJobException = "InvalidJobException"

	// ErrCodeInvalidJobStateException for service response error code
	// "InvalidJobStateException".
	//
	// The job state was specified in an invalid format.
	ErrCodeInvalidJobStateException = "InvalidJobStateException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The next token was specified in an invalid format. Make sure that the next
	// token you provide is the token returned by a previous call.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidNonceException for service response error code
	// "InvalidNonceException".
	//
	// The nonce was specified in an invalid format.
	ErrCodeInvalidNonceException = "InvalidNonceException"

	// ErrCodeInvalidStageDeclarationException for service response error code
	// "InvalidStageDeclarationException".
	//
	// The stage declaration was specified in an invalid format.
	ErrCodeInvalidStageDeclarationException = "InvalidStageDeclarationException"

	// ErrCodeInvalidStructureException for service response error code
	// "InvalidStructureException".
	//
	// The structure was specified in an invalid format.
	ErrCodeInvalidStructureException = "InvalidStructureException"

	// ErrCodeInvalidTagsException for service response error code
	// "InvalidTagsException".
	//
	// The specified resource tags are invalid.
	ErrCodeInvalidTagsException = "InvalidTagsException"

	// ErrCodeInvalidWebhookAuthenticationParametersException for service response error code
	// "InvalidWebhookAuthenticationParametersException".
	//
	// The specified authentication type is in an invalid format.
	ErrCodeInvalidWebhookAuthenticationParametersException = "InvalidWebhookAuthenticationParametersException"

	// ErrCodeInvalidWebhookFilterPatternException for service response error code
	// "InvalidWebhookFilterPatternException".
	//
	// The specified event filter rule is in an invalid format.
	ErrCodeInvalidWebhookFilterPatternException = "InvalidWebhookFilterPatternException"

	// ErrCodeJobNotFoundException for service response error code
	// "JobNotFoundException".
	//
	// The job was specified in an invalid format or cannot be found.
	ErrCodeJobNotFoundException = "JobNotFoundException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The number of pipelines associated with the Amazon Web Services account has
	// exceeded the limit allowed for the account.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeNotLatestPipelineExecutionException for service response error code
	// "NotLatestPipelineExecutionException".
	//
	// The stage has failed in a later run of the pipeline and the pipelineExecutionId
	// associated with the request is out of date.
	ErrCodeNotLatestPipelineExecutionException = "NotLatestPipelineExecutionException"

	// ErrCodeOutputVariablesSizeExceededException for service response error code
	// "OutputVariablesSizeExceededException".
	//
	// Exceeded the total size limit for all variables in the pipeline.
	ErrCodeOutputVariablesSizeExceededException = "OutputVariablesSizeExceededException"

	// ErrCodePipelineExecutionNotFoundException for service response error code
	// "PipelineExecutionNotFoundException".
	//
	// The pipeline execution was specified in an invalid format or cannot be found,
	// or an execution ID does not belong to the specified pipeline.
	ErrCodePipelineExecutionNotFoundException = "PipelineExecutionNotFoundException"

	// ErrCodePipelineExecutionNotStoppableException for service response error code
	// "PipelineExecutionNotStoppableException".
	//
	// Unable to stop the pipeline execution. The execution might already be in
	// a Stopped state, or it might no longer be in progress.
	ErrCodePipelineExecutionNotStoppableException = "PipelineExecutionNotStoppableException"

	// ErrCodePipelineNameInUseException for service response error code
	// "PipelineNameInUseException".
	//
	// The specified pipeline name is already in use.
	ErrCodePipelineNameInUseException = "PipelineNameInUseException"

	// ErrCodePipelineNotFoundException for service response error code
	// "PipelineNotFoundException".
	//
	// The pipeline was specified in an invalid format or cannot be found.
	ErrCodePipelineNotFoundException = "PipelineNotFoundException"

	// ErrCodePipelineVersionNotFoundException for service response error code
	// "PipelineVersionNotFoundException".
	//
	// The pipeline version was specified in an invalid format or cannot be found.
	ErrCodePipelineVersionNotFoundException = "PipelineVersionNotFoundException"

	// ErrCodeRequestFailedException for service response error code
	// "RequestFailedException".
	//
	// The request failed because of an unknown error, exception, or failure.
	ErrCodeRequestFailedException = "RequestFailedException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The resource was specified in an invalid format.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeStageNotFoundException for service response error code
	// "StageNotFoundException".
	//
	// The stage was specified in an invalid format or cannot be found.
	ErrCodeStageNotFoundException = "StageNotFoundException"

	// ErrCodeStageNotRetryableException for service response error code
	// "StageNotRetryableException".
	//
	// Unable to retry. The pipeline structure or stage state might have changed
	// while actions awaited retry, or the stage contains no failed actions.
	ErrCodeStageNotRetryableException = "StageNotRetryableException"

	// ErrCodeTooManyTagsException for service response error code
	// "TooManyTagsException".
	//
	// The tags limit for a resource has been exceeded.
	ErrCodeTooManyTagsException = "TooManyTagsException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The validation was specified in an invalid format.
	ErrCodeValidationException = "ValidationException"

	// ErrCodeWebhookNotFoundException for service response error code
	// "WebhookNotFoundException".
	//
	// The specified webhook was entered in an invalid format or cannot be found.
	ErrCodeWebhookNotFoundException = "WebhookNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ActionNotFoundException":                            newErrorActionNotFoundException,
	"ActionTypeNotFoundException":                        newErrorActionTypeNotFoundException,
	"ApprovalAlreadyCompletedException":                  newErrorApprovalAlreadyCompletedException,
	"ConcurrentModificationException":                    newErrorConcurrentModificationException,
	"ConcurrentPipelineExecutionsLimitExceededException": newErrorConcurrentPipelineExecutionsLimitExceededException,
	"ConflictException":                                  newErrorConflictException,
	"DuplicatedStopRequestException":                     newErrorDuplicatedStopRequestException,
	"InvalidActionDeclarationException":                  newErrorInvalidActionDeclarationException,
	"InvalidApprovalTokenException":                      newErrorInvalidApprovalTokenException,
	"InvalidArnException":                                newErrorInvalidArnException,
	"InvalidBlockerDeclarationException":                 newErrorInvalidBlockerDeclarationException,
	"InvalidClientTokenException":                        newErrorInvalidClientTokenException,
	"InvalidJobException":                                newErrorInvalidJobException,
	"InvalidJobStateException":                           newErrorInvalidJobStateException,
	"InvalidNextTokenException":                          newErrorInvalidNextTokenException,
	"InvalidNonceException":                              newErrorInvalidNonceException,
	"InvalidStageDeclarationException":                   newErrorInvalidStageDeclarationException,
	"InvalidStructureException":                          newErrorInvalidStructureException,
	"InvalidTagsException":                               newErrorInvalidTagsException,
	"InvalidWebhookAuthenticationParametersException":    newErrorInvalidWebhookAuthenticationParametersException,
	"InvalidWebhookFilterPatternException":               newErrorInvalidWebhookFilterPatternException,
	"JobNotFoundException":                               newErrorJobNotFoundException,
	"LimitExceededException":                             newErrorLimitExceededException,
	"NotLatestPipelineExecutionException":                newErrorNotLatestPipelineExecutionException,
	"OutputVariablesSizeExceededException":               newErrorOutputVariablesSizeExceededException,
	"PipelineExecutionNotFoundException":                 newErrorPipelineExecutionNotFoundException,
	"PipelineExecutionNotStoppableException":             newErrorPipelineExecutionNotStoppableException,
	"PipelineNameInUseException":                         newErrorPipelineNameInUseException,
	"PipelineNotFoundException":                          newErrorPipelineNotFoundException,
	"PipelineVersionNotFoundException":                   newErrorPipelineVersionNotFoundException,
	"RequestFailedException":                             newErrorRequestFailedException,
	"ResourceNotFoundException":                          newErrorResourceNotFoundException,
	"StageNotFoundException":                             newErrorStageNotFoundException,
	"StageNotRetryableException":                         newErrorStageNotRetryableException,
	"TooManyTagsException":                               newErrorTooManyTagsException,
	"ValidationException":                                newErrorValidationException,
	"WebhookNotFoundException":                           newErrorWebhookNotFoundException,
}
