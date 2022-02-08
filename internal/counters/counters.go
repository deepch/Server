/*
	Autor Andrey Semochkin
*/

package counters

/*

	Counters descriptions element statistic

*/

//Counters - main level statistic
type Counters struct {
	Devices  map[string]*deviceCounters `json:"Devices,omitempty" groups:"api"`
	Platform platformCounters           `json:"Platform,omitempty" groups:"api"`
	Client   map[string]*clientCounters `json:"Client,omitempty" groups:"api"`
}

//deviceCounters - device statistic
type deviceCounters struct {
	InBytes            int64                      `json:"InBytes,omitempty" groups:"api"`
	OutBytes           int64                      `json:"OutBytes,omitempty" groups:"api"`
	InTrafficTimeLine  map[int64]int64            `json:"InTrafficTimeLine,omitempty" groups:"api"`
	OutTrafficTimeLine map[int64]int64            `json:"OutTrafficTimeLine,omitempty" groups:"api"`
	Channel            map[string]channelCounters `json:"Channel,omitempty" groups:"api"`
}

//deviceCounters - device channel statistic
type channelCounters struct {
	InBytes            int64           `json:"InBytes,omitempty" groups:"api"`
	OutBytes           int64           `json:"OutBytes,omitempty" groups:"api"`
	InTrafficTimeLine  map[int64]int64 `json:"InTrafficTimeLine,omitempty" groups:"api"`
	OutTrafficTimeLine map[int64]int64 `json:"OutTrafficTimeLine,omitempty" groups:"api"`
}

//platformCounters - platform statistic
type platformCounters struct {
	ReadBytes      int64           `json:"ReadBytes,omitempty" groups:"api"`
	WriteBytes     int64           `json:"WriteBytes,omitempty" groups:"api"`
	MemoryTimeLine map[int64]int64 `json:"MemoryTimeLine,omitempty" groups:"api"`
	CPUTimeLine    map[int64]int64 `json:"CpuTimeLine,omitempty" groups:"api"`
}

//clientCounters - client statistic
type clientCounters struct {
	InBytes            int64           `json:"InBytes,omitempty" groups:"api"`
	OutBytes           int64           `json:"OutBytes,omitempty" groups:"api"`
	InTrafficTimeLine  map[int64]int64 `json:"InTrafficTimeLine,omitempty" groups:"api"`
	OutTrafficTimeLine map[int64]int64 `json:"OutTrafficTimeLine,omitempty" groups:"api"`
}
