//Copyright (c) [2021] [YangLei]
//[mogu-go] is licensed under Mulan PSL v2.
//You can use this software according to the terms and conditions of the Mulan PSL v2.
//You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
//EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
//MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PSL v2 for more details.

package page

type IPage struct {
	Records          interface{} `json:"records"`
	Total            int64       `json:"total"`
	Size             int         `json:"size"`
	Current          int         `json:"current"`
	OptimizeCountSql bool        `json:"optimizeCountSql"`
	IsSearchCount    bool        `json:"isSearchCount"`
}
