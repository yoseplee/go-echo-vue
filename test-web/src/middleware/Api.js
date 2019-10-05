import axios from 'axios'

const baseURL = '/api'
const RestApi = class {
    constructor() { }
    static getInstance() {
        if(RestApi.instance === undefined) RestApi.instance = new RestApi()
        return RestApi.instance
    }

    async result(promise) {
        const response = await promise
        const data = response.data
        if(!DataTransfer.success) throw data.err

        return response
    }

    async getTest() {return await this.result(axios.get(`${baseURL}`))}
    // async getTest() {return await this.result(axios.get("http://localhost:1323/api"))}
    // async getTest(params) {return await this.result(axios.get(`${baseURL}`), {params})}
}

export default RestApi.getInstance