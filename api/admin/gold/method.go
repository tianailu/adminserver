package gold

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tianailu/adminserver/api/admin/user/models"
	"github.com/tianailu/adminserver/pkg/common"
	"github.com/tianailu/adminserver/pkg/db/mysql"
	"log"
	"net/http"
	"strconv"
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

func GetGoldInfo(c echo.Context) error {
	var (
		resp common.Response
	)
	userID, err := strconv.Atoi(c.QueryParam("user_id"))
	if err != nil {
		resp.Status = 1
		resp.Msg = err.Error()
		return c.JSON(http.StatusOK, resp)
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		resp.Status = 1
		resp.Msg = err.Error()
		return c.JSON(http.StatusOK, resp)
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		resp.Status = 1
		resp.Msg = err.Error()
		return c.JSON(http.StatusOK, resp)
	}
	types, err := strconv.Atoi(c.QueryParam("types"))
	if err != nil {
		resp.Status = 1
		resp.Msg = err.Error()
		return c.JSON(http.StatusOK, resp)
	}
	status, err := strconv.Atoi(c.QueryParam("status"))
	if err != nil {
		resp.Status = 1
		resp.Msg = err.Error()
		return c.JSON(http.StatusOK, resp)
	}

	var results []GetGoldInfoResp

	// 使用GORM的Preload方法关联交易记录表，并使用Select方法指定要查询的字段
	u := models.User{}
	sqlSelect := fmt.Sprintf("%v.ID AS UserID, %v.Name AS UserName, CASE WHEN t1.UserID IS NOT NULL THEN 2  WHEN %v.RealAuth = true THEN 1  ELSE 0  END AS RealAuthExists, CASE WHEN t2.UserID IS NOT NULL THEN 2  WHEN %v.EducationAuth = true THEN 1 ELSE 0 END AS EducationAuthExists, CASE WHEN t3.UserID IS NOT NULL THEN 2 WHEN %v.WorkAuth = true THEN 1  ELSE 0 END AS WorkAuthExists, CASE WHEN %v.Avatar = true THEN '存在' ELSE '不存在' END AS AvatarExists", u.TableName(), u.TableName(), u.TableName(), u.TableName(), u.TableName(), u.TableName())
	sql := mysql.GetDB().
		Table(u.TableName()).
		Select(sqlSelect).
		Joins(
			fmt.Sprintf("LEFT JOIN (?) t1 ON %v.ID = t1.UserID", u.TableName()),
			mysql.GetDB().Table(TradeDBName).Select("UserID, 1 AS Exists").Where("type = ?", TradeType_RealCert),
		).
		Joins(
			fmt.Sprintf("LEFT JOIN (?) t2 ON %v.ID = t2.UserID", u.TableName()),
			mysql.GetDB().Table(TradeDBName).Select("UserID, 1 AS Exists").Where("type = ?", TradeType_AcademicCert),
		).
		Joins(
			fmt.Sprintf("LEFT JOIN (?) t3 ON %v.ID = t3.UserID", u.TableName()),
			mysql.GetDB().Table(TradeDBName).Select("UserID, 1 AS Exists").Where("type = ?", TradeType_JobCert),
		).
		Joins(
			fmt.Sprintf("LEFT JOIN (?) t4 ON %v.ID = t4.UserID", u.TableName()),
			mysql.GetDB().Table(TradeDBName).Select("UserID, 1 AS Exists").Where("type = ?", TradeType_SoulMate),
		)
	if userID != 0 {
		sql.Where(TradeDBName+".user_id = ?", userID)
	}
	switch types {
	case TradeType_RealCert:
		if status == 2 {
			sql.Where("t1.UserID IS NOT NULL")
		} else {
			sql.Where("t1.UserID IS NULL")
		}
	case TradeType_AcademicCert:
		if status == 2 {
			sql.Where("t2.UserID IS NOT NULL")
		} else {
			sql.Where("t2.UserID IS NULL")
		}
	case TradeType_JobCert:
		if status == 2 {
			sql.Where("t3.UserID IS NOT NULL")
		} else {
			sql.Where("t3.UserID IS NULL")
		}
	case TradeType_SoulMate:
		if status == 1 {
			sql.Where(u.TableName() + ".ta_% IS NOT NULL OR " + u.TableName() + ".be_% IS NOT NULL")
		} else {
			sql.Where(u.TableName() + ".ta_% IS NULL AND " + u.TableName() + ".be_% IS NULL")
		}
	case TradeType_Avatar:
		if status == 1 {
			sql.Where(u.TableName() + ".img IS NOT NULL")
		} else {
			sql.Where(u.TableName() + ".img IS NULL")
		}
	}
	if err := sql.Limit(limit).Offset(page).
		Find(&results).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch users"})
	}
	resp.Status = 0
	resp.Msg = "ok"
	resp.Data = &GetSettingResp{
		Data: results,
	}
	return nil
}
