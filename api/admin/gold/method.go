package gold

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/tealeg/xlsx"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func SetInvite(c echo.Context) error {
	var (
		resp common.ResponseNoData
		req  = SetInviteReq{}
	)
	err := c.Bind(&req)
	if err != nil {
		resp.Status = 1
		resp.Msg = "get param failed"
		return c.JSON(http.StatusOK, resp)
	}

	if err := updateInviteInfo(c, &req); err != nil {
		resp.Status = 1
		resp.Msg = "set invite failed"
		return c.JSON(http.StatusOK, resp)
	}
	resp.Status = 0
	resp.Msg = "ok"
	return c.JSON(http.StatusOK, resp)
}
func SetTrade(c echo.Context) error {
	var (
		resp    common.ResponseNoData
		req     = SetTradeReq{}
		dbSlice = DBSlice[GoldTradeList]{}
	)
	err := c.Bind(&req)
	if err != nil {
		resp.Status = 1
		resp.Msg = "get param failed"
		return c.JSON(http.StatusOK, resp)
	}
	dbSlice = append(dbSlice, GoldTradeList{
		UserID:      req.UserID,
		Amount:      req.Amount,
		Type:        req.Type,
		TotalAmount: req.TotalAmount,
	})
	err = dbSlice.Set(TradeDBName)
	if err != nil {
		resp.Status = 1
		resp.Msg = "set data failed"
		return c.JSON(http.StatusOK, resp)
	}
	resp.Status = 0
	resp.Msg = "ok"
	return c.JSON(http.StatusOK, resp)
}

// 获取数据
func GetTrade(c echo.Context) error {
	var (
		resp    common.Response
		dbSlice DBSlice[GoldTradeList]
		pageNum int64
	)
	page := c.QueryParam("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Println("参数解析失败, err:", err)
		resp.Status = 1
		resp.Msg = "参数解析出错了， err：" + err.Error()
		return c.JSON(http.StatusOK, resp)
	}
	if pageInt <= 0 {
		pageInt = 1
	}

	size := c.QueryParam("size")
	sizeNum, err := strconv.Atoi(size)
	if err != nil {
		log.Println("参数解析失败, err:", err)
		resp.Status = 1
		resp.Msg = "参数解析出错了， err：" + err.Error()
		return c.JSON(http.StatusOK, resp)
	}
	if sizeNum <= 0 {
		sizeNum = 10
	}
	pageNum, err = dbSlice.GetOnePageWithOrder(pageInt, sizeNum, TradeDBName)
	if err != nil {
		log.Println("数据库查询失败, err:", err)
		resp.Status = 1
		resp.Msg = "查询出错了， err：" + err.Error()
		return c.JSON(http.StatusOK, resp)
	}
	resp.Status = 0
	resp.Msg = "ok"
	resp.Data = &GetTradeResp{
		Data: dbSlice,
		Page: pageNum + 1,
	}
	return c.JSON(http.StatusOK, resp)
}

