// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package dashboard

import "encoding/json"
import "fmt"
import "reflect"

// Contains the list of annotations that are associated with the dashboard.
// Annotations are used to overlay event markers and overlay event tags on graphs.
// Grafana comes with a native annotation store and the ability to add annotation
// events directly from the graph panel or via the HTTP API.
// See
// https://grafana.com/docs/grafana/latest/dashboards/build-dashboards/annotate-visualizations/
type AnnotationContainerJson struct {
	// List of annotations
	List []AnnotationQueryJson `json:"list,omitempty" yaml:"list,omitempty" mapstructure:"list,omitempty"`
}

type AnnotationPanelFilterJson struct {
	// Should the specified panels be included or excluded
	Exclude bool `json:"exclude,omitempty" yaml:"exclude,omitempty" mapstructure:"exclude,omitempty"`

	// Panel IDs that should be included or excluded
	Ids []int `json:"ids" yaml:"ids" mapstructure:"ids"`
}

// TODO docs
// FROM: AnnotationQuery in grafana-data/src/types/annotations.ts
type AnnotationQueryJson struct {
	// Datasource corresponds to the JSON schema field "datasource".
	Datasource DataSourceRefJson `json:"datasource" yaml:"datasource" mapstructure:"datasource"`

	// When enabled the annotation query is issued with every dashboard refresh
	Enable bool `json:"enable" yaml:"enable" mapstructure:"enable"`

	// Filter corresponds to the JSON schema field "filter".
	Filter *AnnotationPanelFilterJson `json:"filter,omitempty" yaml:"filter,omitempty" mapstructure:"filter,omitempty"`

	// Annotation queries can be toggled on or off at the top of the dashboard.
	// When hide is true, the toggle is not shown in the dashboard.
	Hide bool `json:"hide,omitempty" yaml:"hide,omitempty" mapstructure:"hide,omitempty"`

	// Color to use for the annotation event markers
	IconColor string `json:"iconColor" yaml:"iconColor" mapstructure:"iconColor"`

	// Name of annotation.
	Name string `json:"name" yaml:"name" mapstructure:"name"`

	// Target corresponds to the JSON schema field "target".
	Target *AnnotationTargetJson `json:"target,omitempty" yaml:"target,omitempty" mapstructure:"target,omitempty"`

	// TODO -- this should not exist here, it is based on the --grafana-- datasource
	Type *string `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
}

// TODO: this should be a regular DataQuery that depends on the selected dashboard
// these match the properties of the "grafana" datasouce that is default in most
// dashboards
type AnnotationTargetJson struct {
	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	Limit int `json:"limit" yaml:"limit" mapstructure:"limit"`

	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	MatchAny bool `json:"matchAny" yaml:"matchAny" mapstructure:"matchAny"`

	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	Tags []string `json:"tags" yaml:"tags" mapstructure:"tags"`

	// Only required/valid for the grafana datasource...
	// but code+tests is already depending on it so hard to change
	Type string `json:"type" yaml:"type" mapstructure:"type"`
}

type DashboardCursorSyncJson int

type DashboardJson struct {
	// Annotations corresponds to the JSON schema field "annotations".
	Annotations *AnnotationContainerJson `json:"annotations,omitempty" yaml:"annotations,omitempty" mapstructure:"annotations,omitempty"`

	// Description of dashboard.
	Description *string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	// Whether a dashboard is editable or not.
	Editable bool `json:"editable" yaml:"editable" mapstructure:"editable"`

	// The month that the fiscal year starts on.  0 = January, 11 = December
	FiscalYearStartMonth int `json:"fiscalYearStartMonth,omitempty" yaml:"fiscalYearStartMonth,omitempty" mapstructure:"fiscalYearStartMonth,omitempty"`

	// ID of a dashboard imported from the https://grafana.com/grafana/dashboards/
	// portal
	GnetId *string `json:"gnetId,omitempty" yaml:"gnetId,omitempty" mapstructure:"gnetId,omitempty"`

	// GraphTooltip corresponds to the JSON schema field "graphTooltip".
	GraphTooltip DashboardCursorSyncJson `json:"graphTooltip" yaml:"graphTooltip" mapstructure:"graphTooltip"`

	// Unique numeric identifier for the dashboard.
	// `id` is internal to a specific Grafana instance. `uid` should be used to
	// identify a dashboard across Grafana instances.
	Id *int `json:"id,omitempty" yaml:"id,omitempty" mapstructure:"id,omitempty"`

	// Links with references to other dashboards or external websites.
	Links []DashboardLinkJson `json:"links,omitempty" yaml:"links,omitempty" mapstructure:"links,omitempty"`

	// When set to true, the dashboard will redraw panels at an interval matching the
	// pixel width.
	// This will keep data "moving left" regardless of the query refresh rate. This
	// setting helps
	// avoid dashboards presenting stale live data
	LiveNow *bool `json:"liveNow,omitempty" yaml:"liveNow,omitempty" mapstructure:"liveNow,omitempty"`

	// List of dashboard panels
	Panels []DashboardJsonPanelsElem `json:"panels,omitempty" yaml:"panels,omitempty" mapstructure:"panels,omitempty"`

	// Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m",
	// "1h", "1d".
	Refresh interface{} `json:"refresh,omitempty" yaml:"refresh,omitempty" mapstructure:"refresh,omitempty"`

	// This property should only be used in dashboards defined by plugins.  It is a
	// quick check
	// to see if the version has changed since the last time.
	Revision *int `json:"revision,omitempty" yaml:"revision,omitempty" mapstructure:"revision,omitempty"`

	// Version of the JSON schema, incremented each time a Grafana update brings
	// changes to said schema.
	SchemaVersion int `json:"schemaVersion" yaml:"schemaVersion" mapstructure:"schemaVersion"`

	// Snapshot corresponds to the JSON schema field "snapshot".
	Snapshot *SnapshotJson `json:"snapshot,omitempty" yaml:"snapshot,omitempty" mapstructure:"snapshot,omitempty"`

	// Theme of dashboard.
	// Default value: dark.
	Style DashboardJsonStyle `json:"style" yaml:"style" mapstructure:"style"`

	// Tags associated with dashboard.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty" mapstructure:"tags,omitempty"`

	// Configured template variables
	Templating *DashboardJsonTemplating `json:"templating,omitempty" yaml:"templating,omitempty" mapstructure:"templating,omitempty"`

	// Time range for dashboard.
	// Accepted values are relative time strings like {from: 'now-6h', to: 'now'} or
	// absolute time strings like {from: '2020-07-10T08:00:00.000Z', to:
	// '2020-07-10T14:00:00.000Z'}.
	Time *DashboardJsonTime `json:"time,omitempty" yaml:"time,omitempty" mapstructure:"time,omitempty"`

	// Configuration of the time picker shown at the top of a dashboard.
	Timepicker *DashboardJsonTimepicker `json:"timepicker,omitempty" yaml:"timepicker,omitempty" mapstructure:"timepicker,omitempty"`

	// Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or
	// "utc".
	Timezone string `json:"timezone,omitempty" yaml:"timezone,omitempty" mapstructure:"timezone,omitempty"`

	// Title of dashboard.
	Title *string `json:"title,omitempty" yaml:"title,omitempty" mapstructure:"title,omitempty"`

	// Unique dashboard identifier that can be generated by anyone. string (8-40)
	Uid *string `json:"uid,omitempty" yaml:"uid,omitempty" mapstructure:"uid,omitempty"`

	// Version of the dashboard, incremented each time the dashboard is updated.
	Version *int `json:"version,omitempty" yaml:"version,omitempty" mapstructure:"version,omitempty"`

	// Day when the week starts. Expressed by the name of the day in lowercase, e.g.
	// "monday".
	WeekStart *string `json:"weekStart,omitempty" yaml:"weekStart,omitempty" mapstructure:"weekStart,omitempty"`
}

