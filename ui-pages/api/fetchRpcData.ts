import {Rpc} from "@/models/rpc";
import axios, {AxiosResponse} from "axios";

const API_URL: string = process.env.API_URL || 'https://api-public-rpc.web3toolz.com';


export async function fetchRpcData({chain, network}: { chain?: string, network?: string }): Promise<Rpc[]> {
    const response: AxiosResponse = await axios.get(API_URL, {params: {chain, network}});
    if (response.status !== 200) {
        throw new Error(`Error fetching data: ${response.statusText}`);
    }
    return await response.data;
}