func GetGoldSetting(c echo.Context) error {
	var (
		resp    common.Response
		dbValue DBValue[GoldSetting]
	)
	namespace := c.QueryParam("namespace")
	if len(namespace) == 0 {
		resp.Status = 1
		resp.Msg = "【namespace】参数不可以为空请重试"
		return c.JSON(http.StatusOK, resp)
	}
	if err := dbValue.GetSettingWithNameSpace(c.Request().Context(), InviteSettingDBName, namespace); err != nil {
		log.Println("数据库查询失败, err:", err)
		resp.Status = 1
		resp.Msg = "查询出错了， err：" + err.Error()
		return c.JSON(http.StatusOK, resp)
	}
	resp.Status = 0
	resp.Msg = "ok"
	switch namespace {
	case "name":
		resp.Data = &GetSettingResp{
			Data: dbValue.Data.GoldName,
		}
	case "info":
		resp.Data = &GetSettingResp{
			Data: dbValue.Data.InfoAmount,
		}
	case "sign":
		sign := &SignInSetting{}
		err := json.Unmarshal([]byte(dbValue.Data.SignIn), sign)
		if err != nil {
			log.Println("数据解析失败, err:", err)
			resp.Status = 1
			resp.Msg = "解析出错了， err：" + err.Error()
			return c.JSON(http.StatusOK, resp)
		}
		resp.Data = &GetSettingResp{
			Data: sign,
		}
	case "invite":
		invite := &InviteSetting{}
		err := json.Unmarshal([]byte(dbValue.Data.Invite), invite)
		if err != nil {
			log.Println("数据解析失败, err:", err)
			resp.Status = 1
			resp.Msg = "解析出错了， err：" + err.Error()
			return c.JSON(http.StatusOK, resp)
		}
		resp.Data = &GetSettingResp{
			Data: invite,
		}
	default:
		resp.Data = &GetSettingResp{
			Data: "namespace 输入错误",
		}
	}

	return nil
}

func SetGoldSetting(c echo.Context) error {
	var (
		resp    common.ResponseNoData
		req     = SetSettingReq{}
		dbValue = DBValue[GoldSetting]{}
	)
	err := c.Bind(&req)
	if err != nil {
		resp.Status = 1
		resp.Msg = "get param failed"
		return c.JSON(http.StatusOK, resp)
	}
	if err := dbValue.SetWithCache(c.Request().Context(), ChatSettingDBName, req.Param, req.Value); err != nil {
		resp.Status = 1
		resp.Msg = "set param failed"
		return c.JSON(http.StatusOK, resp)
	}
	return nil
}

func GetChatSetting(c echo.Context) error {
	var (
		resp    common.Response
		dbValue DBValue[ChatSetting]
	)

	if err := dbValue.GetSetting(c.Request().Context(), ChatSettingDBName); err != nil {
		log.Println("数据库查询失败, err:", err)
		resp.Status = 1
		resp.Msg = "查询出错了， err：" + err.Error()
		return c.JSON(http.StatusOK, resp)
	}
	resp.Status = 0
	resp.Msg = "ok"
	resp.Data = &GetSettingResp{
		Data: dbValue,
	}
	return nil
}

func SetChatSetting(c echo.Context) error {
	var (
		resp    common.ResponseNoData
		req     = SetSettingReq{}
		dbValue = DBValue[ChatSetting]{}
	)
	err := c.Bind(&req)
	if err != nil {
		resp.Status = 1
		resp.Msg = "get param  failed"
		return c.JSON(http.StatusOK, resp)
	}
	if err := dbValue.SetWithCache(c.Request().Context(), ChatSettingDBName, req.Param, req.Value); err != nil {
		resp.Status = 1
		resp.Msg = "set param failed"
		return c.JSON(http.StatusOK, resp)
	}
	return nil
}

func GetSignInExcel(c echo.Context) error {
	var (
		resp common.Response
	)

	// 获取前端传递的id列表，这里假设前端通过GET请求传递id参数
	idList := c.QueryParam("ids")
	if len(idList) == 0 {
		resp.Status = 0
		resp.Msg = ""
		return c.JSON(http.StatusOK, resp)
	}
	ids := strings.Split(idList, ",")

	// 确保id列表的长度不超过300个
	if len(ids) > 300 {
		resp.Status = 1
		resp.Msg = "id列表长度不能超过300个"
		return c.JSON(http.StatusOK, resp)
	}

	// 查询数据
	var data []GoldTradeList
	result := mysql.GetDB().Where("id IN (?)", ids).Find(&data)
	if result.Error != nil {
		resp.Status = 1
		resp.Msg = result.Error.Error()
		return c.JSON(http.StatusOK, resp)
	}

	// 创建Excel文件
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return c.String(http.StatusInternalServerError, "创建Sheet失败")
	}

	// 写入表头
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "ID"
	cell = row.AddCell()
	cell.Value = "用户ID"
	cell = row.AddCell()
	cell.Value = "用户名"
	cell = row.AddCell()
	cell.Value = "签到状态"
	cell = row.AddCell()
	cell.Value = "参与时间"

	// 将查询到的数据写入Excel文件
	for _, d := range data {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = strconv.Itoa(d.ID)
		cell = row.AddCell()
		cell.Value = string(d.UserID)
		cell = row.AddCell()
		cell.Value = string(d.UserID)
		cell = row.AddCell()
		cell.Value = "已完成"
		cell = row.AddCell()
		timeFormat := "2006-01-02 15:04:05" // 示例：年-月-日 时:分:秒
		// 使用 Format 方法将时间类型转换为格式化的字符串
		formattedTime := d.CreatedTime.Format(timeFormat)
		cell.Value = formattedTime
	}

	// 设置HTTP响应头，告诉浏览器下载Excel文件
	c.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set("Content-Disposition", "attachment; filename=export.xlsx")

	// Write the Excel file directly to the HTTP response
	return file.Write(c.Response())
}

