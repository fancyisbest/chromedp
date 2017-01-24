// Package runtime provides the Chrome Debugging Protocol
// commands, types, and events for the Chrome Runtime domain.
//
// Runtime domain exposes JavaScript runtime by means of remote evaluation
// and mirror objects. Evaluation results are returned as mirror object that
// expose object type, string representation and unique identifier that can be
// used for further object reference. Original objects are maintained in memory
// unless they are either explicitly released or are released along with the
// other objects in their object group.
//
// Generated by the chromedp-gen command.
package runtime

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	. "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
)

var (
	_ BackendNode
	_ BackendNodeID
	_ ComputedProperty
	_ ErrorType
	_ Frame
	_ FrameID
	_ LoaderID
	_ Message
	_ MessageError
	_ MethodType
	_ Node
	_ NodeID
	_ NodeType
	_ PseudoType
	_ RGBA
	_ ShadowRootType
	_ Timestamp
)

// EvaluateParams evaluates expression on global object.
type EvaluateParams struct {
	Expression            string             `json:"expression"`                      // Expression to evaluate.
	ObjectGroup           string             `json:"objectGroup,omitempty"`           // Symbolic group name that can be used to release multiple objects.
	IncludeCommandLineAPI bool               `json:"includeCommandLineAPI,omitempty"` // Determines whether Command Line API should be available during the evaluation.
	Silent                bool               `json:"silent,omitempty"`                // In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides setPauseOnException state.
	ContextID             ExecutionContextID `json:"contextId,omitempty"`             // Specifies in which execution context to perform evaluation. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ReturnByValue         bool               `json:"returnByValue,omitempty"`         // Whether the result is expected to be a JSON object that should be sent by value.
	GeneratePreview       bool               `json:"generatePreview,omitempty"`       // Whether preview should be generated for the result.
	UserGesture           bool               `json:"userGesture,omitempty"`           // Whether execution should be treated as initiated by user in the UI.
	AwaitPromise          bool               `json:"awaitPromise,omitempty"`          // Whether execution should wait for promise to be resolved. If the result of evaluation is not a Promise, it's considered to be an error.
}

// Evaluate evaluates expression on global object.
//
// parameters:
//   expression - Expression to evaluate.
func Evaluate(expression string) *EvaluateParams {
	return &EvaluateParams{
		Expression: expression,
	}
}

// WithObjectGroup symbolic group name that can be used to release multiple
// objects.
func (p EvaluateParams) WithObjectGroup(objectGroup string) *EvaluateParams {
	p.ObjectGroup = objectGroup
	return &p
}

// WithIncludeCommandLineAPI determines whether Command Line API should be
// available during the evaluation.
func (p EvaluateParams) WithIncludeCommandLineAPI(includeCommandLineAPI bool) *EvaluateParams {
	p.IncludeCommandLineAPI = includeCommandLineAPI
	return &p
}

// WithSilent in silent mode exceptions thrown during evaluation are not
// reported and do not pause execution. Overrides setPauseOnException state.
func (p EvaluateParams) WithSilent(silent bool) *EvaluateParams {
	p.Silent = silent
	return &p
}

// WithContextID specifies in which execution context to perform evaluation.
// If the parameter is omitted the evaluation will be performed in the context
// of the inspected page.
func (p EvaluateParams) WithContextID(contextId ExecutionContextID) *EvaluateParams {
	p.ContextID = contextId
	return &p
}

// WithReturnByValue whether the result is expected to be a JSON object that
// should be sent by value.
func (p EvaluateParams) WithReturnByValue(returnByValue bool) *EvaluateParams {
	p.ReturnByValue = returnByValue
	return &p
}

// WithGeneratePreview whether preview should be generated for the result.
func (p EvaluateParams) WithGeneratePreview(generatePreview bool) *EvaluateParams {
	p.GeneratePreview = generatePreview
	return &p
}

// WithUserGesture whether execution should be treated as initiated by user
// in the UI.
func (p EvaluateParams) WithUserGesture(userGesture bool) *EvaluateParams {
	p.UserGesture = userGesture
	return &p
}

