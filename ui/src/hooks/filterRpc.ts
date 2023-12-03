import {useEffect, useState} from 'react';

import {Rpc} from "../models/rpc";


export const useFilterRpcData = (data: Rpc[], query: string): Rpc[] => {
    const [filteredData, setFilteredData] = useState<Rpc[]>([]);

    useEffect(() => {
        if (!query) {
            setFilteredData(data);
            return
        }
        const tempData = [];
        for (let i = 0; i < data.length; i++) {
            const {chain, network, provider, ws, http} = data[i];
            const url = http ? http : ws;

            const chainMatch: boolean = chain.toLowerCase().includes(query.toLowerCase());
            const networkMatch: boolean = network.toLowerCase().includes(query.toLowerCase());
            const providerMatch: boolean = provider.toLowerCase().includes(query.toLowerCase());
            const urlMatch: boolean = url.toLowerCase().includes(query.toLowerCase());
            if (chainMatch || networkMatch || providerMatch || urlMatch) {
                tempData.push(data[i])
            }
        }
        setFilteredData(tempData);
    }, [data, query]);

    return filteredData;
}
