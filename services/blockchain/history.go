package blockchain

import (
	"RizhaoLanshanLabourUnion/services/dao"
	"RizhaoLanshanLabourUnion/services/models"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"
)

func CreateHistory(caseId, operation, content string, userId int64) {
	prev, _ := dao.GetLastHistoryOperationHashByCaseId(caseId)
	_, _ = dao.CreateNewHistory(caseId, operation, content, userId, createHash(operation+caseId), prev)
}

func createHash(source string) string {
	timeStr := strconv.Itoa(int(time.Now().Unix()))
	res := sha256.Sum256([]byte(timeStr + source))
	return hex.EncodeToString(res[:])
}

func CreateHistoryByCase(operation string, model *models.Case, userId int64) {
	CreateHistoryByUsingModel(model.CaseID, operation, model, userId)
}

func CreateHistoryByUsingModel(caseId, operation string, model interface{}, userId int64) {
	str, _ := json.Marshal(model)
	CreateHistory(caseId, operation, string(str), userId)
}
