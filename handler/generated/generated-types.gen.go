// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package generated

import (
	"encoding/json"

	"github.com/oapi-codegen/runtime"
)

// ActionMessage defines model for ActionMessage.
type ActionMessage struct {
	Body      *string `json:"body,omitempty"`
	Id        *string `json:"id,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	Type      *string `json:"type,omitempty"`
}

// AnonymousUser defines model for AnonymousUser.
type AnonymousUser struct {
	AccessToken *string `json:"accessToken,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Id          *string `json:"id,omitempty"`
}

// AuthenticationConfig defines model for AuthenticationConfig.
type AuthenticationConfig struct {
	IndieAuthEnabled *bool `json:"indieAuthEnabled,omitempty"`
}

// BaseAPIResponse Simple API response
type BaseAPIResponse struct {
	Message *string `json:"message,omitempty"`
	Success *bool   `json:"success,omitempty"`
}

// BrowserConfig defines model for BrowserConfig.
type BrowserConfig struct {
	Enabled   *bool   `json:"enabled,omitempty"`
	PublicKey *string `json:"publicKey,omitempty"`
}

// ChatMessages defines model for ChatMessages.
type ChatMessages = []ChatMessages_Item

// ChatMessages_Item defines model for ChatMessages.Item.
type ChatMessages_Item struct {
	union json.RawMessage
}

// Emoji Name and url for an emoji
type Emoji struct {
	// Name The name of the emoji
	Name *string `json:"name,omitempty"`

	// Url URL for the emoji image
	Url *string `json:"url,omitempty"`
}

// Emojis defines model for Emojis.
type Emojis = []Emoji

// Error Structure for an error response
type Error struct {
	Error *string `json:"error,omitempty"`
}

// Event defines model for Event.
type Event struct {
	Id        *string `json:"id,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	Type      *string `json:"type,omitempty"`
}

// ExternalAction defines model for ExternalAction.
type ExternalAction struct {
	Color          *string `json:"color,omitempty"`
	Description    *string `json:"description,omitempty"`
	Html           *string `json:"html,omitempty"`
	Icon           *string `json:"icon,omitempty"`
	OpenExternally *bool   `json:"openExternally,omitempty"`
	Title          *string `json:"title,omitempty"`
	Url            *string `json:"url,omitempty"`
}

// FederatedAction defines model for FederatedAction.
type FederatedAction struct {
	Body      *string `json:"body,omitempty"`
	Id        *string `json:"id,omitempty"`
	Image     *string `json:"image,omitempty"`
	Link      *string `json:"link,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	Title     *string `json:"title,omitempty"`
	Type      *string `json:"type,omitempty"`
}

// FederationConfig defines model for FederationConfig.
type FederationConfig struct {
	Account       *string `json:"account,omitempty"`
	Enabled       *bool   `json:"enabled,omitempty"`
	FollowerCount *int    `json:"followerCount,omitempty"`
}

// Follower defines model for Follower.
type Follower struct {
	// DisabledAt When this follower was rejected or disabled.
	DisabledAt *string `json:"disabledAt,omitempty"`

	// Image The avatar image of the follower.
	Image *string `json:"image,omitempty"`

	// Link The IRI of the remote actor.
	Link *string `json:"link,omitempty"`

	// Name The display name of the follower.
	Name *string `json:"name,omitempty"`

	// Timestamp When this follow request was created.
	Timestamp *string `json:"timestamp,omitempty"`

	// Username The account username of the remote actor.
	Username *string `json:"username,omitempty"`
}

// MessageEvent defines model for MessageEvent.
type MessageEvent struct {
	Body *string `json:"body,omitempty"`
}

// NotificationConfig defines model for NotificationConfig.
type NotificationConfig struct {
	Browser *BrowserConfig `json:"browser,omitempty"`
}

// PlaybackMetrics defines model for PlaybackMetrics.
type PlaybackMetrics struct {
	Bandwidth             *float64 `json:"bandwidth,omitempty"`
	DownloadDuration      *float64 `json:"downloadDuration,omitempty"`
	Errors                *float64 `json:"errors,omitempty"`
	Latency               *float64 `json:"latency,omitempty"`
	QualityVariantChanges *float64 `json:"qualityVariantChanges,omitempty"`
}

// SocialHandle defines model for SocialHandle.
type SocialHandle struct {
	Icon     *string `json:"icon,omitempty"`
	Platform *string `json:"platform,omitempty"`
	Url      *string `json:"url,omitempty"`
}

