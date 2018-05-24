package helper

import(
    // "regexp"
    "net/url"
    "time"
    "strings"
    "strconv"
    uuid "github.com/satori/go.uuid"
    "github.com/fecshopsoft/fec-go/security"
)

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

type VueSelectStrOps struct{
    Key string `form:"key" json:"key"`
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


func IsValidDomain(domain string) bool {
    // RegExp := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`)
    // return RegExp.MatchString(domain)
    return true
    /*
    _, err := url.ParseRequestURI("http://"+domain)
    if err == nil {
       return true
    }
    return false
    */
}


func IsValidUrl(toTest string) bool {
    _, err := url.Parse(toTest)
    if err != nil {
        return false
    } else {
        return true
    }
}

// 字符串全部替换
func StrReplace(str string, oldStr string, newStr string) string {
    return strings.Replace(str, oldStr, newStr, -1)
}
 


//截取字符串 start 起点下标 end 终点下标(不包括)
func Substr(str string, start int, end int) string {
    rs := []rune(str)
    length := len(rs)

    if start < 0 || start > length {
        return ""
    }

    if end < 0 || end > length {
        return ""
    }
    return string(rs[start:end])
}

// 字符串包含  strings.Contains("widuu", "wi")
func StrContains(strFull string, strContains string) bool {
    return strings.Contains(strFull, strContains)
} 

// 得到随机uuid
func RandomUUID() string{
    u := uuid.Must(uuid.NewV4())
    return u.String()
}


// 通过随机生成的uuid，得到access token
func GenerateAccessToken() (string, error){
    uuid := RandomUUID()
    return security.JwtSignToken(uuid)
}

// siteUid 是传递的网站的website_uid的字符串
func GenerateAccessTokenBySiteId(siteUid string) (string, error){
    return security.JwtSignAccessToken(siteUid)
}

func GetSiteUIdByAccessToken(accessToken string) (string, error){
    return security.JwtParseAccessToken(accessToken)
}
// 通过字符串DateStr，得到时间戳
func GetTimestampsByDate(dateStr string) int64 {
    time, _ := time.Parse("2006-01-02", dateStr)
    return time.Unix()
}
// 通过字符串dateTimeStr，得到时间戳
func GetTimestampsByDateTime(dateTimeStr string) int64 {
    time, _ := time.Parse("2006-01-02 03:04:05", dateTimeStr)
    return time.Unix()
}

// 得到当前的时间戳
func DateTimestamps() int64 {
    return  time.Now().Unix()
}

// 通过时间戳，得到字符串，如果传递的时间戳为0，则取当前时间
// 时间格式为 Y-m-d H:i:s
func DateTimeUTCStr() string {
    loc, _ := time.LoadLocation("UTC")
    now := time.Now().In(loc)
    return now.Format("2006-01-02 03:04:05")
    
}

// 
// 当前时间字符串时间格式，时间格式为 Y-m-d 
func DateUTCStr() string {
    dateTimeStr := DateTimeUTCStr()
    return dateTimeStr[0:10]
}

// 根据时间戳，得到UTC时区的时间，格式为:2006-01-02 03:04:05
func GetDateUtcByTimestamps(timestamps int64) string {
    loc, _ := time.LoadLocation("UTC")
    nTime := time.Unix(timestamps, 0).In(loc)
    return nTime.Format("2006-01-02 03:04:05")

}
// 根据时间戳，得到UTC时区的时间，格式为:2006-01-02
func GetDateTimeUtcByTimestamps(timestamps int64) string {
    dateStr := GetDateUtcByTimestamps(timestamps)
    return dateStr[0:10]
}
// 字符串转换成数字int
func Int(str string) (int, error) {
    return strconv.Atoi(str)
    
}
// 字符串转换成数字int64
func Int64(str string) (int64, error) {
    c, err := strconv.Atoi(str)
    if err == nil {
        return int64(c), err
    }
    return 0, err
}
// 数字转换成字符串
func Str(c int) (string) {
    return strconv.Itoa(c)
}
// 数字int64 转换成字符串
func Str64(c int64) (string) {
    return strconv.Itoa(int(c))
}
// 字符串转换成数字Float64
func Float64(str string) (float64, error) {
    f, err := strconv.ParseFloat(str, 64)
    if err != nil {
        return 0, err
    }
    return f, nil
}

// 两个slice string的差集
func ArrayDiff(s1 []string, s2 []string) []string{
    var s3 []string 
    var k int
    if len(s1) == 0 {
        return s2
    }
    if len(s2) == 0 {
        return s3
    }
    for i:=0; i<len(s2); i++ {
        s2_s := s2[i]
        k = 0;
        for j:=0; j<len(s1); j++ {
            s1_s := s1[j]
            if s1_s == s2_s {
                k = 1;
                break;
            }
        }
        if k == 0 {
            s3 = append(s3, s2_s)
        }
    }
    return s3
}
// slice 并集，并做唯一处理
func ArrayMergeAndUnique(s1 []string, s2 []string) []string{
    var s3 []string 
    m := make(map[string]string)
    for i:=0; i<len(s1); i++ {
        s1_s := s1[i]
        if _, ok := m[s1_s]; !ok {
            m[s1_s] = s1_s
            s3 = append(s3, s1_s)
        }
    }
    for i:=0; i<len(s2); i++ {
        s2_s := s2[i]
        if _, ok := m[s2_s]; !ok {
            m[s2_s] = s2_s
            s3 = append(s3, s2_s)
        }
    }

    return s3
}
// slice 唯一
func ArrayUnique(s1 []string) []string {
    var s3 []string 
    m := make(map[string]string)
    for i:=0; i<len(s1); i++ {
        s1_s := s1[i]
        if _, ok := m[s1_s]; !ok {
            m[s1_s] = s1_s
            s3 = append(s3, s1_s)
        }
    }
    return s3
}