package main
/**
 * 这个是脚本处理的入口部分
 * 1.初始化log输出
 * 2.远程授权认证（这个已经去掉）
 * 3.设置shell内容的log输出，输出到屏幕还是log文件
 * 4.开始跑统计脚本
 */
import(
    "github.com/fecshopsoft/fec-go/initialization"
    "log"
    "time"
    "github.com/fecshopsoft/fec-go/shell"

    "github.com/fecshopsoft/fec-go/config"
    "github.com/fecshopsoft/fec-go/security"
    "github.com/fecshopsoft/fec-go/helper"
    "io/ioutil"
    "net/http"
    "encoding/json"
)

func main() {
    // 初始化log输出，log.Println("---") 输出的内容将输出到globalLog文件里面
    log.Println("------start：" + time.Now().String())
    // 进行验证，将配置中的信息取出来，然后进行jwt处理，然后远程访问。
    initialization.InitShellLog()
    // 进行远程授权验证
    // 脚本执行部分，不再做授权验证，只对服务提供的go入口添加远程验证。
    // verify()
    log.SetFlags(log.LstdFlags | log.Llongfile)
    // 脚本处理，入口函数。
    shell.GoShell()
    // shell.TestEs()
    log.Println("------end：" + time.Now().String())

}
// 脚本端不做远程安全认证了,
// 该函数废弃。
func verify() {
    userName := config.Get("userName")
    userEmail := config.Get("userEmail")
    userTelephone := config.Get("userTelephone")
    userToken := config.Get("userToken")
    timeStamps := helper.DateTimestamps()
    urlBase := config.Get("userUrl")
    // 加密，传输
    s, err := security.JwtSignToken(map[string]string{
        "name": userName,
        "email": userEmail,
        "telephone": userTelephone,
        "token": userToken,
        "time": helper.Str64(timeStamps),
    })
    if err != nil {
        return
    }
    url := urlBase + "?tk=" + s
    log.Println("url:" + url)
    resData, err := httpGet(url)
    // 进行远程验证
    log.Println(resData)
    // 进行初始化
    // 检测网站的数量，以及数据的个数，满足条件，才进行数据统计
    // 初始化参数
    websiteCount := int(resData.SiteCount)
    pvCount := int(resData.PvCount)
    timestamps := resData.Timestamps

    // 判断时间间隔不要超过5分钟否则，将报错
    currentTimestamps := helper.DateTimestamps()
    if currentTimestamps - timestamps > 300 {
        log.Println("curl url from remot, Time out")
        return
    }

    if websiteCount > shell.WebsiteCount {
        shell.WebsiteCount = websiteCount
    }
    if pvCount > shell.PvCount {
        shell.PvCount = pvCount
    }

    return
}
// 该变量废弃。
type ResData struct{
    Status string `form:"status" json:"status"`
    PvCount int64 `form:"pv_count" json:"pv_count"`
    SiteCount int64 `form:"site_count" json:"site_count"`
    ErrorInfo string `form:"error_info" json:"error_info"`
    Timestamps int64 `form:"timestamps" json:"timestamps"`
}
// 该变量废弃。
func httpGet(url string) (ResData, error) {
    var message ResData
    resp, err := http.Get(url)
    if err != nil {
        return message, err
    }
    defer resp.Body.Close()
    bodyData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return message, err
    }
    log.Println("bodyData")
    log.Println(bodyData)
    // 解析出来
    str, err := security.JwtParseAccessToken(string(bodyData[:]))
    log.Println("str")
    log.Println(str)
    err = json.Unmarshal([]byte(str), &message)
    log.Println("message")
    log.Println(message)
    return message, nil
}