// Status Response for status
type Status struct {
	LastConnectTime    *string `json:"lastConnectTime,omitempty"`
	LastDisconnectTime *string `json:"lastDisconnectTime,omitempty"`
	Online             *bool   `json:"online,omitempty"`
	ServerTime         *string `json:"serverTime,omitempty"`
	VersionNumber      *string `json:"versionNumber,omitempty"`
	ViewerCount        *int    `json:"viewerCount,omitempty"`
}

// SystemMessage defines model for SystemMessage.
type SystemMessage struct {
	Body      *string `json:"body,omitempty"`
	Id        *string `json:"id,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	Type      *string `json:"type,omitempty"`
}

// User defines model for User.
type User struct {
	Authenticated *bool     `json:"authenticated,omitempty"`
	CreatedAt     *string   `json:"createdAt,omitempty"`
	DisabledAt    *string   `json:"disabledAt,omitempty"`
	DisplayColor  *int      `json:"displayColor,omitempty"`
	DisplayName   *string   `json:"displayName,omitempty"`
	Id            *string   `json:"id,omitempty"`
	IsBot         *bool     `json:"isBot,omitempty"`
	NameChangedAt *string   `json:"nameChangedAt,omitempty"`
	PreviousNames *[]string `json:"previousNames,omitempty"`
	Scopes        *[]string `json:"scopes,omitempty"`
}

// UserEvent defines model for UserEvent.
type UserEvent struct {
	ClientId *int    `json:"clientId,omitempty"`
	HiddenAt *string `json:"hiddenAt,omitempty"`
	User     *User   `json:"user,omitempty"`
}

// UserMessage defines model for UserMessage.
type UserMessage struct {
	Body      *string `json:"body,omitempty"`
	ClientId  *int    `json:"clientId,omitempty"`
	HiddenAt  *string `json:"hiddenAt,omitempty"`
	Id        *string `json:"id,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
	Type      *string `json:"type,omitempty"`
	User      *User   `json:"user,omitempty"`
}

// VideoVariant defines model for VideoVariant.
type VideoVariant struct {
	Index *int    `json:"index,omitempty"`
	Name  *string `json:"name,omitempty"`
}

// WebConfig defines model for WebConfig.
type WebConfig struct {
	AppearanceVariables  *map[string]string    `json:"appearanceVariables,omitempty"`
	Authentication       *AuthenticationConfig `json:"authentication,omitempty"`
	ChatDisabled         *bool                 `json:"chatDisabled,omitempty"`
	CustomStyles         *string               `json:"customStyles,omitempty"`
	ExternalActions      *[]ExternalAction     `json:"externalActions,omitempty"`
	ExtraPageContent     *string               `json:"extraPageContent,omitempty"`
	Federation           *FederationConfig     `json:"federation,omitempty"`
	HideViewerCount      *bool                 `json:"hideViewerCount,omitempty"`
	Logo                 *string               `json:"logo,omitempty"`
	MaxSocketPayloadSize *int                  `json:"maxSocketPayloadSize,omitempty"`
	Name                 *string               `json:"name,omitempty"`
	Notifications        *NotificationConfig   `json:"notifications,omitempty"`
	Nsfw                 *bool                 `json:"nsfw,omitempty"`
	OfflineMessage       *string               `json:"offlineMessage,omitempty"`
	SocialHandles        *[]SocialHandle       `json:"socialHandles,omitempty"`
	SocketHostOverride   *string               `json:"socketHostOverride,omitempty"`
	StreamTitle          *string               `json:"streamTitle,omitempty"`
	Summary              *string               `json:"summary,omitempty"`
	Tags                 *[]string             `json:"tags,omitempty"`
	Version              *string               `json:"version,omitempty"`
}

// YPDetails defines model for YPDetails.
type YPDetails struct {
	Description           *string         `json:"description,omitempty"`
	LastConnectTime       *string         `json:"lastConnectTime,omitempty"`
	Logo                  *string         `json:"logo,omitempty"`
	Name                  *string         `json:"name,omitempty"`
	Nsfw                  *bool           `json:"nsfw,omitempty"`
	Online                *bool           `json:"online,omitempty"`
	OverallMaxViewerCount *int            `json:"overallMaxViewerCount,omitempty"`
	SessionMaxViewerCount *int            `json:"sessionMaxViewerCount,omitempty"`
	Social                *[]SocialHandle `json:"social,omitempty"`
	StreamTitle           *string         `json:"streamTitle,omitempty"`
	Tags                  *[]string       `json:"tags,omitempty"`
	ViewerCount           *int            `json:"viewerCount,omitempty"`
}

// AuthorizationHeader defines model for AuthorizationHeader.
type AuthorizationHeader = string

// N400 Simple API response
type N400 = BaseAPIResponse

// N500 Structure for an error response
type N500 = Error

// N501 Structure for an error response
type N501 = Error

