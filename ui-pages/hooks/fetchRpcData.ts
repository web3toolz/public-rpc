import {useEffect, useState} from 'react';
import {Rpc} from "@/models/rpc";
import {fetchRpcData} from "@/api/fetchRpcData";
import {noop} from "@mantine/core";

export const useFetchRpcData = ({initialState, chain, network}: { initialState: Rpc[], chain?: string, network?: string }) => {
    const [rpcData, setData] = useState<Rpc[]>(initialState || []);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<Error | null>(null);
    useEffect(() => {
        const fetchData = async () => {
            try {
                const data: Rpc[] = await fetchRpcData({chain, network});
                setData(data);
            } catch (e: any) {
                setError(e);
            } finally {
                setLoading(false);
            }
        };
        fetchData().then(noop);
    }, []);

    return {rpcData, loading, error};
};