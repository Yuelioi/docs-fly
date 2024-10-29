package controllers

import (
	"docsfly/internal/common"
	"docsfly/internal/config"
	"docsfly/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type OtherController struct{}

func (*OtherController) Register(engine *gin.Engine) {
	engine.POST("/"+config.Instance.App.ApiVersion+"/ip", VisitorInsertLog)

	engine.GET("/"+config.Instance.App.ApiVersion+"/app/version", GetAppVersion)
	engine.GET("/"+config.Instance.App.ApiVersion+"/rand/nickname", GetRndName)
	engine.GET("/"+config.Instance.App.ApiVersion+"/rand/poem", GetRndPoem)
	engine.GET("/"+config.Instance.App.ApiVersion+"/rand/post", GetRndPost)

	engine.GET("/"+config.Instance.App.ApiVersion+"/vendor/yiyan", GetYiYan)
}

func rndName() string {

	el1 := []string{"废墟", "深海", "反应堆", "学园", "腐烂", "东京", "三维", "四次元", "少管所", "流星", "闪光", "南极", "消极", "幽浮", "网路", "暗狱", "离子态", "液态", "黑色", "抱抱", "暴力", "垃圾", "残暴", "残酷", "工口", "原味", "毛茸茸", "香香", "霹雳", "午夜", "美工刀", "爆浆", "机关枪", "无响应", "手术台", "麻风病", "虚拟", "速冻", "智能", "2000", "甜味", "华丽", "玛利亚", "无", "梦之", "蔷薇", "无政府", "酷酷", "西伯利亚", "人造", "法外", "追杀", "通缉", "女子", "微型", "男子", "超", "毁灭", "大型", "绝望", "阴间", "死亡", "坟场", "高科技", "奇妙", "魔法", "极限", "社会主义", "无聊"}
	el2 := []string{"小丑", "仿生", "纳米", "原子", "丧", "电子", "十字架", "咩咩", "赛博", "野猪", "外星", "窒息", "变态", "触手", "小众", "悲情", "飞行", "绿色", "电动", "铁锈", "碎尸", "电音", "蠕动", "酸甜", "虚构", "乱码", "碳水", "内脏", "脑浆", "血管", "绷带", "不合格", "光滑", "标本", "酸性", "碱性", "404", "变身", "反常", "樱桃", "碳基", "矫情", "病娇", "进化", "潮湿", "砂糖", "高潮", "变异", "复合盐", "伏特加", "抑郁", "暴躁", "不爱说话", "废物", "失败", "幻想型", "社恐", "苦涩", "粘液", "浓厚", "快乐", "强制", "中二病", "恶魔", "emo", "激光", "发射", "限量版", "迷因", "堕落", "放射性"}
	el3 := []string{"天使", "精灵", "女孩", "男孩", "宝贝", "小妈咪", "虫", "菇", "公主", "少女", "少年", "1号机", "子", "恐龙", "蜈蚣", "蟑螂", "食人鱼", "小飞船", "舞女", "桃子", "团子", "精", "酱", "废料", "生物", "物质", "奶茶", "搅拌机", "液", "火锅", "祭司", "体", "实验品", "试验体", "小猫咪", "样本", "颗粒", "血块", "汽水", "蛙", "软体", "机器人", "人质", "小熊", "圣母", "胶囊", "乙女", "主义者", "屑", "垢", "污渍", "废人", "毛血旺", "怪人", "肉", "河豚", "豚", "藻类", "唾沫", "咒语", "建筑", "球", "小狗", "碳", "元素", "少先队员", "博士", "糖"}

	selectedEl1 := el1[rand.Intn(len(el1))]
	selectedEl2 := el2[rand.Intn(len(el2))]
	selectedEl3 := el3[rand.Intn(len(el3))]

	// 拼接选取的元素
	return strings.Join([]string{selectedEl1, selectedEl2, selectedEl3}, "")

}

func rndPoem() string {

	poems := []string{
		"醉后不知天在水，满船清梦压星河。",
		"羽衣常带烟霞色，不惹人间桃李花。",
		"满堂花醉三千客，一剑霜寒十四州。",
		"我见青山多妩媚，料青山见我应如是。",
		"折花逢驿使，寄与陇头人。",
		"应是天仙狂醉，乱把白云揉碎。",
		"人间自是有情痴，此恨不关风与月。",
		"晚来天欲雪，能饮一杯无？",
		"愿我如星君如月，夜夜流光相皎洁。",
		"最是人间留不住，朱颜辞镜花辞树。",
		"凤凰台上凤凰游，风去台空江自流。",
		"曾经沧海难为水，除却巫山不是云。",
		"从此无心爱良夜，任他明月下西楼。",
		"觉后不知明月上，满身花影倩人扶。",
		"天上白玉京，十二楼五城。仙人抚我顶，结发受长生。",
		"万一禅关砉然破，美人如玉剑如虹。",
		"吹灭读书灯，一身都是月。",
		"垆边人似月，皓腕凝霜雪。",
	}

	selected := poems[rand.Intn(len(poems))]

	// 拼接选取的元素
	return selected

}

func VisitorInsertLog(c *gin.Context) {

	url := c.Query("url")
	dbContext, exists := c.Get("db")
	if !exists {
		return
	}
	db := dbContext.(*gorm.DB)

	var count int64
	db.Scopes(common.BasicModel, common.MatchUrlPath(url)).Count(&count)

	if count == 0 {
		ReturnFailResponse(c, http.StatusInternalServerError, "Can't find target link")
		return
	}

	urlList := strings.Split(url, "/")

	var category, book, locale string

	if len(urlList) > 2 {
		category = urlList[0]
		book = urlList[1]
		locale = urlList[2]
	} else {
		ReturnFailResponse(c, http.StatusInternalServerError, "Can't find target path")
		return
	}

	today := time.Now().Local()

	vs := models.Visitor{
		IP:       c.ClientIP(),
		URL:      url,
		Time:     today,
		Category: category,
		Book:     book,
		Locale:   locale,
	}

	db.Model(&models.Visitor{}).Create(&vs)

	// 返回 IP 地址给客户端
	ReturnSuccessResponse(c, gin.H{"message": "success"})
}

func GetAppVersion(c *gin.Context) {
	ReturnSuccessResponse(c, config.Instance.App.AppVersion)
}

func GetRndName(c *gin.Context) {
	ReturnSuccessResponse(c, rndName())
}

func GetRndPost(c *gin.Context) {

	dbContext, exists := c.Get("db")
	if !exists {
		return
	}
	var doc models.Entry

	db := dbContext.(*gorm.DB)
	if err := db.Scopes(common.BasicModel, common.FindFile).Order("RANDOM()").First(&doc).Error; err != nil {
		ReturnFailResponse(c, http.StatusInternalServerError, "Could not retrieve a random post")
		return
	}

	ReturnSuccessResponse(c, doc.MetaData)
}

func GetRndPoem(c *gin.Context) {
	ReturnSuccessResponse(c, rndPoem())
}

type Hitokoto struct {
	ID         int    `json:"id"`
	UUID       string `json:"uuid"`
	Hitokoto   string `json:"hitokoto"`
	Type       string `json:"type"`
	From       string `json:"from"`
	FromWho    string `json:"from_who"`
	Creator    string `json:"creator"`
	CreatorUID int    `json:"creator_uid"`
	Reviewer   int    `json:"reviewer"`
	CommitFrom string `json:"commit_from"`
	CreatedAt  string `json:"created_at"`
	Length     int    `json:"length"`
}

func yiyan() (Hitokoto, error) {
	url := "https://v1.hitokoto.cn/?c=b"

	// 发送 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("请求失败:", err)
		return Hitokoto{}, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体失败:", err)
		return Hitokoto{}, err
	}

	var hitokoto Hitokoto
	json.Unmarshal(body, &hitokoto)

	// 输出响应体
	return hitokoto, nil
}

func GetYiYan(c *gin.Context) {
	hitokoto, _ := yiyan()
	ReturnSuccessResponse(c, hitokoto)
}