// WithAwaitPromise whether execution should wait for promise to be resolved.
// If the result of evaluation is not a Promise, it's considered to be an error.
func (p EvaluateParams) WithAwaitPromise(awaitPromise bool) *EvaluateParams {
	p.AwaitPromise = awaitPromise
	return &p
}

// EvaluateReturns return values.
type EvaluateReturns struct {
	Result           *RemoteObject     `json:"result,omitempty"`           // Evaluation result.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details.
}

// Do executes Runtime.evaluate.
//
// returns:
//   result - Evaluation result.
//   exceptionDetails - Exception details.
func (p *EvaluateParams) Do(ctxt context.Context, h FrameHandler) (result *RemoteObject, exceptionDetails *ExceptionDetails, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return nil, nil, err
	}

	// execute
	ch := h.Execute(ctxt, CommandRuntimeEvaluate, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return nil, nil, ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r EvaluateReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return nil, nil, ErrInvalidResult
			}

			return r.Result, r.ExceptionDetails, nil

		case error:
			return nil, nil, v
		}

	case <-ctxt.Done():
		return nil, nil, ErrContextDone
	}

	return nil, nil, ErrUnknownResult
}

// AwaitPromiseParams add handler to promise with given promise object id.
type AwaitPromiseParams struct {
	PromiseObjectID RemoteObjectID `json:"promiseObjectId"`           // Identifier of the promise.
	ReturnByValue   bool           `json:"returnByValue,omitempty"`   // Whether the result is expected to be a JSON object that should be sent by value.
	GeneratePreview bool           `json:"generatePreview,omitempty"` // Whether preview should be generated for the result.
}

// AwaitPromise add handler to promise with given promise object id.
//
// parameters:
//   promiseObjectId - Identifier of the promise.
func AwaitPromise(promiseObjectId RemoteObjectID) *AwaitPromiseParams {
	return &AwaitPromiseParams{
		PromiseObjectID: promiseObjectId,
	}
}

// WithReturnByValue whether the result is expected to be a JSON object that
// should be sent by value.
func (p AwaitPromiseParams) WithReturnByValue(returnByValue bool) *AwaitPromiseParams {
	p.ReturnByValue = returnByValue
	return &p
}

// WithGeneratePreview whether preview should be generated for the result.
func (p AwaitPromiseParams) WithGeneratePreview(generatePreview bool) *AwaitPromiseParams {
	p.GeneratePreview = generatePreview
	return &p
}

// AwaitPromiseReturns return values.
type AwaitPromiseReturns struct {
	Result           *RemoteObject     `json:"result,omitempty"`           // Promise result. Will contain rejected value if promise was rejected.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details if stack strace is available.
}

// Do executes Runtime.awaitPromise.
//
// returns:
//   result - Promise result. Will contain rejected value if promise was rejected.
//   exceptionDetails - Exception details if stack strace is available.
func (p *AwaitPromiseParams) Do(ctxt context.Context, h FrameHandler) (result *RemoteObject, exceptionDetails *ExceptionDetails, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return nil, nil, err
	}

	// execute
	ch := h.Execute(ctxt, CommandRuntimeAwaitPromise, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return nil, nil, ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r AwaitPromiseReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return nil, nil, ErrInvalidResult
			}

			return r.Result, r.ExceptionDetails, nil

		case error:
			return nil, nil, v
		}

	case <-ctxt.Done():
		return nil, nil, ErrContextDone
	}

	return nil, nil, ErrUnknownResult
}

// CallFunctionOnParams calls function with given declaration on the given
// object. Object group of the result is inherited from the target object.
type CallFunctionOnParams struct {
	ObjectID            RemoteObjectID  `json:"objectId"`                  // Identifier of the object to call function on.
	FunctionDeclaration string          `json:"functionDeclaration"`       // Declaration of the function to call.
	Arguments           []*CallArgument `json:"arguments,omitempty"`       // Call arguments. All call arguments must belong to the same JavaScript world as the target object.
	Silent              bool            `json:"silent,omitempty"`          // In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides setPauseOnException state.
	ReturnByValue       bool            `json:"returnByValue,omitempty"`   // Whether the result is expected to be a JSON object which should be sent by value.
	GeneratePreview     bool            `json:"generatePreview,omitempty"` // Whether preview should be generated for the result.
	UserGesture         bool            `json:"userGesture,omitempty"`     // Whether execution should be treated as initiated by user in the UI.
	AwaitPromise        bool            `json:"awaitPromise,omitempty"`    // Whether execution should wait for promise to be resolved. If the result of evaluation is not a Promise, it's considered to be an error.
}