func GetInviteExcel(c echo.Context) error {
	var (
		resp common.Response
	)

	// 获取前端传递的id列表，这里假设前端通过GET请求传递id参数
	idList := c.QueryParam("ids")
	if len(idList) == 0 {
		resp.Status = 0
		resp.Msg = ""
		return c.JSON(http.StatusOK, resp)
	}
	ids := strings.Split(idList, ",")

	// 确保id列表的长度不超过300个
	if len(ids) > 300 {
		resp.Status = 1
		resp.Msg = "id列表长度不能超过300个"
		return c.JSON(http.StatusOK, resp)
	}

	// 查询数据
	var data []InviteInfo
	result := mysql.GetDB().Where("id IN (?)", ids).Find(&data)
	if result.Error != nil {
		resp.Status = 1
		resp.Msg = result.Error.Error()
		return c.JSON(http.StatusOK, resp)
	}

	// 创建Excel文件
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return c.String(http.StatusInternalServerError, "创建Sheet失败")
	}

	// 写入表头
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "ID/用户ID"
	cell = row.AddCell()
	cell.Value = "邀请码"
	cell = row.AddCell()
	cell.Value = "邀请用户的用户名"
	cell = row.AddCell()
	cell.Value = "邀请用户获取金币数"
	cell = row.AddCell()
	cell.Value = "助力用户用户名"
	cell = row.AddCell()
	cell.Value = "助力用户获取金币数"
	cell = row.AddCell()
	cell.Value = "参与时间"

	// 将查询到的数据写入Excel文件
	for _, d := range data {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = strconv.Itoa(d.ID)
		cell = row.AddCell()
		cell.Value = string(d.InviteCode)
		cell = row.AddCell()
		cell.Value = string(d.InviteCode)
		cell = row.AddCell()
		cell.Value = "已完成"
		cell = row.AddCell()
		timeFormat := "2006-01-02 15:04:05" // 示例：年-月-日 时:分:秒
		// 使用 Format 方法将时间类型转换为格式化的字符串
		formattedTime := d.CreatedTime.Format(timeFormat)
		cell.Value = formattedTime
	}

	// 设置HTTP响应头，告诉浏览器下载Excel文件
	c.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set("Content-Disposition", "attachment; filename=export.xlsx")

	// Write the Excel file directly to the HTTP response
	return file.Write(c.Response())
}

