package model

import (
	"bytes"
	"errors"
	"fmt"
	"gin-myboot/global"
	"gin-myboot/modules/common/model/request"
	"gin-myboot/utils"
	"gorm.io/gorm"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// Search 条件搜索
func Search(searchParams []request.SearchParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		for _, search := range searchParams {
			var op string
			name := search.Name
			value := search.Value
			operator := search.Operator
			if name == "" {
				continue
			}

			if operator == "" {
				operator = "="
			}

			if utils.IsNil(value) || value == "" {
				continue
			}

			k := reflect.ValueOf(value)

			if k.IsValid() == false {
				continue
			}

			if k.IsZero() {
				continue
			}
			b := ""
			e := ""
			switch k.Kind() {
			case reflect.Array, reflect.Slice:
				if k.Len() == 0 {
					continue
				}
				for i := 0; i < k.Len(); i++ {
					global.Error("=========>i", k.Index(i))
				}

				if k.Len() >= 2 && (operator == "between" || operator == "notBetween") {
					b = k.Index(0).Elem().String()
					e = k.Index(1).Elem().String()
				}
			default:
				fmt.Println("unknown")
			}

			name = utils.Camel2Case(name)
			switch operator {
			case ">":
				op = "> ?"
			case ">=":
				op = ">= ?"
			case "<":
				op = "< ?"
			case "<=":
				op = "<= ?"
			case "in":
				op = "in ?"
			case "notIn":
				op = "not in ?"
			case "notLike":
				op = "not like ?"
			case "like":
				op = "like %?%"
			case "likeLeft":
				op = "like %?"
			case "likeRight":
				op = "like ?%"
			case "between":
				op = "between ? AND ?"
			case "notBetween":
				op = "not between ? AND ?"
			case "isNull":
				op = "is null"
			case "isNotNull":
				op = "is not null"
			default:
				op = "= ?"
			}

			opStr := name + " " + op
			if operator == "between" || operator == "notBetween" {
				db.Where(opStr, b, e)
			} else {
				db.Where(opStr, value)
			}
			global.Debug("=======>", opStr)

		}

		global.Error("=======>", searchParams)
		return db
	}
}

// SortOrder 排序条件
func SortOrder(sortOrderParams request.SortOrderParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if sortOrderParams.Column != "" {
			name := utils.Camel2Case(sortOrderParams.Column)
			if sortOrderParams.Order == "ascend" {
				db.Order(name + " ASC")
			} else {
				db.Order(name + " DESC")
			}
		}

		return db
	}
}

func getKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for s, _ := range m {
		keys = append(keys, s)
	}

	return keys
}

// m := make(map[string]interface{})

//var m []map[string]interface{}
//
//cc := make(map[string]interface{})
//cc["id"] = 1
//cc["customer_name"] = "你好"
//cc["customer_phone_data"] = "不好"
//
//cc1 := make(map[string]interface{})
//cc1["id"] = 2
//cc1["customer_name"] = "你好2"
//cc1["customer_phone_data"] = "不好2"
//
//cc2 := make(map[string]interface{})
//cc2["id"] = 3
//cc2["customer_name"] = "你好3"
//cc2["customer_phone_data"] = "不好3"
//
//m = append(m, cc)
//m = append(m, cc1)
//m = append(m, cc2)
//
//global.Error("=====>", m)
//ee := model.MapBuildBatchUpdate("exa_customers", m, "id")
//global.Error("==>", ee)

// MapBuildBatchUpdate 批量更新
func MapBuildBatchUpdate(tableName string, dataList []map[string]interface{}, keyName string) (err error) {
	array, err := MapBuildBatchUpdateSQLArray(tableName, dataList, keyName)
	if err != nil {
		return err
	}

	err = global.GormDB.Transaction(func(tx *gorm.DB) error {
		for _, sql := range array {
			global.Error("=========MapBuildBatchUpdate", sql)
			txErr := tx.Exec(sql).Error
			if txErr != nil {
				global.Error("=========MapBuildBatchUpdate", txErr.Error())
				return txErr
			}
		}

		return nil
	})

	return err
}

//UPDATE book
//     SET book_id = CASE id
//         WHEN 1 THEN 3
//         WHEN 2 THEN 4
//         WHEN 3 THEN 5
//     END,
//	      name = CASE id
//         WHEN 1 THEN 'New Title 1'
//         WHEN 2 THEN 'New Title 2'
//         WHEN 3 THEN 'New Title 3'
//     END
// WHERE id IN (1,2,3)

func MapBuildBatchUpdateSQLArray(tableName string, dataList []map[string]interface{}, keyName string) (s []string, e error) {
	if keyName == "" {
		keyName = "id"
	}

	var IDList []string
	var SQLArray []string
	var record bytes.Buffer

	updateMap := make(map[string]interface{})
	buffers := make(map[string]bytes.Buffer)
	length := len(dataList)
	//fields := getKeys(dataList[0])

	k := 0
	for i := 0; i < length; i++ {
		updateMap = dataList[i]

		id := updateMap[keyName]
		_, idStr := interfaceToString(id)

		IDList = append(IDList, idStr)

		k++
		for fieldName, fieldValue := range updateMap {

			if fieldName != keyName {

				buffer := buffers[fieldName]
				if buffer.Len() == 0 {
					buffer.WriteString(fieldName + " = CASE " + keyName + " ")
				}

				_, valueStr := interfaceToString(fieldValue)
				buffer.WriteString(" WHEN " + idStr + " THEN " + valueStr)

				if k >= length {
					buffer.WriteString(" END,")
				}
				buffers[fieldName] = buffer
			}
		}
	}

	record.WriteString("UPDATE " + tableName + " SET ")

	i := 0
	for _, sql := range buffers {
		i++
		str := sql.String()
		if i == len(buffers) {
			str = strings.TrimRight(str, ",")
		}

		record.WriteString(str)
	}

	record.WriteString(" WHERE " + keyName + " IN (")
	record.WriteString(strings.Join(IDList[:], ","))
	record.WriteString(");")

	SQLArray = append(SQLArray, record.String())

	return SQLArray, nil

}