// CallFunctionOn calls function with given declaration on the given object.
// Object group of the result is inherited from the target object.
//
// parameters:
//   objectId - Identifier of the object to call function on.
//   functionDeclaration - Declaration of the function to call.
func CallFunctionOn(objectId RemoteObjectID, functionDeclaration string) *CallFunctionOnParams {
	return &CallFunctionOnParams{
		ObjectID:            objectId,
		FunctionDeclaration: functionDeclaration,
	}
}

// WithArguments call arguments. All call arguments must belong to the same
// JavaScript world as the target object.
func (p CallFunctionOnParams) WithArguments(arguments []*CallArgument) *CallFunctionOnParams {
	p.Arguments = arguments
	return &p
}

// WithSilent in silent mode exceptions thrown during evaluation are not
// reported and do not pause execution. Overrides setPauseOnException state.
func (p CallFunctionOnParams) WithSilent(silent bool) *CallFunctionOnParams {
	p.Silent = silent
	return &p
}

// WithReturnByValue whether the result is expected to be a JSON object which
// should be sent by value.
func (p CallFunctionOnParams) WithReturnByValue(returnByValue bool) *CallFunctionOnParams {
	p.ReturnByValue = returnByValue
	return &p
}

// WithGeneratePreview whether preview should be generated for the result.
func (p CallFunctionOnParams) WithGeneratePreview(generatePreview bool) *CallFunctionOnParams {
	p.GeneratePreview = generatePreview
	return &p
}

// WithUserGesture whether execution should be treated as initiated by user
// in the UI.
func (p CallFunctionOnParams) WithUserGesture(userGesture bool) *CallFunctionOnParams {
	p.UserGesture = userGesture
	return &p
}

// WithAwaitPromise whether execution should wait for promise to be resolved.
// If the result of evaluation is not a Promise, it's considered to be an error.
func (p CallFunctionOnParams) WithAwaitPromise(awaitPromise bool) *CallFunctionOnParams {
	p.AwaitPromise = awaitPromise
	return &p
}

// CallFunctionOnReturns return values.
type CallFunctionOnReturns struct {
	Result           *RemoteObject     `json:"result,omitempty"`           // Call result.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details.
}

// Do executes Runtime.callFunctionOn.
//
// returns:
//   result - Call result.
//   exceptionDetails - Exception details.
func (p *CallFunctionOnParams) Do(ctxt context.Context, h FrameHandler) (result *RemoteObject, exceptionDetails *ExceptionDetails, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return nil, nil, err
	}

	// execute
	ch := h.Execute(ctxt, CommandRuntimeCallFunctionOn, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return nil, nil, ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r CallFunctionOnReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return nil, nil, ErrInvalidResult
			}

			return r.Result, r.ExceptionDetails, nil

		case error:
			return nil, nil, v
		}

	case <-ctxt.Done():
		return nil, nil, ErrContextDone
	}

	return nil, nil, ErrUnknownResult
}

// GetPropertiesParams returns properties of a given object. Object group of
// the result is inherited from the target object.
type GetPropertiesParams struct {
	ObjectID               RemoteObjectID `json:"objectId"`                         // Identifier of the object to return properties for.
	OwnProperties          bool           `json:"ownProperties,omitempty"`          // If true, returns properties belonging only to the element itself, not to its prototype chain.
	AccessorPropertiesOnly bool           `json:"accessorPropertiesOnly,omitempty"` // If true, returns accessor properties (with getter/setter) only; internal properties are not returned either.
	GeneratePreview        bool           `json:"generatePreview,omitempty"`        // Whether preview should be generated for the results.
}

// GetProperties returns properties of a given object. Object group of the
// result is inherited from the target object.
//
// parameters:
//   objectId - Identifier of the object to return properties for.
func GetProperties(objectId RemoteObjectID) *GetPropertiesParams {
	return &GetPropertiesParams{
		ObjectID: objectId,
	}
}

