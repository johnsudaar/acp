| Field        | Type              | Description                                                                          |
| ------------ | ----------------- | ------------------------------------------------------------------------------------ |
| id           | string            | unique ID for this device                                                            |
| name         | string            | Name of this device                                                                  |
| state        | string            | Human readable state of this device                                                  |
| type         | string            | Type of this device according to the [device type API](#list-available-device-types) |
| types        | string            | Types applicable to this device. See type specific APIs (TODO!)                      |
| input_ports  | []string          | List of the device inputs                                                            |
| output_ports | []string          | List of the device outputs                                                           |
| display_opts | DeviceDisplayOpts | Display option of the device                                                         |
