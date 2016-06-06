package data_types

type ResLoadBalancer struct {
	AccountId             int    `json:"accountId"`
	ConnectionLimit       int    `json:"connectionLimit"`
	DedicatedFlag         bool   `json:"dedicatedFlag"`
	Id                    int    `json:"id"`
	IpAddressId           int    `json:"ipAddressId"`
	HighAvailabilityFlag  bool   `json:"highAvailabilityFlag"`
}
