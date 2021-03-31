package main

// UserTag 用户&标签
type UserTag struct {
	UserID        uint  // 用户ID
	TagID         uint  // 标签ID
	CategoryTagID int32 // 标签分类
}

//
//// IDSet 定义
//type IDSet struct {
//	Data *strset.Set
//}
//
//// UserIDGroup 标签组-用户ID集合
//type UserIDGroup struct {
//	CategoryTagID int32 // 标签分类
//	UserIDs       IDSet // 用户ID集合
//}
//
//// Scan scan value into Object
//func (m *IDSet) Scan(value interface{}) error {
//	bytes, ok := value.([]byte)
//	if !ok {
//		return fmt.Errorf("IDSet Scan() only support []byte")
//	}
//	var ids []string
//	if len(bytes) > 0 {
//		ids = strings.Split(string(bytes), ",")
//	}
//	m.Data = strset.New(ids...)
//	return nil
//}
//
//// Value return serialized value
//func (m IDSet) Value() (driver.Value, error) {
//	return []byte(""), nil
//}
