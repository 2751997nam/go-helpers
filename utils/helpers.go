package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/2751997nam/go-helpers/logs"
	"github.com/2751997nam/go-helpers/message"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ResponseSuccess(c *gin.Context, result any, status int) {
	c.JSON(status, Response{
		Status: "successful",
		Result: result,
	})
}

func ResponseSuccessWithMessage(c *gin.Context, result any, message string) {
	c.JSON(http.StatusOK, Response{
		Status:  "successful",
		Result:  result,
		Message: message,
	})
}

func ResponseWithMeta(c *gin.Context, result any, meta Meta) {
	c.JSON(http.StatusOK, Response{
		Status: "successful",
		Result: result,
		Meta:   meta,
	})
}

func ResponseFail(c *gin.Context, message string, status int) {
	c.JSON(status, Response{
		Status:  "fail",
		Message: message,
	})
}

func ArrayChunk[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

func Join[T any](items []T, sep string) string {
	strs := []string{}
	for _, item := range items {
		strs = append(strs, fmt.Sprint(item))
	}

	return strings.Join(strs, sep)
}

func IsNumeric(value any) bool {
	valueType := reflect.TypeOf(value).String()
	return strings.Contains(valueType, "int") || strings.Contains(valueType, "float") || strings.Contains(valueType, "complex")
}

func IsNumericString(value any) bool {
	if reflect.TypeOf(value).String() == "string" {
		return regexp.MustCompile(`\d+`).MatchString(value.(string))
	}
	valueType := reflect.TypeOf(value).String()
	return strings.Contains(valueType, "int") || strings.Contains(valueType, "float") || strings.Contains(valueType, "complex")
}

func RegexReplace(pattern string, replaceStr string, str string) string {
	rex := regexp.MustCompile(pattern)
	return rex.ReplaceAllString(str, replaceStr)
}

func GetUrlFields(url string) []string {
	retVal := []string{}
	regex := *regexp.MustCompile(`\/:(\w+)($|\/)`)
	res := regex.FindAllStringSubmatch(url, -1)
	if len(res) > 0 && len(res[0]) > 1 {
		retVal = append(retVal, res[0][1])
	}
	return retVal
}

func GetRequestData(c *gin.Context) (map[string]any, error) {
	data := map[string]any{}
	if c.Request.ContentLength > 0 {
		bodyAsByteArray, err := io.ReadAll(c.Request.Body)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal([]byte(bodyAsByteArray), &data); err != nil {
			return nil, err
		}
	}

	if data == nil {
		data = map[string]any{}
	}

	fields := GetUrlFields(c.FullPath())

	for _, field := range fields {
		if IsNumeric(c.Param(field)) {
			value, _ := strconv.ParseFloat(c.Param(field), 64)
			data[field] = value
		} else {
			data[field] = c.Param(field)
		}
	}

	query := c.Request.URL.Query()
	for field, value := range query {
		data[field] = value[0]
	}

	return data, nil
}

func GetInput[T any](key string, data map[string]any, defaultValue T) T {
	value, ok := data[key]
	if ok {
		if value == nil {
			return defaultValue
		}
		return value.(T)
	}

	return defaultValue
}

func ArrayDiff[T comparable](left []T, right []T) (result []T) {
	rightByValue := map[T]bool{}
	for _, value := range right {
		rightByValue[value] = true
	}
	for _, value := range left {
		if _, ok := rightByValue[value]; !ok {
			result = append(result, value)
		}
	}

	return result
}

func ArrayUnique[T comparable](array []T) []T {
	retVal := map[T]T{}

	for _, item := range array {
		if _, ok := retVal[item]; !ok {
			retVal[item] = item
		}
	}

	return maps.Values(retVal)
}

func ExistInMap[ValueType any](key any, items map[any]ValueType) bool {
	if _, ok := items[key]; ok {
		return true
	}

	return false
}

func LogJson(prefix string, data any) {
	str, _ := json.MarshalIndent(data, "", "\t")
	log.Println(prefix, string(str))
}

func Log(data ...any) {
	log.Println(data...)
}

func QuickLog(data any, fields ...any) {
	logEntry := LogEntry{
		Data: data,
	}

	if len(fields) > 0 {
		logEntry.Target = fmt.Sprint(fields[0])
	}
	if len(fields) > 1 {
		logEntry.Type = fmt.Sprint(fields[1])
	}
	if len(fields) > 2 {
		logEntry.Action = fmt.Sprint(fields[2])
	}
	if len(fields) > 3 {
		logEntry.Actor = fmt.Sprint(fields[3])
	}

	go LogViaGRPC(logEntry)
}

func LogViaGRPC(logEntry LogEntry) {
	conn, err := grpc.Dial("logger-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Printf("ERROR Logging %v", err)
	}

	defer conn.Close()

	c := logs.NewLogServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	logData := logs.Log{
		Target: logEntry.Target,
		Type:   logEntry.Type,
		Action: logEntry.Action,
		Actor:  logEntry.Actor,
	}

	data, err := json.Marshal(logEntry.Data)
	if err != nil {
		log.Printf("ERROR Logging %v", err)
	} else {
		logData.Data = string(data)
		_, err = c.WriteLog(ctx, &logs.LogRequest{
			LogEntry: &logData,
		})

		if err != nil {
			log.Printf("ERROR Logging %v", err)
		}
	}

}

func GetMessageResponse(data any) (message.MessageResponse, error) {
	bytes, err := json.Marshal(data)

	return message.MessageResponse{
		Result: bytes,
	}, err
}

func SendMessage(service string, messageType string, messageMethod string, data any) (*message.MessageResponse, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s-service:50001", service), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Printf("ERROR Send Message %v", err)
		return nil, err
	}

	defer conn.Close()

	c := message.NewMessageServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	msgData, err := json.Marshal(data)
	if err != nil {
		log.Printf("ERROR Send Message %v", err)
		return nil, err
	}

	result, err := c.HandleMessage(ctx, &message.MessageRequest{
		MessageEntry: &message.Message{
			Type:   messageType,
			Method: messageMethod,
			Data:   msgData,
		},
	})

	if err != nil {
		log.Printf("ERROR Send Message %v", err)
		return nil, err
	}

	return result, nil
}

