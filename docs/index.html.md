---
title: API Reference

language_tabs: # must be one of https://git.io/vQNgJ
  - shell

toc_footers:
  - <a href='https://github.com/johnsucaar/acp'>Source code </a>
  - <a href='https://github.com/slatedocs/slate'>Documentation Powered by Slate</a>

includes:
  - errors

search: true

code_clipboard: true
---

# Introduction

Welcome to the ACP Controp Panel API. You can use this API to control the ACP server and the devices linked to this server.

The source code of this project can be found on [the Github Repository](https://github.com/johnsudaar/acp). Feel free to explore the code and propose pull requests!

# Health Check

## Ping the server

```shell
curl "http://127.0.0.1:8081/api/ping"
```

> The above command returns a JSON structured like this:

```json
{
    "response": "pong"
}
```

This endpoint is used to check the connection to the server.

### HTTP Request

`GET http://127.0.0.1:8081/api/ping`


# Devices

## List all devices

```shell
curl "http://127.0.0.1:8081/api/devices"
```

> The above command returns a JSON structured like this:

```json
[
   {
      "id" : "60b56fc92527fd009002b359",
      "input_ports" : [
         "Input_1",
         "Input_2",
      ],
      "name" : "ATEM",
      "output_ports" : [
         "PGM",
         "AUX",
         "MV"
      ],
      "state" : "Not connected",
      "type" : "ATEM",
      "types" : [
         "switcher"
      ]
   },
   {
      "id" : "60b56fd72527fd009002b35a",
      "input_ports" : [],
      "name" : "CAM 1",
      "output_ports" : [
         "SDI OUT"
      ],
      "state" : "Not connected",
      "type" : "JVC_HM_660",
      "types" : [
         "tally",
         "ptz"
      ]
   }
]
```

### HTTP Request

`GET /api/devices`

### Response body

The response body will be an array of devices.
Each device will have the following fields:

| Field        | Type     | Description                                                     |
| ------------ | -------- | --------------------------------------------------------------- |
| id           | string   | unique ID for this device                                       |
| name         | string   | Name of this device                                             |
| state        | string   | Human readable state of this device                             |
| type         | string   | Type of this device according to the device type API (TODO!)    |
| types        | string   | Types applicable to this device. See type specific APIs (TODO!) |
| input_ports  | []string | List of the device inputs                                       |
| output_ports | []string | List of the device outputs                                      |

## Get a specific device

```shell
curl "http://127.0.0.1:8081/api/devices/[ID]"
```

> The above command returns a JSON structured like this:

```json
{
   "display_opts" : {
      "position" : {
         "x" : 194,
         "y" : 260
      }
   },
   "id" : "60b56fd72527fd009002b35a",
   "input_ports" : [],
   "name" : "CAM 1",
   "output_ports" : [
      "SDI OUT"
   ],
   "state" : "Not connected",
   "type" : "JVC_HM_660",
   "types" : [
      "tally",
      "ptz"
   ]
}
```

### HTTP Request

`GET /api/devices`

### Response body

The response body will be an array of devices.
Each device will have the following fields:

| Field        | Type              | Description                                                     |
| ------------ | ----------------- | --------------------------------------------------------------- |
| id           | string            | unique ID for this device                                       |
| name         | string            | Name of this device                                             |
| state        | string            | Human readable state of this device                             |
| type         | string            | Type of this device according to the device type API (TODO!)    |
| types        | string            | Types applicable to this device. See type specific APIs (TODO!) |
| input_ports  | []string          | List of the device inputs                                       |
| output_ports | []string          | List of the device outputs                                      |
| display_opts | DeviceDisplayOpts | Display option of the device                                    |

### DeviceDisplayOpts

| Field      | Type | Description                                  |
| ---------- | ---- | -------------------------------------------- |
| position.x | int  | X Position of the device in the network view |
| position.y | int  | Y Position of the device in the network view |