// WithOwnProperties if true, returns properties belonging only to the
// element itself, not to its prototype chain.
func (p GetPropertiesParams) WithOwnProperties(ownProperties bool) *GetPropertiesParams {
	p.OwnProperties = ownProperties
	return &p
}

// WithAccessorPropertiesOnly if true, returns accessor properties (with
// getter/setter) only; internal properties are not returned either.
func (p GetPropertiesParams) WithAccessorPropertiesOnly(accessorPropertiesOnly bool) *GetPropertiesParams {
	p.AccessorPropertiesOnly = accessorPropertiesOnly
	return &p
}

// WithGeneratePreview whether preview should be generated for the results.
func (p GetPropertiesParams) WithGeneratePreview(generatePreview bool) *GetPropertiesParams {
	p.GeneratePreview = generatePreview
	return &p
}

// GetPropertiesReturns return values.
type GetPropertiesReturns struct {
	Result             []*PropertyDescriptor         `json:"result,omitempty"`             // Object properties.
	InternalProperties []*InternalPropertyDescriptor `json:"internalProperties,omitempty"` // Internal object properties (only of the element itself).
	ExceptionDetails   *ExceptionDetails             `json:"exceptionDetails,omitempty"`   // Exception details.
}

// Do executes Runtime.getProperties.
//
// returns:
//   result - Object properties.
//   internalProperties - Internal object properties (only of the element itself).
//   exceptionDetails - Exception details.
func (p *GetPropertiesParams) Do(ctxt context.Context, h FrameHandler) (result []*PropertyDescriptor, internalProperties []*InternalPropertyDescriptor, exceptionDetails *ExceptionDetails, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return nil, nil, nil, err
	}

	// execute
	ch := h.Execute(ctxt, CommandRuntimeGetProperties, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return nil, nil, nil, ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r GetPropertiesReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return nil, nil, nil, ErrInvalidResult
			}

			return r.Result, r.InternalProperties, r.ExceptionDetails, nil

		case error:
			return nil, nil, nil, v
		}

	case <-ctxt.Done():
		return nil, nil, nil, ErrContextDone
	}

	return nil, nil, nil, ErrUnknownResult
}

// ReleaseObjectParams releases remote object with given id.
type ReleaseObjectParams struct {
	ObjectID RemoteObjectID `json:"objectId"` // Identifier of the object to release.
}

// ReleaseObject releases remote object with given id.
//
// parameters:
//   objectId - Identifier of the object to release.
func ReleaseObject(objectId RemoteObjectID) *ReleaseObjectParams {
	return &ReleaseObjectParams{
		ObjectID: objectId,
	}
}

// Do executes Runtime.releaseObject.
func (p *ReleaseObjectParams) Do(ctxt context.Context, h FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, CommandRuntimeReleaseObject, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ErrContextDone
	}

	return ErrUnknownResult
}

// ReleaseObjectGroupParams releases all remote objects that belong to a
// given group.
type ReleaseObjectGroupParams struct {
	ObjectGroup string `json:"objectGroup"` // Symbolic object group name.
}

// ReleaseObjectGroup releases all remote objects that belong to a given
// group.
//
// parameters:
//   objectGroup - Symbolic object group name.
func ReleaseObjectGroup(objectGroup string) *ReleaseObjectGroupParams {
	return &ReleaseObjectGroupParams{
		ObjectGroup: objectGroup,
	}
}

// Do executes Runtime.releaseObjectGroup.
func (p *ReleaseObjectGroupParams) Do(ctxt context.Context, h FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, CommandRuntimeReleaseObjectGroup, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ErrContextDone
	}

	return ErrUnknownResult
}

// RunIfWaitingForDebuggerParams tells inspected instance to run if it was
// waiting for debugger to attach.
type RunIfWaitingForDebuggerParams struct{}

// RunIfWaitingForDebugger tells inspected instance to run if it was waiting
// for debugger to attach.
func RunIfWaitingForDebugger() *RunIfWaitingForDebuggerParams {
	return &RunIfWaitingForDebuggerParams{}
}

// Do executes Runtime.runIfWaitingForDebugger.
func (p *RunIfWaitingForDebuggerParams) Do(ctxt context.Context, h FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, CommandRuntimeRunIfWaitingForDebugger, Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ErrContextDone
	}

	return ErrUnknownResult
}

