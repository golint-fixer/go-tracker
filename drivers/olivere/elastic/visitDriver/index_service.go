package visitDriver

import (
	"time"
	"encoding/json"
	"github.com/index0h/go-tracker/visit/entity"
	"github.com/index0h/go-tracker/uuid"
)

type IndexService struct {
	uuid uuid.UuidMaker
	timestampLayout string

	namePrefix string
	nameDateSuffix string
	typeName string
}

func NewIndexService(uuid uuid.UuidMaker) *IndexService {

	return &IndexService{
		uuid: uuid,
		timestampLayout: "2006-01-02 15:04:05",
		namePrefix: "tracker-visit",
		nameDateSuffix: "2006-01",
		typeName: "visit",
	}
}

func (index *IndexService) Name(timestamp int64) string {
	return index.namePrefix + "-" + time.Unix(timestamp, 0).Format(index.nameDateSuffix)
}

func (index *IndexService) Type() string {
	return index.typeName
}

func (index *IndexService) Body() string {

	return `{
   "mapping":{
      "visit":{
         "properties":{
            "_id":{
               "index":"not_analyzed",
               "stored":true,
               "type":"string"
            },
            "@timestamp":{
               "format":"YYYY-MM-DD HH:mm:ss",
               "type":"date"
            },
            "clientId":{
               "index":"not_analyzed",
               "type":"string"
            },
            "dataList":{
               "include_in_parent":true,
               "properties":{
                  "key":{
                     "index":"not_analyzed",
                     "type":"string"
                  },
                  "value":{
                     "index":"not_analyzed",
                     "type":"string"
                  }
               },
               "type":"nested"
            },
            "sessionId":{
               "index":"not_analyzed",
               "type":"string"
            },
            "warnings":{
               "index":"not_analyzed",
               "type":"string"
            }
         }
      }
   }
}`
}

func (index *IndexService) Marshal(visit *entity.Visit) (string, []byte, error) {
	visitID := index.uuid.ToString(visit.VisitID())
	model := mapVisit{
		VisitID: visitID,
		Timestamp: time.Unix(visit.Timestamp(), 0).Format(index.timestampLayout),
		SessionID: index.uuid.ToString(visit.SessionID()),
		ClientID: visit.ClientID(),
		WarningList: visit.Warnings(),
	}

	dataFromVisit := visit.Data()
	model.DataList = make([]mapDataList, len(dataFromVisit))
	var i uint
	for key, value := range dataFromVisit {
		model.DataList[i] = mapDataList{Key: key, Value: value}

		i++
	}

	result, err := json.Marshal(model)

	return visitID, result, err
}

func (index *IndexService) MarshalID(visitID uuid.Uuid) string {
	return index.uuid.ToString(visitID)
}

func (index *IndexService) Unmarshal(data []byte) (visit *entity.Visit, err error) {
	var raw mapVisit
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}


	return nil, nil
}

type mapVisit struct {
	VisitID string `json:"_id"`
	Timestamp string `json:"@timestamp"`
	SessionID string `json:"sessionId"`
	ClientID string `json:"clientId"`
	DataList []mapDataList `json:"dataList"`
	WarningList []string `json:"warningList"`
}

type mapDataList struct {
	Key string `json:"key"`
	Value string `json:"value"`
}