type DashboardJsonPanelsElem map[string]interface{}

type DashboardJsonStyle string

const DashboardJsonStyleDark DashboardJsonStyle = "dark"
const DashboardJsonStyleLight DashboardJsonStyle = "light"

// Configured template variables
type DashboardJsonTemplating struct {
	// List of configured template variables with their saved values along with some
	// other metadata
	List []VariableModelJson `json:"list,omitempty" yaml:"list,omitempty" mapstructure:"list,omitempty"`
}

// Time range for dashboard.
// Accepted values are relative time strings like {from: 'now-6h', to: 'now'} or
// absolute time strings like {from: '2020-07-10T08:00:00.000Z', to:
// '2020-07-10T14:00:00.000Z'}.
type DashboardJsonTime struct {
	// From corresponds to the JSON schema field "from".
	From string `json:"from" yaml:"from" mapstructure:"from"`

	// To corresponds to the JSON schema field "to".
	To string `json:"to" yaml:"to" mapstructure:"to"`
}

// Configuration of the time picker shown at the top of a dashboard.
type DashboardJsonTimepicker struct {
	// Whether timepicker is collapsed or not. Has no effect on provisioned dashboard.
	Collapse bool `json:"collapse" yaml:"collapse" mapstructure:"collapse"`

	// Whether timepicker is enabled or not. Has no effect on provisioned dashboard.
	Enable bool `json:"enable" yaml:"enable" mapstructure:"enable"`

	// Whether timepicker is visible or not.
	Hidden bool `json:"hidden" yaml:"hidden" mapstructure:"hidden"`

	// Interval options available in the refresh picker dropdown.
	RefreshIntervals []string `json:"refresh_intervals" yaml:"refresh_intervals" mapstructure:"refresh_intervals"`

	// Selectable options available in the time picker dropdown. Has no effect on
	// provisioned dashboard.
	TimeOptions []string `json:"time_options" yaml:"time_options" mapstructure:"time_options"`
}

