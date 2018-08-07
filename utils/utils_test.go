package utils

import "testing"

func TestParamsOrDefault(t *testing.T) {
	t.Log(ParamsOrDefault("cdscsd/xsaxsacds/cdsxsa", 4), len(ParamsOrDefault("cdscsd/xsaxsacds/cdsxsa", 4)))
}