// EnableParams enables reporting of execution contexts creation by means of
// executionContextCreated event. When the reporting gets enabled the event will
// be sent immediately for each existing execution context.
type EnableParams struct{}

// Enable enables reporting of execution contexts creation by means of
// executionContextCreated event. When the reporting gets enabled the event will
// be sent immediately for each existing execution context.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Runtime.enable.
func (p *EnableParams) Do(ctxt context.Context, h FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, CommandRuntimeEnable, Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ErrContextDone
	}

	return ErrUnknownResult
}

// DisableParams disables reporting of execution contexts creation.
type DisableParams struct{}

// Disable disables reporting of execution contexts creation.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Runtime.disable.
func (p *DisableParams) Do(ctxt context.Context, h FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, CommandRuntimeDisable, Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ErrContextDone
	}

	return ErrUnknownResult
}

// DiscardConsoleEntriesParams discards collected exceptions and console API
// calls.
type DiscardConsoleEntriesParams struct{}

// DiscardConsoleEntries discards collected exceptions and console API calls.
func DiscardConsoleEntries() *DiscardConsoleEntriesParams {
	return &DiscardConsoleEntriesParams{}
}

// Do executes Runtime.discardConsoleEntries.
func (p *DiscardConsoleEntriesParams) Do(ctxt context.Context, h FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, CommandRuntimeDiscardConsoleEntries, Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ErrContextDone
	}

	return ErrUnknownResult
}

type SetCustomObjectFormatterEnabledParams struct {
	Enabled bool `json:"enabled"`
}

// parameters:
//   enabled
func SetCustomObjectFormatterEnabled(enabled bool) *SetCustomObjectFormatterEnabledParams {
	return &SetCustomObjectFormatterEnabledParams{
		Enabled: enabled,
	}
}

// Do executes Runtime.setCustomObjectFormatterEnabled.
func (p *SetCustomObjectFormatterEnabledParams) Do(ctxt context.Context, h FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, CommandRuntimeSetCustomObjectFormatterEnabled, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return ErrContextDone
	}

	return ErrUnknownResult
}

// CompileScriptParams compiles expression.
type CompileScriptParams struct {
	Expression         string             `json:"expression"`                   // Expression to compile.
	SourceURL          string             `json:"sourceURL"`                    // Source url to be set for the script.
	PersistScript      bool               `json:"persistScript"`                // Specifies whether the compiled script should be persisted.
	ExecutionContextID ExecutionContextID `json:"executionContextId,omitempty"` // Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
}

// CompileScript compiles expression.
//
// parameters:
//   expression - Expression to compile.
//   sourceURL - Source url to be set for the script.
//   persistScript - Specifies whether the compiled script should be persisted.
func CompileScript(expression string, sourceURL string, persistScript bool) *CompileScriptParams {
	return &CompileScriptParams{
		Expression:    expression,
		SourceURL:     sourceURL,
		PersistScript: persistScript,
	}
}

// WithExecutionContextID specifies in which execution context to perform
// script run. If the parameter is omitted the evaluation will be performed in
// the context of the inspected page.
func (p CompileScriptParams) WithExecutionContextID(executionContextId ExecutionContextID) *CompileScriptParams {
	p.ExecutionContextID = executionContextId
	return &p
}

// CompileScriptReturns return values.
type CompileScriptReturns struct {
	ScriptID         ScriptID          `json:"scriptId,omitempty"`         // Id of the script.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details.
}

// Do executes Runtime.compileScript.
//
// returns:
//   scriptId - Id of the script.
//   exceptionDetails - Exception details.
func (p *CompileScriptParams) Do(ctxt context.Context, h FrameHandler) (scriptId ScriptID, exceptionDetails *ExceptionDetails, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return "", nil, err
	}

	// execute
	ch := h.Execute(ctxt, CommandRuntimeCompileScript, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return "", nil, ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r CompileScriptReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return "", nil, ErrInvalidResult
			}

			return r.ScriptID, r.ExceptionDetails, nil

		case error:
			return "", nil, v
		}

	case <-ctxt.Done():
		return "", nil, ErrContextDone
	}

	return "", nil, ErrUnknownResult
}