// Links with references to other dashboards or external resources
type DashboardLinkJson struct {
	// If true, all dashboards links will be displayed in a dropdown. If false, all
	// dashboards links will be displayed side by side. Only valid if the type is
	// dashboards
	AsDropdown bool `json:"asDropdown" yaml:"asDropdown" mapstructure:"asDropdown"`

	// Icon name to be displayed with the link
	Icon string `json:"icon" yaml:"icon" mapstructure:"icon"`

	// If true, includes current template variables values in the link as query params
	IncludeVars bool `json:"includeVars" yaml:"includeVars" mapstructure:"includeVars"`

	// If true, includes current time range in the link as query params
	KeepTime bool `json:"keepTime" yaml:"keepTime" mapstructure:"keepTime"`

	// List of tags to limit the linked dashboards. If empty, all dashboards will be
	// displayed. Only valid if the type is dashboards
	Tags []string `json:"tags" yaml:"tags" mapstructure:"tags"`

	// If true, the link will be opened in a new tab
	TargetBlank bool `json:"targetBlank" yaml:"targetBlank" mapstructure:"targetBlank"`

	// Title to display with the link
	Title string `json:"title" yaml:"title" mapstructure:"title"`

	// Tooltip to display when the user hovers their mouse over it
	Tooltip string `json:"tooltip" yaml:"tooltip" mapstructure:"tooltip"`

	// Type corresponds to the JSON schema field "type".
	Type DashboardLinkTypeJson `json:"type" yaml:"type" mapstructure:"type"`

	// Link URL. Only required/valid if the type is link
	Url string `json:"url" yaml:"url" mapstructure:"url"`
}

type DashboardLinkTypeJson string

const DashboardLinkTypeJsonDashboards DashboardLinkTypeJson = "dashboards"
const DashboardLinkTypeJsonLink DashboardLinkTypeJson = "link"

// Ref to a DataSource instance
type DataSourceRefJson struct {
	// The plugin type-id
	Type *string `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`

	// Specific datasource instance
	Uid *string `json:"uid,omitempty" yaml:"uid,omitempty" mapstructure:"uid,omitempty"`
}