// RegisterAnonymousChatUserJSONBody defines parameters for RegisterAnonymousChatUser.
type RegisterAnonymousChatUserJSONBody struct {
	DisplayName *string `json:"displayName,omitempty"`
}

// RegisterAnonymousChatUserParams defines parameters for RegisterAnonymousChatUser.
type RegisterAnonymousChatUserParams struct {
	XForwardedUser *string `json:"X-Forwarded-User,omitempty"`
}

// GetFollowersParams defines parameters for GetFollowers.
type GetFollowersParams struct {
	// Offset The number of items to skip before starting to collect the result set
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit The numbers of items to return
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// SendSystemMessageParams defines parameters for SendSystemMessage.
type SendSystemMessageParams struct {
	// Authorization Bearer access token
	Authorization *AuthorizationHeader `json:"Authorization,omitempty"`
}

// PostNotificationsRegisterJSONBody defines parameters for PostNotificationsRegister.
type PostNotificationsRegisterJSONBody struct {
	// Channel Name of notification channel
	Channel *string `json:"channel,omitempty"`

	// Destination Target of the notification in the channel
	Destination *string `json:"destination,omitempty"`
}

// PostNotificationsRegisterParams defines parameters for PostNotificationsRegister.
type PostNotificationsRegisterParams struct {
	AccessToken string `form:"accessToken" json:"accessToken"`
}

// RemoteFollowJSONBody defines parameters for RemoteFollow.
type RemoteFollowJSONBody struct {
	Account *string `json:"account,omitempty"`
}

// RegisterAnonymousChatUserJSONRequestBody defines body for RegisterAnonymousChatUser for application/json ContentType.
type RegisterAnonymousChatUserJSONRequestBody RegisterAnonymousChatUserJSONBody

// SendSystemMessageJSONRequestBody defines body for SendSystemMessage for application/json ContentType.
type SendSystemMessageJSONRequestBody = SystemMessage

// PostMetricsPlaybackJSONRequestBody defines body for PostMetricsPlayback for application/json ContentType.
type PostMetricsPlaybackJSONRequestBody = PlaybackMetrics

// PostNotificationsRegisterJSONRequestBody defines body for PostNotificationsRegister for application/json ContentType.
type PostNotificationsRegisterJSONRequestBody PostNotificationsRegisterJSONBody

// RemoteFollowJSONRequestBody defines body for RemoteFollow for application/json ContentType.
type RemoteFollowJSONRequestBody RemoteFollowJSONBody

// AsUserMessage returns the union data inside the ChatMessages_Item as a UserMessage
func (t ChatMessages_Item) AsUserMessage() (UserMessage, error) {
	var body UserMessage
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromUserMessage overwrites any union data inside the ChatMessages_Item as the provided UserMessage
func (t *ChatMessages_Item) FromUserMessage(v UserMessage) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeUserMessage performs a merge with any union data inside the ChatMessages_Item, using the provided UserMessage
func (t *ChatMessages_Item) MergeUserMessage(v UserMessage) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsSystemMessage returns the union data inside the ChatMessages_Item as a SystemMessage
func (t ChatMessages_Item) AsSystemMessage() (SystemMessage, error) {
	var body SystemMessage
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSystemMessage overwrites any union data inside the ChatMessages_Item as the provided SystemMessage
func (t *ChatMessages_Item) FromSystemMessage(v SystemMessage) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSystemMessage performs a merge with any union data inside the ChatMessages_Item, using the provided SystemMessage
func (t *ChatMessages_Item) MergeSystemMessage(v SystemMessage) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsActionMessage returns the union data inside the ChatMessages_Item as a ActionMessage
func (t ChatMessages_Item) AsActionMessage() (ActionMessage, error) {
	var body ActionMessage
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromActionMessage overwrites any union data inside the ChatMessages_Item as the provided ActionMessage
func (t *ChatMessages_Item) FromActionMessage(v ActionMessage) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeActionMessage performs a merge with any union data inside the ChatMessages_Item, using the provided ActionMessage
func (t *ChatMessages_Item) MergeActionMessage(v ActionMessage) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFederatedAction returns the union data inside the ChatMessages_Item as a FederatedAction
func (t ChatMessages_Item) AsFederatedAction() (FederatedAction, error) {
	var body FederatedAction
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFederatedAction overwrites any union data inside the ChatMessages_Item as the provided FederatedAction
func (t *ChatMessages_Item) FromFederatedAction(v FederatedAction) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFederatedAction performs a merge with any union data inside the ChatMessages_Item, using the provided FederatedAction
func (t *ChatMessages_Item) MergeFederatedAction(v FederatedAction) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t ChatMessages_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ChatMessages_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
