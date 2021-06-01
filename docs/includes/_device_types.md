# Device Types

The DeviceType lets you list all hardware and software integrations available with ACP.

## List available device types

```shell
curl "http://127.0.0.1:8081/api/device_types"
```

> The above command returns a JSON structured like this:

```json
["TALLY_RASP","SMARTVIEW_DUO","ATEM","JVC_HM_660","JVC_REMOTE","HS_50","HYPERDECK","RESTREAM","DISCORD"]
```

### HTTP Request

`GET /api/device_types`

### Response body

The response body will be an array of all the device types available on this instance of ACP.

## Get Device types params

Device types params are parameters needed to request a new device.

```shell
curl "http://127.0.0.1:8081/api/device_types/[ID]"
```

> The above command returns a JSON structured like this:

```json
{
   "ip" : {
      "type" : "ip",
      "required" : true,
      "description" : "Camera IP"
   },
   "password" : {
      "type" : "string",
      "required" : true,
      "description" : "Password for the cam API",
      "default" : "0000"
   },
   "port" : {
      "type" : "number",
      "required" : true,
      "description" : "Camera Port",
      "default" : 80,
      "max" : 65535,
      "min" : 1
   },
   "user" : {
      "type" : "string",
      "required" : true,
      "description" : "Username for the cam API",
      "default" : "prohd"
   }
}
```

### HTTP Request

`GET /api/device_types/[ID]/params`

### Response body

The response body will be an object having required field names as keys and an `Input` as value.

### Input

| Field       | Type           | Available On Types | Description                                                    |
| ----------- | -------------- | ------------------ | -------------------------------------------------------------- |
| type        | Enum           | *                  | Type of this input. One of `ip`, `number`, `string`, `select`. |
| required    | bool           | *                  | True if this field is required to create a device.             |
| description | string?        | *                  | A human readable description of the input.                     |
| default     | any?           | *                  | Default value for this input                                   |
| min         | number         | number             | Minimum value for this input                                   |
| max         | number         | number             | Maximum value for this input                                   |
| Options     | []SelectOption | select             | Options for the select input                                   |

### SelectOption

| Field | Type   | Description                                                   |
| ----- | ------ | ------------------------------------------------------------- |
| value | string | Value that must be sent to the API if this option is selected |
| name  | string | Human readable text for this value                            |
