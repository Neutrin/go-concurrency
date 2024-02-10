package main

type DBLoader struct {
	mp map[string]string
}

func NewDBLoader() CacheLoader {
	return &DBLoader{mp: map[string]string{
		"key one":   "key one",
		"key two ":  "key two",
		"key three": "key three"}}
}

func (loader *DBLoader) Load(key string) string {
	if loader.mp != nil {
		return loader.mp[key]
	}
	return ""
}