// RunScriptParams runs script with given id in a given context.
type RunScriptParams struct {
	ScriptID              ScriptID           `json:"scriptId"`                        // Id of the script to run.
	ExecutionContextID    ExecutionContextID `json:"executionContextId,omitempty"`    // Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ObjectGroup           string             `json:"objectGroup,omitempty"`           // Symbolic group name that can be used to release multiple objects.
	Silent                bool               `json:"silent,omitempty"`                // In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides setPauseOnException state.
	IncludeCommandLineAPI bool               `json:"includeCommandLineAPI,omitempty"` // Determines whether Command Line API should be available during the evaluation.
	ReturnByValue         bool               `json:"returnByValue,omitempty"`         // Whether the result is expected to be a JSON object which should be sent by value.
	GeneratePreview       bool               `json:"generatePreview,omitempty"`       // Whether preview should be generated for the result.
	AwaitPromise          bool               `json:"awaitPromise,omitempty"`          // Whether execution should wait for promise to be resolved. If the result of evaluation is not a Promise, it's considered to be an error.
}

// RunScript runs script with given id in a given context.
//
// parameters:
//   scriptId - Id of the script to run.
func RunScript(scriptId ScriptID) *RunScriptParams {
	return &RunScriptParams{
		ScriptID: scriptId,
	}
}

// WithExecutionContextID specifies in which execution context to perform
// script run. If the parameter is omitted the evaluation will be performed in
// the context of the inspected page.
func (p RunScriptParams) WithExecutionContextID(executionContextId ExecutionContextID) *RunScriptParams {
	p.ExecutionContextID = executionContextId
	return &p
}

// WithObjectGroup symbolic group name that can be used to release multiple
// objects.
func (p RunScriptParams) WithObjectGroup(objectGroup string) *RunScriptParams {
	p.ObjectGroup = objectGroup
	return &p
}

// WithSilent in silent mode exceptions thrown during evaluation are not
// reported and do not pause execution. Overrides setPauseOnException state.
func (p RunScriptParams) WithSilent(silent bool) *RunScriptParams {
	p.Silent = silent
	return &p
}

// WithIncludeCommandLineAPI determines whether Command Line API should be
// available during the evaluation.
func (p RunScriptParams) WithIncludeCommandLineAPI(includeCommandLineAPI bool) *RunScriptParams {
	p.IncludeCommandLineAPI = includeCommandLineAPI
	return &p
}

// WithReturnByValue whether the result is expected to be a JSON object which
// should be sent by value.
func (p RunScriptParams) WithReturnByValue(returnByValue bool) *RunScriptParams {
	p.ReturnByValue = returnByValue
	return &p
}

// WithGeneratePreview whether preview should be generated for the result.
func (p RunScriptParams) WithGeneratePreview(generatePreview bool) *RunScriptParams {
	p.GeneratePreview = generatePreview
	return &p
}

// WithAwaitPromise whether execution should wait for promise to be resolved.
// If the result of evaluation is not a Promise, it's considered to be an error.
func (p RunScriptParams) WithAwaitPromise(awaitPromise bool) *RunScriptParams {
	p.AwaitPromise = awaitPromise
	return &p
}

// RunScriptReturns return values.
type RunScriptReturns struct {
	Result           *RemoteObject     `json:"result,omitempty"`           // Run result.
	ExceptionDetails *ExceptionDetails `json:"exceptionDetails,omitempty"` // Exception details.
}

// Do executes Runtime.runScript.
//
// returns:
//   result - Run result.
//   exceptionDetails - Exception details.
func (p *RunScriptParams) Do(ctxt context.Context, h FrameHandler) (result *RemoteObject, exceptionDetails *ExceptionDetails, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return nil, nil, err
	}

	// execute
	ch := h.Execute(ctxt, CommandRuntimeRunScript, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return nil, nil, ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r RunScriptReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return nil, nil, ErrInvalidResult
			}

			return r.Result, r.ExceptionDetails, nil

		case error:
			return nil, nil, v
		}

	case <-ctxt.Done():
		return nil, nil, ErrContextDone
	}

	return nil, nil, ErrUnknownResult
}