// A dashboard snapshot shares an interactive dashboard publicly.
// It is a read-only version of a dashboard, and is not editable.
// It is possible to create a snapshot of a snapshot.
// Grafana strips away all sensitive information from the dashboard.
// Sensitive information stripped: queries (metric, template,annotation) and panel
// links.
type SnapshotJson struct {
	// Time when the snapshot was created
	Created string `json:"created" yaml:"created" mapstructure:"created"`

	// Time when the snapshot expires, default is never to expire
	Expires string `json:"expires" yaml:"expires" mapstructure:"expires"`

	// Is the snapshot saved in an external grafana instance
	External bool `json:"external" yaml:"external" mapstructure:"external"`

	// external url, if snapshot was shared in external grafana instance
	ExternalUrl string `json:"externalUrl" yaml:"externalUrl" mapstructure:"externalUrl"`

	// Unique identifier of the snapshot
	Id int `json:"id" yaml:"id" mapstructure:"id"`

	// Optional, defined the unique key of the snapshot, required if external is true
	Key string `json:"key" yaml:"key" mapstructure:"key"`

	// Optional, name of the snapshot
	Name string `json:"name" yaml:"name" mapstructure:"name"`

	// org id of the snapshot
	OrgId int `json:"orgId" yaml:"orgId" mapstructure:"orgId"`

	// last time when the snapshot was updated
	Updated string `json:"updated" yaml:"updated" mapstructure:"updated"`

	// url of the snapshot, if snapshot was shared internally
	Url *string `json:"url,omitempty" yaml:"url,omitempty" mapstructure:"url,omitempty"`

	// user id of the snapshot creator
	UserId int `json:"userId" yaml:"userId" mapstructure:"userId"`
}

type VariableHideJson int

// A variable is a placeholder for a value. You can use variables in metric queries
// and in panel titles.
type VariableModelJson struct {
	// Format to use while fetching all values from data source, eg: wildcard, glob,
	// regex, pipe, etc.
	AllFormat *string `json:"allFormat,omitempty" yaml:"allFormat,omitempty" mapstructure:"allFormat,omitempty"`

	// Current corresponds to the JSON schema field "current".
	Current *VariableOptionJson `json:"current,omitempty" yaml:"current,omitempty" mapstructure:"current,omitempty"`

	// Datasource corresponds to the JSON schema field "datasource".
	Datasource *DataSourceRefJson `json:"datasource,omitempty" yaml:"datasource,omitempty" mapstructure:"datasource,omitempty"`

	// Description of variable. It can be defined but `null`.
	Description *string `json:"description,omitempty" yaml:"description,omitempty" mapstructure:"description,omitempty"`

	// Hide corresponds to the JSON schema field "hide".
	Hide VariableHideJson `json:"hide" yaml:"hide" mapstructure:"hide"`

	// Unique numeric identifier for the variable.
	Id string `json:"id" yaml:"id" mapstructure:"id"`

	// Optional display name
	Label *string `json:"label,omitempty" yaml:"label,omitempty" mapstructure:"label,omitempty"`

	// Whether multiple values can be selected or not from variable value list
	Multi bool `json:"multi,omitempty" yaml:"multi,omitempty" mapstructure:"multi,omitempty"`

	// Name of variable
	Name string `json:"name" yaml:"name" mapstructure:"name"`

	// Options that can be selected for a variable.
	Options []VariableOptionJson `json:"options,omitempty" yaml:"options,omitempty" mapstructure:"options,omitempty"`

	// Query used to fetch values for a variable
	Query interface{} `json:"query,omitempty" yaml:"query,omitempty" mapstructure:"query,omitempty"`

	// Refresh corresponds to the JSON schema field "refresh".
	Refresh *VariableRefreshJson `json:"refresh,omitempty" yaml:"refresh,omitempty" mapstructure:"refresh,omitempty"`

	// Whether the variable value should be managed by URL query params or not
	SkipUrlSync bool `json:"skipUrlSync" yaml:"skipUrlSync" mapstructure:"skipUrlSync"`

	// Type corresponds to the JSON schema field "type".
	Type VariableTypeJson `json:"type" yaml:"type" mapstructure:"type"`
}

// Option to be selected in a variable.
type VariableOptionJson struct {
	// Whether the option is selected or not
	Selected *bool `json:"selected,omitempty" yaml:"selected,omitempty" mapstructure:"selected,omitempty"`

	// Text to be displayed for the option
	Text interface{} `json:"text" yaml:"text" mapstructure:"text"`

	// Value of the option
	Value interface{} `json:"value" yaml:"value" mapstructure:"value"`
}

type VariableRefreshJson int

type VariableTypeJson string

const VariableTypeJsonAdhoc VariableTypeJson = "adhoc"
const VariableTypeJsonConstant VariableTypeJson = "constant"
const VariableTypeJsonCustom VariableTypeJson = "custom"
const VariableTypeJsonDatasource VariableTypeJson = "datasource"
const VariableTypeJsonInterval VariableTypeJson = "interval"
const VariableTypeJsonQuery VariableTypeJson = "query"
const VariableTypeJsonSystem VariableTypeJson = "system"
const VariableTypeJsonTextbox VariableTypeJson = "textbox"

