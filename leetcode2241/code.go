package leetcode2241

type ATM struct {
	orderList []int
}

func Constructor() ATM {
	return ATM{
		make([]int, 5),
	}
}

var price [5]int = [5]int{20, 50, 100, 200, 500}

func (this *ATM) Deposit(banknotesCount []int) {
	if len(banknotesCount) != 5 {
		return // 需要确保输入的 banknotesCount 长度为 5
	}
	for i, count := range banknotesCount {
		this.orderList[i] += count
	}
}

func (this *ATM) Withdraw(amount int) []int {
	tmp := make([]int, 5)
	for i := 4; i >= 0; i-- {
		tmp[i] = min(amount/price[i], this.orderList[i])
		amount -= tmp[i] * price[i]
	}
	if amount > 0 {
		return []int{-1} // 无法提取所需金额
	}
	for idx, v := range tmp {
		this.orderList[idx] -= v
	}
	return tmp // 返回提取的面额数量
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 判断 现有的钱 怎么最大的找零
/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */
