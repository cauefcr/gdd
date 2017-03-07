package serviceworker


import (
    "../shared"
)

type EnableCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type EnableReturn struct {
}

func (c *EnableCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *EnableCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type DisableCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
}
type DisableReturn struct {
}

func (c *DisableCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DisableCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type UnregisterCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ScopeURL string `json:"scopeURL"`
}
type UnregisterReturn struct {
}

func (c *UnregisterCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *UnregisterCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type UpdateRegistrationCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ScopeURL string `json:"scopeURL"`
}
type UpdateRegistrationReturn struct {
}

func (c *UpdateRegistrationCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *UpdateRegistrationCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type StartWorkerCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ScopeURL string `json:"scopeURL"`
}
type StartWorkerReturn struct {
}

func (c *StartWorkerCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StartWorkerCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SkipWaitingCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ScopeURL string `json:"scopeURL"`
}
type SkipWaitingReturn struct {
}

func (c *SkipWaitingCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SkipWaitingCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type StopWorkerCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    VersionId string `json:"versionId"`
}
type StopWorkerReturn struct {
}

func (c *StopWorkerCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *StopWorkerCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type InspectWorkerCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    VersionId string `json:"versionId"`
}
type InspectWorkerReturn struct {
}

func (c *InspectWorkerCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *InspectWorkerCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type SetForceUpdateOnPageLoadCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    ForceUpdateOnPageLoad bool `json:"forceUpdateOnPageLoad"`
}
type SetForceUpdateOnPageLoadReturn struct {
}

func (c *SetForceUpdateOnPageLoadCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *SetForceUpdateOnPageLoadCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type DeliverPushMessageCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Origin string `json:"origin"`
    RegistrationId string `json:"registrationId"`
    Data string `json:"data"`
}
type DeliverPushMessageReturn struct {
}

func (c *DeliverPushMessageCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DeliverPushMessageCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}

type DispatchSyncEventCommand struct {
    DestinationTargetID string
    responseId int64
    conn *shared.Connection
    Origin string `json:"origin"`
    RegistrationId string `json:"registrationId"`
    Tag string `json:"tag"`
    LastChance bool `json:"lastChance"`
}
type DispatchSyncEventReturn struct {
}

func (c *DispatchSyncEventCommand) RespondWithError(code shared.ResponseErrorCodes, message string) {
    c.conn.SendErrorResult(c.responseId, c.DestinationTargetID, code, message)
}

func (c *DispatchSyncEventCommand) Respond() {
    c.conn.SendResult(c.responseId, c.DestinationTargetID, struct{}{})
}