var enumValues_DashboardCursorSyncJson = []interface{}{
	0,
	1,
	2,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableHideJson) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_VariableHideJson {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_VariableHideJson, v)
	}
	*j = VariableHideJson(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SnapshotJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["created"]; !ok || v == nil {
		return fmt.Errorf("field created in SnapshotJson: required")
	}
	if v, ok := raw["expires"]; !ok || v == nil {
		return fmt.Errorf("field expires in SnapshotJson: required")
	}
	if v, ok := raw["external"]; !ok || v == nil {
		return fmt.Errorf("field external in SnapshotJson: required")
	}
	if v, ok := raw["externalUrl"]; !ok || v == nil {
		return fmt.Errorf("field externalUrl in SnapshotJson: required")
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id in SnapshotJson: required")
	}
	if v, ok := raw["key"]; !ok || v == nil {
		return fmt.Errorf("field key in SnapshotJson: required")
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name in SnapshotJson: required")
	}
	if v, ok := raw["orgId"]; !ok || v == nil {
		return fmt.Errorf("field orgId in SnapshotJson: required")
	}
	if v, ok := raw["updated"]; !ok || v == nil {
		return fmt.Errorf("field updated in SnapshotJson: required")
	}
	if v, ok := raw["userId"]; !ok || v == nil {
		return fmt.Errorf("field userId in SnapshotJson: required")
	}
	type Plain SnapshotJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SnapshotJson(plain)
	return nil
}

var enumValues_VariableTypeJson = []interface{}{
	"query",
	"adhoc",
	"constant",
	"datasource",
	"interval",
	"textbox",
	"custom",
	"system",
}
var enumValues_VariableRefreshJson = []interface{}{
	0.0,
	1.0,
	2.0,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DashboardJsonStyle) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DashboardJsonStyle {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DashboardJsonStyle, v)
	}
	*j = DashboardJsonStyle(v)
	return nil
}

var enumValues_DashboardJsonStyle = []interface{}{
	"dark",
	"light",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableOptionJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["text"]; !ok || v == nil {
		return fmt.Errorf("field text in VariableOptionJson: required")
	}
	if v, ok := raw["value"]; !ok || v == nil {
		return fmt.Errorf("field value in VariableOptionJson: required")
	}
	type Plain VariableOptionJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableOptionJson(plain)
	return nil
}

var enumValues_VariableHideJson = []interface{}{
	0.0,
	1.0,
	2.0,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableTypeJson) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_VariableTypeJson {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_VariableTypeJson, v)
	}
	*j = VariableTypeJson(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DashboardLinkJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["icon"]; !ok || v == nil {
		return fmt.Errorf("field icon in DashboardLinkJson: required")
	}
	if v, ok := raw["tags"]; !ok || v == nil {
		return fmt.Errorf("field tags in DashboardLinkJson: required")
	}
	if v, ok := raw["title"]; !ok || v == nil {
		return fmt.Errorf("field title in DashboardLinkJson: required")
	}
	if v, ok := raw["tooltip"]; !ok || v == nil {
		return fmt.Errorf("field tooltip in DashboardLinkJson: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type in DashboardLinkJson: required")
	}
	if v, ok := raw["url"]; !ok || v == nil {
		return fmt.Errorf("field url in DashboardLinkJson: required")
	}
	type Plain DashboardLinkJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["asDropdown"]; !ok || v == nil {
		plain.AsDropdown = false
	}
	if v, ok := raw["includeVars"]; !ok || v == nil {
		plain.IncludeVars = false
	}
	if v, ok := raw["keepTime"]; !ok || v == nil {
		plain.KeepTime = false
	}
	if v, ok := raw["targetBlank"]; !ok || v == nil {
		plain.TargetBlank = false
	}
	*j = DashboardLinkJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableRefreshJson) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_VariableRefreshJson {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_VariableRefreshJson, v)
	}
	*j = VariableRefreshJson(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DashboardCursorSyncJson) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DashboardCursorSyncJson {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DashboardCursorSyncJson, v)
	}
	*j = DashboardCursorSyncJson(v)
	return nil
}

