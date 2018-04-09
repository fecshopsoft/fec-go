package helper

import(
    "regexp"
    "net/url"
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
    RegExp := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`)
    return RegExp.MatchString(domain)
}


func IsValidUrl(toTest string) bool {
    _, err := url.Parse(toTest)
    if err != nil {
        return false
    } else {
        return true
    }
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



