package messageprocessor

import (
	"github.com/carbonblack/cb-event-forwarder/internal/sensor_events"
)

//both jsonpb.Marshalers
type ConvertedCbMessage struct {
	OriginalMessage *sensor_events.CbEventMsg
	CbServerURL     string // server of origin
	Type            string // routing key
}



type CbEventForwarderMessage struct {
	CbServerURL string
	Type string
	*CbProcessMessage

}
//type = childproc
type CbProcessMessage struct {

}
type CbModuleLoadMessage struct {}
type CbFileModMessage struct {}
type CbNetConnMessage struct {}
type CbRegModMessage struct {}
type CbStatisticsMessage struct {}
type CbEnvironmentMessage struct{}
type CbCrossProcessMessage struct{}
type CbTamperAlertMessage struct{}
type CbProcessBlockedMessage struct {}
type CbEmetMitigationMessage struct{}
type CbNetConnBlockedMessage struct{}
type CbProcessMetadataMessage struct {}
type CbNetConnMsgv2 struct {}
type NetconnBlockedv2 struct {}

/*
type struct {
	Strings              []*CbStringMsg         `protobuf:"bytes,2,rep,name=strings" json:"strings,omitempty"`
	Process              *CbProcessMsg          `protobuf:"bytes,3,opt,name=process" json:"process,omitempty"`
	Modload              *CbModuleLoadMsg       `protobuf:"bytes,4,opt,name=modload" json:"modload,omitempty"`
	Filemod              *CbFileModMsg          `protobuf:"bytes,5,opt,name=filemod" json:"filemod,omitempty"`
	Network              *CbNetConnMsg          `protobuf:"bytes,6,opt,name=network" json:"network,omitempty"`
	Regmod               *CbRegModMsg           `protobuf:"bytes,7,opt,name=regmod" json:"regmod,omitempty"`
	Stats                *CbStatisticsMsg       `protobuf:"bytes,8,opt,name=stats" json:"stats,omitempty"`
	Module               *CbModuleInfoMsg       `protobuf:"bytes,9,opt,name=module" json:"module,omitempty"`
	Vtwrite              *CbVtWriteMsg          `protobuf:"bytes,10,opt,name=vtwrite" json:"vtwrite,omitempty"`
	Vtload               *CbVtLoadMsg           `protobuf:"bytes,11,opt,name=vtload" json:"vtload,omitempty"`
	Childproc            *CbChildProcessMsg     `protobuf:"bytes,12,opt,name=childproc" json:"childproc,omitempty"`
	Env                  *CbEnvironmentMsg      `protobuf:"bytes,13,opt,name=env" json:"env,omitempty"`
	Crossproc            *CbCrossProcessMsg     `protobuf:"bytes,14,opt,name=crossproc" json:"crossproc,omitempty"`
	TamperAlert          *CbTamperAlertMsg      `protobuf:"bytes,15,opt,name=tamperAlert" json:"tamperAlert,omitempty"`
	Blocked              *CbProcessBlockedMsg   `protobuf:"bytes,16,opt,name=blocked" json:"blocked,omitempty"`
	Emet                 *CbEmetMitigationMsg   `protobuf:"bytes,17,opt,name=emet" json:"emet,omitempty"`
	NetconnBlocked       *CbNetConnBlockedMsg   `protobuf:"bytes,18,opt,name=netconnBlocked" json:"netconnBlocked,omitempty"`
	ProcessMeta          *CbProcessMetadataMsg  `protobuf:"bytes,19,opt,name=processMeta" json:"processMeta,omitempty"`
	Networkv2            *CbNetConnMsgv2        `protobuf:"bytes,20,opt,name=networkv2" json:"networkv2,omitempty"`
	NetconnBlockedv2     *CbNetConnBlockedMsgv2 `
}*/