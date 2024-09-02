package models
type SmsRecordData struct {
	//Supi			string		     	`json:"supi,omitempty"`
	smsRecordId       	string		  	`json:"smsRecordId,omitempty" bson:"smsRecordId"`
	smsPayLoad        	string  		`json:"smsPayLoad,omitempty" bson:"smsPayLoad"`
	Gpsi			string  		`json:"Gpsi,omitempty" bson:"Gpsi"`
        Pei			string 	 		`json:"pei,omitempty" bson:"pei"`
        //Guami			*Guami          	`json:"guami" bson:"guami"`
        AccessType		string  		`json:"AccessType,omitempty" bson:"AccessType"`
        //RatType			RatType         	`json:"ratType" bson:"ratType"`
        //AmfInstanceId		string 		 	`json:"amfInstanceId" bson:"amfInstanceId"`
        UeTimeZone        	string   		`json:"ueTimeZone,omitempty"`
	UeLocation		*UserLocation		`json:"ueLocation,omitempty"`
        //SupportedFeatures	string  		`json:"supportedFeatures,omitempty" bson:"supportedFeatures"`
	//N1MessageContainer	*N1MessageContainer	`json:"n1MessageContainer,omitempty"`
	//BackupAmfInfo		[]BackupAmfInfo 	`json:"backupAmfInfo,omitempty" bson:"backupAmfInfo"`
        //DrFlag			bool            	`json:"drFlag,omitempty" bson:"drFlag"`
}
