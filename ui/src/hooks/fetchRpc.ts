import {useEffect, useState} from 'react';

const API_URL = process.env.API_URL || 'https://api-public-rpc.web3toolz.com/';


export interface Rpc {
    id: string;
    http: string;
    ws: string;
    provider: string;
    chain: string;
    network: string;
    status: string;
    checkedAt: string;
}

export const useFetchRpcData = () => {
    const [data, setData] = useState<Rpc[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<Error | null>(null);
    useEffect(() => {
        const fetchData = async () => {
            try {
                const response: Response = await fetch(API_URL);
                if (!response.ok) {
                    setError(new Error('Failed to fetch data'));
                    return
                }
                const data: Rpc[] = await response.json();
                setData(data);
            } catch (e: Error) {
                setError(e);
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, []);

    return {data, loading, error};
};