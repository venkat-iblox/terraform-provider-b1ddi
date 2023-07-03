package b1ddi

import (
	"terraform-provider-b1ddi/b1ddi/util"
)

func deepUpdateDataRecordRData(d interface{}, recordType string) (interface{}, error) {
	if d == nil {
		return nil, nil
	}

	in := d.(map[string]interface{})
	switch recordType {
	case "MX":
		toInt, fieldExists, err := util.ToInt(in, "preference")
		if err != nil {
			return nil, err
		}
		if fieldExists {
			in["preference"] = toInt
		}
	case "CAA":
		flags, fieldExists, err := util.ToInt(in, "flags")
		if err != nil {
			return nil, err
		}
		if fieldExists {
			in["flags"] = flags
		}

	case "NAPTR":
		order, fieldExists, err := util.ToInt(in, "order")
		if err != nil {
			return nil, err
		}
		if fieldExists {
			in["order"] = order
		}

		toInt, fieldExists, err := util.ToInt(in, "preference")
		if err != nil {
			return nil, err
		}
		if fieldExists {
			in["preference"] = toInt
		}

	case "SOA":
		serial, fieldExists, err := util.ToInt(in, "serial")
		if err != nil {
			return nil, err
		}
		if fieldExists {
			in["serial"] = serial
		}

	case "SRV":
		port, fieldExists, err := util.ToInt(in, "port")
		if err != nil {
			return nil, err
		}
		if fieldExists {
			in["port"] = port
		}

		priority, fieldExists, err := util.ToInt(in, "priority")
		if err != nil {
			return nil, err
		}
		if fieldExists {
			in["priority"] = priority
		}

		weight, fieldExists, err := util.ToInt(in, "weight")
		if err != nil {
			return nil, err
		}
		if fieldExists {
			in["weight"] = weight
		}

	default:
		return d, nil
	}

	return in, nil
}
