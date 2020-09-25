package main

import (
	"fmt"
	"strings"
)

/**
	算法描述
	给定一个公式 提取里面的变量
	*变量以@开头
	*变量@ab@+@c 不合法的 因为@ab@后面没有值即一个变量不能包含多个@
	*假设操作符只有+,-,*,/

	例如:
	@a + @b / @c  =>["a","b","c"]

	@a + @b@   //=>["a"]
 */
 /**
 	函数体
  */
 func getVariable(formula string)[]string{
	//声明返回
	var result []string

	//将" " => ""
	formula = strings.Replace(formula," ","",-1)

	// 验证参数
	if len(formula) == 0 || len(formula) == 1{
		return result
	}

	//定义操作符
	operator := make(map[byte]int)  //默认operator[key] = 0(int默认值)
	operator['+'] = 1
	operator['-'] = 1
	operator['*'] = 1
	operator['/'] = 1

	//添加哨兵
	operator['@'] = 1

	//双指针解法
	i := 0
	length := len(formula)

	//寻找
	for i < length {
		//判断
		if formula[i] == '@' {
			//快指针
			j := i + 1

			//遍历
			for j < length && operator[formula[j]] != 1 {
				j ++
			}

			//判断是否合法---遇到哨兵
			if j < length && formula[j] == '@' {
				break
			}
			//是否为@+这种情况
			if j - i == 1 {
				break
			}

			//获取变量
			result = append(result,string(formula[i + 1:j]))

			//i 下移
			i = j + 1 //j此时为符号故 j要下移

		} else {
			i ++
		}
	}

	//返回
	return result
 }

func  main()()  {
	formula := "@bujianjijia_NWduikai_NWduikaicost+@bujianjijia_NWbakai_NWbakaicost+@bujianjijia_DYduikai_dyduikaicost+@bujianjijia_DYbakai_DYBKcost+@bujianjijia_zbinfo_ZBcost+@bujianjijia_biaomiaogongyijf_SCbiaomain+@bujianjijia_zhuangjia+@bujianjijia_Zn+@bujianjijia_yinjia+noRpnIf(@bujianjijia_yinshuaseshuoming==0,0,0)+@bujianjijia_shangxia+@bujianjijia_zuoyou+@bujianjijia_banshu+if(@bujianjijia_zhankaichic==0,0,0)+@bujianjijia_NWzheye_NWzheyecost+@bujianjijia_produandao_duandaocost+@bujianjijia_NWZYzhizuo+@bujianjijia_blzhizuoshul+@bujianjijia_bailiaocaiq_bailiaocost+if(@bujianjijia_yinshuagongyizs==0,0,0)+@bujianjijia_ysjijia_seqty+if(@bujianjijia_DYduikai_yinzhangmingxi==0,0,0)+if(@bujianjijia_DYbakai_yinzhangmingxi==0,0,0)+if(@bujianjijia_NWduikai_mingxi==0,0,0)+if(@bujianjijia_NWbakai_yinzhangmingxi==0,0,0)+if(@bujianjijia_NWduikai_seshu==0,0,0)+if(@bujianjijia_NWbakai_seshu==0,0,0)+if(@bujianjijia_DYduikai_seshu==0,0,0)+if(@bujianjijia_DYbakai_seshu==0,0,0)+if(@bujianjijia_zbinfo_banbie==0,0,0)+if(@bujianjijia_NWduikai_tao==0,0,0)+if(@bujianjijia_NWbakai_tao==0,0,0)+if(@bujianjijia_DYduikai_tao==0,0,0)+if(@bujianjijia_DYbakai_tao==0,0,0)+@bujianjijia_ysjijia_ksqty+if(@bujianjijia_bailiaocaiq_KLsize==0,0,0)+if(@bujianjijia_NWduikai_zhizhangdazhangshu==0,0,0)+if(@bujianjijia_NWbakai_dazhangshu==0,0,0)+if(@bujianjijia_DYduikai_zhizhangdazhangshu==0,0,0)+if(@bujianjijia_DYbakai_zhizhangdazhangshu==0,0,0)+if(@bujianjijia_NWduikai_laiyuan==0,0,0)+if(@bujianjijia_NWbakai_laiyuan==0,0,0)+if(@bujianjijia_DYduikai_laiyuan==0,0,0)+if(@bujianjijia_DYbakai_liayuan==0,0,0)+@bujianjijia_kaishu@bujianjijia_NWduikai_lingshu+@bujianjijia_NWbakai_lingshu+@bujianjijia_DYduikai_lingshu+@bujianjijia_DYbakai_lingshu+if(@bujianjijia_produandao_duandaozheye==0,0,0)+if(@bujianjijia_NWzheye_zheyefangshi==0,0,0)+if(@bujianjijia_bailiaocaiq_wuliao==0,0,0)"

	fmt.Println(getVariable(formula))
}