package messageprocessor

import (
	"github.com/carbonblack/cb-event-forwarder/internal/sensor_events"
         "time"
)

type ConvertedCbMessage struct {
	OriginalMessage *sensor_events.CbEventMsg
	CbServerURL     string // server of origin
	Type            string // routing key
}

type CbEventForwarderMessage struct {
        CbServerURL string `json:'cb_server_url'`
        Type string `json:'type'`
        CbServer string `json:'cb_server'`
        CbVersion string `json:'cb_version'`
        SchemaVersion int `json:'schema_version'`
	*CbWatchlistHitProcessStorage `json:"omitemtpy"`
	*CbWatchlistHitProcess `json:"omitempty"`
	*CbWatchlistHitBinaryStorage `json:'omitemtpy'`
	*CbWatchlistHitBinary `json:'omitempty'`
	*CbFeedHitProcess `json:'omitempty'`
	*CbFeedHitProcessStorage `json:'omitempty'`
        *CbAlertWatchlistHitProcess `json:'omitempty'`
}

type CbFeedHitProcess struct {
    Timestamp float64 `json:'timestamp'`
    ReportScore int `json:'report_score'`
    ReportId string `json:'report_id'`
    Hostname string `json:'hostname'`
    Group string `json:'group'`
    FromFeedSearch bool `json:'from_freed_serach'`
    FeedName string `json:'feed_name'`
    FeedId int `json:'feed_id'`
    IOCAttr map[string] string `json:'ioc_attr'`
    IOCType string `json:'ioc_type'`
    IOCValue string `json:'ioc_value'`
    LinkProcess string `json:'link_process'`
    LinkSensor string `json:'link_sensor'`
    OsType string `json:'os_type'`
    ProcessGuid string `json:'process_guid'`
    ProcessId string `json:'process_id'`
}

type CbFeedHitProcessStorage struct {
    TimeStamp float64   `json:'timestamp'`
    SensorId int64 `json:'sensor_id'`
    ReportScore int `json:'report_score'`
    ReportId string `json:'report_Id'`
    FromFeedSearch bool `json:'from_feed_search'`
    FeedName string `json:'feed_name'`
    FeedId int `json:'feed_id'`
    Username string `json:'username'`
    ProcessGUID string  `json:'process_guid'`
    LinkParentMd5 string `json:'link_parent_md5'`
    LinkParent string   `json:'link_parent'`
    LastUpdate time.Time `json:'link_update'`
    HostType string `json:'host_type'`
    FileModCount int64 `json:'filemod_count'`
    CrossProcCount int64 `json:'crossproc_count'`
    CmdLine string `'json:'cmdline'`
    ChildProcCount int64 `json:'childproc_count'`
    *AllianceDataMap `json:'omitempty'`
    LinkProcessMd5 string `json:'link_process_md5'`
    ModLoadCount int64 `json:'modload_count'`
    NetconnCount int64 `json:'netconn_count`
    OsType string `json:'os_type'`
    ParentGuid string `json:'parent_guid'`
    ParentMd5 string `json:'parent_md5'`
    ParentName string `json:'parent_name'`
    ParentPid int64 `json:'parent_pid'`
    ParentGUID string `json:'parent_guid'`
    Path string    `json:'path'`
    ProcessMd5 string `'son:'process_md5'`
    ProcessName string `json:'process_name'`
    ProcessPid int64 `json:'process_pid'`
    RegmodCount int64 `json:'regmod_count'`
    Start time.Time `json:'start'`
    ComputerName string `json:'computer_name'`
    CommsIP string `json:'comms_ip'`
    Group string `json:'group'`
    Hostname string `json:'hostname'`
    InterfaceIP string `json:'inerface_ip'`
    IOCAttr map[string] string `json:'ioc_attr'`
    IOCType string `json:'ioc_type'`
    IOCValue string `json:'ioc_value'`
    LinkProcess string `json:'link_process'`
    LinkSensor string `json:'link_sensor'`
}

