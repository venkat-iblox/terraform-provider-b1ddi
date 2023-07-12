package b1ddi

import (
	"terraform-provider-b1ddi/b1ddi/util"
)

/*
	updateDataRecordOptions helps convert string options(supposed to be boolean) values into boolean
    Introduced to fix issues where terraform converts boolean values to string in rendered config
*/
func updateDataRecordOptions(d interface{}, recordType string) (interface{}, error) {
	if d == nil {
		return nil, nil
	}
	in := d.(map[string]interface{})
	switch recordType {
	case "A", "AAAA":
		createPtr, exists, err := util.ToBool(in, "create_ptr")
		if err != nil {
			return nil, err
		}
		if exists {
			in["create_ptr"] = createPtr
		}
		checkRmz, exists, err := util.ToBool(in, "check_rmz")
		if err != nil {
			return nil, err
		}
		if exists {
			in["check_rmz"] = checkRmz
		}
	}

	return in, nil
}
