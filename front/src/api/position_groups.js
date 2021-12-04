import axios from 'axios'
import { cleanError } from './utils'

export default class PositionGroupsClient {
    constructor(url) {
        this._url = url;
    }

    all() {
        return new Promise((resolve, reject) => {
            axios.get(`${this._url}/api/position_groups`)
                .then((response) => {
                    resolve(response.data)
                })
                .catch((error) => {
                    reject(error);
                })
        })
    }

    create(params) {
        return new Promise((resolve, reject) => {
            axios.post(`${this._url}/api/position_groups`, params)
                .then((response) => {
                    resolve(response.data)
                })
                .catch((error) => {
                    reject(cleanError(error))
                })
        })
    }

    destroy(id, destroyMembers) {
        return new Promise((resolve, reject) => {
        let data = {destroy_members: destroyMembers};
        return axios.delete(`${this._url}/api/position_groups/${id}`, {data: data})
            .then(resolve)
            .catch((error) => {
                reject(cleanError(error))
            })
        })
    }
}