func GetInviteCodeExcel(c echo.Context) error {
	var (
		resp common.Response
	)

	// 获取前端传递的id列表，这里假设前端通过GET请求传递id参数
	idList := c.QueryParam("ids")
	if len(idList) == 0 {
		resp.Status = 0
		resp.Msg = ""
		return c.JSON(http.StatusOK, resp)
	}
	ids := strings.Split(idList, ",")

	// 确保id列表的长度不超过300个
	if len(ids) > 300 {
		resp.Status = 1
		resp.Msg = "id列表长度不能超过300个"
		return c.JSON(http.StatusOK, resp)
	}

	// 查询数据
	var data []InviteInfo
	result := mysql.GetDB().Where("id IN (?)", ids).Find(&data)
	if result.Error != nil {
		resp.Status = 1
		resp.Msg = result.Error.Error()
		return c.JSON(http.StatusOK, resp)
	}

	// 创建Excel文件
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return c.String(http.StatusInternalServerError, "创建Sheet失败")
	}

	// 写入表头
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "ID"
	cell = row.AddCell()
	cell.Value = "用户ID"
	cell = row.AddCell()
	cell.Value = "用户名"
	cell = row.AddCell()
	cell.Value = "邀请码"
	cell = row.AddCell()
	cell.Value = "人数"
	cell = row.AddCell()
	cell.Value = "参与时间"

	// 将查询到的数据写入Excel文件
	for _, d := range data {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = strconv.Itoa(d.ID)
		cell = row.AddCell()
		cell.Value = string(d.InviteCode)
		cell = row.AddCell()
		cell.Value = string(d.InviteCode)
		cell = row.AddCell()
		cell.Value = "已完成"
		cell = row.AddCell()
		timeFormat := "2006-01-02 15:04:05" // 示例：年-月-日 时:分:秒
		// 使用 Format 方法将时间类型转换为格式化的字符串
		formattedTime := d.CreatedTime.Format(timeFormat)
		cell.Value = formattedTime
	}

	// 设置HTTP响应头，告诉浏览器下载Excel文件
	c.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set("Content-Disposition", "attachment; filename=export.xlsx")

	// Write the Excel file directly to the HTTP response
	return file.Write(c.Response())
}

func GetShareExcel(c echo.Context) error {
	var (
		resp common.Response
	)

	// 获取前端传递的id列表，这里假设前端通过GET请求传递id参数
	idList := c.QueryParam("ids")
	if len(idList) == 0 {
		resp.Status = 0
		resp.Msg = ""
		return c.JSON(http.StatusOK, resp)
	}
	ids := strings.Split(idList, ",")

	// 确保id列表的长度不超过300个
	if len(ids) > 300 {
		resp.Status = 1
		resp.Msg = "id列表长度不能超过300个"
		return c.JSON(http.StatusOK, resp)
	}

	// 查询数据
	var data []InviteInfo
	result := mysql.GetDB().Where("id IN (?)", ids).Find(&data)
	if result.Error != nil {
		resp.Status = 1
		resp.Msg = result.Error.Error()
		return c.JSON(http.StatusOK, resp)
	}

	// 创建Excel文件
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		return c.String(http.StatusInternalServerError, "创建Sheet失败")
	}

	// 写入表头
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "ID"
	cell = row.AddCell()
	cell.Value = "用户ID"
	cell = row.AddCell()
	cell.Value = "用户名"
	cell = row.AddCell()
	cell.Value = "邀请码"
	cell = row.AddCell()
	cell.Value = "分享类型"
	cell = row.AddCell()
	cell.Value = "参与时间"

	// 将查询到的数据写入Excel文件
	for _, d := range data {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = strconv.Itoa(d.ID)
		cell = row.AddCell()
		cell.Value = string(d.InviteCode)
		cell = row.AddCell()
		cell.Value = string(d.InviteCode)
		cell = row.AddCell()
		cell.Value = "已完成"
		cell = row.AddCell()
		timeFormat := "2006-01-02 15:04:05" // 示例：年-月-日 时:分:秒
		// 使用 Format 方法将时间类型转换为格式化的字符串
		formattedTime := d.CreatedTime.Format(timeFormat)
		cell.Value = formattedTime
	}

	// 设置HTTP响应头，告诉浏览器下载Excel文件
	c.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set("Content-Disposition", "attachment; filename=export.xlsx")

	// Write the Excel file directly to the HTTP response
	return file.Write(c.Response())
}