func SendAsyncMessage(service string, messageType string, messageMethod string, data any) error {
	conn, err := grpc.Dial(fmt.Sprintf("%s-service:50001", service), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Printf("ERROR Send Message %v", err)
		return err
	}
	defer conn.Close()

	msgData, err := json.Marshal(data)
	if err != nil {
		log.Printf("ERROR Send Message %v", err)
		return err
	}

	c := message.NewMessageServiceClient(conn)
	ctx := context.Background()

	go func() {
		c.HandleMessage(ctx, &message.MessageRequest{
			MessageEntry: &message.Message{
				Type:   messageType,
				Method: messageMethod,
				Data:   msgData,
			},
		})
	}()

	return nil
}

func ResponseMessage(response any) (message.MessageResponse, error) {
	bytes, err := json.Marshal(response)

	if err != nil {
		return message.MessageResponse{
			Message: err.Error(),
			Status:  "fail",
		}, err
	}
	return message.MessageResponse{
		Status:     "successful",
		Result:     bytes,
		StatusCode: 200,
	}, nil
}

func ResponseCustomMessage(response Response) (message.MessageResponse, error) {
	bytes, err := json.Marshal(response.Result)

	if err != nil {
		return message.MessageResponse{
			Message:    err.Error(),
			Status:     "fail",
			StatusCode: 500,
		}, err
	}

	messageResponse := message.MessageResponse{
		Result:     bytes,
		Message:    response.Message,
		Status:     response.Status,
		StatusCode: response.StatusCode,
	}

	if response.Meta != nil {
		metaBytes, err := json.Marshal(response.Meta)

		if err == nil {
			messageResponse.Meta = metaBytes
		}
	}

	return messageResponse, nil
}

func MissingRequiredKeys[T comparable](data map[T]any, keys []T) *T {
	for _, key := range keys {
		if _, ok := data[key]; !ok {
			return &key
		}
	}
	return nil
}
