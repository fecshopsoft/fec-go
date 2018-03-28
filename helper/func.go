package helper


// 三种map类型，方便使用
type MapStrInterface map[string]interface{}
type MapIntStr map[int]string
type MapStrInt map[string]int
type MapInt64Str map[int64]string
type MapStrInt64 map[string]int64

type VueMutilSelect map[string][]MapStrInterface

type VueSelectOps struct{
    Key int64 `form:"key" json:"key"`
    DisplayName string `form:"display_name" json:"display_name"`
}

type DeleteIds struct{
    Ids []int `form:"ids" json:"ids"`
}

type DeleteId struct{
    Id int `form:"id" json:"id"`
}

// 通过两重循环过滤重复元素
func SliceInt64Unique(slc []int64) []int64 {
    result := []int64{}  // 存放结果
    for i := range slc{
        flag := true
        for j := range result{
            if slc[i] == result[j] {
                flag = false  // 存在重复元素，标识为false
                break
            }
        }
        if flag {  // 标识为false，不添加进结果
            result = append(result, slc[i])
        }
    }
    return result
}