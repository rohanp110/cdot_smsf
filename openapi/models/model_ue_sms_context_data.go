package models

type UeSmsContextData struct {
	Supi			string		     	`json:"supi,omitempty"`
	Gpsi			string  		`json:"gpsi,omitempty" bson:"gpsi"`
        Pei			string 	 		`json:"pei,omitempty" bson:"pei"`
        Guamis			[]Guami          	`json:"guamis" bson:"guamis"`
        AccessType		string  		`json:"accessType,omitempty" bson:"accessType"`
        AdditionalAccessType	string  		`json:"additionalAccessType,omitempty" bson:"additionalAccessType"`
        RatType			RatType         	`json:"ratType" bson:"ratType"`
        AmfId			string 		 	`json:"amfId" bson:"amfId"`
	UeLocation		*UserLocation		`json:"ueLocation,omitempty"`
        UeTimeZone        	string   		`json:"ueTimeZone,omitempty"`
	TraceData		*TraceData		`json:"traceData,omitempty" yaml:"traceData" bson:"traceData"`
	BackupAmfInfo		[]BackupAmfInfo 	`json:"backupAmfInfo,omitempty" bson:"backupAmfInfo"`
	UdmGroupId              string			`json:"udmGroupId,omitempty"`
	RoutingIndicator        string                  `json:"routingIndicator,omitempty"`
}
