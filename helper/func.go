package helper


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