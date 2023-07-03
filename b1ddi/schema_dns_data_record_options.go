package b1ddi

import (
	"terraform-provider-b1ddi/b1ddi/util"
)

func deepUpdateDataRecordOptions(d interface{}, recordType string) (interface{}, error) {
	if d == nil {
		return nil, nil
	}
	in := d.(map[string]interface{})
	switch recordType {
	case "A", "AAAA":
		createPtr, fieldExists, err := util.ToBool(in, "create_ptr")
		if err != nil {
			return nil, err
		}
		if fieldExists {
			in["create_ptr"] = createPtr
		}

		checkRmz, fieldExists, err := util.ToBool(in, "check_rmz")
		if err != nil {
			return nil, err
		}
		if fieldExists {
			in["check_rmz"] = checkRmz
		}
	}

	return in, nil
}