var enumValues_DashboardLinkTypeJson = []interface{}{
	"link",
	"dashboards",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableModelJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["hide"]; !ok || v == nil {
		return fmt.Errorf("field hide in VariableModelJson: required")
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name in VariableModelJson: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type in VariableModelJson: required")
	}
	type Plain VariableModelJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		plain.Id = "00000000-0000-0000-0000-000000000000"
	}
	if v, ok := raw["multi"]; !ok || v == nil {
		plain.Multi = false
	}
	if v, ok := raw["skipUrlSync"]; !ok || v == nil {
		plain.SkipUrlSync = false
	}
	*j = VariableModelJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DashboardLinkTypeJson) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DashboardLinkTypeJson {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DashboardLinkTypeJson, v)
	}
	*j = DashboardLinkTypeJson(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AnnotationQueryJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["datasource"]; !ok || v == nil {
		return fmt.Errorf("field datasource in AnnotationQueryJson: required")
	}
	if v, ok := raw["iconColor"]; !ok || v == nil {
		return fmt.Errorf("field iconColor in AnnotationQueryJson: required")
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name in AnnotationQueryJson: required")
	}
	type Plain AnnotationQueryJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["enable"]; !ok || v == nil {
		plain.Enable = true
	}
	if v, ok := raw["hide"]; !ok || v == nil {
		plain.Hide = false
	}
	*j = AnnotationQueryJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DashboardJsonTime) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain DashboardJsonTime
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["from"]; !ok || v == nil {
		plain.From = "now-6h"
	}
	if v, ok := raw["to"]; !ok || v == nil {
		plain.To = "now"
	}
	*j = DashboardJsonTime(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AnnotationTargetJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["limit"]; !ok || v == nil {
		return fmt.Errorf("field limit in AnnotationTargetJson: required")
	}
	if v, ok := raw["matchAny"]; !ok || v == nil {
		return fmt.Errorf("field matchAny in AnnotationTargetJson: required")
	}
	if v, ok := raw["tags"]; !ok || v == nil {
		return fmt.Errorf("field tags in AnnotationTargetJson: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type in AnnotationTargetJson: required")
	}
	type Plain AnnotationTargetJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AnnotationTargetJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DashboardJsonTimepicker) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain DashboardJsonTimepicker
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["collapse"]; !ok || v == nil {
		plain.Collapse = false
	}
	if v, ok := raw["enable"]; !ok || v == nil {
		plain.Enable = true
	}
	if v, ok := raw["hidden"]; !ok || v == nil {
		plain.Hidden = false
	}
	if v, ok := raw["refresh_intervals"]; !ok || v == nil {
		plain.RefreshIntervals = []string{
			"5s",
			"10s",
			"30s",
			"1m",
			"5m",
			"15m",
			"30m",
			"1h",
			"2h",
			"1d",
		}
	}
	if v, ok := raw["time_options"]; !ok || v == nil {
		plain.TimeOptions = []string{
			"5m",
			"15m",
			"1h",
			"6h",
			"12h",
			"24h",
			"2d",
			"7d",
			"30d",
		}
	}
	*j = DashboardJsonTimepicker(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AnnotationPanelFilterJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["ids"]; !ok || v == nil {
		return fmt.Errorf("field ids in AnnotationPanelFilterJson: required")
	}
	type Plain AnnotationPanelFilterJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["exclude"]; !ok || v == nil {
		plain.Exclude = false
	}
	*j = AnnotationPanelFilterJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DashboardJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["graphTooltip"]; !ok || v == nil {
		return fmt.Errorf("field graphTooltip in DashboardJson: required")
	}
	type Plain DashboardJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["editable"]; !ok || v == nil {
		plain.Editable = true
	}
	if v, ok := raw["fiscalYearStartMonth"]; !ok || v == nil {
		plain.FiscalYearStartMonth = 0.0
	}
	if v, ok := raw["schemaVersion"]; !ok || v == nil {
		plain.SchemaVersion = 36.0
	}
	if v, ok := raw["style"]; !ok || v == nil {
		plain.Style = "dark"
	}
	if v, ok := raw["timezone"]; !ok || v == nil {
		plain.Timezone = "browser"
	}
	*j = DashboardJson(plain)
	return nil
}
