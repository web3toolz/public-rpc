import {useEffect, useState} from 'react';
import {Rpc} from "@/models/rpc";
import {fetchRpcData} from "@/api/fetchRpcData";

export const useFetchRpcData = ({initialState}: { initialState: Rpc[] }) => {
    const [data, setData] = useState<Rpc[]>(initialState || []);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<Error | null>(null);
    useEffect(() => {
        const fetchData = async () => {
            try {
                const data: Rpc[] = await fetchRpcData();
                setData(data);
            } catch (e: any) {
                setError(e);
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, []);

    return {data, loading, error};
};