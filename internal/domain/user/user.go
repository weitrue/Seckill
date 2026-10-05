/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2021/4/12 上午11:50
 * Description: user 领域对象
 *   当前只持有一个请求鉴权后的用户信息值对象 Info,
 *   后续阶段补 User 聚合(昵称、注册时间、角色、权限等) + UserRepository 接口
 **/

package user

// Info 请求鉴权后的用户信息
//
//	由鉴权中间件(interfaces/api/middleware)从 token 解出,
//	通过 gin.Context 的 "userInfo" key 传给 handler
type Info struct {
	UID        string `json:"uid"`        // 用户 ID
	LoginTime  int64  `json:"loginTime"`  // 登录时间(秒)
	ExpireTime int64  `json:"expireTime"` // 过期时间(秒)
}
