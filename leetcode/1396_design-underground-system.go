package leetcode

type user struct {
	id               int
	startTime        int
	endTime          int
	startStationName string
	endStatationName string
}

type UndergroundSystem struct {
	userMap map[int]*user
	pathMap map[string][]int // 存放对应路程的用时，用于计算平均时间
}

func Constructor() UndergroundSystem {
	// 初始化
	return UndergroundSystem{make(map[int]*user), make(map[string][]int)}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	//userMap 添加
	this.userMap[id] = &user{id: id} // 重点！！！ 这里多层嵌套指针的时候现需要新建指针初始化，否则会报错找不到该地址
	this.userMap[id].startTime = t
	this.userMap[id].startStationName = stationName

}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	//出站时，更新user.end* ;append pathmap
	useTime := 0
	if _, ok := this.userMap[id]; ok {
		this.userMap[id].endTime = t
		this.userMap[id].endStatationName = stationName
	}

	mapKey := this.userMap[id].startStationName + "->" + this.userMap[id].endStatationName
	useTime = this.userMap[id].endTime - this.userMap[id].startTime
	this.pathMap[mapKey] = append(this.pathMap[mapKey], useTime)

}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	mapKey := startStation + "->" + endStation
	return average(this.pathMap[mapKey])
}

func average(s []int) float64 {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return float64(sum) / float64(len(s))
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