type CbWatchlistHitBinary struct {
    WatchlistName string `json:'watchlist_name'`
    UserName string `json:'username'`
    ProcessGuid string `json:'process_guid'`
    LinkProcess string `json:'link_process'`
    Start time.Time `json:'start'`
    SensorId int `json:'sensor_id'`
    RegmodCount	int `json:'regmod_count'`
    ProcessBlockCount int `json:'process_block_count'`
    LinkParent string `json:'link_parent'`
    LastUpdate time.Time `json:'last_update'`
    LastServerUpdate int `json:'last_server_update'`
    InterfaceIP string `json:'interface_ip'`
    Hostname string `json:'hostname'`
    HostType string `json:'host_type'`
    Group string `json:'group'`
    ChildprocCount int `json:'childproc_count'`
    CmdLine string `json:'cmdline'`
    CommsIP string `json:'comms_ip'`
    CrossProcCount int `json:'crossproc_count'`
    EmetCount int   `json:'emet_count'`
    FileModCount int `json:'filemod_count'`
    FilteringKnownDlls bool `json:'filtering_known_dlls'`
    LinkProcessMd5 string `json:'link_process_md5'`
    LinkSensor string `'json:'link_sensor'`
    ModloadCount int   `json:'modload_count'`
    NetconnCount int `json:'netconn_count'`
    OsType string `json:'os_type'`
    ParentMd5 string `json:'parent_md5'`
    ParentName string `json:'parent_name'`
    ParentPid int `json:'parent_pid'`
    ParentGuid string `json:'parent_guid'`
    Path string `json:'path'`
    ProcessMd5 string `json:'process_md5'`
    ProcessName string `json:'process_name'`
    ProcessPid int `json:'process_pid'`
    EventTimestamp float64 `json:'event_timestamp'`
    WatchlistII int `json:'watchlist_id'`
}

type CbWatchlistHitBinaryStorage struct {
    WatchlistName string `json:'watchlist_name'`
    Usermame string `json:'username'`
    ProcessGuid string `json:'process_guid'`
    LinkProcess string `json:'link_process'`
    Start time.Time `json:'start'`
    SensorId int `json:'sensor_id'`
    RegmodCount	int `json:'regmod_count'`
    ProcessBlockCount int `json:'process_block_count'`
    LinkParent string `json:'link_parent'`
    LastUpdate time.Time `json:'link_update'`
    InterfaceIP string `json:'interface_ip'`
    Hostname string `json:'hostname'`
    HostType string `json:'host_type'`
    ChildprocCount int `json:'childproc_count'`
    Cmdline string `json:'cmdline'`
    CrossprocCount int  `json:'crossproc_count'`
    EmetCount int `json:'emetcount'`
    FileModCount int `json:'filemod_count'`
    FilteringKnownDlls bool `json:'filtering_known_dlls'`
    LinkProcessMd5 string `json:'link_process_md5'`
    LinkSensor string `json:'link_sensor'`
    ModloadCount int `'json:'modload_count'`
    NetconnCount int `json:'netconn_count'`
    OsType string `json:'os_type'`
    ParentMd5 string `json:'parent_md5`
    ParentName string `json:'parent_name'`
    ParentPid int `json:'parent_pid'`
    ParentGuid string `json:'parent_guid'`
    Path string `json:'path'`
    ProcessMd5 string `json:'process_md5'`
    ProcessName string `json:'process_name'`
    ProcessPid int `json:'process_pid'`
    EventTimestamp float64 `json:'event_timestamp'`
    WatchlistID int `json:'watchlist_id'`
}

type CbWatchlistHitProcess struct {
    WatchlistName string `json:'watchlist_name'`
    Signed	string `json:'signed'`
    ServerAddedTimestamp time.Time `json:'server_added_tiemstamp'`
    ProductVersion	string `json:'product_version'`
    ProductName string `json:'product_name'`
    OSType string `json:'os_Type'`
    OriginalFilename string `json:'original_filename'`
    OrigModLen int64 `json:'orig_modlen'`
    ObservedFileName []	string `json:'observerd_filename'`
    Md5 string `json:'md5'`
    LinkMd5 string `json:'link_md5'`
    LastSeen time.Time `json:'last_seen'`
    IsExecutableImage bool `json:'is_exectuable_image'`
    Is64Bit bool `json:'is_64bit'`
    DigsigSubject string `json:'digsig_subject'`
    DigsigResultCode string `json:'digsig_resultcode'`
    DigsigResult string `json:'digsig_result'`
    DigsigProgName string `json:'digsig_progname'`
    DigsigIssuer string `json:'digsig_issuer'`
    CopiedModLen int64 `json:'copied_modlen'`
    CompanyName string `json:'company_name'`
    Endpoint [] string `json:'endpoint'`
    EventPartitionId [] int64 `json:'event_partition_id'`
    FacetId int64 `json:'facet_id'`
    FileDesc string `json:'file_desc'`
    FileVersion string `json:'file_version'`
    Group [] string `json:'group'`
    *AllianceDataMap `json:omitempty`
    HostCount int `json:'host_count'`
    InternalName string `json:'internal_name'`
    EventTimestamp float64 `json:'event_timestamp'`
    WatchlistId int `json:'watchlist_id'`
}


type AllianceDataMap map[string] AllianceData

