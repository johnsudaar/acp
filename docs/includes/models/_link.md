| Field            | Type     | Description                                                                   |
| ---------------- | -------- | ----------------------------------------------------------------------------- |
| id               | string   | unique ID for this link                                                       |
| created_at       | DateTime | date and time of the link creation                                            |
| updated_at       | DateTime | date and time of the last link modification                                   |
| input.device_id  | string   | unique ID of the device that contains the input port connected to this link   |
| input.port       | string   | name of the input port on which the link is connected                         |
| output.device_id | string   | unique ID of the device that contains the output poirt connected to this link |
| output.port      | string   | name of the output port on which the link is connected                        |