func interfaceToString(value interface{}) (err error, s string) {
	elem := reflect.ValueOf(value)
	var temp string
	switch elem.Kind() {
	case reflect.Int:
		temp = strconv.FormatInt(elem.Int(), 10)
	case reflect.Int64:
		temp = strconv.FormatInt(elem.Int(), 10)
	case reflect.String:
		if strings.Contains(elem.String(), "'") {
			temp = fmt.Sprintf("'%v'", strings.ReplaceAll(elem.String(), "'", "\\'"))
		} else {
			temp = fmt.Sprintf("'%v'", elem.String())
		}
	case reflect.Float64:
		temp = strconv.FormatFloat(elem.Float(), 'f', -1, 64)
	case reflect.Bool:
		temp = strconv.FormatBool(elem.Bool())
	default:
		return errors.New("type conversion error"), ""
	}

	return nil, temp

}

// tableName表的名字，itemList你定义的数组类型的结构体，[]*Demo

func GormBuildBatchUpdateSQLArray(tableName string, dataList interface{}, keyName string) (s []string, e error) {

	if keyName == "" {
		keyName = "id"
	}

	fieldValue := reflect.ValueOf(dataList)
	fieldType := reflect.TypeOf(dataList).Elem().Elem()
	sliceLength := fieldValue.Len()
	fieldNum := fieldType.NumField()

	// 检验结构体标签是否为空和重复
	verifyTagDuplicate := make(map[string]string)
	for i := 0; i < fieldNum; i++ {
		fieldTag := fieldType.Field(i).Tag.Get("gorm")

		fieldName := GetFieldName(fieldTag)
		if len(strings.TrimSpace(fieldName)) == 0 {
			return nil, errors.New("the structure attribute should have tag")
		}

		if !strings.HasPrefix(fieldName, "id;") {
			return nil, errors.New("the structure attribute should have primary_key")
		}

		_, ok := verifyTagDuplicate[fieldName]
		if !ok {
			verifyTagDuplicate[fieldName] = fieldName
		} else {
			return nil, errors.New("the structure attribute %v tag is not allow duplication" + fieldName)
		}

	}

	var IDList []string
	updateMap := make(map[string][]string)
	for i := 0; i < sliceLength; i++ {
		// 得到某一个具体的结构体的
		structValue := fieldValue.Index(i).Elem()
		for j := 0; j < fieldNum; j++ {
			elem := structValue.Field(j)

			var temp string
			switch elem.Kind() {
			case reflect.Int64:
				temp = strconv.FormatInt(elem.Int(), 10)
			case reflect.String:
				if strings.Contains(elem.String(), "'") {
					temp = fmt.Sprintf("'%v'", strings.ReplaceAll(elem.String(), "'", "\\'"))
				} else {
					temp = fmt.Sprintf("'%v'", elem.String())
				}
			case reflect.Float64:
				temp = strconv.FormatFloat(elem.Float(), 'f', -1, 64)
			case reflect.Bool:
				temp = strconv.FormatBool(elem.Bool())
			default:
				return nil, errors.New("type conversion error, param is %v" + fieldType.Field(j).Tag.Get("json"))
			}

			gormTag := fieldType.Field(j).Tag.Get("gorm")

			fieldTag := GetFieldName(gormTag)

			if strings.HasPrefix(fieldTag, "id;") {
				id, err := strconv.ParseInt(temp, 10, 64)
				if err != nil {
					return nil, err
				}
				// id 的合法性校验
				if id < 1 {
					return nil, errors.New("this structure should have a primary key and gt 0")
				}
				IDList = append(IDList, temp)
				continue
			}

			valueList := append(updateMap[fieldTag], temp)
			updateMap[fieldTag] = valueList
		}
	}

	length := len(IDList)
	size := 200
	SQLQuantity := getSQLQuantity(length, size)
	var SQLArray []string
	k := 0

	for i := 0; i < SQLQuantity; i++ {
		count := 0

		var record bytes.Buffer
		record.WriteString("UPDATE " + tableName + " SET ")

		for fieldName, fieldValueList := range updateMap {
			record.WriteString(fieldName)
			record.WriteString(" = CASE " + "id")

			for j := k; j < len(IDList) && j < len(fieldValueList) && j < size+k; j++ {
				record.WriteString(" WHEN " + IDList[j] + " THEN " + fieldValueList[j])
			}
			count++
			if count != fieldNum-1 {
				record.WriteString(" END, ")
			}
		}

		record.WriteString(" END WHERE ")
		record.WriteString("id" + " IN (")
		min := size + k
		if len(IDList) < min {
			min = len(IDList)
		}
		record.WriteString(strings.Join(IDList[k:min], ","))
		record.WriteString(");")

		k += size
		SQLArray = append(SQLArray, record.String())
	}

	return SQLArray, nil
}

func getSQLQuantity(length, size int) int {
	SQLQuantity := int(math.Ceil(float64(length) / float64(size)))
	return SQLQuantity
}

func GetFieldName(fieldTag string) string {
	fieldTagArr := strings.Split(fieldTag, ":")
	if len(fieldTagArr) == 0 {
		return ""
	}

	fieldName := fieldTagArr[len(fieldTagArr)-1]

	return fieldName
}

// Paginate 分页
func Paginate(r request.PageInfo) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := r.Page
		if page == 0 {
			page = 1
		}

		pageSize := r.PageSize
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
