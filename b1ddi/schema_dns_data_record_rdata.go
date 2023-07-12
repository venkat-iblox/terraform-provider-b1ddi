package b1ddi

import (
	"terraform-provider-b1ddi/b1ddi/util"
)

/*
	updateDataRecordRData helps convert rDATA record fields of type integer from string value
    Introduced to fix issues where terraform converts integer values to string in rendered config
*/
func updateDataRecordRData(d interface{}, recordType string) (interface{}, error) {
	if d == nil {
		return nil, nil
	}

	in := d.(map[string]interface{})
	switch recordType {
	case "MX":
		toInt, exists, err := util.ToInt(in, "preference")
		if err != nil {
			return nil, err
		}
		if exists {
			in["preference"] = toInt
		}

	case "CAA":
		flags, exists, err := util.ToInt(in, "flags")
		if err != nil {
			return nil, err
		}
		if exists {
			in["flags"] = flags
		}

	case "NAPTR":
		order, exists, err := util.ToInt(in, "order")
		if err != nil {
			return nil, err
		}
		if exists {
			in["order"] = order
		}

		toInt, exists, err := util.ToInt(in, "preference")
		if err != nil {
			return nil, err
		}
		if exists {
			in["preference"] = toInt
		}

	case "SOA":
		serial, exists, err := util.ToInt(in, "serial")
		if err != nil {
			return nil, err
		}
		if exists {
			in["serial"] = serial
		}

	case "SRV":
		port, exists, err := util.ToInt(in, "port")
		if err != nil {
			return nil, err
		}
		if exists {
			in["port"] = port
		}

		priority, exists, err := util.ToInt(in, "priority")
		if err != nil {
			return nil, err
		}
		if exists {
			in["priority"] = priority
		}

		weight, exists, err := util.ToInt(in, "weight")
		if err != nil {
			return nil, err
		}
		if exists {
			in["weight"] = weight
		}
	default:
		return d, nil
	}

	return in, nil
}
