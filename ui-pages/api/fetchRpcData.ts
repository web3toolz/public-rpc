import {Rpc} from "@/models/rpc";

const API_URL: string = process.env.API_URL || 'https://api-public-rpc.web3toolz.com/';


export async function fetchRpcData(): Promise<Rpc[]> {
    const response: Response = await fetch(API_URL);
    if (!response.ok) {
        throw new Error(`Error fetching data: ${response.statusText}`);
    }
    return await response.json();
}