/*
func (ad AllianceData) MarshalJSON() ([]byte, error) {
    buffer := bytes.NewBufferString("{")
    buffer.WriteString(fmt.Sprintf("\"%s%s\":%s", "alliance_link_",ad.Ally, ad.Link))
    buffer.WriteString(",")
    buffer.WriteString(fmt.Sprintf("\"%s%s:%s", "alliance_score_",ad.Ally, ad.Link))
    buffer.WriteString(",")
    buffer.WriteString(fmt.Sprintf("\"%s%s:%s", "alliance_updated_",ad.Ally, ad.Link))
    buffer.WriteString(",")
    buffer.WriteString(fmt.Sprintf("\"%s%s:%s", "alliance_data_",ad.Ally, json.Marshal(ad.Data)))
    buffer.WriteString(",")
    buffer.WriteString("}")
    return buffer.Bytes(), nil
}

func (ad * AllianceData) UnmarshalJSON(data []byte) error {

} */

type AllianceData struct {
	Data [] string `json:'data'`
	Link string `'json:'link'`
	Score int `json:'score'`
	Updated time.Time `json:'updated'`
}

type CbWatchlistHitProcessStorage struct {
    WatchlistName string `json:'watchlist_name'`
    WatchlistID int `json:'watchlist_id'`
    FileVersion string `json:'file_version'`
    FileDesc string `json:'file_desc'`
    FacetId int `json:'facet_id'`
    EventParitionId int64 `json:'event_partition_id'`
    Hostname string `json:'hostname'`
    SensorId int `json:'sensor_id'`
    DigsigSignTime time.Time `json:'digsig_sign_time'`
    DigsigResultCode string `json:'digsig_result_code'`
    DigsigResult string `json:'digsig_result'`
    *AllianceDataMap `json:omitemtpy`
    CompanyName string `json:'company_name'`
    CopiedModLen int64 `json:'copied_modlen'`
    DigsigPublisher string `json:'digsig_publisher'`
    Group string `json:'group'`
    HostCount int64 `json:'host_count'`
    InternalName string `json:'internal_name'`
    Is64bit bool `json:'is_64bit'`
    IsExecutableImage bool `json:'is_exectuable_image'`
    LastSeen time.Time `json:'last_seen'`
    LegalCopyright string `json:'legal_copyright'`
    LinkMd5 string `json:'link_md5'`
    Md5 string `json:'md5'`
    ObservedFilename string `json:'observerd_filenaem'`
    OrigModLen int `json:'orig_modlen'`
    OrignalFilename string `json:'original_filename'`
    OsType string `json:'os_type'`
    ProductName string `json:'product_name'`
    ProductVersion string `json:'product_version'`
    ServerAddedTimestamp time.Time `json:'server_added_timestamp'`
    ServerName string `json:'server_name'`
    EventTimestamp time.Time `json:'event_timestamp'`
}

type CbAlertWatchlistHitProcess struct {
    WatchlistName string `json:'watchlist_name'`
    WatchlistID string `json:'watchlist_Id'`
    Username string `json:'username'`
    UniqueID string `json:'unique_id'`
    Timestamp float64 `json:'timestamp'`
    Status string `json:'status'`
    SensorID int64 `json:'sensor_id'`
    SensorCriticality int64 `json:'sensor_criticality'`
    ReportScore int64 `json:'report_score'`
    RegmodCount int64 `json:'regmod_count'`
    IOCAttr map[string ] string `json:'ioc_attr'`
    InterfaceIP string `json:'interface_ip'`
    Hostname string `json:'hostname'`
    Group string `json:'group'`
    FilemodCount int64 `json:'filemod_count'`
    FeedRating int64  `json:'file_rating'`
    FeedName string `json:'feed_name'`
    FeedId int64 `json:'feed_id'`
    AlertSeverity float64 `json:'alert_serverity'`
    AlertType string `json:'alert_type'`
    ChildProcCount int64 `json:'childproc_count'`
    CommsIP string `json:'comms_ip'`
    ComputerName string `json:'computer_name'`
    CreatedTime time.Time `json:'created_time'`
    CrossProcCount int64 `json:'crossproc_count'`
    IOCConfidence float64 `json:'ioc_confidence'`
    IOCType string `json:'ioc_type'`
    IOCValue string `json:'ioc_value'`
    IOCValueFacet string `json:'ioc_value_facet'`
    LinkMd5 string `json:'link_md5'`
    LinkProcess string `json:'link_process'`
    LinkSensor string `json:'link_sensor'`
    Md5 string `json:'md5'`
    ModloadCount int64 `json:'modload_count'`
    NetconnCount int64 `json:'netconn_count'`
    OSType string `json:'os_type'`
    ProcessName string `json:'process_name'`
    ProcessPath string `json:'process_path'`
    ProcessGuid string `json:'procss_guid'`
}

type CbFeedSynchronized struct {
     FeedId int64 `json:'feed_id`
     FeedName string `json:'feed_name'`
     FeedUpdateTIme time.Time `json:'feed_update_time'`
     TimeStamp float64 `json:timestamp'`
     ScanStartTime int64 `json:'scan_start_time'`
}